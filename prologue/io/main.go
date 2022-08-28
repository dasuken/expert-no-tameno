package main

import (
	"bytes"
	"fmt"
	"io"
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
	}

	//　シークの挙動
	// io.Reader型の値はSeekによって動いてしまう
}

func Factory() []io.Reader {
	return []io.Reader{
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