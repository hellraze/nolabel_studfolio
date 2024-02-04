package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"nolabel_studfolio/internal/usecase/tokens"
	"nolabel_studfolio/internal/usecase/users"
	"os"
)

type POSTUserHandler struct {
	createUserUseCase *users.CreateUserUseCase
	readUserUseCase   *users.ReadUserUseCase
}

func NewPOSTUserHandler(createUseCase *users.CreateUserUseCase, readUseCase *users.ReadUserUseCase) *POSTUserHandler {
	return &POSTUserHandler{
		createUserUseCase: createUseCase,
		readUserUseCase:   readUseCase,
	}
}

type POSTUserRequest struct {
	Name         string
	Email        string
	Password     []byte
	About        string
	Education    string
	Experience   string
	Urls         []string
	RoleID       int
	ProjectsList []int
}

type POSTUserResponse struct {
	SignedToken string `json:"signedToken"`
}

func (handler *POSTUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var body POSTUserRequest
	secretKey := os.Getenv("SECRET_KEY")
	ctx := request.Context()
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	command := &users.CreateUserCommand{
		Name:         body.Name,
		Email:        body.Email,
		Password:     body.Password,
		About:        body.About,
		Education:    body.Education,
		Experience:   body.Experience,
		Urls:         body.Urls,
		RoleID:       body.RoleID,
		ProjectsList: body.ProjectsList,
	}
	readCommand := &users.ReadUserCommand{
		Email: body.Email,
	}

	user, err := handler.readUserUseCase.ReadUserHandler(ctx, readCommand)
	if err != nil {
		user, err = handler.createUserUseCase.CreateUserHandler(ctx, command)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		signedToken, err := tokens.NewSignedToken(user.ID(), []byte(secretKey))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		response := &POSTUserResponse{
			SignedToken: signedToken,
		}
		err = json.NewEncoder(writer).Encode(response)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	} else {
		err = errors.New("users already exists")
		http.Error(writer, err.Error(), http.StatusConflict)
	}
}
