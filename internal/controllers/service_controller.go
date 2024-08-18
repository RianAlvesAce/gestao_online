package controllers

import (
	"context"
	"time"

	"github.com/RianAlvesAce/gestao_online/internal/db"
	"github.com/RianAlvesAce/gestao_online/internal/db/models"
	"github.com/RianAlvesAce/gestao_online/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateService(c *gin.Context) {
	var service models.Service

	if err := c.BindJSON(&service); err != nil {
		c.JSON(400, gin.H{
			"msg": "Fora do padrão esperado",
		})
		return
	}

	service.DateSV = time.Now()

	if err := utils.IsEmpty(service); err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	service.ID = primitive.NewObjectID()

	if _, err := db.ServiceCollection.InsertOne(db.Ctx, service); err != nil {
		c.JSON(400, gin.H{
			"msg": "ocorreu algum erro ao efetuar a atualização",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "dados inseridos com sucesso",
	})
}

func DiagramData(c *gin.Context) {
	var list []models.Service

	cursor, err := db.ServiceCollection.Find(db.Ctx, bson.D{})

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "deu errado",
		})
	}

	if err = cursor.All(context.TODO(), &list); err != nil {
		c.JSON(400, gin.H{
			"msg": "erro",
		})
		return 
	}

	c.JSON(200, gin.H{
		"services": list,
	})
}