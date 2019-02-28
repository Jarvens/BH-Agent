# BH-Agent


## 推送插件服务

![](https://github.com/Jarvens/image-resource/blob/master/%E6%96%B0%E7%89%88%E6%8E%A8%E9%80%81%E6%9E%B6%E6%9E%84.png)
- 账户资产推送
- 订单推送
- 聊天推送

## Websocket 与Socket

当前Socket推送的问题：
- 指令集不够细化
- 需要手动编解码
- 处理客户端突然断开异常不够及时
- 数据结构不够清晰

新版本采用gRPC:

- 需要手动编写协议文件*.proto
- 不需要手动编解码
- 双向数据流
- pb序列化性能更高
