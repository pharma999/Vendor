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

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "product")
//var venderCollection = database.OpenCollection("vender")

// GetProductDetails godoc
// @Summary Get all products
// @Description Fetch all product details
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} model.ProductDetail
// @Failure 500 {object} map[string]string
// @Router /product_detail [get]
func GetProductDetails() gin.HandlerFunc{
	return func(c *gin.Context){
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		cursor, err := productCollection.Find(ctx, bson.D{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products."})
		}
		defer cursor.Close(ctx)

		var venderProfiles []model.ProductDetail

		if err = cursor.All(ctx, &venderProfiles); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode products."})
			return
		}

		c.JSON(http.StatusOK, venderProfiles)
	}
}

// GetProductDetail godoc
// @Summary Get product by vender ID
// @Description Fetch product details using vender_id
// @Tags Product
// @Accept json
// @Produce json
// @Param vender_id path string true "Vender ID"
// @Success 200 {object} model.ProductDetail
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /product_detail/{vender_id} [get]
func GetProductDetail() gin.HandlerFunc{
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
		var vender model.ProductDetail

		err = productCollection.FindOne(ctx, bson.D{{Key: "vender_id", Value: objecID}}).Decode(&vender)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Vender not found"})
			return
		}
		
		c.JSON(http.StatusOK, vender)
	}
}

// CreateProductDetail godoc
// @Summary Create product
// @Description Create a new product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body model.ProductDetail true "Product payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product_detail [post]
func CreateProductDetail() gin.HandlerFunc{
	return func(c *gin.Context){
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var vender model.ProductDetail
		if err := c.ShouldBindJSON(&vender); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
			return
		}
		var validate = validator.New()
		if err := validate.Struct(vender); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
			return
		}
		vender.CreatedAt = time.Now()
		vender.UpdatedAt = time.Now()

		result, err := productCollection.InsertOne(ctx, vender)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add vender"})
			return
		}

		c.JSON(http.StatusCreated, result)
	}
}

// UpdateProductDetail godoc
// @Summary Update product
// @Description Update product by vender_id
// @Tags Product
// @Accept json
// @Produce json
// @Param vender_id path string true "Vender ID"
// @Param product body model.ProductDetail true "Updated product data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product_detail/{vender_id} [patch]
func UpdateProductDetail() gin.HandlerFunc{
	return func(c *gin.Context){
        ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		// Get the VenderID from the URL or request body
		venderID := c.Param("vender_id") // Assuming you pass `vender_id` as a URL parameter.

		var updatedVenderProfile model.ProductDetail
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
		result, err := productCollection.UpdateOne(ctx, filter, update)
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