package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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

type RespMsg struct {
	Message string `json:"message"`
}

var ClientList = []Cliente{}
var SecretKey = "superawesomesecketkey"

func main() {
	router := gin.Default()
	auth := router.Group("")
	auth.Use(CheckSignature())
	{
		auth.GET("/clients", GetClients)
		auth.GET("/clients/:id", GetClientByID)
		auth.POST("/clients", PostClient)
	}
	router.POST("/signature", SignPayload)
	router.Run(Port())
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func CheckSignature() gin.HandlerFunc {
	return func(c *gin.Context) {
		hsign := c.Request.Header.Get("X-CB-Signature")
		if hsign == "" {
			respondWithError(c, http.StatusUnauthorized, "X-CB-Signature required")
			return
		}

		// Obtiene el payload del request
		buf := make([]byte, 1024)
		num, _ := c.Request.Body.Read(buf)
		payload := string(buf[0:num])

		// Verifica si el payload es igual a la firma
		h := hmac.New(sha256.New, []byte(SecretKey))
		h.Write([]byte(payload))
		calculated := h.Sum(nil)
		decoded, _ := hex.DecodeString(hsign)
		if hmac.Equal(calculated, decoded) == false {
			respondWithError(c, http.StatusUnauthorized, "Invalid X-CB-Signature")
			return
		}
		c.Next()
	}
}

func GetClientByID(c *gin.Context) {
	clientID := c.Param("id")
	for _, cli := range ClientList {
		if cli.ID == clientID {
			c.JSON(http.StatusOK, cli)
			return
		}
	}
	msg := RespMsg{
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
		fmt.Println(err)
		msg := RespMsg{
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

func SignPayload(c *gin.Context) {
	// Obtine el payload del request
	// Obtiene el payload del request
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	payload := string(buf[0:num])

	// Crea un hmac pasandole la llave privada
	h := hmac.New(sha256.New, []byte(SecretKey))

	// le agrega el json al hmac
	h.Write([]byte(payload))

	// obtiene el resultado y lo encodea como hexadecimal
	sha := hex.EncodeToString(h.Sum(nil))
	msg := RespMsg{
		Message: fmt.Sprintf("Payload signature with hmac is: %v", sha),
	}
	c.JSON(http.StatusCreated, msg)
}

func Port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"

	}
	return ":" + port
}
