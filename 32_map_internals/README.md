# 32_map_internals

补齐 map 高频面试点。

覆盖：

- key 必须可比较
- map 遍历顺序不可靠
- 顺序读取 map 的常见写法
- 遍历时删除元素
- map 删除元素后的长度变化

说明：

底层 `hmap/bucket` 结构这里先不直接复刻 runtime 源码，而是先用代码把“外部可观察行为”讲清楚。
