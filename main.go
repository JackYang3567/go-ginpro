package main

import (
	 "io"
	 "os"
	 "time"
	 "fmt"
	"net/http"
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"ginpro/conf"
	_"ginpro/docs"
    _"ginpro/models"
	"ginpro/controllers"
	"ginpro/middleware"
	
)

func init(){

	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
    gin.DisableConsoleColor()
    // 记录到文件。
    f, _ := os.Create("gin.log")
    //gin.DefaultWriter = io.MultiWriter(f)

    // 如果需要同时将日志写入文件和控制台，请使用以下代码。
    gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	
	
}
// 模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}



// @title XXXX项目名称-Swagger Example API
// @version 1.0
// @description XXXX项目描述
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name token
// @host localhost:9090
// @BasePath /v1 
func main() {
	

	 r := gin.Default() //不使用中间件		
	// r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", 
									"Accept","Access-Control-Allow-Origin","MB-Property-ID"}, //"*",
		ExposeHeaders:    []string{"Content-Length", "Content-Type","Access-Control-Allow-Origin","Access-Control-Allow-Headers"},
		AllowCredentials: true,		
		AllowOriginFunc: func(origin string) bool {
		  return origin == "*"
		},
		 MaxAge: 12 * time.Hour,
	  }))

	 r.LoadHTMLGlob("templates/*")
	//swagger:
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
    
	r.GET("/index", func(c *gin.Context) {
		for k,v :=range c.Request.Header {
			fmt.Println(k,v)
		} 
		c.JSON(http.StatusOK, gin.H{
			"title": "Main website",
			"data":c.Request.Header,
		})
		/*
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
			"data":c.Request.Header,
		})
		*/
	})

//	r.GET("/captcha", func(c *gin.Context)

	captchaRouter := r.Group("/v1")	
	{
		
		captchaRouter.GET("/captcha", controllers.GetCaptcha)
		captchaRouter.GET("/captcha/:captchaId", controllers.GetCaptchaImage)
		captchaRouter.GET("/captcha/verify/:captchaId/:value", controllers.VerifyCaptcha)
	}
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	/*
	// 认证路由组
	// 路由组中间件! 在此例中，我们在 "authorized" 路由组中使用自定义创建的 
	 authorized := r.Group("/v1", AuthRequired())
	 {
		authorized.POST("/singup", controllers.SingUp)
		authorized.POST("/singin", controllers.SingIn)
		authorized.PUT("/changepass", controllers.ChangePass)

		authorized.POST("/token", controllers.CreatedToken)

		// 嵌套路由组
		//==testing := authorized.Group("testing")
		//==testing.GET("/analytics", analyticsEndpoint)
	}
    */
	// 简单的路由组: v1
	tenantRouter := r.Group("/v1")
	tenantRouter.Use(middleware.JWTAuth())	
	{
		tenantRouter.POST("/tenant", controllers.CreatedTenant)
		tenantRouter.PUT("/tenant", controllers.UpdateTenant)
		tenantRouter.GET("/tenant", controllers.GetTenantList)
		tenantRouter.GET("/tenant/:tenant_id", controllers.GetTenantById)
		tenantRouter.GET("/tenant/name/:tenant_name", controllers.GetTenantByName)
		tenantRouter.DELETE("/tenant/:tenant_id", controllers.DeleteTenantById)
	}

	loginRouter := r.Group("/v1")
	{
		loginRouter.POST("/login", controllers.LogIn)
		loginRouter.POST("/singup", controllers.SingUp)
		loginRouter.POST("/singin", controllers.SingIn)
		loginRouter.PUT("/changepass", controllers.ChangePass)

		loginRouter.POST("/token", controllers.CreatedToken)
	}

	userRouter := r.Group("/v1/user")
	{
		userRouter.GET("/:id", controllers.GetUserById)
		userRouter.GET("/email/:email", controllers.GetUserByEmail)
		userRouter.GET("/name/:name", controllers.GetUserByName)
		userRouter.GET("/tenant/:tenant_id", controllers.GetUserByTenantId)
	}
/*
	// 此 handler 将匹配 /user/john 但不会匹配 /user/ 或者 /user
	r.GET("/v1/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		message := name + " is " + conf.RunMode
		c.String(http.StatusOK, message)
		//c.String(http.StatusOK, "Hello %s", message)
	})

	// 此 handler 将匹配 /user/john/ 和 /user/john/send
	// 如果没有其他路由匹配 /user/john，它将重定向到 /user/john/
	r.GET("/v1/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	*/

	// 绑定 JSON ({"name": "manu", "password": "123"})
	r.GET("/v1/test", controllers.Test)

	
    port := ":"+conf.HTTPPort
	r.Run(port) // 监听并在 0.0.0.0:8080 上启动服务
}
