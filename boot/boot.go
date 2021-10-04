package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Boot struct {
	*gin.Engine
	g           *gin.RouterGroup
	beanFactory *BeanFactory
}

// Ignite 构造函数
func Ignite() *Boot {
	b := &Boot{Engine: gin.New(), beanFactory: NewBeanFactory()}
	b.Use(ErrorHandler())
	config := InitConfig()
	b.beanFactory.setBean(config)
	if config.Server.Html != "" {
		b.LoadHTMLGlob(config.Server.Html)
	}
	return b
}

// Server 启动函数
func (this *Boot) Server() {
	var port int32 = 8080
	if config := this.beanFactory.GetBean(new(SysConfig)); config != nil {
		port = config.(*SysConfig).Server.Port
	}
	getCronTask().Start()
	err := this.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}

// Handle 封装group
func (this *Boot) Handle(httpMethod, relativePath string, handler interface{}) *Boot {
	if h := Convert(handler); h != nil {
		this.g.Handle(httpMethod, relativePath, h)
	}
	return this
}

// Mount 调用Build路由
func (this *Boot) Mount(group string, routes ...IRoute) *Boot {
	this.g = this.Group(group)
	for _, route := range routes {
		route.Build(this)
		// 反射注入
		this.beanFactory.inject(route)
		this.Beans(route)
	}
	return this
}

// Attach 中间件封装
func (this *Boot) Attach(handler Handler) *Boot {
	this.Use(func(context *gin.Context) {
		err := handler.OnRequest(context)
		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.Next()
	})
	return this
}

func (this *Boot) Beans(beans ...interface{}) *Boot {
	this.beanFactory.setBean(beans...)
	return this
}

// Task 定时任务
func (this *Boot) Task(expr string, f func()) *Boot {
	_, err := getCronTask().AddFunc(expr, f)
	if err != nil {
		log.Println(err)
	}
	return this
}
