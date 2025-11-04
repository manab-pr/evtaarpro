package users

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/config"
	"github.com/manab-pr/evtaarpro/internal/datastore"
	"github.com/manab-pr/evtaarpro/modules/users/data/postgresql/repository"
	"github.com/manab-pr/evtaarpro/modules/users/domain/usecases"
	"github.com/manab-pr/evtaarpro/modules/users/presentation/http/handlers"
	"github.com/manab-pr/evtaarpro/modules/users/presentation/http/routes"
)

// RegisterRoutes registers users module routes
func RegisterRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	// Infrastructure
	userRepo := repository.NewUserRepository(pgStore.DB)

	// Use cases
	getUserUC := usecases.NewGetUserUseCase(userRepo)
	listUsersUC := usecases.NewListUsersUseCase(userRepo)
	updateUserUC := usecases.NewUpdateUserUseCase(userRepo)

	// Handlers
	userHandlers := handlers.NewUserHandlers(getUserUC, listUsersUC, updateUserUC)

	// Register routes
	routes.RegisterRoutes(rg, userHandlers, cfg.JWT.Secret)
}
