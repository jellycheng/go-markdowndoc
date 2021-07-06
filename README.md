
## md to html
```


```

## 使用mac或linux系统打包
```
windows包：
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o output/go-markdowndoc.exe main.go  

mac包：
    GOOS=darwin GOARCH=amd64 go build -o output/go-markdowndoc-mac main.go

linux包：
    GOOS=linux GOARCH=amd64 go build -o output/go-markdowndoc main.go

```
