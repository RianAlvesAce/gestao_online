package controllers

import (
	"context"
	"net/http"

	"github.com/RianAlvesAce/gestao_online/internal/db"
	"github.com/RianAlvesAce/gestao_online/internal/db/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateSchedule(c *gin.Context) {
	var schedule models.Schedule

	if err := c.BindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "fora do padrão esperado",
		})
		return
	}

	schedule.Id = primitive.NewObjectID()
	
	if _, err := db.ScheduleCollection.InsertOne(db.Ctx, schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "algo de errado ao criar  o agendamento",
		})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg": "agendamento criado com sucesso",
	})
}

func GetSchedule(c *gin.Context) {
	var schedule_list []models.Schedule

	cursor, err := db.ScheduleCollection.Find(db.Ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "houve algum problema ao adquirir os dados",
		})
		return 
	}

	if err := cursor.All(context.TODO(), &schedule_list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "problema organizar os dados",
		})
		return 
	}

	c.JSON(http.StatusOK, schedule_list)
}

func DeleteSchedule(c *gin.Context) {
	schedule_id := c.Query("id")
	objId, _ := primitive.ObjectIDFromHex(schedule_id)

	if err := db.ScheduleCollection.FindOneAndDelete(db.Ctx, bson.M{"_id": objId}); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "não foi encontrado um agendamento para deletar",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "agendamento deletado com sucesso",
	})
}