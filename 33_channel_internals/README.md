# 33_channel_internals

补齐 channel 高频面试点。

覆盖：

- 无缓冲 channel 的同步语义
- 有缓冲 channel 的容量与长度
- 关闭 channel 后的读取行为
- 向已关闭 channel 发送的 panic
- `select + timeout` 的典型模式
