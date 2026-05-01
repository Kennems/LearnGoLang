# 22_unsafe_escape

补齐 `unsafe.Pointer`、`uintptr` 和逃逸分析入口。

覆盖：

- `unsafe.Pointer` 与 `uintptr` 的相互转换
- `unsafe.Offsetof`
- 逃逸分析命令提示

建议亲自运行：

`go build -gcflags='-m -m -l' ./22_unsafe_escape`
