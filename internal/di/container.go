package di

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"nolabel_studfolio/internal/domain"
	"nolabel_studfolio/internal/handler"
	"nolabel_studfolio/internal/handler/middleware"
	"nolabel_studfolio/internal/usecase/users"
	"os"
)

type Container struct {
	router http.Handler

	userRepository    *domain.UserRepository
	postUserHandler   *handler.POSTUserHandler
	createUserUseCase *users.CreateUserUseCase
}

func (c *Container) HTTPRouter() http.Handler {
	if c.router != nil {
		return c.router
	}
	router := mux.NewRouter()
	router.Use(middleware.Recover)

	publicRouter := router.PathPrefix("/api").Subrouter()
	publicRouter.Handle("/users", c.POSTUserHandler()).Methods(http.MethodPost)
	publicRouter.Handle("/tokens", c.POSTTokenHandler()).Methods(http.MethodPost)

	securedRouter := router.PathPrefix("/api").Subrouter()
	securedRouter.Use(middleware.AuthMiddleware)
	securedRouter.Handle("/accounts", c.POSTAccountHandler()).Methods(http.MethodPost)
	securedRouter.Handle("/accounts", c.GETUserAccountsHandler()).Methods(http.MethodGet)
	securedRouter.Handle("/deposit", c.POSTDepositAccountHandler()).Methods(http.MethodPost)
	c.router = router
	return c.router

}

func CreateConnection(ctx context.Context) (*pgxpool.Pool, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env file not found")
	}
	dns := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(ctx, dns)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return pool, err
}
