package handler

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/tst/backend/database"
	"github.com/tst/backend/model/request"
	"github.com/tst/backend/model/response"
	"github.com/tst/backend/utils"
)

func LoginHandler(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	// Validasi Request
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// Check Available Username
	var user response.UserResponse
	err := database.DB.Debug().Table("users").Find(&user, "username = ?", loginRequest.Username).Error
	if err != nil || user.UserID == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	// // Check Available Password
	// isValid := utils.CheckPasswordHas(loginRequest.Password, user.Password)
	// if !isValid {
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"message": "wrong credential",
	// 	})
	// }

	//Generate JWT

	claims := jwt.MapClaims{}
	claims["userID"] = user.UserID
	claims["username"] = user.Username
	claims["fullname"] = user.Fullname
	claims["userType"] = user.UserType
	claims["parentUserID"] = user.ParentUserID
	claims["exp"] = time.Now().Add(time.Minute * 60 * 60).Unix()

	if user.Username != "mon" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
		"userName": user.Username,
		"userID": user.UserID,
	})

}
