# ð¦¸ð¼ââï¸ go-iris-sample

[Iris](https://github.com/kataras/iris) ï¼Goè¨èªãã¬ã¼ã ã¯ã¼ã¯ï¼ã®ä½¿ãæ¹åèç¨ãã³ãã¬ã¼ããªãã¸ããª  


### ð¢ announce

ããã¾ã§Irisã®ä½¿ãæ¹ã"åè"ããããã®ãªãã¸ããªã§ããï¼ãã£ã¬ã¯ããªåã§ä½¿ç¨ããã¦ãã

- å¤æ°
- ã¡ã½ãã
- ãã£ã¬ã¯ããªæ§é 
- ãã¡ã¤ã«
- ãã®ä»è«¸ã...

ã«ã¤ãã¦ï¼ãèªç±ã«å¤æ´ãã¦ãä½¿ããã ãã



## ð Install

### 1. Download [golang](https://go.dev)

### 2. Clone this repository on your working directory

```shell
$ cd $WORK_DIR
$ git clone https://github.com/mkosakana/go-iris-sample.git
$ cd go-iris-sample
$ make install
```



## ð² Get Start


### Route 1 : Without Docker

Only be able to use `/_example-basic-api` or `/_example-basic-view`, because they have NO Data Base connection.

```shell
go-iris-sample $ cd _example-basic-api/ or _example-basic-view/
$ go run main.go
```


### Route 2 : With Docker (Data Base connection â¨ï¼

#### 1. ç«ã¡ä¸ããã³ã³ããã®æå®  

`.air.toml` ãã¡ã¤ã«å `cmd = "go build -o ./tmp/main ./_example-mvc-api/main.go"` ã®ï¼  
`"./_example-mvc-api/main.go"` ã§æå®ããããã£ã¬ã¯ããªï¼ããã©ã«ãã§ã¯ã_example-mvc-apiãï¼ãã³ã³ããã¨ãã¦ç«ã¡ä¸ããããï¼  
**ã_example-mvc-apiã** ã®ç®æã Docker ãä½¿ã£ã¦ä½æ¥­ããããã£ã¬ã¯ããªåã«ç½®æããï¼

#### 2. start Docker

```shell
$ cd go-iris-sample

// 1. docker-composeããã¤ã¡ã¼ã¸ã®ãã«ã
$ make build

// 2. ã³ã³ããã¼ã»DBãããã¯ã°ã©ã¦ã³ãã§ç«ã¡ä¸ã
$ make up
```



## ð TODO
 - [x] make Docker environment
 - [x] DB connection
 - [x] test - `/_example-basic-api`
 - [x] test - `/_example-basic-view`
 - [ ] test - `/_example-mvc-api`


