package controllers

import (
	// "fmt"
	database "final-project/config/postgres"
	"final-project/helpers"
	"final-project/models"
	"net/http"

	// "encoding/json"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

// UserRegister godoc
// @Summary Register a new user
// @Description Register a new user and create an associated photo and profile
// @Tags Users
// @Accept json
// @Produce json
// @Param registerReq body models.RegisterReq true "User and Photo Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /users/register [post]
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	registerReq := models.RegisterReq{}

	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Begin Transaction
	tx := db.Begin()

	// Create User
	user := models.User{
		Username:   registerReq.User.Username,
		Email:      registerReq.User.Email,
		FullName:   registerReq.User.FullName,
		Address:    registerReq.User.Address,
		Password:   registerReq.User.Password,
		Age:        uint(registerReq.User.Age),
		StatusUser: registerReq.User.StatusUser,
	}

	// Save user and rollback if error occurs
	if err := tx.Debug().Create(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request user",
			"message": err.Error(),
		})
		return
	}

	// Create Photo associated with the User
	photo := models.Photo{
		Title:    registerReq.Photo.Title,
		Caption:  registerReq.Photo.Caption,
		PhotoUrl: registerReq.Photo.PhotoUrl,
		UserID:   user.Id,
	}

	// Save photo and rollback if error occurs
	if err := tx.Debug().Create(&photo).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request photo",
			"message": err.Error(),
		})
		return
	}

	// Automatically Create Profile
	profile := models.Profile{
		User:       user,
		TotalSwipe: 0, // default value
	}

	// Save profile and rollback if error occurs
	if err := tx.Debug().Create(&profile).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request profile",
			"message": err.Error(),
		})
		return
	}

	// Commit transaction if everything is successful
	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{
		"id":         user.Id,
		"username":   user.Username,
		"email":      user.Email,
		"fullname":   user.FullName,
		"address":    user.Address,
		"age":        user.Age,
		"statusUser": user.StatusUser,
		"photos":     photo.PhotoUrl,
		"profileID":  profile.Id,
		"createdAt":  user.CreatedAt,
	})
}

// UserLogin godoc
// @Summary Log in a user
// @Description Log in a user and return a JWT token
// @Tags Users
// @Accept json
// @Produce json
// @Param login body models.LoginReq true "Login Data"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /users/login [post]
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	loginReq := models.LoginReq{}

	// Binding and validation
	if contentType == appJSON {
		if err := c.ShouldBindJSON(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	// Retrieve the user based on the email
	user := models.User{}
	err := db.Debug().Where("email = ?", loginReq.Email).Take(&user).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	// Compare the provided password with the hashed password in the database
	comparePass := helpers.ComparePass([]byte(user.Password), []byte(loginReq.Password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	// Generate JWT token
	token := helpers.GenerateToken(user.Id, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// UpdateUser godoc
// @Summary Update user information
// @Description Update user information by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User Data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	db := database.GetDB()

	// Check data exist
	var User models.User

	err := db.First(&User, "id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "request not found",
			"message": err.Error(),
		})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	errUpdate := db.Debug().Model(&User).Updates(input).Error

	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": errUpdate.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         User.Id,
		"username":   User.Username,
		"password":   User.Password,
		"email":      User.Email,
		"fullname":   User.FullName,
		"address":    User.Address,
		"age":        User.Age,
		"statusUser": User.StatusUser,
		"createdAt":  User.CreatedAt,
		"updatedAt":  User.UpdatedAt,
	})
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags Users
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	db := database.GetDB()
	var User models.User

	err := db.First(&User, "id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "request not found",
			"message": err.Error(),
		})
		return
	}

	errDelete := db.Debug().Delete(&User).Error

	if errDelete != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
