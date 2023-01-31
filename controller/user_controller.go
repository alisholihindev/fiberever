package controller

import (
	"fiberever/biz/user"
	"fiberever/model"
	"strconv"

	lib "github.com/alisholihindev/go-lib"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	if uid != n {
		return false
	}

	return true
}

func validUser(id string, p string) bool {
	var user model.User
	lib.DBConn.First(&user, id)
	if user.Username == "" {
		return false
	}
	if !CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}

// GetUser godoc
// @tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} lib.OutputFormat{Data=model.User}
// @Security BearerAuth
// @Router /user/{id} [get]
func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user, err := user.Get(id)
	if err != nil {
		lib.CommonLogger().Error(err)
		return c.Status(500).JSON(lib.ResponseFormat(false, err.Error(), nil, nil))
	}
	return c.JSON(lib.ResponseFormat(true, "ok", user, nil))
}

// AddUser godoc
// @tags User
// @Accept  mpfd
// @Produce  json
// @Param email formData string false  "Email" Format(email)
// @Param username formData string false  "Username"
// @Param password formData string false  "Password" Format(password)
// @Success 200 {object} lib.OutputFormat{Data=model.User}
// @Security BearerAuth
// @Router /user/ [post]
func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})

	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})

	}

	user.Password = hash
	if err := lib.DBConn.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	newUser := NewUser{
		Email:    user.Email,
		Username: user.Username,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

// EditUser godoc
// @tags User
// @accept  json
// @produce  json
// @Param user_id path int true "User ID"
// @Param name formData string false "Name"
// @success 200 {object} lib.OutputFormat{data=[]model.User}
// @security BearerAuth
// @Router /user/{user_id} [put]
func UpdateUser(c *fiber.Ctx) error {
	type UpdateUserInput struct {
		Names string `json:"names"`
	}
	var uui UpdateUserInput
	if err := c.BodyParser(&uui); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}

	db := lib.DBConn
	var user model.User

	db.First(&user, id)
	user.Names = uui.Names
	db.Save(&user)

	return c.JSON(fiber.Map{"status": "success", "message": "User successfully updated", "data": user})
}

// Delete User godoc
// @tags User
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Param password formData string false  "Password" Format(password)
// @Success 200 {object} lib.OutputFormat{Data=model.User}
// @Security BearerAuth
// @Router /user/{user_id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	type PasswordInput struct {
		Password string `json:"password"`
	}
	var pi PasswordInput
	if err := c.BodyParser(&pi); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})

	}

	if !validUser(id, pi.Password) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Not valid user", "data": nil})

	}

	db := lib.DBConn
	var user model.User

	db.First(&user, id)

	db.Delete(&user)
	return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
}
