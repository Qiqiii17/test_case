package middleware

import (
	"net/http"
	"time"

	"test_case/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RateLimitMiddleware adalah middleware untuk membatasi jumlah login yang gagal
func RateLimitMiddleware(db *gorm.DB) gin.HandlerFunc {
	const maxFailedAttempts = 3
	const lockoutDuration = 5 * time.Minute // Durasi kunci akun setelah mencapai batas gagal login

	return func(c *gin.Context) {
		email := c.PostForm("email")

		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
			c.Abort()
			return
		}

		// Mengecek apakah akun terkunci
		var lastLog models.Log
		err := db.Where("email = ? AND expiredlock > ?", email, time.Now()).Order("idlog desc").First(&lastLog).Error
		if err == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Too many failed login attempts. Please try again later."})
			c.Abort()
			return
		}

		// Hitung jumlah percobaan login gagal
		var failedAttempts int64
		err = db.Model(&models.Log{}).Where("email = ? AND islogin = 0 AND created_at > ?", email, time.Now().Add(-lockoutDuration)).Count(&failedAttempts).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}

		if failedAttempts >= maxFailedAttempts {
			// Set expiredlock jika gagal login melebihi batas
			expiredLock := time.Now().Add(lockoutDuration)
			logEntry := models.Log{
				Email:       email,
				IsLogin:     0,
				ExpiredLock: expiredLock,
			}
			if err := db.Create(&logEntry).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Too many failed login attempts. Please try again later."})
			c.Abort()
			return
		}

		c.Next()
	}
}
