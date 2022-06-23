## 执行命令

```bash
[iarno@xxx json]$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: ej
cpu: Intel(R) Xeon(R) CPU E5-2630 v2 @ 2.60GHz
BenchmarkMarshalStd-16                        39          28729452 ns/op         1584273 B/op       7748 allocs/op
BenchmarkMarshalEasyjson-16                   72          15596759 ns/op          841571 B/op       7741 allocs/op
BenchmarkMarshalSonic-16                      66          17749983 ns/op         1868336 B/op       7743 allocs/op
BenchmarkMarshalJsoniter-16                   64          16630098 ns/op         1455315 B/op       7745 allocs/op
BenchmarkUnmarshalStd-16                      94          12446730 ns/op          263219 B/op       5279 allocs/op
BenchmarkUnmarshalEasyjson-16                312           3494945 ns/op          255307 B/op       5273 allocs/op
BenchmarkUnmarshalSonic-16                   235           5323271 ns/op          852935 B/op       5275 allocs/op
BenchmarkUnmarshalJsoniter-16                156           7237369 ns/op         1185091 B/op      20554 allocs/op
PASS
ok      ej      11.866s
[iarno@xxx json]$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: ej
cpu: Intel(R) Xeon(R) CPU E5-2630 v2 @ 2.60GHz
BenchmarkMarshalStd-16                        43          27610997 ns/op         1489298 B/op       7746 allocs/op
BenchmarkMarshalEasyjson-16                   72          14568265 ns/op          840124 B/op       7741 allocs/op
BenchmarkMarshalSonic-16                      74          17366276 ns/op         2028366 B/op       7744 allocs/op
BenchmarkMarshalJsoniter-16                   74          15122561 ns/op         1417441 B/op       7745 allocs/op
BenchmarkUnmarshalStd-16                      99          12098387 ns/op          262674 B/op       5279 allocs/op
BenchmarkUnmarshalEasyjson-16                368           3138297 ns/op          254820 B/op       5273 allocs/op
BenchmarkUnmarshalSonic-16                   249           4625963 ns/op          850076 B/op       5276 allocs/op
BenchmarkUnmarshalJsoniter-16                172           7091668 ns/op         1184619 B/op      20555 allocs/op
PASS
ok      ej      11.929s
[iarno@xxx json]$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: ej
cpu: Intel(R) Xeon(R) CPU E5-2630 v2 @ 2.60GHz
BenchmarkMarshalStd-16                        51          23397591 ns/op         1462250 B/op       7744 allocs/op
BenchmarkMarshalEasyjson-16                   74          14626428 ns/op          842940 B/op       7741 allocs/op
BenchmarkMarshalSonic-16                      68          16152159 ns/op         1631867 B/op       7741 allocs/op
BenchmarkMarshalJsoniter-16                   78          15430692 ns/op         1468790 B/op       7746 allocs/op
BenchmarkUnmarshalStd-16                     100          12138661 ns/op          262563 B/op       5279 allocs/op
BenchmarkUnmarshalEasyjson-16                405           3254779 ns/op          254568 B/op       5273 allocs/op
BenchmarkUnmarshalSonic-16                   255           4638503 ns/op          836535 B/op       5276 allocs/op
BenchmarkUnmarshalJsoniter-16                164           7007530 ns/op         1184839 B/op      20554 allocs/op
PASS
ok      ej      13.568s
```

## 总结
`easyjson`总体性能比较好，但使用需预编译生成对应的json文件。想方便使用的话`Unmarshal`场景推荐使用`sonic`, `marshal`场景推荐使用`jsoniter`。

## 链接
https://www.iarno.cn/article/golang-json/