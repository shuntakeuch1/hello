# testingパッケージの紹介

## usage

```
// 基本的な実行
go test
// 詳細出力
go -v test
// カバレッジプロファイルを出力
go test -coverprofile=cover.out
// HTMLに出力
go tool cover -html=cover.out -o cover.html
```


Godoc
```
godoc -http=:8080
```
Example: http://localhost:8080/pkg/github.com/shuntakeuch1/hello/
