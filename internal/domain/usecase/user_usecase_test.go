package usecase_test

import (
	"context"
	"testing"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/entity"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
	"github.com/google/uuid"
)

func buildUserTest(t *testing.T) (*usecase.UserUsecase, uuid.UUID, func()) {
	t.Helper()

	db := usecase.OpenTestDB(t)
	userUsc, userID, _ := usecase.CreateTestUser(t, db, "Test User", "user-test-", []entity.Role{entity.RoleBuyer})

	cleanup := func() {
		db.Close()
	}

	return userUsc, userID, cleanup
}

func TestUserUsecase(t *testing.T) {
	usc, userID, cleanup := buildUserTest(t)
	defer cleanup()

	t.Run("GetUserById", func(t *testing.T) {
		input := &dto.GetUserByIdDTO{ID: userID}
		user, err := usc.GetUserById(context.Background(), input)
		if err != nil {
			t.Fatalf("GetUserById falhou: %v", err)
		}
		if user == nil || user.ID != userID {
			t.Fatalf("esperado id de usuario %v, obtido %v", userID, user)
		}
	})

	t.Run("GetUserByRole", func(t *testing.T) {
		input := &dto.GetUserByRoleDTO{Role: string(entity.RoleBuyer)}
		users, err := usc.GetUserByRole(context.Background(), input)
		if err != nil {
			t.Fatalf("GetUserByRole falhou: %v", err)
		}
		if len(users) == 0 {
			t.Fatal("pelo menos um usuario com a role de comprador deveria existir")
		}
	})

	t.Run("GetAllUsers", func(t *testing.T) {
		input := &dto.GetAllUsersDTO{ID: uuid.New()}
		users, err := usc.GetAllUsers(context.Background(), input)
		if err != nil {
			t.Fatalf("GetAllUsers falhou: %v", err)
		}
		if len(users) == 0 {
			t.Fatal("pelo menos um usuario deveria ser retornado")
		}
	})

	t.Run("UpdateUser", func(t *testing.T) {
		newName := "Updated User Name"
		roles := []entity.Role{entity.RoleSeller}
		updateDTO := &dto.UpdateUserDTO{
			ID:    userID.String(),
			Name:  &newName,
			Roles: &roles,
		}
		if err := usc.UpdateUser(context.Background(), updateDTO); err != nil {
			t.Fatalf("UpdateUser falhou: %v", err)
		}

		updated, err := usc.GetUserById(context.Background(), &dto.GetUserByIdDTO{ID: userID})
		if err != nil {
			t.Fatalf("GetUserById apos atualizacao falhou: %v", err)
		}
		if updated.Name != newName {
			t.Fatalf("nome atualizado esperado %q, obtido %q", newName, updated.Name)
		}
	})

	t.Run("DeleteUser", func(t *testing.T) {
		if err := usc.DeleteUser(context.Background(), &dto.DeleteUserDTO{ID: userID.String()}); err != nil {
			t.Fatalf("DeleteUser falhou: %v", err)
		}

		_, err := usc.GetUserById(context.Background(), &dto.GetUserByIdDTO{ID: userID})
		if err == nil {
			t.Fatal("o get user nao devia funcionar apos deletar o usuario")
		}
	})
}
