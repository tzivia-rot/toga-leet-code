package exerciseService

import (
	"context"
	model "go-lenguage/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExerciseService struct {
	Collection *mongo.Collection
}

func NewExerciseService(collection *mongo.Collection) *ExerciseService {
	return &ExerciseService{Collection: collection}
}

func (s *ExerciseService) CreateExercise(c *gin.Context) {
	var exercise model.Exercise
	if err := c.ShouldBindJSON(&exercise); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := s.Collection.InsertOne(context.TODO(), exercise)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating exercise"})
		return
	}

	c.JSON(http.StatusCreated, result.InsertedID)
}

func (s *ExerciseService) GetExerciseById(c *gin.Context) {
	id := c.Param("id")
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

	c.JSON(http.StatusOK, exercise)
}
func (s *ExerciseService) GetAllExercises(c *gin.Context) {
	var exercises []model.Exercise

	cursor, err := s.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching exercises"})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var exercise model.Exercise
		cursor.Decode(&exercise)
		exercises = append(exercises, exercise)
	}

	c.JSON(http.StatusOK, exercises)
}
func (s *ExerciseService) UpdateExercise(c *gin.Context) {
	id := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var exercise model.Exercise
	if err := c.ShouldBindJSON(&exercise); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"name":        exercise.Name,
			"description": exercise.Description,
		},
	}

	_, err = s.Collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating exercise"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exercise updated successfully"})
}

func (s *ExerciseService) DeleteExercise(c *gin.Context) {
	id := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	_, err = s.Collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting exercise"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exercise deleted successfully"})
}
