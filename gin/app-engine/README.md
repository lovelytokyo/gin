# Run
default is debug mode

```
$ goapp serve app-engine/
INFO     2016-05-26 05:00:39,343 devappserver2.py:769] Skipping SDK update check.
INFO     2016-05-26 05:00:39,384 api_server.py:205] Starting API server at: http://localhost:55873
INFO     2016-05-26 05:00:39,387 dispatcher.py:197] Starting module "default" running at: http://localhost:8080
INFO     2016-05-26 05:00:39,388 admin_server.py:116] Starting admin server at: http://localhost:8000
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)
```

release mode
```
$ export GIN_MODE=release                                                                                    (develop)
$ goapp serve app-engine

INFO     2016-05-26 05:07:30,823 module.py:788] default: "GET /ping HTTP/1.1" 200 4
INFO     2016-05-26 05:07:36,633 module.py:788] default: "GET / HTTP/1.1" 200 12
```
