
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

## 启动服务
```
以mac包启动为例：
    ./output/go-markdowndoc-mac web -c ./.env.example
    其中 -c是指定配置文件，默认是.env

```

## 浏览器访问
```
http://127.0.0.1:9909/

```
