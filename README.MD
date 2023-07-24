#  星途旅程

## COSMOS
Chronicles of Stellar Mythology and Otherworldly Sagas

### （酝酿中的）游戏特色
- 太空歌剧
- 分布式微服务引擎
- 时间回溯，即时因果推测

### 变因系统
拟成为一个支持时间回溯的多人游戏模式

### 架构
- 内核
  使用异步Python实现，对接前端和后端计算集群，处理“变因”和“历史”
- 后端算力
  将计算密集性功能部署到分布式gRPC微服务中
- 数据服务
  Redis+MongoDB+时序数据库，实现高速缓存和存档功能

### 进度
- [x] 时间系统，用以验证分布式计算
- [x] 位移系统，验证算力
- [ ] 更多丰富游戏性的系统
- [ ] 时间回溯+即时因果（只实现人因部分）

### 如何开发新标量系统

1. 编写runner/something.go
   实现一个Runner结构体
   Runner需要实现v2/type中的Runner接口
   并在type中显式地进行编译时检查
2. 将该Runner添加到NewHandler的列表中去
