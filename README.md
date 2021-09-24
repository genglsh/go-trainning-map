# go-trainning-map
问题：不应该，该错误为查询不到相关数据，应该降级处理，记录错误，返回nil

本仓库实现了一个简单的people 插入和查询服务，在app/people/service/internal/data/people.go文件中写了相关逻辑。


第三周作业：

作业内容在app/people/service/cmd/people/main.go文件当中。
errorgroup是将error waitgroup 和context的功能做了融合，使其能够在wait go协程的同时，提供利用context.cancel函数推出所有协程的方法，同时还能够将遇到的第一个error报告出来。

第四周作业：
该库则为根据krato框架实现的学员注册和查询的服务。
