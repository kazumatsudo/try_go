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


## 読み物

- [他言語プログラマがgolangの基本を押さえる為のまとめ](https://qiita.com/tfrcm/items/e2a3d7ce7ab8868e37f7)
- [動的言語だけやってた僕が、38日間Go言語を書いて学んだこと](https://qiita.com/suin/items/22662f43b6a6e8728798)
- [さあGoを始めよう！環境構築，”Hello World”から簡単なバックエンドサーバーまで](https://postd.cc/how-i-start-go/)