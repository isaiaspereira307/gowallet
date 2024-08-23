package handlers

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/internal/db"
	"golang.org/x/crypto/bcrypt"
)

var (
	lastID int32
	mu     sync.Mutex
)

func generateUniqueID() int32 {
	mu.Lock()
	defer mu.Unlock()
	lastID++
	return lastID
}

// @BasePath /api/v1
// @Summary Create an user
// @Description Create an user
// @Tags user
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "Create User Params"
// @Success 200 {object} CreateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [post]
func CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		sendErr(ctx, err, http.StatusBadRequest)
		return
	}
	err := req.Validate()
	if err != nil {
		sendErr(ctx, err, http.StatusBadRequest)
		return
	}
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
		return
	}

	user := db.CreateUserParams{
		ID:        generateUniqueID(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = queries.CreateUser(ctx, user)
	if err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
		return
	}

	sendSuccess(ctx, "createUser", req)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
