# Performance

rqlite replicates SQLite for fault-tolerance. It does not replicate it for performance. In fact performance is reduced due to a standalone SQLite database due to nature of distributed systems. 

rqlite performance -- defined as the number of database updates performed in a given period of time -- is primarily determined by two factors:
- Disk performance
- Network latency

Depending on your machine (particularly its IO performance) and network, individual INSERT performance could be anything from 10 operations per second to more than 200 operations per second.

## Disk
Disk performance is the single biggest determinant of rqlite performance. This is because every change to the system must go through the Raft subsystem, and the Raft subsystem calls `fsync()` after every write to its log. Raft does this to ensure that the change is safely persisted in permanent storage before applying those changes to the SQLite database. This is why rqlite runs with an in-memory database by default, as using as on-disk SQLite database would put even more load on the disk, reducing the disk throughput available to Raft.

## Network
When running a rqlite cluster, network latency is also a factor. This is because Raft must contact every node **twice** before a change is committed to the Raft log. Obviously the faster your network, the shorter time contact each node will take.

# Improving Performance

There are a few ways to improve performance, but not all will be suitable for a given application.

## Batching
The more SQLite statements you can include in a single request to a rqlite node, the better the system will perform. 

By using the [bulk API](https://github.com/rqlite/rqlite/blob/master/DOC/BULK.md), transactions, or both, throughput will increase significantly, often by 2 orders of magnitude. This speed-up is due to the way Raft and SQLite work. So for high throughput, execute as many operations as possible within a single transaction.

## Use more powerful hardware
Obviously running rqlite on better disks, better networks, or both, will improve performance.

## Use a memory-backed filesystem
It is possible to run rqlite entirely on-top of a memory-backed file system. This means that **both** the Raft log and SQLite database would be stored in memory only. For example, on Linux you can create a memory-based filesystem like so:
```bash
mount -t tmpfs -o size=512m tmpfs /mnt/ramdisk
```
**This comes with risks, however**. The behavior of rqlite when a node fails, but the Raft log has not been permanently persisted, **is not defined**. But if your policy is to completely deprovision your rqlite node, or rqlite cluster, in the event of any node failure, this option might work for you. Testing shows that using rqlite with a memory-only file system can result in 100x improvement in performance.

## Improving read-write concurrency
SQLite can offer better concurrent read and write support when using an on-disk database, compared to in-memory databases. But as explained above, using an on-disk SQLite database can significant impact performance. Of course, the database-update performance will be so much better with an in-memory database, that improving read-write concurrency may not be needed in practise.

However if you enable an on-disk SQLite database, but then place the SQLite database on a memory-backed file system, you can have the best of both worlds. You can dedicate your disk to the Raft log, but still get better read-write concurrency with SQLite.

An alternative approach would be to place the SQLite on-disk database on a different disk, but this is unlikely to be as performant as an in-memory file system.