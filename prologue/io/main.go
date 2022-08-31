package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
)

type MyWriter struct {
	w io.StringWriter
}

func (m *MyWriter) WriteString(s string) (n int, err error) {
	return m.w.WriteString(s)
}

func main() {
	// MyWriter構造体にWriteStringメソッドを実装
	// 引数に書き込みたいstringを渡す
	// メンバ引数としてio.Writerを持たせる

	// io.Readerとio.Writerって中身にデータあるのどっちだっけ
	// てかbufferってどっちの性質も持ってたよね？あれ？バイト配列からバッファ作ってるんだっけ？

	// io.ReadAtLeastで引数のio.Readerのバイト配列の値が期待された物かいなかを確かめよう
	ss := Factory()
	for _, s := range ss {
		fmt.Println(IsHelloWorld(s))
	}

	for _, s := range ss {
		fmt.Println(IsHelloWorld(s))

		Seek(s)

		fmt.Println(IsHelloWorld(s))
	}

	// MultiWriter
	// hash fileに同時に hello worldを書き込んでみる
	// 結果はh.Sum(nil)


}

func SampleMW() {
	f, _ := os.Create("sample.txt")
	h := sha256.New()
	w := io.MultiWriter(f, h)
	w.Write([]byte("hello world"))

	fmt.Printf("%x", h.Sum(nil))
}

func Seek(rs io.ReadSeeker) (int64, error){
	return rs.Seek(0, io.SeekStart)
}

func Factory() []io.ReadSeeker {
	return []io.ReadSeeker{
		strings.NewReader("hello world"),
		strings.NewReader("hello"),
		strings.NewReader("world"),
	}
}

func IsHelloWorld(r io.Reader) (bool, error) {
	expected := []byte("hello world")
	buf := make([]byte, len(expected))
	_, err := io.ReadAtLeast(r, buf, len(expected))
	if err != nil {
		return false, err
	}

	return bytes.Equal(buf, expected), nil
}

// repl
// `package main;import"fmt"; func main(){%s;fmt.Println(%s)}`

// io/fs
// fs.File構造体にはos.Fileと異なり、読み取り専用メソッドしか定義されていない
// os.FileInfoのエイリアスとしてfs.FileInfoが定義されている(あー確かにos.FileとFileInfoの扱いややこしかったかも)
// ? go:embed? そういえば調べて見たかったもの




