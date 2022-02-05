# worker_manager

Manager who manipulates N workers (can activate/deactivate, increase/decrease count of the workers)

```go
workersManager := manager.NewWorkerManager(someService, 10, true)
go workersManager.Process(ctx)
defer workersManager.Close(ctx)
```

### example of work
```json lines
{"level":"info","msg":"[worker-manager] is_active = true","time":"2022-02-05T15:10:20+03:00"}
{"level":"info","msg":"worker 10 started","time":"2022-02-05T15:10:20+03:00"}
{"level":"info","msg":"worker 10 do smth","time":"2022-02-05T15:10:20+03:00"}
{"level":"info","msg":"worker 8 started","time":"2022-02-05T15:10:20+03:00"}
{"level":"info","msg":"worker 8 do smth","time":"2022-02-05T15:10:20+03:00"}
{"level":"info","msg":"worker 9 started","time":"2022-02-05T15:10:20+03:00"}
{"level":"info","msg":"worker 9 do smth","time":"2022-02-05T15:10:20+03:00"}
```

### set new workers count
```go
workersManager.Deactivate()
workersManager.SetNewWorkerCount(1)
workersManager.Activate()
```

```json lines
{"level":"info","msg":"[worker-manager] is_active = false","time":"2022-02-05T15:23:25+03:00"}
{"level":"info","msg":"worker 2 do smth","time":"2022-02-05T15:23:25+03:00"}
{"level":"info","msg":"worker 4 done","time":"2022-02-05T15:23:26+03:00"}
{"level":"info","msg":"worker 1 do smth","time":"2022-02-05T15:23:26+03:00"}
{"level":"info","msg":"worker 2 done","time":"2022-02-05T15:23:28+03:00"}
{"level":"info","msg":"worker 3 done","time":"2022-02-05T15:23:28+03:00"}
{"level":"info","msg":"worker 1 done","time":"2022-02-05T15:23:30+03:00"}
{"level":"info","msg":"[worker-manager] worker pool deactivated","time":"2022-02-05T15:23:30+03:00"}
{"level":"info","msg":"[worker-manager] is_active = true","time":"2022-02-05T15:23:30+03:00"}
{"level":"info","msg":"worker 1 started","time":"2022-02-05T15:23:30+03:00"}
{"level":"info","msg":"worker 1 do smth","time":"2022-02-05T15:23:30+03:00"}
```
### close manager
```json lines
{"level":"info","msg":"worker 8 do smth","time":"2022-02-05T15:09:10+03:00"}
{"level":"info","msg":"worker 5 do smth","time":"2022-02-05T15:09:10+03:00"}
{"level":"info","msg":"worker 6 do smth","time":"2022-02-05T15:09:10+03:00"}
{"level":"info","msg":"worker 6 do smth","time":"2022-02-05T15:09:10+03:00"}
{"level":"info","msg":"worker 6 do smth","time":"2022-02-05T15:09:10+03:00"}
{"level":"info","msg":"[worker-manager] worker pool was closed","time":"2022-02-05T15:09:11+03:00"}
```