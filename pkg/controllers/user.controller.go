package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nayeem01/Go-Chat/pkg/database"
	"github.com/nayeem01/Go-Chat/pkg/models"
)

type UserResFormat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func UserResFormater(user models.User) UserResFormat {
	return UserResFormat{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Db.Create(&user)
	res := UserResFormater(user)

	return c.Status(200).JSON(res)

}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Db.Find(&users)
	responseUsers := []UserResFormat{}
	for _, user := range users {
		res := UserResFormater(user)
		responseUsers = append(responseUsers, res)
	}

	return c.Status(200).JSON(responseUsers)
}

func GetUser(c *fiber.Ctx) error {

	var user models.User

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("id not found")
	}

	database.Db.Find(&user, "id=?", id)

	res := UserResFormater(user)

	if user.ID == 0 {

		return c.Status(200).JSON("user not found")
	}
	return c.Status(200).JSON(res)
}

func UpdateUser(c *fiber.Ctx) error {

	var user models.User

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("id not found")
	}

	database.Db.Find(&user, "id=?", id)

	res := UserResFormater(user)

	if user.ID == 0 {

		return c.Status(200).JSON("user not found")
	}
	return c.Status(200).JSON(res)
}
