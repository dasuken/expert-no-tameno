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
  

* go:embed
    * 参考文献　https://future-architect.github.io/articles/20210208/
    * 役割: 外部ファイルを単一バイナリに含める事でアセットを用意する必要がなくなる
    * `//go:embed finename`を宣言して対象ファイルをバイナリ化
        * 本来`os`と`io`使ってバイナリ化する手間が省けるっぽい
    * 単一ファイルの場合
        * `_ embed`でimportするのが慣しらしい
        * `//go:embed filename`で対象ファイルをバイナリに変換
        * (推測)次に現れた変数`[]byte`もしくは`embed.FS`に代入
    * 複数ファイルの場合
        * `embed.FS`構造体の変数に代入する。例では変数名を`static`としていた
            * 静的ファイルとしてバンドルする感じ？
        * 特定のファイルのバイナリデータを取得する場合は、この変数に`RradFile`メソッドを実行すればいいっぽい。
            * イメージはkeyのファイル名、valueにバイナリデータの入ったkvs?
     * go:embedの宣言はグローバルに行う必要があり、main関数内部など閉じたスコープではエラーを吐く
   
## skipした項目
* walk
* tmp file
* tmp dir
