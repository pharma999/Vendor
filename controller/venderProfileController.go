package controller

import(
	"context"
	// "errors"
	// "log"
	"net/http"
	// "os"
	// "strconv"
	// "strings"
	"time"
	"github.com/pharma999/vender/enum"
	"github.com/gin-gonic/gin"
	"github.com/pharma999/vender/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"github.com/pharma999/vender/model"
)
var validate = validator.New()
var venderCollection *mongo.Collection = database.OpenCollection(database.Client, "vender")
//var venderCollection = database.OpenCollection("vender")

// GetVenders godoc
// @Summary Get all venders
// @Description Fetch all vender profiles
// @Tags Vender
// @Accept json
// @Produce json
// @Success 200 {array} model.VenderProfile
// @Failure 500 {object} map[string]string
// @Router /vender [get]
func GetVenders() gin.HandlerFunc{
	return func(c *gin.Context){
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		cursor, err := venderCollection.Find(ctx, bson.D{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies."})
		}
		defer cursor.Close(ctx)

		var venderProfiles []model.VenderProfile

		if err = cursor.All(ctx, &venderProfiles); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies."})
			return
		}

		c.JSON(http.StatusOK, venderProfiles)
	}
}


// GetVender godoc
// @Summary Get vender by ID
// @Description Fetch a vender using vender_id
// @Tags Vender
// @Accept json
// @Produce json
// @Param vender_id path string true "Vender ID"
// @Success 200 {object} model.VenderProfile
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /vender/{vender_id} [get]
func GetVender() gin.HandlerFunc{
	return func(c *gin.Context){
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		venderID := c.Param("vender_id")

		if venderID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Vender ID is required"})
			return
		}

		objecID, err := bson.ObjectIDFromHex(venderID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vender ID"})
			return
		}
		var vender model.VenderProfile

		err = venderCollection.FindOne(ctx, bson.D{{Key: "vender_id", Value: objecID}}).Decode(&vender)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Vender not found"})
			return
		}
		
		c.JSON(http.StatusOK, vender)
	}
}


// CreateVender godoc
// @Summary Create a new vender
// @Description Add a new vender profile
// @Tags Vender
// @Accept json
// @Produce json
// @Param vender body model.VenderProfile true "Vender Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /vender [post]
func CreateVender() gin.HandlerFunc{
	return func(c *gin.Context){
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var vender model.VenderProfile
		if err := c.ShouldBindJSON(&vender); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := validate.Struct(vender); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
			return
		}
		vender.CreatedAt = time.Now()
		vender.UpdatedAt = time.Now()
		vender.Status = enum.Active

		result, err := venderCollection.InsertOne(ctx, vender)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add vender"})
			return
		}

		c.JSON(http.StatusCreated, result)
	}
}

// UpdateVender godoc
// @Summary Update vender profile
// @Description Update vender details by vender_id
// @Tags Vender
// @Accept json
// @Produce json
// @Param vender_id path string true "Vender ID"
// @Param vender body model.VenderProfile true "Updated vender data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /vender/{vender_id} [patch]
func UpdateVender() gin.HandlerFunc{
	return func(c *gin.Context){
        ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		// Get the VenderID from the URL or request body
		venderID := c.Param("vender_id") // Assuming you pass `vender_id` as a URL parameter.

		var updatedVenderProfile model.VenderProfile
		// Bind the request body to the updatedVenderProfile struct
		if err := c.ShouldBindJSON(&updatedVenderProfile); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
			return
		}

		// Validation logic (optional but recommended)
		if updatedVenderProfile.FirstName == "" || updatedVenderProfile.LastName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "First name and last name are required"})
			return
		}

		// Convert venderID from string to bson.ObjectID (if you're using bson.ObjectID for the database)
		objectID, err := bson.ObjectIDFromHex(venderID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vender ID"})
			return
		}

		// Create a filter to find the document
		filter := bson.M{"vender_id": objectID}

		// Prepare the update data
		update := bson.M{
			"$set": bson.M{
				"first_name":    updatedVenderProfile.FirstName,
				"last_name":     updatedVenderProfile.LastName,
				"email":         updatedVenderProfile.Email,
				"phone_number":  updatedVenderProfile.PhoneNumber,
				"updated_at":    time.Now(),
				"token":         updatedVenderProfile.Token,
				"refresh_token": updatedVenderProfile.RefreshToken,
			},
		}

		// Perform the update
		result, err := venderCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vender profile"})
			return
		}
		if result.MatchedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Vender profile not found"})
			return
		}

		// Send a success response
		c.JSON(http.StatusOK, gin.H{"message": "Vender profile updated successfully"})

	}
}