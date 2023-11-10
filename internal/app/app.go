package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/http"
	"smolathon/config"
	"smolathon/internal/api"
	"smolathon/internal/repository"
	"smolathon/internal/service"
	"smolathon/pkg/sctx"
	"smolathon/pkg/slogger"
	"smolathon/pkg/spostgres"
)

type App struct {
	Server *api.Server

	logger          slogger.Logger
	config          config.Config
	contextProvider sctx.DefaultContextProviderFunc
	notify          chan struct{}

	httpClient *http.Client

	// db
	pgDb     *spostgres.Postgres
	s3client *minio.Client

	// repos
	questRepo       *repository.QuestRepository
	settlementRepo  *repository.SettlementRepository
	achievementRepo *repository.AchievementRepository

	// services
	questService      *service.QuestService
	settlementService *service.SettlementService

	imageService *service.ImageService

	achievementService *service.AchievementService
}

func NewApp(l slogger.Logger, config config.Config, contextProvider sctx.DefaultContextProviderFunc) (*App, error) {
	app := &App{
		logger:          l,
		config:          config,
		contextProvider: contextProvider,
		notify:          make(chan struct{}, 1),
	}

	if err := app.initDatabases(); err != nil {
		return app, err
	}

	if err := app.initServices(); err != nil {
		return app, err
	}

	return app, nil
}

func (a *App) initDatabases() error {
	var err error

	// Postgres
	if a.pgDb, err = spostgres.NewWithMigration(
		a.contextProvider(),
		a.config.Databases.Postgres.ConnectionString,
		a.config.Databases.Postgres.MigrationPath,
	); err != nil {
		return fmt.Errorf("database: %v", err)
	}

	// S3
	if a.s3client, err = minio.New(a.config.S3.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(a.config.S3.AccessKeyId, a.config.S3.SecretKey, ""),
		Secure: false,
	}); err != nil {
		return fmt.Errorf("s3: %w", err)
	}

	a.questRepo = repository.NewQuestRepository(a.pgDb)
	a.settlementRepo = repository.NewSettlementRepository(a.pgDb)
	a.achievementRepo = repository.NewAchievementRepository(a.pgDb)

	return nil
}

func (a *App) initServices() error {
	a.imageService = service.NewService(a.s3client, a.config.S3.Buckets.Main)

	a.questService = service.NewQuestService(a.imageService, a.questRepo)
	a.settlementService = service.NewSettlementService(a.settlementRepo)
	a.achievementService = service.NewAchievementService(a.achievementRepo)
	a.Server = api.NewServer(
		a.config.Port,
		a.logger,
		a.contextProvider,
		a.questService,
		a.settlementService,
		a.achievementService,
	)

	return nil
}

func (a *App) Run() {
	a.Server.Start()

	go func() {
		select {
		case err := <-a.Server.Notify():
			if !errors.Is(err, http.ErrServerClosed) {
				a.logger.Error(err.Error())
			}
		}
		a.notify <- struct{}{}
	}()
}

func (a *App) Notify() <-chan struct{} {
	return a.notify
}

func (a *App) Stop(ctx context.Context) {
	// Services
	if a.Server != nil {
		if err := a.Server.Shutdown(ctx); err != nil {
			a.logger.Errorf("Server: %v", err)
		}
	}

	// DB
	if a.pgDb != nil {
		if err := a.pgDb.Close(); err != nil {
			a.logger.Errorf("pg database: %v", err)
		}
	}
}
