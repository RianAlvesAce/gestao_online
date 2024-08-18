package controllers

import (
	"errors"

	"github.com/RianAlvesAce/gestao_online/internal/db"
	"github.com/RianAlvesAce/gestao_online/internal/db/models"
	"github.com/RianAlvesAce/gestao_online/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func SearchUser(login string) (models.User, error) {
	var result models.User

	err := db.UserCollection.FindOne(db.Ctx, bson.M{"login": login}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, mongo.ErrNoDocuments
		}
		return models.User{}, errors.New("erro ao buscar usuário")
	}

	return result, nil
}


func UserLogin(c *gin.Context) {

	uLogin := c.Query("login")
	uPass := c.Query("pass")

	user, err := SearchUser(uLogin)

	if err != nil {
		c.JSON(400, gin.H{
			"msg": err,
		})
		return
	}

	passError := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(uPass))

	if passError != nil {
		c.JSON(401, gin.H{
			"msg": "Senha incorreta",
		})
		return
	}

	token, tokenErr := jwt.GenerateToken(user.ID.String())

	if tokenErr != nil {
		c.JSON(500, gin.H{
			"msg": "Problema ao gerar o token",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":   "Login bem sucedido",
		"token": token,
	})
}

// func VerifyToken(c *gin.Context) {
// 	uToken := c.Query("token")

// 	if !jwt.ValidateToken(uToken) {
// 		c.JSON(400, gin.H{
// 			"msg": "Token inválido",
// 		})
// 	} else {
// 		c.JSON(200, gin.H{
// 			"msg": "Token válido",
// 		})
// 	}
// }
