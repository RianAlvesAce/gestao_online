package controllers

import (
	"net/http"

	"github.com/RianAlvesAce/gestao_online/internal/db"
	"github.com/RianAlvesAce/gestao_online/internal/db/models"
	"github.com/RianAlvesAce/gestao_online/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePacient(c *gin.Context) {
	var pacient models.Pacient

	if err := c.BindJSON(&pacient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "fora do padrão esperado",
		})
		return
	}
	
	if err := utils.IsEmpty(pacient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	
	if pacientExists("cpf", pacient.CPF) {
		c.JSON(400, gin.H{
			"msg": "paciente já existe no banco de dados",
		})
		return
	}

	pacient.Id = primitive.NewObjectID()

	if _, err := db.PacientCollection.InsertOne(db.Ctx, pacient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "ocorreu algo errado na criação do paciente",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg": "paciente criado com sucesso",
	})
}

func UpdatePacient(c *gin.Context) {
	var pacientUpdated models.Pacient

	pacientId := c.Query("id")

	if err := c.BindJSON(&pacientUpdated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "padrão incorreto",
		})
		return
	}

	objId, _ := primitive.ObjectIDFromHex(pacientId)
	pacientUpdated.Id = objId

	var pacientTeste models.Pacient

	if err := db.PacientCollection.FindOne(db.Ctx, bson.M{"_id": pacientUpdated.Id}).Decode(&pacientTeste); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "paciente não encontrado",
		})
		return
	}

	c.JSONP(http.StatusOK, pacientTeste)

}

func pacientExists(key string, value string) bool {
	var pacient models.Pacient
	err := db.PacientCollection.FindOne(db.Ctx, bson.M{key: value}).Decode(&pacient)
	return err == nil
}