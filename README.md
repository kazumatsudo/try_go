## 実行

```
// コンパイル && 実行
$ go run hello.go

// コンパイル
$ go build hello.go

// 実行
$ ./hello
```

## コードフォーマット

```
$ go fmt ./
```

## Tips

```
// defer は関数スコープを離れるときに実行する
defer response.Body.Close()

// エラーハンドリング
// 1番目の引数に成功時の返り値、2番目の引数に error インタフェースを返却する
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
// do something with the open *File f

// 並行処理 ゴルーチン
go xxx()

// ゴルーチン間でデータ受け渡し
var ch chan int
var ch <- chan int // 受信専用
var ch chan<- int // 送信専用
```
