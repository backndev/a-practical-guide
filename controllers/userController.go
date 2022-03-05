package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/main/db"
	"go-admin/main/middlewares"
	"go-admin/main/models"
	"strconv"
)

func AllUsers(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	take, _ := strconv.Atoi(c.Query("take", "10"))

	return c.JSON(models.Paginate(db.DB, &models.User{}, page, take))
}

func CreateUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	db := db.DB
	user := new(models.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(models.Result{
			Status:  "error",
			Code:    500,
			Message: "Review your input",
			Data:    err,
		})
	}

	hash, err := user.SetPassword(string(user.Password))
	if err != nil {
		return c.Status(500).JSON(models.Result{
			Status:  "error",
			Code:    500,
			Message: "Couldn't hash password",
			Data:    err,
		})

	}
	if user.LastName == "" || user.FirstName == "" {
		return c.Status(500).JSON(models.Result{
			Status:  "error",
			Code:    500,
			Message: "input your user",
			Data:    err,
		})
	}

	user.Password = []byte(hash)
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(models.Result{
			Status:  "error",
			Code:    500,
			Message: "Couldn't create user",
			Data:    err,
		})
	}

	newUser := models.User{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		RoleId:    user.RoleId,
		Role:      user.Role,
	}

	return c.Status(200).JSON(models.Result{
		Status:  "success",
		Code:    200,
		Message: "Created user",
		Data:    newUser,
	})
}

func GetUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	db.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	db.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	db.DB.Delete(&user)

	return nil
}
