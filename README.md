## 常用命令



### api

```shell
goctl api -o order.api  # 通过goctl生成api文件

goctl api go -api order.api -dir .
```


### rpc

```shell
goctl rpc template -o order.proto

goctl rpc proto -src order.proto -dir .
```

#### model 

```shell
goctl model mysql ddl -src="./*.sql" -dir="./sql/model" -c

```


## 端口



服务 | 端口
---|---
usercenter-rpc | 127.0.0.1:8710
usercenter     | 127.0.0.1:8711
identity.rpc   | 127.0.0.1:8720
identity       | 127.0.0.1:8721










## 例子：

> https://github.com/zeromicro/zero-doc/blob/main/docs/zero/bookstore.md