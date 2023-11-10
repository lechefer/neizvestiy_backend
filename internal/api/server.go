package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"smolathon/internal/api/handler"
	"smolathon/internal/service"
	"time"

	_ "smolathon/docs"
	"smolathon/pkg/sctx"
	"smolathon/pkg/slogger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	_defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	Server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func NewServer(
	port int,
	logger slogger.Logger,
	contextProvider sctx.DefaultContextProviderFunc,
	questService *service.QuestService,
	settlementService *service.SettlementService,
	achievementService *service.AchievementService,
) *Server {
	r := gin.New()

	// cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	r.Use(cors.New(corsConfig))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	apiGroup := r.Group("/api")
	{
		serviceGroup := apiGroup.Group("/")
		{
			serviceGroup.GET("/ping", handler.Ping())
		}

		settlementGroup := apiGroup.Group("/settlements")
		{
			settlementGroup.GET("/search", handler.SearchSettlements(logger, settlementService))
		}

		questGroup := apiGroup.Group("/quests")
		{
			questGroup.GET("/list", handler.ListQuests(logger, questService))
			questGroup.GET("/:questId", handler.GetQuest(logger, questService))
		}

		achievementGroup := apiGroup.Group("/achievements")
		{
			achievementGroup.GET("/:accountId/:achievementId", handler.GetAchievement(logger, achievementService))
			achievementGroup.GET("/list", handler.ListAchievements(logger, achievementService))
		}
	}

	return &Server{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: r,
			BaseContext: func(listener net.Listener) context.Context {
				return contextProvider()
			},
		},
		shutdownTimeout: _defaultShutdownTimeout,
	}
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.Server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown(ctx context.Context) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, s.shutdownTimeout)
	defer cancel()

	return s.Server.Shutdown(shutdownCtx)
}
