# gin-api

gin-api是一个基于gin开发的API基础框架，让您把更多时间专注在业务开发上。

封装就是要简单、简单、简单！！！

# 第三方
- [gin](https://github.com/gin-gonic/gin)
- [bun](https://github.com/uptrace/bun)
- [zerolog](https://github.com/rs/zerolog)
- [redis](https://github.com/redis/go-redis)
- [swagger](https://github.com/swaggo/gin-swagger)

# 安装
```bash
git clone git@github.com:iaping/gin-api.git
cd gin-api
make
```
或自行go build

# API文档
请先安装[swag](https://github.com/swaggo/swag)
```bash
# swag init
make doc
```
访问：http://127.0.0.1:8080/_doc/index.html

# 解释

### 上下文

ctx.Ctx为整个系统的桥梁（全局资源的上下文）！！！已经初始常用的资源，如数据库、缓存的操作等。自己需要的资源自己注入并初始化即可，方便随时调用。

### API
自己的API请在http.Serve.router()中注入，请查看已存在的简单示例。

# 最后
如果对你有帮助，欢迎star或提交代码完善此项目。