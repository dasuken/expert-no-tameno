# context
* 親子関係にコンテキスト木を作成
* 実行環境の情報を共有するためにしよう
* 中断処理等を伝播し、複数のgoroutine間で一貫した処理を施す

# ルール
```go
func Do(ctx context.Context, arg Arg) error {
}
```

* 構造体に含めない。関数の引数として使う
* 辺数名はctx
* 渡すべきコンテキストが判定できない場合、nilを渡さずにcontext.TODOを渡すべき
* コンテキストに保存する値はリクエストスコープに治る値にする

# example
```go
func main() {
	// cancel付きのコンテキスト生成
	emptyCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(emptyCtx)
	defer cancel()
}
```

* root contextを生成
* withCancel, withDedline / withTimeout, WithValueで３種類のコンテキストを生成
* WithTodoは`ctx.Context`型を満たすのみ

# WithCancel
* `cancel()`関数で`cancelCtx`自身が持つチャネルを閉じる
    * チャネルがとじラエrたことがメッセージ受診をまつすべてのゴルーチンにブロードキャストされる
    * Doneメソッドを参照している全てのゴルーチンにキャンセルされた情報が送られる
 
```go
// キャンセルctx作って
// エラーがあれば拾ってcancel
// 複数のgoroutineで並列処理
// エラー拾って最初のerrorだけ表示
func doSomeThingParallel(workerNum int) error {
	ctx := context.Background()
	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	errCh := make(chan error, workerNum)

	var wg sync.WaitGroup
	for i := 0; i < workerNum ; i++  {
		wg.Add(1)

		go func(i int) {
			if err := doSomeThingWithContext(cancelCtx, 1); err != nil {
				cancel()
				errCh <- err
			}
		}(i)
	}

	wg.Wait()

	close(errCh)
	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}

func doSomeThingWithContext(ctx context.Context, num int) error {
	// 関数のしょっぱなに死活問題
	// 閉じられていなければ通常
	select {
	case <- ctx.Done():
		return ctx.Err()
	default:
	}

	fmt.Println(num)
	return nil
}
```
* workerNumの数だけ並列処理
* エラーが出た場合、キャンセルで全てのgoorutineにブロードキャスト
    * エラーはエラー用のチャネルに入れる
* チャネルを閉じてエラーがあれば先頭

コンテキストはチャネルを通じて処理を伝播。キャンセルするとブロードキャストで全てのコンテキストいん広まるらしい
 
実行される関数は「キャンセル処理が行われた場合」の処理を先頭に記述
* SELECT　caseに特殊終了条件を記述し、引っかからなければ通常の挙動

# Withdeadline
* Deadlineに到達

```go
// 2022 09 11に契約が切れる
// 2022 09 12になるまでまつ
// またはcancel
func exampleWithDeadline() {
	ctx := context.Background()
	d := time.Date(2022, 9, 11, 0, 0, 0, 0, time.UTC)
	deadlineCtx, cancel := context.WithDeadline(ctx, d)
	defer cancel()
	
	nd := d.AddDate(0,0,1)

	select {
	case <- time.Tick(time.Until(nd)):
		fmt.Println("期限切れですー")
	case <- deadlineCtx.Done():
		fmt.Println(deadlineCtx.Err())
	}
}
```
time.Timeとtime.Duratio

# WithTimeout
```go
func exampleWithTimeout() {
	ctx := context.Background()
	d := 15 * time.Second
	timeoutCtx, cancel := context.WithTimeout(ctx, d)
	defer cancel()

	select {
	case <- time.Tick(10*time.Second):
		fmt.Println("10秒たちました")
	case <- timeoutCtx.Done():
		fmt.Println("end")
 	}
}
```

# WithValue
*　メソッドチェーンよろしく、親contextの値しか参照できない
* usecaseとしてはリクエストidなど。基本的に1リクエスト1contextを割り振るため、ライフサイクルは同じ。