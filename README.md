# solitude
solitude 这是一个长链转化为短链，支持短链跳转，并提供了数据的缓存和持久化的项目。

## 使用方式
1.进入项目目录，使用命令`go run main.go`启动项目

2.设置短链:在浏览器中访问`localhost:5555/shorten?url=https://news.cctv.com/2020/12/08/ARTItP6OrqV93zTT8kxMqKl2201208.shtml?spm=C94212.P4YnMod9m2uD.EfOoEZcMXuiv.9`
得到以下信息{"message":"eexrTL"}

3.拼接url:在浏览器中访问短链`localhost:5555/expand/eexrTL`，结果为跳转到对应长链，DONE！

## 参数配置
Web设置、MySQL设置、Redis设置

配置如下
```yaml
GIN_MODE: debug
Addr: 5555 # 端口号

MySQL:
  User: root
  Password: root
  Host: 127.0.0.1:3306
  DBname: test_data
Redis:
  Host: 127.0.0.1:6379
  Password:
  DB: 0  # use default DB
```

## 项目结构
```
.
├── LICENSE
├── README.md
├── config
│   └── config.go
├── config.yaml
├── dao
│   ├── mysql.go
│   ├── redis.go
│   └── shorturl.go
├── go.mod
├── go.sum
├── handler
│   └── router.go
├── logic
│   └── shorten.go
└── main.go
```
