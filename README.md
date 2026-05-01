# LearnGoLang

这个仓库按主题整理 Go 学习示例。

## 结构

- `examples/`:
  从根目录抽离出的独立示例模块，可单独运行和测试。
- `im_system/`:
  基于 TCP 的即时通讯示例。
- `gorm_example/`, `jwt_example/`, `mock_example/`, `zap_example/`:
  依赖第三方库的独立示例，默认通过 `external` build tag 隔离。
- `01_hello` 到 `26_go_interview_coding`, `goroutine*`:
  语言基础和并发练习。
- `XIAOLIN_GO_COVERAGE.md`:
  对照小林 Go 面试题清单的覆盖说明。
- `LONG_TASK_ROADMAP.md`:
  后续长期补全路线图，默认按这个清单推进。

## 测试

- 根模块:
  `GOCACHE=$(pwd)/.gocache GOPROXY=off go test ./08_slice ./10_oop ./11_reflect ./12_goroutine ./13_error_handling ./14_context ./15_select_channel_patterns ./16_sync_advanced ./17_net_http ./18_database_sql ./19_strings_runes_fmt ./20_make_new_nil_struct ./21_interface_advanced ./22_unsafe_escape ./23_map_advanced ./24_channel_advanced ./25_gmp_memory_gc ./26_go_interview_coding ./im_system ./picture ./go_test1`
- examples 模块:
  `cd examples && GOCACHE=$(pwd)/../.gocache go test ./...`
