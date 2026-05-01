# 31_slice_internals

补齐 slice 底层行为与高频面试点。

覆盖：

- slice header 的三个核心信息：指针、长度、容量
- append 前后底层数组是否共享
- 子切片修改对原切片的影响
- 扩容后别名关系断开
- 简单 benchmark
