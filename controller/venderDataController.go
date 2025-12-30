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

	"github.com/gin-gonic/gin"
	"github.com/pharma999/vender/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"github.com/pharma999/vender/model"
)

var indvisulVenderCollection *mongo.Collection = database.OpenCollection(database.Client, "invender")
//var venderCollection = database.OpenCollection("vender")

// GetIndvisualVenders godoc
// @Summary Get all individual venders
// @Description Fetch all individual vender profiles
// @Tags IndividualVender
// @Accept json
// @Produce json
// @Success 200 {array} model.IndvisulVenderProfile
// @Failure 500 {object} map[string]string
// @Router /invender [get]
func GetIndvisualVenders() gin.HandlerFunc{
	return func(c *gin.Context){
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		cursor, err := indvisulVenderCollection.Find(ctx, bson.D{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies."})
		}
		defer cursor.Close(ctx)

		var venderProfiles []model.IndvisulVenderProfile

		if err = cursor.All(ctx, &venderProfiles); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies."})
			return
		}

		c.JSON(http.StatusOK, venderProfiles)
	}
}

// GetIndvisualVender godoc
// @Summary Get individual vender by ID
// @Description Fetch individual vender profile using vender_id
// @Tags IndividualVender
// @Accept json
// @Produce json
// @Param vender_id path string true "Vender ID"
// @Success 200 {object} model.IndvisulVenderProfile
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /invender/{vender_id} [get]
func GetIndvisualVender() gin.HandlerFunc{
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
		var vender model.IndvisulVenderProfile

		err = indvisulVenderCollection.FindOne(ctx, bson.D{{Key: "vender_id", Value: objecID}}).Decode(&vender)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Vender not found"})
			return
		}
		
		c.JSON(http.StatusOK, vender)
	}
}

// CreateIndvisualVender godoc
// @Summary Create vender by type
// @Description Create vender (INDIVISUAL / CLINIC / HOSPITAL)
// @Tags IndividualVender
// @Accept json
// @Produce json
// @Param vender_type path string true "Vender Type"
// @Param vender body model.IndvisulVenderProfile true "Individual vender payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /invender/{vender_type} [post]
func CreateIndvisualVender() gin.HandlerFunc{
	return func(c *gin.Context){
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()
		venderType := c.Param("vender_type")

		if venderType == "INDIVISUAL" {
			var vender model.IndvisulVenderProfile
			if err := c.ShouldBindJSON(&vender); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
				return
			}
			var validate = validator.New()
			if err := validate.Struct(vender); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
				return
			}
			vender.CreatedAt = time.Now()
			vender.UpdatedAt = time.Now()
			result, err := indvisulVenderCollection.InsertOne(ctx, vender)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add vender"})
				return
			}
			c.JSON(http.StatusCreated, result)
			return
		}else if venderType == "CLINIC" {
			var vender model.ClinicVenderProfile
			if err := c.ShouldBindJSON(&vender); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
				return
			}
			var validate = validator.New()
			if err := validate.Struct(vender); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
				return
			}
			vender.CreatedAt = time.Now()
			vender.UpdatedAt = time.Now()
			result, err := indvisulVenderCollection.InsertOne(ctx, vender)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add vender"})
				return
			}
			c.JSON(http.StatusCreated, result)
			return			
		}else if venderType == "HOSPITAL" {
			var vender model.HospitalVenderProfile
			if err := c.ShouldBindJSON(&vender); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
				return
			}
			var validate = validator.New()
			if err := validate.Struct(vender); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
				return
			}
			vender.CreatedAt = time.Now()
			vender.UpdatedAt = time.Now()
			result, err := indvisulVenderCollection.InsertOne(ctx, vender)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add vender"})
				return
			}
			c.JSON(http.StatusCreated, result)
			return		
		}else{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vender type"})
			return
		}
		

		
	}
}

// UpdateIndvisualVender godoc
// @Summary Update individual vender
// @Description Update individual vender profile by ID
// @Tags IndividualVender
// @Accept json
// @Produce json
// @Param vender_id path string true "Vender ID"
// @Param vender body model.IndvisulVenderProfile true "Updated vender data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /invender/{vender_id} [patch]
func UpdateIndvisualVender() gin.HandlerFunc{
	return func(c *gin.Context){
        ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		// Get the VenderID from the URL or request body
		venderID := c.Param("vender_id") // Assuming you pass `vender_id` as a URL parameter.

		var updatedVenderProfile model.IndvisulVenderProfile
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
		result, err := indvisulVenderCollection.UpdateOne(ctx, filter, update)
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