package router

import (
	"Jinshuzhai-Bookstore/handler/user"
	"errors"
	"github.com/spf13/viper"
	swag "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"net/http"
	"time"

	_ "Jinshuzhai-Bookstore/docs" // docs is generated by Swag CLI, you have to import it.
	"Jinshuzhai-Bookstore/handler/state"
	"Jinshuzhai-Bookstore/router/middleware"

	"github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() {
	gin.SetMode(viper.GetString("mode"))
	g := gin.New()
	Load(
		// Cores
		g,
		/*		middleware.GinLogger(logger),
				middleware.GinRecovery(logger, true),*/
		// Middlewares
		ginzap.Ginzap(zap.L(), time.RFC3339, true),
		ginzap.RecoveryWithZap(zap.L(), true),
		middleware.RequestId(),
	)

	go testPing()

	startListen(g)
}

func testPing() {
	if err := pingServer(); err != nil {
		zap.L().Error("The router has no response, or it might took too long to start up.", zap.Error(err))
	}
	zap.L().Info("The router has been deployed successfully.")
}

func startListen(g *gin.Engine) {
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			zap.L().Info("Start to listening the incoming requests on https address", zap.String("tls.addr", viper.GetString("tls.addr")))
			zap.L().Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}
	zap.L().Info("Start to listening the incoming requests on http address", zap.String("addr", viper.GetString("addr")))
	zap.L().Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/state/health`.
		resp, err := http.Get(viper.GetString("url") + "/state/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		zap.L().Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("connect to the router failed")
}

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The requested resource is not found.")
	})

	// swagger api docs
	g.GET("/swagger/*any", swag.WrapHandler(swaggerFiles.Handler))

	// pprof router
	pprof.Register(g)

	// api for authentication functionalities
	g.POST("/login", user.Login)
	// api for register functionalities
	g.POST("/register", user.Create)
	// The user handlers, requiring authentication
	u := g.Group("/v1/user")
	// use authentication middleware
	u.Use(middleware.AuthMiddleware())
	// user permission middleware
	u.Use(middleware.HasPermission())
	{
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}

	// The health check handlers
	st := g.Group("/state")
	{
		st.GET("/health", state.HealthCheck)
		st.GET("/disk", state.DiskCheck)
		st.GET("/cpu", state.CPUCheck)
		st.GET("/ram", state.RAMCheck)
	}

	return g
}
