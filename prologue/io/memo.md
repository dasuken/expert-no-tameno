思い出し

# IO
* バッファ
```go
make([]byte, n)
```

* ReadAtLeast(r io.Reader, buffer, length)

`r.Read(buf, length)`で、指定の長さreaderから読み取りbufferに書き込む。usecaseとしてimgの先頭4バイトを確認し、それがpdfか確認す　など

* 