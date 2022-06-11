# 环境搭建

## 1、克隆Server

```shell
cd $GOPATH/src/
mkdir -p github.com/Path-IM
cd github.com/Path-IM
git clone https://github.com/Path-IM/Path-IM-Server.git -b main --depth 1
```

## 2、依赖

### x86架构

> arm架构参考deploy/docker-compose/x86/dependencies目录

### arm架构

> arm架构参考deploy/docker-compose/arm64/dependencies目录

### 修改配置文件

> 我的内网ip是`192.168.1.98`; 将`10.1.3.12`全局替换为`192.168.1.98`

### 运行

```shell
docker-compose up -d
```

### 进入kafka-ui 创建topic

- im_msg
- im_msg_push_single
- im_msg_push_group
- kick_conn

## 3、Server

### x86架构

> arm架构参考deploy/docker-compose/x86/server目录

### arm架构

> arm架构参考deploy/docker-compose/arm64/server目录

### 修改配置文件

> 我的内网ip是`192.168.1.98`; 将`10.1.3.12`全局替换为`192.168.1.98`

### 运行

```shell
docker-compose up -d
```

> 此时发现msgpush-rpc启动不起来 因为imuser-rpc没有启动

---

# 下一章：[编写imuser-rpc](imuser-rpc.md)