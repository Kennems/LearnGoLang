# Examples Module

这个目录是从根目录散落示例中整理出来的独立模块，先保留原文件，再提供一份可单独维护的结构化版本。

## 对应关系

- `bench.go` / `bench_test.go` -> `benchmark/`
- `Goroutine.go` -> `concurrency/`
- `demo.go` -> `cmd/print-goroutines/`
- `main.go` -> `database/`
- `Pic.go` -> `image/`
- `pprof.go` -> `pprof/`

## 运行方式

- benchmark 示例: `cd examples && go test ./benchmark -bench .`
- goroutine 打印示例: `cd examples && go run ./cmd/print-goroutines`
- goroutine 平方示例: `cd examples && go run ./cmd/square-workers`
- pprof http 示例: `cd examples && go run ./cmd/pprof-server`
- 图片转换示例资源: `image/Demo.jpg`

`image/` 依赖第三方库，默认通过 `external` build tag 隔离，避免影响仓库常规测试。
