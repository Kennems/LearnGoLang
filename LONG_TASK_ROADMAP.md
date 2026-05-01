# Long Task Roadmap

这个文件是仓库后续长期补全计划。

目标不是一次性“补齐所有题目”，而是把这个仓库稳定推进成：

1. 语言基础清晰
2. 面试高频题有代码支撑
3. runtime / GC / 调度有可运行观察样例
4. 每个专题至少包含：
   - `README.md`
   - `*.go` 示例
   - `*_test.go` 测试
   - 必要时补 benchmark / pprof / trace 观测入口

## 当前状态

当前仓库已经覆盖：

- 基础语法：`01_hello` ~ `12_goroutine`
- 第一批补全：`13_error_handling` ~ `18_database_sql`
- 第二批补全：`19_strings_runes_fmt` ~ `26_go_interview_coding`
- 小林清单对照文件：[XIAOLIN_GO_COVERAGE.md](./XIAOLIN_GO_COVERAGE.md)

当前仓库还没有系统补透的核心模块：

- `for range` 细节与 Go 版本差异
- map / channel / sync 的 runtime 级原理
- interface / nil / 比较规则的边界行为
- GMP 调度模型
- 内存分配与逃逸分析
- GC 机制、观测与调优
- Go 代码题完整题单化
- 工程化目录布局、性能分析、race / fuzz / trace

## 执行原则

后续推进遵循这几个规则：

- 每轮优先补“一个完整专题”，不要同时横跳很多小点
- 先补“可运行代码 + 测试”，再补“底层原理说明”
- 每个专题至少要有一个“容易考”的版本和一个“稍深入”的版本
- 优先把高频面试模块做厚：map、channel、sync、context、interface、GMP、GC
- 原理类专题不要只写笔记，必须有可观测程序

## 阶段总览

### Phase 1: Go基础缺口收尾

目标：

- 把小林基础题里还没单独成专题的内容补齐
- 重点处理边界行为和容易混淆的问题

任务：

- [ ] `27_for_range_semantics`
  - 覆盖：
    - range 变量地址
    - Go 1.22+ 行为变化
    - 闭包捕获循环变量
  - 交付：
    - 示例代码
    - 单测
    - 版本差异说明

- [ ] `28_nil_comparison`
  - 覆盖：
    - typed nil
    - interface nil
    - 两个 nil 为什么可能不相等
    - interface 比较 panic 场景
  - 交付：
    - 示例代码
    - 单测

- [ ] `29_function_returns_params`
  - 覆盖：
    - 值传递本质
    - slice/map/channel 传参
    - 多返回值
    - named return + defer
  - 交付：
    - 示例代码
    - table-driven tests

- [ ] `30_fmt_tags_blank_identifier`
  - 覆盖：
    - `%v %+v %#v`
    - struct tag 常见用途
    - `_` 忽略值与匿名导入
  - 交付：
    - 示例代码
    - 单测

### Phase 2: Slice / Map / Channel 深化

目标：

- 把“会用”升级成“知道底层行为和面试答法”

任务：

- [ ] `31_slice_internals`
  - 覆盖：
    - slice header
    - 扩容规则
    - 截断、共享底层数组
    - append 导致的别名问题
  - 交付：
    - 示例代码
    - benchmark
    - 内存别名观察测试

- [ ] `32_map_internals`
  - 覆盖：
    - hmap / bucket 概念
    - 遍历无序
    - key 可比较
    - 并发不安全
    - 扩容时机与渐进扩容
    - 删除后内存语义
  - 交付：
    - 示例代码
    - README 中原理图解
    - 单测

- [ ] `33_channel_internals`
  - 覆盖：
    - hchan 概念
    - 无缓冲/有缓冲
    - 关闭语义
    - select 行为
    - 泄漏场景
  - 交付：
    - 示例代码
    - 单测
    - 若干“错误示例”

### Phase 3: Sync / 并发原语深化

目标：

- 把并发原语从 API 认知提升到“适用场景 + runtime 行为”

任务：

- [ ] `34_sync_mutex_rwmutex`
  - 覆盖：
    - Mutex
    - RWMutex
    - 锁与原子操作对比
    - 争用和错误用法
  - 交付：
    - 示例代码
    - benchmark
    - race 示例

- [ ] `35_sync_once_waitgroup_cond`
  - 覆盖：
    - Once
    - WaitGroup
    - Cond
    - 常见死锁和误用
  - 交付：
    - 示例代码
    - 单测

- [ ] `36_sync_map_atomic_pool`
  - 覆盖：
    - sync.Map 适用场景
    - read/dirty 概念
    - atomic
    - sync.Pool
  - 交付：
    - 示例代码
    - benchmark
    - README 对比说明

### Phase 4: Interface / 反射 / Unsafe

目标：

- 补齐语言层“抽象与运行时表示”的核心知识

任务：

- [ ] `37_interface_runtime`
  - 覆盖：
    - eface / iface
    - 动态类型与动态值
    - 断言与转换
    - 比较规则
  - 交付：
    - 示例代码
    - README 术语解释
    - 单测

- [ ] `38_reflect_advanced`
  - 覆盖：
    - CanSet / Elem
    - 深度比较
    - 标签读取
    - 反射修改字段
  - 交付：
    - 示例代码
    - 单测

- [ ] `39_unsafe_patterns`
  - 覆盖：
    - Pointer / uintptr
    - Offsetof / Sizeof / Alignof
    - 风险示例
  - 交付：
    - 示例代码
    - 单测
    - README 风险说明

### Phase 5: Context / net/http / database/sql 实战化

目标：

- 把“概念题”变成“服务端场景题”

任务：

- [ ] `40_context_patterns`
  - 覆盖：
    - cancellation tree
    - timeout
    - value 使用边界
    - 请求链路传递

- [ ] `41_http_server_patterns`
  - 覆盖：
    - middleware
    - request body
    - timeout
    - graceful shutdown
    - pprof endpoint

- [ ] `42_database_sql_patterns`
  - 覆盖：
    - QueryRow / Query
    - rows.Close
    - transaction
    - context-aware db calls
    - sql.ErrNoRows

### Phase 6: GMP 调度专题

目标：

- 补齐 Go 面试里最难、最值钱的一块

任务：

- [ ] `43_gmp_overview`
  - 覆盖：
    - G / M / P
    - 调度器作用
    - work stealing

- [ ] `44_scheduler_triggers`
  - 覆盖：
    - 发生调度的时机
    - syscall
    - channel 阻塞
    - time.Sleep
    - 抢占

- [ ] `45_g0_m0_p_runtime_notes`
  - 覆盖：
    - m0
    - g0
    - 栈切换概念
    - 为什么需要 P

交付要求：

- 每个目录都要有简化说明
- 至少一个“观察程序”
- 必要时通过 `runtime`、`trace`、`GODEBUG` 辅助观测

### Phase 7: 内存管理与逃逸分析

目标：

- 把“栈/堆/逃逸/分配器”连起来理解

任务：

- [ ] `46_escape_analysis`
  - 覆盖：
    - 常见逃逸场景
    - `-gcflags=-m -m -l`
    - 接口、闭包、返回指针

- [ ] `47_memory_allocator`
  - 覆盖：
    - mcache / mcentral / mheap
    - 小对象 / 大对象
    - tiny allocator

- [ ] `48_memory_leaks`
  - 覆盖：
    - goroutine 泄漏
    - slice 持有大数组
    - timer / ticker
    - map 长期膨胀

### Phase 8: GC 机制、观测与调优

目标：

- 把 GC 从“背答案”变成“能观察、能解释”

任务：

- [ ] `49_gc_basics`
  - 覆盖：
    - tracing vs 引用计数
    - Go 的并发标记清扫
    - STW

- [ ] `50_tricolor_barrier`
  - 覆盖：
    - 三色标记
    - 对象消失问题
    - 写屏障 / 混合写屏障

- [ ] `51_gc_observability`
  - 覆盖：
    - `GODEBUG=gctrace=1`
    - `runtime.ReadMemStats`
    - `debug.ReadGCStats`
    - `runtime.GC`

- [ ] `52_gc_tuning`
  - 覆盖：
    - GOGC
    - 减少分配
    - 对象复用
    - 分析 pause / heap / allocs

### Phase 9: Go 代码题题单化

目标：

- 把代码题集中整理成独立题库，而不是散落在多个目录

任务：

- [ ] `53_coding_printing_patterns`
  - 覆盖：
    - 交替打印
    - abc 顺序打印
    - 奇偶数打印

- [ ] `54_coding_worker_pool`
  - 覆盖：
    - 固定 worker 数
    - 限流执行
    - 超时控制

- [ ] `55_coding_producer_consumer`
  - 覆盖：
    - 多生产者多消费者
    - sync.Cond
    - channel 实现

- [ ] `56_coding_interview_set`
  - 覆盖：
    - 小林清单 12.1 ~ 12.9 全部归档
    - 每题一份 README
    - 每题至少 1 个测试

### Phase 10: 工程化与性能分析

目标：

- 让这个仓库不仅像“题库”，也像“工程训练场”

任务：

- [ ] `57_testing_advanced`
  - 覆盖：
    - subtest
    - helper
    - fuzz
    - coverage

- [ ] `58_benchmark_pprof`
  - 覆盖：
    - benchmark
    - allocs
    - CPU profile
    - heap profile

- [ ] `59_project_layout`
  - 覆盖：
    - `cmd/`
    - `internal/`
    - `pkg/`
    - 配置与依赖注入

- [ ] `60_tooling`
  - 覆盖：
    - `go vet`
    - `race`
    - `trace`
    - `pprof`
    - `embed`

## 每轮推进模板

以后每次继续时，默认按这个顺序推进：

1. 选择路线图中“最前面还没完成的一个目录”
2. 创建目录与 README
3. 写最小示例代码
4. 写测试
5. 跑测试
6. 更新覆盖清单

除非有特殊说明，否则不要重新讨论方向。

## 下一个默认动作

如果没有新的优先级指令，下一轮默认从这里开始：

- `27_for_range_semantics`
- `28_nil_comparison`
- `29_function_returns_params`

## 完成定义

当以下条件满足时，才算这个仓库“第一版补全完成”：

- 小林 Go 清单里的高频模块都有对应专题目录
- runtime / GC / GMP 至少各有一个可运行观测程序
- 代码题 12.1 ~ 12.9 有独立归档
- 根 README 能清楚表达学习路径
- 覆盖清单里 `todo` 只剩下少量可选深化项
