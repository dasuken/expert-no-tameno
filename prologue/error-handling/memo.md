知らない事が多すぎる
* errorのラップ
    * `errors.As` `errors.Is`
    * `fmt.Errorf("%w", err)`
        * %w でUnwrapメソッドが生える。Asで遡れる
    * `errors.Unwrap`
    * スタックトレース
* これらのエラーを正しく設計し、エラーの根元を正しく識別する事でわかり易いログとエラー設計が実現できる
    * エラーをグローバル変数で宣言する理由。テスタビリティの向上 + 原因特定容易性

エラー方針に関する記事
* [https://zenn.dev/nekoshita/articles/097e00c6d3d1c9](https://zenn.dev/nekoshita/articles/097e00c6d3d1c9https://zenn.dev/nekoshita/articles/097e00c6d3d1c9)
* [Goでスタックトレースを上書きせずにエラーをラップする方法](https://tech.liquid.bio/entry/2021/07/02/135816https://tech.liquid.bio/entry/2021/07/02/135816)   
* https://twitter.com/mattn_jp/status/1444675642461605890

# ラップする

```go
type Error struct {
	op string
	err error
}

func (e *Error) Error() string {
	return fmt.Sprintf("op: %s, err: %s", e.op, e.err.Error())	
}
```
*　`type error interface{}`を満たす形で構造体を定義

エラーハンドリングの際はキャストして判定

```go
func handleSignup() error {
	if err := createUser(); err != nil {
		if _, ok := err.(*Error); ok {
			return &Error{
				op: "signup",
				err: err,
			}
		} 
		
		// 自分で定義したエラー
		return err
	}
	
	return nil
}
```

## errors.As
```go
if err := createUser(); err != nil {
    var e *Error
    if errors.As(err, e) {
        return e
    }

    return err
}

func (e *Error) Unwrap() error {
	return e.err
}
```

`Unwrap`が実装されたerror型が、第２引数に渡されたエラーと一致するかを確認。
一致した場合は第２匹数に値が入る

## errors.Is
二つのエラーの値