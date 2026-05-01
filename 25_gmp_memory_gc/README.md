# 25_gmp_memory_gc

补齐 GMP、内存管理、GC 的入门观察目录。

这个目录先不追求“解释 runtime 全部细节”，而是提供可以实际跑起来的观察入口：

- `runtime.NumGoroutine`
- `runtime.ReadMemStats`
- 主动触发 GC

后续可以继续拆分成：

- `25_gmp_scheduler`
- `26_memory_management`
- `27_gc`
