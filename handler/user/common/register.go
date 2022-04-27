package common

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/handler/user"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/berror"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Register registers a new user account.
//
// @Summary Register a new user
// @Description Register a new user by username and password
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.RegisterRequest true "user information include username and password"
// @Success 200 {object} user.SwaggerRegisterResponse "{"code":0,"message":"OK","data":{"userId":12,"username":"汤桂英","role":"business"}}"
// @Router /user/register [post]
func Register(c *gin.Context) {
	log.RegisterCalled(c)

	// bind request body
	var r user.RegisterRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.ErrBind(err)
		SendResponse(c, berror.ErrBind, nil)
		return
	}
	u := model.User{
		Username: r.Username,
		Password: r.Password,
		Role:     "general",
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		log.ErrValidate(err)
		SendResponse(c, berror.ErrValidation, nil)
		return
	}

	// Validate if the user exists
	_, deleted, err := model.GetUser(r.Username)
	// if user data exists and deleted is false, send an error
	if deleted == false && err == nil {
		fmt.Println(deleted, err)
		log.ErrUserExists()
		SendResponse(c, berror.ErrUserExists, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		log.ErrEncrypt(err)
		SendResponse(c, berror.ErrEncrypt, nil)
		return
	}

	// Insert the user to the database.
	if err := u.CreateUser(deleted); err != nil {
		log.ErrCreateUser(err)
		SendResponse(c, berror.ErrDatabase, nil)
		return
	}

	rsp := user.RegisterResponse{
		UserId:   u.ID,
		Username: u.Username,
	}

	// Return the user id, username and role.
	SendResponse(c, nil, rsp)
}
