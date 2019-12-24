# Status and Diagnostics API
A status API exists, which dumps some diagnostic and statistical information, as well as basic information about the underlying Raft node. Assuming the rqlite node is started with default settings, node status is available like so:

```bash
curl localhost:4001/status?pretty
```

The use of the URL param `pretty` is optional, and results in pretty-printed JSON responses.

You can also request the same status information via the CLI:
```
$ ./rqlite 
Welcome to the rqlite CLI. Enter ".help" for usage hints.
127.0.0.1:4001> .status
build:
  build_time: unknown
  commit: unknown
  version: 5
  branch: unknown
http:
  addr: 127.0.0.1:4001
  auth: disabled
  redirect: 
node:
  start_time: 2019-12-23T22:34:46.215507011-05:00
  uptime: 16.963009139s
runtime:
  num_goroutine: 9
  version: go1.13
 ```

## expvar support
rqlite also exports [expvar](http://godoc.org/pkg/expvar/) information. The standard expvar information, as well as some custom information, is exposed. This data can be retrieved like so (assuming the node is started in its default configuration):

```bash
curl localhost:4001/debug/vars
```

You can also request the same expvar information via the CLI:
```
$ rqlite
127.0.0.1:4001> .expvar
cmdline: [./rqlited data]
db:
  execute_transactions: 0
  execution_errors: 1
  executions: 1
  queries: 0
  query_transactions: 0
http:
  backups: 0
  executions: 0
  queries: 0
memstats:
  Mallocs: 8950
  HeapSys: 2.588672e+06
  StackInuse: 557056
  LastGC: 0...
 ```

## pprof support
pprof information is available by default and can be retrieved as follows:

```bash
curl localhost:4001/debug/pprof/cmdline
curl localhost:4001/debug/pprof/profile
curl localhost:4001/debug/pprof/symbol
```
