package usecase

import (
	"context"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/rules"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
	"github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore/repository"
)

type UserUsecase struct {
	repo *repository.UserRepository
}

func NewUserUsecase(repo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}
func (usc *UserUsecase) CreateUser(ctx context.Context, user *dto.CreateUserDTO) error {
	err := rules.ValidateCreateUser(user)
	if err != nil {
		return err
	}
	userEntity := dto.ToUserEntity(*user)
	return usc.repo.CreateUser(ctx, *userEntity)
}
func (usc *UserUsecase) UpdateUser(ctx context.Context, updateIt *dto.UpdateUserDTO) error {
	err := rules.ValidateUpdateUser(updateIt)
	if err != nil {
		return err
	}
	return usc.repo.UpdateUser(ctx, *updateIt)
}
func (usc *UserUsecase) DeleteUser(ctx context.Context, deleteIt *dto.DeleteUserDTO) error {
	return usc.repo.DeleteUser(ctx, *deleteIt)
}
func (usc *UserUsecase) GetUserById(ctx context.Context, input *dto.GetUserByIdDTO) (*dto.UserResponseDTO, error) {
	userEntity, err := usc.repo.GetUserById(ctx, input.ID.String())
	if err != nil {
		return nil, err
	}
	user := dto.ToUserResponseDTO(userEntity)
	return &user, nil
}
func (usc *UserUsecase) GetUserByRole(ctx context.Context, input *dto.GetUserByRoleDTO) ([]*dto.UserResponseDTO, error) {
	userEntities, err := usc.repo.GetUserByRole(ctx, entity.Role(input.Role))
	if err != nil {
		return nil, err
	}
	var users []*dto.UserResponseDTO
	for _, userEntity := range userEntities {
		user := dto.ToUserResponseDTO(userEntity)
		users = append(users, &user)
	}
	return users, nil
}
func (usc *UserUsecase) GetAllUsers(ctx context.Context, input *dto.GetAllUsersDTO) ([]*dto.UserResponseDTO, error) {
	userEntities, err := usc.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	var users []*dto.UserResponseDTO
	for _, userEntity := range userEntities {
		user := dto.ToUserResponseDTO(userEntity)
		users = append(users, &user)
	}
	return users, nil
}
