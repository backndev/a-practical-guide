package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/main/db"
	"go-admin/main/models"
	"strconv"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	db.DB.Find(&permissions)

	return c.JSON(permissions)
}

func CreatePermission(c *fiber.Ctx) error {
	var permission models.Permission

	if err := c.BodyParser(&permission); err != nil {
		return err
	}

	db.DB.Create(&permission)

	return c.JSON(permission)
}

func GetPermission(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	permission := models.Permission{
		Id: uint(id),
	}

	db.DB.Find(&permission)

	return c.JSON(permission)
}

func UpdatePermission(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	permission := models.Permission{
		Id: uint(id),
	}

	if err := c.BodyParser(&permission); err != nil {
		return err
	}

	db.DB.Model(&permission).Updates(permission)

	return c.JSON(permission)
}

func DeletePermission(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	permission := models.Permission{
		Id: uint(id),
	}

	db.DB.Delete(&permission)

	return nil
}
