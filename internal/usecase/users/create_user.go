package users

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"nolabel_studfolio/internal/domain"
	"nolabel_studfolio/internal/pkg"
)

type CreateUserUseCase struct {
	userRepository domain.UserRepository
}

func NewCreateUserUseCase(userRepository domain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}

type CreateUserCommand struct {
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

func (useCase *CreateUserUseCase) CreateUserHandler(ctx context.Context, command *CreateUserCommand) (*domain.User, error) {
	password, err := bcrypt.GenerateFromPassword(command.Password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	passwordHash := string(password)

	user := domain.NewUser(pkg.GenerateID(), command.Name, command.Email, passwordHash, command.About, command.Education, command.Experience, command.Urls, command.RoleID, command.ProjectsList)
	err = useCase.userRepository.Save(ctx, user)
	return user, err
}
