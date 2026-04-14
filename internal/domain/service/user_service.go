package service

import (
	"context"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/rules"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service/dto"
	"github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}
func (srv *UserService) CreateUser(ctx context.Context, user *dto.CreateUserDTO) error {
	err := rules.ValidateCreateUser(user)
	if err != nil {
		return err
	}
	userEntity := dto.ToUserEntity(*user)
	return srv.repo.CreateUser(ctx, *userEntity)
}
func (srv *UserService) UpdateUser(ctx context.Context, updateIt *dto.UpdateUserDTO) error {
	err := rules.ValidateUpdateUser(updateIt)
	if err != nil {
		return err
	}
	return srv.repo.UpdateUser(ctx, *updateIt)
}
func (srv *UserService) DeleteUser(ctx context.Context, deleteIt *dto.DeleteUserDTO) error {
	return srv.repo.DeleteUser(ctx, *deleteIt)
}
func (srv *UserService) GetUserById(ctx context.Context, input *dto.GetUserByIdDTO) (*dto.UserResponseDTO, error) {
	userEntity, err := srv.repo.GetUserById(ctx, input.ID.String())
	if err != nil {
		return nil, err
	}
	user := dto.ToUserResponseDTO(userEntity)
	return &user, nil
}
func (srv *UserService) GetUserByRole(ctx context.Context, input *dto.GetUserByRoleDTO) ([]*dto.UserResponseDTO, error) {
	userEntities, err := srv.repo.GetUserByRole(ctx, entity.Role(input.Role))
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
func (srv *UserService) GetAllUsers(ctx context.Context, input *dto.GetAllUsersDTO) ([]*dto.UserResponseDTO, error) {
	userEntities, err := srv.repo.GetAllUsers(ctx)
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
