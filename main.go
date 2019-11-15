package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Cliente struct {
	ID               string `json:"id"`
	Rut              string `json:"rut"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	SecondLastName   string `json:"secondLastName"`
	Pep              bool   `json:"pep"`
	Gender           string `json:"gender"`
	DateOfBirth      string `json:"dateOfBirth"`
	Nationality      string `json:"nationality"`
	Phone            int    `json:"phone"`
	ResidenceCountry string `json:"residenceCountry"`
	Address          string `json:"address"`
	City             string `json:"city"`
	Commune          string `json:"commune"`
	PostalCode       int    `json:"postalCode"`
	MaritalStatus    string `json:"maritalStatus"`
	Occupation       string `json:"occupation"`
	Degree           string `json:"degree"`
}

type ErrMsg struct {
	Message string `json:"message"`
}

var ClientList = []Cliente{}

func main() {
	router := gin.Default()
	router.GET("/clients", GetClients)
	router.GET("/clients/:id", GetClientByID)
	router.POST("/clients", PostClient)
	router.Run(Port())

}

func GetClientByID(c *gin.Context) {
	clientID := c.Param("id")
	for _, cli := range ClientList {
		if cli.ID == clientID {
			c.JSON(http.StatusOK, cli)
			return
		}
	}
	msg := ErrMsg{
		Message: fmt.Sprintf("client not found with id: %v", clientID),
	}
	c.JSON(http.StatusNotFound, msg)
}

func GetClients(c *gin.Context) {
	c.JSON(http.StatusOK, ClientList)
}

func PostClient(c *gin.Context) {
	cli := Cliente{}
	err := c.BindJSON(&cli)
	if err != nil {
		msg := ErrMsg{
			Message: "Invalid JSON object",
		}
		c.JSON(http.StatusBadRequest, msg)
	} else {
		id := uuid.New()
		cli.ID = id.String()
		ClientList = append(ClientList, cli)
		c.JSON(http.StatusCreated, cli)
	}
}

func Port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"

	}
	return ":" + port
}
