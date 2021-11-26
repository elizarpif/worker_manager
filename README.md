# worker_manager

Manager who manipulates N workers (can activate/deactivate, increase/decrease count of the workers)

```go
workersManager := manager.NewWorkerManager(10, true)
go workersManager.Process(context.Background())
defer workersManager.Close()
```

### example of work
```json lines
{"level":"info","msg":"worker 7 do smth","time":"2021-11-26T13:20:19+03:00"}
{"level":"info","msg":"worker 4 do smth","time":"2021-11-26T13:20:19+03:00"}
{"level":"info","msg":"worker 10 do smth","time":"2021-11-26T13:20:19+03:00"}
{"level":"info","msg":"worker 8 do smth","time":"2021-11-26T13:20:19+03:00"}
{"level":"info","msg":"worker 10 do smth","time":"2021-11-26T13:20:19+03:00"}
```

### close manager
```json lines
{"level":"info","msg":"worker 4 done","time":"2021-11-26T13:20:22+03:00"}
{"level":"info","msg":"worker 6 done","time":"2021-11-26T13:20:22+03:00"}
{"level":"info","msg":"worker 1 do smth","time":"2021-11-26T13:20:21+03:00"}
{"level":"info","msg":"worker 1 done","time":"2021-11-26T13:20:22+03:00"}
{"level":"info","msg":"worker 5 do smth","time":"2021-11-26T13:20:21+03:00"}
{"level":"info","msg":"worker 5 done","time":"2021-11-26T13:20:22+03:00"}
{"level":"info","msg":"worker manager deactivated","time":"2021-11-26T13:20:22+03:00"}
{"level":"info","msg":"worker manager was closed","time":"2021-11-26T13:20:22+03:00"}
```