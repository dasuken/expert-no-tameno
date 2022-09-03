package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
}

func exampleWithValue() {
	ctx := context.Background()
	valueCtx := context.WithValue(ctx, "key1", "value1")
	fmt.Println(valueCtx.Value("key1").(string)) // interface型が帰ってくる
}

// 10秒ごとにstdoutに出力
// 15秒で終了
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