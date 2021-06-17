package controller

import (
	"strconv"
	"time"

	"github.com/anhht1999/go_admin/database"
	"github.com/anhht1999/go_admin/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil{
		return err
	}

	if data["Password"] != data["Password_confirm"]{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password do not match",
		})
	}

	Password,_ := bcrypt.GenerateFromPassword([]byte(data["Password"]),14)

	user := models.User{
		FirstName: data["FirstName"],
		LastName: data["LastName"],
		Email: data["Email"],
		Password: Password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
  }
//login
func Login(c *fiber.Ctx) error{
	
	var data map[string]string
	if err := c.BodyParser(&data); err != nil{
		return err
	}

	var user models.User

	database.DB.Where("Email = ?", data["Email"]).First(&user)

	if user.Id == 0{
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Not Found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password,[]byte(data["Password"])); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour*24).Unix(),
	})

	token, err:= claims.SignedString([]byte("secret"))

	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
type Claims struct{
	jwt.StandardClaims
}
func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token,err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		}) 
	} 

	claims := token.Claims.(*Claims)

	var user models.User

	database.DB.Where("Id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message":"success",
	})
}