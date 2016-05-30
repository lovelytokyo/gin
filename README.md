#gin
https://github.com/gin-gonic/gin

- [httprouter](https://github.com/julienschmidt/httprouter)をカスタマイズしている
 - ルーティング処理がもっとも早いフレームワーク
     - [Benchmark](https://github.com/gin-gonic/gin/blob/develop/BENCHMARKS.md)
 - unitテストが完了していて、安定している
 - APIが凍結されてるため、バージョンが上がったことによる被害を回避できる 
 - ~~Graceful restart、stopができる~~ [endless](https://github.com/fvbock/endless) 使ってる
 - リクエストGET, POST, PUT, DELETE,PATCH, OPTIONS が使える
 - Middlewareのカスタマイズができる

## サーバ起動

``` test.go
    package main
    
    import "github.com/gin-gonic/gin"

    func main() {
    	router := gin.Default();
    	router.GET("/ping", func(c *gin.Context) {
    		c.JSON(200, gin.H{
    			"message": "pong",
    		})
    	})
    	router.Run();
    }
```

## API

```api_get.go
	router.GET("welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
```
```api_post.go
     /* Multipart/Urlencoded Form */
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})

    /* query + post form */
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		c.JSON(200, gin.H{
			"id": id,
			"page": page,
			"name": name,
			"message": message,
		})
	})
```

```api_put.go
    router.PUT("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		name := c.PostForm("name")

		// なんらかの更新処理

		c.JSON(200, gin.H{
			"status": "putted",
			"id": id,
			"name": name,
		})
	})
```
```api_delete.go
    router.DELETE("/user/:id", func(c *gin.Context) {
		id := c.Param("id")

		// なんらかの削除処理

		c.JSON(200, gin.H{
			"status": "deleted",
			"id": id,
		})
	})
```

```file.go
    router.POST("/upload", func(c *gin.Context){
		file, header, err := c.Request.FormFile("upload")

		filename := header.Filename
		fmt.Println(header.Filename)
		out, err := os.Create("./tmp/"+filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
	})
```

## ルーティング

``` routing.go
	// match /user/john but will not match neither /user/ or /user
	router.GET("user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	//  match /user/john/ and also /user/john/send
	router.GET("/user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)

	})
```

グルーピング・ルーティング

```group.go
    v1 := router.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			fmt.Println("router group 1 /login ")
		})
		v1.POST("/submit", func(c *gin.Context) {
			fmt.Println("router group 1 /submit ")
		})
		v1.POST("/read", func(c *gin.Context) {
			fmt.Println("router group 1 /read ")
		})
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			fmt.Println("router group 2 /login ")
		})
		v2.POST("/submit", func(c *gin.Context) {
			fmt.Println("router group 2 /submit ")
		})
		v2.POST("/read", func(c *gin.Context) {
			fmt.Println("router group 2 /read ")
		})
	}
```

## ミドルウェアのカスタマイズ

```middleware.go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func (c *gin.Context) {
		t := time.Now()

		// set example valiable
		c.Set("example", "12345")

		c.Next()

		// after request
		latency := time.Since(t)
		log.Println(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main()  {
	r := gin.New()
	r.Use(Logger()) // custom middlware

	r.GET("/test", func (c *gin.Context) {
		example := c.MustGet("example").(string)

		log.Println(example)
	})

	r.Run(":8080")
}

```

実行
```
$ go run gin/middleware.go                                                                                                                                                                   
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /test                     --> main.main.func1 (2 handlers)
[GIN-debug] Listening and serving HTTP on :8080
12345
18.133µs
200
```

## 認証
```auth.go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var secrets = gin.H{
	"foo": gin.H{
		"email": "foo@bar.com",
		"phone": "123433",
	},
	"austin": gin.H{
		"email": "austin@bar.com",
		"phone": "66666",
	},
}

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "token required"})
			c.Abort()
		}

		os.Setenv("TOKEN", "tokentest")
		if token != os.Getenv("TOKEN") {
			c.Redirect(http.StatusMovedPermanently, "/readme")
			c.Abort()
		}
		c.Next()
	}
}

func main () {
	r := gin.Default()

	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo": "bar",
		"austin": "1234",
		"lena": "hello2",
		"manu": "4321",
	}))

	 // basic auth
	authorized.GET("/secrets", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// custom auth
	designer := r.Group("/designer", TokenAuth())
	designer.GET("/news", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"news": "hogehoge"})
	})

	// allowed no auth
	r.GET("/readme", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"readme": "thanks"})
	})

	r.Run(":8080")
}
```

## Custom HTTP Configuration

```http.go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()

	s := &http.Server {
		Addr: ":3000",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 <<20,
	}
	router.GET("/", func (c *gin.Context) {
		time.Sleep(2000)

		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	s.ListenAndServe()

}
```

## ユーザ入力値のモデルバインディングとバリデーション
```model.go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User		string `form:"user" json:"user" binding:"required"`
	Password	string `form:"password" json: "password" binding:"required"`
}

func main()  {
	router := gin.Default()

	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if c.BindJSON(&json) == nil {
			if json.User == "manu" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			}else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})

	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		if c.Bind(&form) == nil {
			if form.User == "manu" && form.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			}else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})

	router.Run()
}

```

FORM POSTリクエストで確認
```
$ curl http://localhost:8080/loginForm -X POST -d "user=manu&password=123"
{"status":"you are logged in"}
```

JSON POSTリクエストで確認
```
$ curl -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"user": "manu", "password": "123"}' http://localhost:8080/loginJSON
{"status":"you are logged in"}
```


## XML, JSON, YAML rendering

```rendering.go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// gin.H is a shortcut for map[string]interface{}
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123

		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.Run()
}
```

## HTML rendering, 静的ファイルのルーティング

```
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	router.Run()
}

```
