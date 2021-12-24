 ## 用户服务
 

### api 

```shell
goctl api go -api usercenter.api -dir ./../
```


### rpc

```shell
goctl rpc proto -src usercenter.proto -dir ./../
```

社区大佬的写法

```shell
goctl rpc proto -src rpc/pb/*.proto -dir ./rpc -style=goZero


sed -i 's/,omitempty//g' ./rpc/pb/*.pb.go
```