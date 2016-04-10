rqlite [![Circle CI](https://circleci.com/gh/otoolep/rqlite/tree/master.svg?style=svg)](https://circleci.com/gh/otoolep/rqlite/tree/master) [![GoDoc](https://godoc.org/github.com/otoolep/rqlite?status.svg)](https://godoc.org/github.com/otoolep/rqlite)
======

*Detailed background on rqlite can be found on [these blog posts](http://www.philipotoole.com/tag/rqlite/)*.

*rqlite* is a distributed system that provides a replicated [SQLite](https://www.sqlite.org/) database. rqlite is written in [Go](http://golang.org/) and uses [Raft](http://raftconsensus.github.io/) to achieve consensus across all the instances of the SQLite databases. rqlite ensures that every change made to the database is made to a quorum of SQLite files, or none at all.

### Why?
rqlite gives you the functionality of a [rock solid](http://www.sqlite.org/testing.html), fault-tolerant, replicated relational database, but with very easy installation, deployment, and operation. With it you can build lightweight and reliable central stores for relational data.

## Getting started
The quickest way to get running on OSX and Linux is to download a pre-built release binary. You can find these binaries on the [Github releases page](https://github.com/otoolep/rqlite/releases). Once installed, you can start a single rqlite node like so:
```bash
rqlited ~/node.1
```
This single node automatically becomes the leader.

If you want to build rqlite, either because you want the latest code or a pre-built binary for platform is not available, take a look at [CONTRIBUTING.md](https://github.com/otoolep/rqlite/blob/master/CONTRIBUTING.md).

### Forming a Cluster
While not strictly necessary to run rqlite, running multiple nodes means the SQLite database is replicated.

Start a second and third node (so a majority can still form in the event of a single node failure) like so:

```bash
rqlited -http localhost:4003  -raft :4004 -join :4001 ~/node.2
rqlited -http localhost:4005  -raft :4006 -join :4001 ~/node.3
```

*(This assumes you've started the first node as instructed in the Getting Started section*

Under each node will be an SQLite file, which should remain in consensus. You can create clusters of any size, but clusters of 3, 5, and 7 nodes are most practical.

### Restarting a node
If a node needs to be restarted, perhaps because of failure, don't pass the `-join` option. Using the example nodes above, if node 2 needed to be restarted, do so as follows:

```bash
rqlited -http localhost:4005 -raft :4006 ~/node.3
```

On restart it will rejoin the cluster and apply any changes to the local sqlite database that took place while it was down. Depending on the number of changes in the Raft log, restarts may take a little while.

## Data API
rqlite exposes an HTTP API allowing the database to be modified such that the changes are replicated. Queries are also executed using the HTTP API, though the SQLite database could be queried directly. Modifications go through the Raft log, ensuring only changes committed by a quorum of rqlite nodes are actually executed against the SQLite database. Queries do not __necessarily__ go through the Raft log, however, since they do not change the state of the database, and therefore do not need to be captured in the log. More on this later.

All responses from rqlite are in the form of JSON.

### Writing Data
To write data successfully to the database, you must create at least 1 table. To do this, perform a HTTP POST, with a `CREATE TABLE` SQL command encapsulated in a JSON array, in the body of the request. For example:

```bash
curl -XPOST 'localhost:4001/db/execute?pretty&timings' -H "Content-Type: application/json" -d '[
    "CREATE TABLE foo (id integer not null primary key, name text)"
]'
```

where `curl` is the [well known command-line tool](http://curl.haxx.se/).

To insert an entry into the database, execute a second SQL command:

```bash
curl -XPOST 'localhost:4001/db/execute?pretty&timings' -H "Content-Type: application/json" -d '[
    "INSERT INTO foo(name) VALUES(\"fiona\")"
]'
```

The response is of the form:

```json
{
    "results": [
        {
            "last_insert_id": 1,
            "rows_affected": 1,
            "time": 0.00886
        }
    ],
    "time": 0.0152
}
```

The use of the URL param `pretty` is optional, and results in pretty-printed JSON responses. Time is measured in seconds. If you do not want timings, do not pass `timings` as a URL parameter.

You can confirm that the data has been writen to the database by accessing the SQLite database directly.

```bash
$ sqlite3 ~/node.3/db.sqlite
SQLite version 3.7.15.2 2013-01-09 11:53:05
Enter ".help" for instructions
Enter SQL statements terminated with a ";"
sqlite> select * from foo;
1|fiona
```

Note that this is the SQLite file that is under `node 3`, which is not the node that accepted the `INSERT` operation.

### Bulk Updates
Bulk updates are supported. To execute multipe statements in one HTTP call, simply include the statements in the JSON array:

```bash
curl -XPOST 'localhost:4001/db/execute?pretty&timings' -H "Content-Type: application/json" -d "[
    \"INSERT INTO foo(name) VALUES('fiona')\",
    \"INSERT INTO foo(name) VALUES('sinead')\"
]"
```

The response is of the form:

```json
{
    "results": [
        {
            "last_insert_id": 1,
            "rows_affected": 1,
            "time": 0.00759015
        },
        {
            "last_insert_id": 2,
            "rows_affected": 1,
            "time": 0.00669015
        }
    ],
    "time": 0.869015
}
```

A bulk update is contained within a single Raft log entry, so the network round-trips between nodes in the cluster are amortized over the bulk update. This should result in better throughput, if it is possible to use this kind of update.

### Querying Data
Querying data is easy. The most important thing to know is that, by default, queries must go through the leader node. More on this later.

For a single query simply perform a HTTP GET, setting the query statement as the query parameter `q`:

```bash
curl -G 'localhost:4001/db/query?pretty&timings' --data-urlencode 'q=SELECT * FROM foo'
```

The response is of the form:

```json
{
    "results": [
        {
            "columns": [
                "id",
                "name"
            ],
            "types": [
                "integer",
                "text"
            ],
            "values": [
                [
                    1,
                    "fiona"
                ],
                [
                    2,
                    "sinead"
                ]
            ],
            "time": 0.0150043
        }
    ],
    "time": 0.0220043
}
```

The behaviour of rqlite when more than 1 query is passed via `q` is undefined. If you want to execute more than one query per HTTP request, perform a POST, and place the queries in the body of the request as a JSON array. For example:

```bash
curl -XPOST 'localhost:4001/db/query?pretty' -H "Content-Type: application/json" -d '[
    "SELECT * FROM foo",
    "SELECT * FROM bar"
]'
```

Another approach is to read the database file directly via `sqlite3`, the command-line tool that comes with SQLite. As long as you can be sure the file you access is under the leader, the records returned will be accurate and up-to-date.

**If you use the query API to execute a command that modifies the database, those changes will not be replicated**. Always use the write API for inserts and updates.

#### Read Consistency
Even though serving queries does not require consensus (because the database is not changed), [queries should generally be served by the leader](https://github.com/otoolep/rqlite/issues/5). Why is this? Because without this check queries on a node could return out-of-date results.  This could happen for one of two reasons:

 * The node, which still part of the cluster, has fallen behind the leader.
 * The node is no longer part of the cluster, and has stopped receiving Raft log updates.

This is why rqlite offers read consistency levels of _none_, _weak_, and _strong_. Each is explained below.

With _none_, the node simply queries its local SQLite file, and does not even check if it is leader. This offers the fastest query response, but suffers from the problems listed above. _Weak_ instructs the node to check that it is the leader, before querying the local SQLite file. Checking leader state only involves checking local state, so is still very fast. There is, however, a very small window of time (milliseconds by default) during which the node may return stale data. This is because after the leader check, but before the local SQLite file is read, another node could be elected leader. As result the node may not be up-to-date with the rest of cluster. To avoid even this possibility, rqlite also offers _strong_. In this mode, rqlite sends the query through Raft consensus system, ensuring that the node remains the leader throughout query processing. However, this will involve the leader contacting at least a quorum of nodes, and will therefore increase query response times.

_Weak_ is probably sufficient for most applications, and is the default read consistency level. To explicitly select consistency, set the query param `level`. Examples of enabling each read consistency level for a simple query is shown below.

```bash
curl -G 'localhost:4001/db/query?level=none' --data-urlencode 'q=SELECT * FROM foo'
curl -G 'localhost:4001/db/query?level=weak' --data-urlencode 'q=SELECT * FROM foo'
curl -G 'localhost:4001/db/query?level=strong' --data-urlencode 'q=SELECT * FROM foo'
```

### Transactions
Transactions are supported. To execute statements within a transaction, add `transaction` to the URL. An example of the above operation executed within a transaction is shown below.

```bash
curl -XPOST 'localhost:4001/db/execute?pretty&transaction' -H "Content-Type: application/json" -d "[
    \"INSERT INTO foo(name) VALUES('fiona')\",
    \"INSERT INTO foo(name) VALUES('sinead')\"
]"
```

When a transaction takes place either both statements will succeed, or neither. Performance is *much, much* better if multiple SQL INSERTs or UPDATEs are executed via a transaction. Note that processing of the request ceases the moment any single query results in an error.

The behaviour of rqlite when using `BEGIN`, `COMMIT`, or `ROLLBACK` to control transactions is **not defined**. It is important to control transactions only through the query parameters shown above.

### Handling Errors
If an error occurs while processing a statement, it will be marked as such in the response. For example.

```bash
curl -XPOST 'localhost:4001/db/execute?pretty&timings' -H "Content-Type: application/json" -d "[
    \"INSERT INTO foo(name) VALUES('fiona')\",
    \"INSERT INTO nonsense\"
]"
```
```json
{
    "results": [
        {
            "last_insert_id": 3,
            "rows_affected": 1,
            "time": 182.033
        },
        {
            "error": "near \"nonsense\": syntax error"
        }
    ],
    "time": 2.478862
}
```

## Performance
rqlite replicates SQLite for fault-tolerance. It does not replicate it for performance. In fact performance is reduced somewhat due to the network round-trips.

Depending on your machine, individual INSERT performance could be anything from 1 operation per second to more than 100 operations per second. However, by using transactions, throughput will increase significantly, often by 2 orders of magnitude. This speed-up is due to the way SQLite works. So for high throughput, execute as many operations as possible within a single transaction.

### In-memory databases
You can also try using an [in-memory database](https://www.sqlite.org/inmemorydb.html) to increase performance. In this mode no actual SQLite file is created and the entire database is stored in memory.

#### Will this put my data at risk?
No.

Using an in-memory does not put your data at risk. Since the Raft log is the authoritative store for all data, and it is written to disk, an in-memory database can be fully recreated on start-up.

Pass `-mem` to `rqlited` at start-up to enable an in-memory database.

## Status API
A status API exists, which dumps some basic diagnostic and statistical information, as well as basic information about the underlying Raft node. Assuming rqlite is started with default settings, rqlite status is available like so:

```bash
curl localhost:4001/status?pretty
```

The use of the URL param `pretty` is optional, and results in pretty-printed JSON responses.

### expvar support
rqlite also exports [expvar](http://godoc.org/pkg/expvar/) information. The standard, and some custom information, is exposed. This data can be retrieved like so:

```bash
curl localhost:4001/debug/vars
```

## Backups
rqlite supports hot-backing up a node. You can retrieve and write a consistent snapshot of the underlying SQLite database to a file like so:

```bash
curl localhost:4001/db/backup -o bak.sqlite3
```

The node can then be restored by loading this database file via `sqlite3` and executing `.dump`. You can then use the output of the dump to replay the entire database back into brand new node (or cluster), *with the exception* of `BEGIN TRANSACTION` and `COMMIT` commands. You should ignore those commands in the `.dump` output.

By default a backup can only be retrieved from the leader, though this check can be disabled by adding `noleader` to the URL as a query param.

## Log Compaction
rqlite automatically performs log compaction. After a fixed number of changes rqlite snapshots the SQLite database, and truncates the Raft log. This is a technical feature of the Raft consensus system, and most users of rqlite need not be concerned with this.

## Limitations
 * SQLite commands such as `.schema` are not handled.
 * The supported types are those supported by [go-sqlite3](http://godoc.org/github.com/mattn/go-sqlite3).

This is new software, so it goes without saying it has bugs. It's by no means finished -- issues are now being tracked, and I plan to develop this project further. Pull requests are also welcome.

## Pronunciation?
How do I pronounce rqlite? For what it's worth I pronounce it "ree-qwell-lite".

## Credits
This project uses the [Hashicorp](https://github.com/hashicorp/raft) implementation of the Raft consensus protocol, and was inspired by the [raftd](https://github.com/goraft/raftd) reference implementation. rqlite also uses [go-sqlite3](http://godoc.org/github.com/mattn/go-sqlite3) to talk to the SQLite database.
