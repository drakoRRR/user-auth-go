package handlers

import (
	"github.com/drakoRRR/user-auth-go/internal/models"
	"github.com/drakoRRR/user-auth-go/pkg/utils"
	"go.uber.org/zap"
	"net/http"
)

// CreateUserHandler godoc
// @Summary Create a new user
// @Description Create a new user in the system
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.RegisterUserPayload true "User data"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /users [post]
func (h *Handlers) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	h.Log.Info("Create User endpoint hit")
	var user models.RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		h.Log.Error("Failed to parse JSON", zap.Error(err))
		utils.ErrorJSON(w, err)
		return
	}

	ipAddress := r.Header.Get("X-Forwarded-For")
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}

	createdUser, err := h.Services.Users.CreateUser(r.Context(), user, ipAddress)
	if err != nil {
		h.Log.Error("Failed to create user", zap.String("email", user.Email), zap.Error(err))
		utils.ErrorJSON(w, err)
		return
	}
	h.Log.Info("Created user in DB", zap.String("email", createdUser.Email), zap.String("user_id", createdUser.ID))

	utils.WriteJSON(w, http.StatusCreated, createdUser)
}
