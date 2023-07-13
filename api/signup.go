package api

import (
	"log"
	"net/http"
	database "shopping/database/implement"
	model "shopping/model"

	"github.com/gin-gonic/gin"
)

// API to sign up as a user
func SignUp(c *gin.Context) {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	var user model.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	_, err = db.Exec("INSERT INTO user (name,email,password) VALUES ($1, $2,$3)", user.Name, user.Email, user.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to signup"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
