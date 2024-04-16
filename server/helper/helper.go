package helper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type content struct {
	email    string
	password string
	username string
}

func Register(c *gin.Context) {
	fmt.Println(c.Request.Body)
	var test1 content
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&test1)
	if err != nil {
		fmt.Printf("error %s", err)
		c.JSON(501, gin.H{"error": err})
	}

	var test2 content
	err = c.BindJSON(&test2)

	fmt.Printf("Bind Email - Why not Set? %v\n", test2.email)
	fmt.Printf("Bind Password  - Why not Set? %v\n", test2.password)
	fmt.Printf("Bind Username  - Why not Set? %v\n", test2.username)

	c.JSON(http.StatusCreated, gin.H{"success": "success"})
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, "This is Login page")
}
