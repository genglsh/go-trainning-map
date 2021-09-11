# go-trainning-map
问题：不应该，该错误为查询不到相关数据，应该降级处理，记录错误，返回nil

本仓库实现了一个简单的people 插入和查询服务，在app/people/service/internal/data/people.go文件中写了相关逻辑。
