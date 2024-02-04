package users

import (
	"context"
	"nolabel_studfolio/internal/domain"
)

type ReadUserUseCase struct {
	userRepository domain.UserRepository
}

func NewReadUserUseCase(userRepository domain.UserRepository) *ReadUserUseCase {
	return &ReadUserUseCase{
		userRepository: userRepository,
	}
}

type ReadUserCommand struct {
	Email string
}

func (useCase ReadUserUseCase) ReadUserHandler(ctx context.Context, command *ReadUserCommand) (*domain.User, error) {
	user, err := useCase.userRepository.FindByEmail(ctx, command.Email)
	return user, err
}
