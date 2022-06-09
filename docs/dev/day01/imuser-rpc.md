# 上一章：[搭建开发环境](env.md)

---

# 初始化项目

```shell
go mod init github.com/showurl/Path-IM-Server-OICQ
```

## 1、编写imuser-rpc

### 初始化imuser-rpc

```shell
mkdir -p app/imuser/rpc
cp -r ~/go/src/github.com/Path-IM/Path-IM-Server/goctl .
cp -r ~/go/src/github.com/Path-IM/Path-IM-Server/app/im-user/cmd/rpc/pb app/imuser/rpc
cd app/imuser/rpc/pb 
bash gencode.sh
go mod tidy
```

### 转移配置文件目录

```shell
mkdir etc
mv app/imuser/rpc/etc/imuser.yaml etc/
```

### 运行imuser-rpc测试

```shell
go run app/imuser/rpc/imuser.go
#showurl@bogon Path-IM-Server-OICQ % go run app/imuser/rpc/imuser.go
#Starting rpc server at 127.0.0.1:8080...
#{"level":"warn","ts":"2022-06-09T10:26:57.388+0800","logger":"etcd-client","caller":"v3@v3.5.4/retry_interceptor.go:62","msg":"retrying of unary invoker failed","target":"etcd-endpoints://0x1400061bc00/127.0.0.1:2379","attempt":0,"error":"rpc error: code = DeadlineExceeded desc = latest balancer error: last connection error: connection error: desc = \"transport: Error while dialing dial tcp 127.0.0.1:2379: connect: connection refused\""}
#{"@timestamp":"2022-06-09T10:26:57.388+08:00","caller":"zrpc/server.go:90","content":"context deadline exceeded","level":"error"}
#panic: context deadline exceeded
#
#goroutine 1 [running]:
#github.com/zeromicro/go-zero/zrpc.(*RpcServer).Start(0x1042f7028?)
#        /Users/showurl/go/pkg/mod/github.com/zeromicro/go-zero@v1.3.4/zrpc/server.go:91 +0x84
#main.main()
#        /Users/showurl/go/src/github.com/showurl/Path-IM-Server-OICQ/app/imuser/rpc/imuser.go:39 +0x26c
#exit status 2
#showurl@bogon Path-IM-Server-OICQ % 
```

### 修改配置文件

> jaeger地址根据实际情况修改

```yaml
Name: imuser-rpc
ListenOn: :10011
Log:
  ServiceName: imuser-rpc
  Level: info
Telemetry:
  Name: imuser-rpc
  Endpoint: http://192.168.1.98:14268/api/traces
  Sampler: 1.0
  Batcher: jaeger
```

### 修改msgpush-rpc.yaml

```yaml
ImUserRpc:
  Endpoints:
    - 192.168.1.98:10011
```

### 启动imuser-rpc

```shell
go run app/imuser/rpc/imuser.go
```

### 尝试重启docker中的msgpush-rpc

```shell
docker restart server-msgpush-rpc-1
```

> 启动成功
 
---

# 下一章：[编写msg-rpc](msg-rpc.md)