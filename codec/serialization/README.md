#### proto 协议工具

#### GoLang benchmark 压测命令
注意执行 bash 命令需要在 proto 同级目录下执行


jsonCodec 压测:
```json
go test -bench='BenchmarkJsonCodec' -benchtime=5000000x -benchmem
```
protobufCodec 压测:
```json
go test -bench='BenchmarkProtobufCodec' -benchtime=5000000x -benchmem
```

