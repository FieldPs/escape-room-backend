package routes

import (
	"net/http"

	"github.com/FieldPs/escape-room-backend/internal/stats"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterPuzzleRoutes sets up puzzle and stats endpoints
func RegisterPuzzleRoutes(r gin.IRouter, db *gorm.DB) {
	// Protected routes under /api
	authGroup := r.Group("/", AuthMiddleware())
	{
		authGroup.GET("/stats", statsHandler(db))
	}
}

// statsHandler retrieves user statistics
func statsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("userID").(int)
		stats, err := stats.GetUserStats(db, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stats"})
			return
		}
		c.JSON(http.StatusOK, stats)
	}
}
