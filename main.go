package main

import (
    "github.com/gin-gonic/gin"
    "gopkg.in/go-playground/validator.v9"
    "net/http"
)

// User represents the structure of the request payload
type User struct {
    Email string `json:"email" binding:"required,email"`
}

// Define a global validator
var validate *validator.Validate

func main() {
    r := gin.Default()

    // Initialize the validator
    validate = validator.New()

    r.POST("/validateEmail", func(c *gin.Context) {
        var user User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := validate.Struct(user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Email is valid"})
    })

    r.Run(":8080")
}



// Binding JSON to a Struct:

// When you call c.ShouldBindJSON(&user), Gin attempts to read the JSON body of the request and populate the fields of the user struct with the corresponding values from the JSON.


// gin.H
// gin.H is a shortcut provided by Gin for creating map[string]interface{}. This is commonly used to build JSON responses in a concise and readable way.
// gin.H{"message": "Email is valid"} creates a JSON object with a single key-value pair: {"message": "Email is valid"}.