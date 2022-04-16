#### proto 协议工具

#### GoLang benchmark 压测命令
注意执行 bin
jsonCodec 压测:
```json
go test -bench='BenchmarkJsonCodec' -benchtime=5000000x -benchmem
```
protobufCodec 压测:
```json
go test -bench='BenchmarkProtobufCodec' -benchtime=5000000x -benchmem
```

