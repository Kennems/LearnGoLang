# Xiaolin Go Coverage

这份文件用来对照“小林 Golang 面试题”清单，标记当前仓库覆盖情况。

状态说明：

- `done`: 已有对应目录或示例
- `partial`: 有基础示例，但还没覆盖到底层原理或典型面试追问
- `todo`: 当前仓库基本缺失

## 1. Go基础面试题

- 1.1~1.5: `done`
  对应 `01_hello` 到 `08_slice`
- 1.6 for-range 地址变化: `done`
- 1.7 高效拼接字符串: `todo`
- 1.8 defer 执行顺序: `done`
  对应 `07_defer`
- 1.9 rune: `todo`
- 1.10 tag: `done`
  对应 `11_reflect`
- 1.11 `%v %+v %#v`: `done`
- 1.12~1.13 空 struct{}: `todo`
- 1.14 init 执行时机: `done`
  对应 `05_init`
- 1.15~1.17 interface 比较 / nil / 传参: `done`
- 1.18 逃逸分析: `todo`
- 1.19 多返回值实现: `done`
- 1.20 `_` 用法: `done`
- 1.21~1.22 unsafe.Pointer / uintptr: `todo`

## 2. Slice

- 2.1~2.4: `done`
  对应 `08_slice`
- 底层扩容原理仍偏概念化: `partial`

## 3. Map

- 3.1~3.11: `partial`
  有 `09_map`，但缺底层实现、扩容、并发安全、删除与遍历等专题化示例

## 4. Channel

- 4.1~4.11: `partial`
  当前已有 `12_goroutine`、`channel`、`goroutine_channel`、`15_select_channel_patterns`
  但缺关闭语义、阻塞过程、select 机制、内存泄漏等系统化整理

## 5. Sync

- 5.1~5.13: `partial`
  当前已有 `16_sync_advanced`、`goroutine_mutex`、`goroutine_workers`
  但缺 WaitGroup 原理、sync.Map、Mutex 模式、原子操作与锁对比

## 6. Context

- 6.1~6.4: `done`
  对应 `14_context`

## 7. Interface

- 7.1~7.5: `partial`
  有基础接口和反射示例，但还缺 iface/eface、断言 vs 转换、比较规则

## 8. 反射

- 8.1~8.4: `done`
  对应 `11_reflect`

## 9. GMP

- 9.1~9.10: `todo`

## 10. 内存管理

- 10.1~10.6: `todo`

## 11. GC

- 11.1~11.14: `todo`

## 12. Go代码面试题

- 12.1~12.9: `partial`
  当前仓库有若干 goroutine / channel 示例，但没有集中收录成题单目录

## 建议优先级

最优先补齐：

1. `19_strings_runes_fmt`
2. `20_make_new_nil_struct`
3. `21_interface_advanced`
4. `22_unsafe_escape`
5. `23_map_advanced`
6. `24_channel_advanced`
7. `25_gmp_memory_gc`
8. `26_go_interview_coding`
