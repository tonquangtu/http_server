package handler

import (
	"errors"
	"fmt"
	"github.com/tonquangtu/http_server/internal/repository"
	"github.com/tonquangtu/http_server/internal/util"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userRepository *repository.UserRepository
}

func NewUserHandler(userRepository *repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepository: userRepository,
	}
}

func (handler *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		util.RenderJson(w, util.BadRequest(errors.New("id field is not allow empty")))
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.RenderJson(w, util.BadRequest(fmt.Errorf("id field is invalid %s", err.Error())))
		return
	}

	user, err := handler.userRepository.GetUser(id)
	if err != nil {
		util.RenderJson(w, util.UnprocessableEntity(err))
		return
	}

	util.RenderJson(w, util.Ok(user))
}
