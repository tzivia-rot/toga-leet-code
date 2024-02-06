package exerciseRouter

import (
	exerciseService "go-lenguage/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(router *gin.Engine, collection *mongo.Collection) {
	exerciseServiceObj := exerciseService.NewExerciseService(collection)

	router.POST("/exercises/create", exerciseServiceObj.CreateExercise)
	router.GET("/exercises/getById/:id", exerciseServiceObj.GetExerciseById)
	router.GET("/exercises/getAll", exerciseServiceObj.GetAllExercises)
	router.PUT("/exercises/update/:id", exerciseServiceObj.UpdateExercise)
	router.DELETE("/exercises/delete/:id", exerciseServiceObj.DeleteExercise)
	router.POST("/exercises/check/:id", exerciseServiceObj.CheckExercise)

}
