package controllers

import (
	"final-project/config/postgres"
	"final-project/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// SwipeProfile godoc
// @Summary Swipe a profile
// @Description Perform a swipe action on another profile and check for a match
// @Tags Profiles
// @Accept json
// @Produce json
// @Param profileID path int true "ID of the profile that is performing the swipe"
// @Param otherProfileID path int true "ID of the profile being swiped"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /swipe/{profileID}/{otherProfileID} [post]
func SwipeProfile(c *gin.Context) {
	db := postgres.GetDB()
	profileID := c.Param("profileID")           // Mengambil ProfileID dari parameter URL
	otherProfileID := c.Param("otherProfileID") // Mengambil ID profile lain yang di-swipe dari parameter URL

	// Ambil informasi pengguna yang sedang login dari konteks
	userData := c.MustGet("userData").(jwt.MapClaims)
	loggedInUserID := uint(userData["id"].(float64))

	// Verifikasi bahwa profileID adalah milik pengguna yang sedang login
	var profile models.Profile
	err := db.Where("id = ? AND user_id = ?", profileID, loggedInUserID).First(&profile).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Profile not found or not authorized",
			"message": err.Error(),
		})
		return
	}

	// Temukan Profile lain yang di-swipe berdasarkan OtherProfileID
	var otherProfile models.Profile
	err = db.First(&otherProfile, "id = ?", otherProfileID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Other profile not found",
			"message": err.Error(),
		})
		return
	}

	// Logic untuk Profile Free (dibatasi 10 swipe)
	if profile.User.StatusUser == "free" && profile.TotalSwipe >= 10 {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Swipe limit reached",
			"message": "Free users are limited to 10 swipes.",
		})
		return
	}

	// Lakukan operasi swipe (misalnya, menambah TotalSwipe)
	profile.TotalSwipe++
	err = db.Save(&profile).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update swipe count",
			"message": err.Error(),
		})
		return
	}

	// Cek apakah otherProfile sudah swipe profile ini sebelumnya
	var match models.MatchProfile
	err = db.Where("profile_id = ? AND other_profile_id = ?", otherProfileID, profileID).First(&match).Error
	if err == nil {
		// Jika otherProfile sudah swipe profile ini, maka match terjadi
		match.StatusMatch = 1
		db.Save(&match)
		c.JSON(http.StatusOK, gin.H{
			"message":      "It's a match!",
			"total_swipes": profile.TotalSwipe,
			"match":        true,
		})
		return
	}

	// Jika belum ada match, buat entry baru di MatchProfile dengan status 0
	newMatch := models.MatchProfile{
		ProfileID:      profile.Id,
		OtherProfileID: otherProfile.Id,
		StatusMatch:    0,
	}
	err = db.Create(&newMatch).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create match profile",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Swipe successful",
		"total_swipes": profile.TotalSwipe,
		"match":        false,
	})
}
