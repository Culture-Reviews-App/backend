package controllers

import (
	"time"

	"github.com/Culture-Reviews-App/backend/app/models"
	"github.com/google/uuid"

	"github.com/Culture-Reviews-App/backend/pkg/utils"
	"github.com/Culture-Reviews-App/backend/platform/database"

	"github.com/gofiber/fiber/v2"
)

// CreateReview method to create reviews for signed in user.
// @Description Create a review related to given token.
// @Summary create review for the user.
// @Tags Review
// @Accept json
// @Produce json
// @Param review body models.ReviewCreate true "Review Data"
// @Success 201 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/review/create [post]
func CreateReview(c *fiber.Ctx) error {
	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a reviewCreate struct.
	reviewCreate := &models.ReviewCreate{}

	// Checking received data from JSON body.
	if err := c.BodyParser(reviewCreate); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Review model.
	validate := utils.NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(reviewCreate); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection("reviews")
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a review struct.
	review := &models.Review{}

	// Set review object.
	review.ID = uuid.New()
	review.CreatedAt = time.Now()
	review.UpdatedAt = review.CreatedAt
	review.UserID = claims.UserID
	review.Title = reviewCreate.Title
	review.Content = reviewCreate.Content
	review.Likes = 0

	// Validate review fields.
	if err := validate.Struct(review); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create a new review with validated data.
	if err := db.CreateReview(review); err != nil {
		// Return status 500 and create process error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 201 OK.
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"review": review,
	})
}

// ListReviews method to list reviews for signed in user.
// @Description List reviews related to given token.
// @Summary list all reviews for the user.
// @Tags Review
// @Accept json
// @Produce json
// @Success 200 {object} models.ReviewList "List of reviews"
// @Security ApiKeyAuth
// @Router /v1/review/list [get]
func ListReviews(c *fiber.Ctx) error {
	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection("reviews")
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	reviews, err := db.ListReviews(claims.UserID)
	if err != nil {
		// Return status 500 and create process error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return reviews
	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"reviews": reviews,
	})
}
