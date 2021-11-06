package main

import (
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/fvbock/endless"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/hiagomf/bank-api/server/config"
	"github.com/hiagomf/bank-api/server/config/database"
	"github.com/hiagomf/bank-api/server/logger"
	"github.com/hiagomf/bank-api/server/middleware"
	"github.com/hiagomf/bank-api/server/validations"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var (
		err  error
		logg *zap.Logger
	)

	config.LoadConfig()

	if logg, err = logger.SetupLogger(); err != nil {
		log.Fatal(err)
	}
	defer func() { _ = logg.Sync() }()
	zap.ReplaceGlobals(logg)

	if err = database.OpenConnection(); err != nil {
		zap.L().Fatal("Não foi possível conectar-se ao banco de dados", zap.Error(err))
	}
	defer database.CloseConnections()

	validations.ConfigValidators()

	// configuramos o ambiente do gin antes de fazer qualquer configuracao
	if config.GetConfig().Production {
		gin.SetMode(gin.ReleaseMode)
	}

	group := errgroup.Group{}
	group.Go(func() error {
		return endless.ListenAndServe(config.GetConfig().InternalAddress, internalRouter(logg))
	})
	group.Go(func() error {
		return endless.ListenAndServe(config.GetConfig().ExternalAddress, externalRouter(logg))
	})
	group.Go(func() error {
		return endless.ListenAndServe(config.GetConfig().ExternalPublicAddress, externalPublicRouter(logg))
	})

	if err = group.Wait(); err != nil {
		zap.L().Error("Erro ao inicializar aplicação", zap.Error(err))
	}
}

func externalRouter(logg *zap.Logger) http.Handler {
	r := gin.New()
	r.Use(
		middleware.RequestIdentifier(),
		middleware.GinZap(logg),
		ginzap.RecoveryWithZap(logg, true),
	)

	// v1 := r.Group("v1")

	// records.Router(v1.Group("records"))
	// reports.Router(v1.Group("reports"))

	return r
}

func externalPublicRouter(logg *zap.Logger) http.Handler {
	r := gin.New()
	r.Use(
		middleware.RequestIdentifier(),
		ginzap.RecoveryWithZap(logg, true),
	)

	public := r.Group("public")
	public.Use()

	return r
}

func internalRouter(logg *zap.Logger) http.Handler {
	r := gin.New()
	r.Use(
		middleware.RequestIdentifier(),
		ginzap.RecoveryWithZap(logg, true),
	)

	api := r.Group("api")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

func pprofRouter(r *gin.RouterGroup) {
	prefixRouter := r.Group("debug/pprof")
	prefixRouter.GET("/", pprofHandler(pprof.Index))
	prefixRouter.GET("/cmdline", pprofHandler(pprof.Cmdline))
	prefixRouter.GET("/profile", pprofHandler(pprof.Profile))
	prefixRouter.POST("/symbol", pprofHandler(pprof.Symbol))
	prefixRouter.GET("/symbol", pprofHandler(pprof.Symbol))
	prefixRouter.GET("/trace", pprofHandler(pprof.Trace))
	prefixRouter.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
	prefixRouter.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
	prefixRouter.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
	prefixRouter.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
	prefixRouter.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
	prefixRouter.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
