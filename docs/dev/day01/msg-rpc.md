# 上一章：[编写imuser-rpc](imuser-rpc.md)

---

# 编写msg-rpc

## 1、复制代码

```shell
mkdir -p app/msg/rpc
cp -r ~/go/src/github.com/Path-IM/Path-IM-Server/app/msg/cmd/rpc/* app/msg/rpc
```

## 2、修改代码

### 1、包

- `github.com/Path-IM/Path-IM-Server/app/msg/cmd` 替换为你自己的package

> 我替换为`github.com/showurl/Path-IM-Server-OICQ/app/msg`

### 2、msgcallback-rpc

- `app/msg/rpc/internal/svc/serviceContext.go`
  注释代码 `MsgCallback: msgcallbackservice.NewMsgcallbackService(zrpc.MustNewClient(c.MsgCallbackRpc)),`

## 3、修改配置文件

> 修改位置

```shell
mv app/msg/rpc/etc/chat.yaml etc/
```

> 修改文件

```yaml
Name: msg-rpc
ListenOn: :10012
Log:
  ServiceName: msg-rpc
  Level: info
Telemetry:
  Name: msg-rpc
  Endpoint: http://192.168.1.98:14268/api/traces
  Sampler: 1.0
  Batcher: jaeger

ImUserRpc:
  Endpoints:
    - 127.0.0.1:10011
MsgCallbackRpc:
  Endpoints:
    - 127.0.0.1:10030
MessageVerify:
  FriendVerify: false # 只有好友可以发送消息
Callback:
  CallbackBeforeSendGroupMsg:
    Enable: false # 开启群消息发送前回调
    ContinueOnError: true # 开启群消息发送前回调时，如果出错，是否继续发送
  CallbackAfterSendGroupMsg:
    Enable: false # 开启群消息发送后回调
    ContinueOnError: true # 无意义
  CallbackBeforeSendSingleMsg:
    Enable: false # 开启私聊消息发送前回调
    ContinueOnError: true # 开启私聊消息发送前回调时，如果出错，是否继续发送
  CallbackAfterSendSingleMsg:
    Enable: false # 开启私聊消息发送后回调
    ContinueOnError: true # 无意义
Kafka:
  Brokers:
    - 192.168.1.98:9092
  Topic: im_msg
RedisConfig:
  Conf:
    Host: 192.168.1.98:6379
    Pass: "123456"
    Type: node
  DB: 0
Mongo:
  Uri: mongodb://192.168.1.98/admin
  SingleChatMsgCollectionName: "single_chat_msg"
  GroupChatMsgCollectionName: "group_chat_msg"
  DBDatabase: "zeroim"
  DBTimeout: 30
Cassandra:
  Hosts:
    - 192.168.1.98
  Port: 9042
  Keyspace: "zeroim"
  Username: "cassandra"
  Password: "cassandra"
  Consistency: ONE
  SingleChatMsgTableName: "single_chat_msg"
  GroupChatMsgTableName: "group_chat_msg"
  TimeoutSecond: 5
HistoryDBType: cassandra # mongo, cassandra
```

## 4、启动服务

```shell
go run app/msg/rpc/chat.go
```

---

# 下一章：[简单测试](simple-test.md)