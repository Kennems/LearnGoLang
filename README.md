# LearnGoLang

这个仓库按主题整理 Go 学习示例。

## 结构

- `examples/`:
  从根目录抽离出的独立示例模块，可单独运行和测试。
- `im_system/`:
  基于 TCP 的即时通讯示例。
- `gorm_example/`, `jwt_example/`, `mock_example/`, `zap_example/`:
  依赖第三方库的独立示例，默认通过 `external` build tag 隔离。
- `01_hello` 到 `12_goroutine`, `goroutine*`:
  语言基础和并发练习。

## 测试

- 根模块:
  `GOCACHE=$(pwd)/.gocache GOPROXY=off go test ./08_slice ./10_oop ./11_reflect ./12_goroutine ./im_system ./picture ./go_test1`
- examples 模块:
  `cd examples && GOCACHE=$(pwd)/../.gocache go test ./...`
