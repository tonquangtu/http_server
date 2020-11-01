package repository

import (
	"github.com/tonquangtu/data_server/pkg/rpc/dto"
	"github.com/tonquangtu/http_server/internal/model"
	"net/rpc"
)

type UserRepository struct {
	client *rpc.Client
}

func NewUserRepository(client *rpc.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (userRepo *UserRepository) GetUser(id int) (*model.User, error) {
	var getUserResult dto.GetUserDto
	err := userRepo.client.Call("UserHandler.GetUser", id, &getUserResult)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:           getUserResult.ID,
		Username:     getUserResult.Username,
		PasswordHash: getUserResult.PasswordHash,
		CreatedAt:    getUserResult.CreatedAt,
		UpdatedAt:    getUserResult.UpdatedAt,
	}, nil
}
