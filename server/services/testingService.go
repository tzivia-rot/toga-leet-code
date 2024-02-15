package exerciseService

import (
	"context"
	"fmt"
	model "go-lenguage/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ExerciseService) CheckExercise(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("idd", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var exercise model.Exercise
	err = s.Collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&exercise)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exercise not found" + err.Error()})
		return
	}

	var answer model.Answer
	if err := c.ShouldBindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Print("\nfuncion\n", answer.Function)
	newUUID := uuid.New()
	imageName := fmt.Sprintf("tziviarot/%s:latest", newUUID)
	if answer.Lenguage == "node.js" {
		answerWhitBasicCode := addStringAfterWord(exercise.BasisOperationNodeJS, answer.Function)
		fmt.Print("\nanswerWhitBasicCode\n", answerWhitBasicCode)

		checkExerciseNode, err := CheckExerciseNode(answerWhitBasicCode, exercise.Examples, newUUID.String(), imageName)
		c.JSON(http.StatusOK, gin.H{"response": checkExerciseNode})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"response": err})
		}

	}

	if answer.Lenguage == "GO" {
		answerWhitBasicCode := addStringAfterWord(exercise.BasisOperationGO, answer.Function)
		fmt.Print("\nanswerWhitBasicCode\n", answerWhitBasicCode)
		checkExerciseGO, err := CheckExerciseGO(answerWhitBasicCode, exercise.Examples, newUUID.String(), imageName)
		c.JSON(http.StatusOK, gin.H{"response": checkExerciseGO})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"response": err})
		}
	}

}
func addStringAfterWord(firstStr string, secondStr string) string {
	closingBracketIndex := strings.Index(firstStr, "{}")
	if closingBracketIndex != -1 {
		return firstStr[:closingBracketIndex+1] + secondStr + firstStr[closingBracketIndex+1:]
	}
	return firstStr
}
