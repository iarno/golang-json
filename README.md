## 执行命令

```bash
[iarno@xxxx json]$ go test -bench=. -benchmem -count=3
goos: linux
goarch: amd64
pkg: ej
cpu: Intel(R) Xeon(R) CPU E5-2630 v2 @ 2.60GHz
BenchmarkMarshalStd-16                        50          28561013 ns/op         1505537 B/op       7745 allocs/op
BenchmarkMarshalStd-16                        49          24422802 ns/op         1539423 B/op       7747 allocs/op
BenchmarkMarshalStd-16                        44          24585770 ns/op         1510562 B/op       7746 allocs/op
BenchmarkMarshalEasyjson-16                   78          15516738 ns/op          840841 B/op       7741 allocs/op
BenchmarkMarshalEasyjson-16                   67          17627544 ns/op          840016 B/op       7742 allocs/op
BenchmarkMarshalEasyjson-16                   69          15311299 ns/op          840791 B/op       7741 allocs/op
BenchmarkMarshalSonic-16                      67          18134417 ns/op         1965089 B/op       7744 allocs/op
BenchmarkMarshalSonic-16                      72          17689474 ns/op         1865333 B/op       7742 allocs/op
BenchmarkMarshalSonic-16                      70          17238276 ns/op         1791502 B/op       7743 allocs/op
BenchmarkMarshalJsoniter-16                   63          16762821 ns/op         1451630 B/op       7746 allocs/op
BenchmarkMarshalJsoniter-16                   70          17972752 ns/op         1403577 B/op       7744 allocs/op
BenchmarkMarshalJsoniter-16                   66          16069798 ns/op         1416747 B/op       7744 allocs/op


BenchmarkUnmarshalStd-16                      93          12613253 ns/op          263340 B/op       5279 allocs/op
BenchmarkUnmarshalStd-16                      97          12310017 ns/op          262881 B/op       5279 allocs/op
BenchmarkUnmarshalStd-16                      82          13161649 ns/op          264768 B/op       5279 allocs/op
BenchmarkUnmarshalEasyjson-16                394           3476542 ns/op          254637 B/op       5273 allocs/op
BenchmarkUnmarshalEasyjson-16                348           3558817 ns/op          254982 B/op       5273 allocs/op
BenchmarkUnmarshalEasyjson-16                322           3393389 ns/op          255211 B/op       5273 allocs/op
BenchmarkUnmarshalSonic-16                   250           5647755 ns/op          849839 B/op       5275 allocs/op
BenchmarkUnmarshalSonic-16                   232           4930591 ns/op          846819 B/op       5276 allocs/op
BenchmarkUnmarshalSonic-16                   238           5219573 ns/op          819086 B/op       5276 allocs/op
PASS
ok      ej      29.946s
```

## 总结
`easyjson`总体性能比较好，但使用需预编译生成对应的json文件。`jsoniter`umarshal使用不方便，如果要方便使用推荐`sonic`。

## 链接
https://www.iarno.cn/article/golang-json/