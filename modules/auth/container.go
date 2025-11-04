package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/config"
	"github.com/manab-pr/evtaarpro/internal/datastore"
	"github.com/manab-pr/evtaarpro/modules/auth/domain/usecases"
	"github.com/manab-pr/evtaarpro/modules/auth/infra/postgresql"
	"github.com/manab-pr/evtaarpro/modules/auth/infra/redis"
	"github.com/manab-pr/evtaarpro/modules/auth/infra/security"
	"github.com/manab-pr/evtaarpro/modules/auth/presentation/http/handlers"
	"github.com/manab-pr/evtaarpro/modules/auth/presentation/http/routes"
)

// RegisterRoutes registers auth module routes
func RegisterRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	// Infrastructure
	userRepo := postgresql.NewUserRepository(pgStore.DB)
	passwordHasher := security.NewBcryptHasher()
	tokenGenerator := security.NewJWTGenerator(
		cfg.JWT.Secret,
		cfg.JWT.Issuer,
		cfg.JWT.AccessTokenExpiry,
		cfg.JWT.RefreshTokenExpiry,
	)
	sessionStore := redis.NewSessionStore(redisStore, cfg.JWT.RefreshTokenExpiry)

	// Use cases
	registerUseCase := usecases.NewRegisterUseCase(userRepo, passwordHasher)
	loginUseCase := usecases.NewLoginUseCase(userRepo, passwordHasher, tokenGenerator, sessionStore)
	logoutUseCase := usecases.NewLogoutUseCase(sessionStore)

	// Handlers
	registerHandler := handlers.NewRegisterHandler(registerUseCase)
	loginHandler := handlers.NewLoginHandler(loginUseCase)
	logoutHandler := handlers.NewLogoutHandler(logoutUseCase)

	// Register routes
	routes.RegisterRoutes(rg, registerHandler, loginHandler, logoutHandler, cfg.JWT.Secret)
}
