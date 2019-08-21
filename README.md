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
```
