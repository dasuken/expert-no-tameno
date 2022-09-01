# channelとmutex
* mutex: 複数のgoroutineでクリティカルセクションを保護し、原子性を担保するための仕組み
* channel: メモリアクセスではなく、communicate

何回やっても忘れる自信がある。usecaseが欲しい

# channel

## バッファなし
* 受け取りでブロック
* goroutineの中で送信
* closeで終了

```go
func main() {
	ch := make(chan int)
	
	go func() {
		ch <- 100
	}()
	
	fmt.Println(<-ch)
	close(ch)
}
```

select

```go
select {
case n, ok := <-ch1:
    fmt.Println(n,ok)
case n,ok := <-ch2:
    fmt.Println(n, ok)
}
```

range
* closeするまで繰り返し受診してくれる

```go
func main() {
	ch1 := make(chan int)

	go func() {
		ch1 <- 100
		ch1 <- 200
		close(ch1)
	}()

	for n := range ch1  {
		fmt.Println(n)
	}
}
```

一方向
```go

func main() {
	ch := make(chan int)
	go send(ch, 100)
	fmt.Println(receive(ch))
}

func send(ch chan <- int, n int) {
	ch <- n
}

func receive(ch <- chan int) int {
	return <- ch
}
```