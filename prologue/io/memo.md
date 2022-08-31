思い出し

# IO

## example
* バッファ
```go
make([]byte, n)
```

* ReadAtLeast(r io.Reader, buffer, length)

`r.Read(buf, length)`で、指定の長さreaderから読み取りbufferに書き込む。usecaseとしてimgの先頭4バイトを確認し、それがpdfか確認す　など

* MultiReader
複数のio.Readerをまとめる(skip)

* 

## 知見
* Streamで扱っている物はStreamのまま扱う
    * [Goでのstreamの扱い方を学ぶ](!https://christina04.hatenablog.com/entry/2017/01/06/190000)
        * ReadAllで全てのデータをメモリにダンプするのはアンチパターン
        * json.Unmarshal => json.NewDecoder
            * やりがちだったーwww
            * gzip.NewEncoder　等
        * io.ReadAll => io.Copy(w, r)
        * bytes.Buffer => io.Pipe
            * 可読性が落ちる場合もあるっぽい
```go
func pipe1() {
    var buf bytes.Buffer

    err := json.NewEncoder(&buf).Encode(&v)

    resp, err := http.Post("example.com", "application/json", &buf)
}

// 内部バッファを持たない
func pipe2(v User) {
    pr, pw := io.Pipe()

    go func() {
        err := json.NewEncoder(pw).Encode(&v)
        pw.Close()
    }()

    resp, err := http.Post("example.com", "application/json", pr)
}
```
  
* Zip爆弾の様に、想定外のサイズのファイルを読み込んでしまうケースも考えられる。LimitReaderしかり上限を決めて読み込む

* io/fsの存在
    * fs.File interfaceには読み込み用のメソッドのみ定義されている
        * `os.File`にはwrite, seekも実装
    * `fs.FileInfo`=`io.FileInfo`のエイリアス
        * 今までややこしかった！
    * `embed.FS`形はfs.FS型を実装しているため、バイナリに埋め込まれたデータをファイルとして扱うことができる。これが`go:embed`らしい
        * そういえば試してなかった。渋川さんがなんか作ってた気がする。チェックしてみる
  
## skipした項目
* walk
* tmp file
* tmp dir
