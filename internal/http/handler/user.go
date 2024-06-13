package handler

import (
	"Trello/internal/hash"
	"Trello/internal/model"
	"Trello/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type User struct {
	repo repository.UserRepository
}

func NewUser(repo repository.UserRepository) *User {
	return &User{repo: repo}
}

func (h *User) Register(g *echo.Group) {
	g.POST("/auth/signup", h.SignUp)
	g.POST("/auth/login", h.Login)
	g.GET("/", h.GetUserList)
	g.GET("/:id", h.GetUser)
	g.POST("/", h.CreateUser)
	g.PUT("/:id", h.UpdateUser)
	g.DELETE("/:id", h.DeleteUser)
}

func (h *User) SignUp(c echo.Context) error {
	var newUser model.User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input data"})
	}

	hashedPassword := hash.PasswordHash(newUser.Password)
	newUser.Password = hashedPassword

	if err := h.repo.Create(&newUser); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create user"})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"registered": newUser})
}

func (h *User) Login(c echo.Context) error {
	return nil
}

func (h *User) GetUserList(c echo.Context) error {
	userList, err := h.repo.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"userList": userList})
}

func (h *User) GetUser(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return echo.ErrBadRequest
	}

	user, err := h.repo.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"user": user})
}

func (h *User) CreateUser(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	hashedPassword := hash.PasswordHash(user.Password)
	user.Password = hashedPassword

	if err := h.repo.Create(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{"created": user})
}

func (h *User) UpdateUser(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return echo.ErrBadRequest
	}

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	user.ID = uint(id)

	hashedPassword := hash.PasswordHash(user.Password)
	user.Password = hashedPassword

	if err := h.repo.Update(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{"updated": user})
}

func (h *User) DeleteUser(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return echo.ErrBadRequest
	}
	if err := h.repo.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "delete user " + c.Param("id")})
}
