// main.go

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var names []string

func getAllNames(c *gin.Context) {
	c.JSON(http.StatusOK, names)
}

func addName(c *gin.Context) {
	var input struct {
		Name string `json:"name"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entrada invalida"})
		return
	}

	names = append(names, input.Name)
	c.Status(http.StatusNoContent)
}

func main() {
	r := gin.Default()

	r.GET("/getallnames", getAllNames)
	r.POST("/postonenames", addName)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
