# Cloudgo

**1. 概述**  
    
开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理。

**任务目标**

熟悉 go 服务器工作原理
基于现有 web 库，编写一个简单 web 应用类似 cloudgo。
使用 curl 工具访问 web 程序
对 web 执行压力测试

**2. 开发**

首先，选择一个框架，这里选择的是 martini框架，martini 是一个非常新的 Go 语言的 Web 框架  
martini有许多特性：
使用非常简单
无侵入设计
可与其他 Go 的包配合工作
超棒的路径匹配和路由
模块化设计，可轻松添加工具
大量很好的处理器和中间件
很棒的开箱即用特性
完全兼容 http.HandlerFunc 接口  
   
首先利用go get 安装martini  
   
完成后进行 main 和 server的测试

·main.go  
这里使用老师的main.go的代码 （更改sever的位置）

``` 

    package main

    import (
    "os"

    "./server"  // 更改sever文件夹的位置
    flag "github.com/spf13/pflag"
    )

    const (
      PORT string = "8080"
    )

    func main() {
        port := os.Getenv("PORT")
        if len(port) == 0 {
          port = PORT
     }
        pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
        flag.Parse()
        if len(*pPort) != 0 {
         port = *pPort
     }
        server := server.NewServer()
     server.Run(":" + port)
    }
``` 


·server.go  

```

    package server
        import (
         "github.com/codegangsta/martini" 
    )
    func NewServer(port string) {   
         m := martini.Classic()

    m.Get("/", func(params martini.Params) string {
        return "bingo"
    })

    m.RunOnAddr(":"+port)   
}

```

**3. 运行测试**    

首先进行go run main 然后地址为8080  
会得到一个正在listen的结果  
![main](https://github.com/saltydong/cloudgo/blob/master/pic/main.JPG)  
   
然后打开浏览器访问localhost  , 显示bingo
![bingo](https://github.com/saltydong/cloudgo/blob/master/pic/bingo.JPG)  

·curl测试  
打开另一个控制台，利用curl访问localhost：8080  
![curl](https://github.com/saltydong/cloudgo/blob/master/pic/curl.JPG)   
   
·ab测试  
首先安装ab  
![安装](https://github.com/saltydong/cloudgo/blob/master/pic/ab%20%E5%AE%89%E8%A3%85.JPG)  
   
然后进行测试  
![ab](https://github.com/saltydong/cloudgo/blob/master/pic/ab%E6%B5%8B%E8%AF%95.JPG)  


