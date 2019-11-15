package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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
var Signature = ""
var SecretKey = "superawesomesecketkey"

func main() {
	router := gin.Default()
	router.GET("/clients", GetClients)
	router.GET("/clients/:id", GetClientByID)
	// Solo en este metodo se verifica la cabecera.
	router.POST("/clients", CheckHeaders(), PostClient)
	router.POST("/signature", SignPayload)
	router.Run(Port())
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

// Valida que venga la cabecera
func CheckHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		Signature = c.Request.Header.Get("X-CB-Signature")
		if Signature == "" {
			respondWithError(c, http.StatusUnauthorized, "X-CB-Signature required")
			return
		}
		c.Next()
	}
}

// Con el payload, la firma y el secrete verifica que sea un request conocido.
func VerifySignature(data string) bool {
	h := hmac.New(sha256.New, []byte(SecretKey))
	h.Write([]byte(data))
	calculated := h.Sum(nil)
	decoded, _ := hex.DecodeString(Signature)
	if hmac.Equal(calculated, decoded) == false {
		return false
	}
	return true
}

// Obtiene cliente por ID
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

// Obtiene todos los clientes
func GetClients(c *gin.Context) {
	c.JSON(http.StatusOK, ClientList)
}

// Crea un cliente
func PostClient(c *gin.Context) {
	cli := Cliente{}

	// Valida el objeto json
	err := c.BindJSON(&cli)
	if err != nil {
		msg := RespMsg{
			Message: "Invalid JSON object",
		}
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	// Valida que la firma corresponda al payload que viene en el request.
	payload, _ := json.Marshal(cli)
	if VerifySignature(string(payload)) == false {
		msg := RespMsg{
			Message: "Invalid X-CB-Signature",
		}
		c.JSON(http.StatusUnauthorized, msg)
		return
	}

	// Valida que no se cree un cliente con el mismo rut.
	for _, client := range ClientList {
		if client.Rut == cli.Rut {
			msg := RespMsg{
				Message: fmt.Sprintf("A client witch rut %v already exists", cli.Rut),
			}
			c.JSON(http.StatusConflict, msg)
			return
		}
	}

	// Crea el cliente
	id := uuid.New()
	cli.ID = id.String()
	ClientList = append(ClientList, cli)
	c.JSON(http.StatusCreated, cli)
}

// Genera un firma con el payload
func SignPayload(c *gin.Context) {
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

// Levanta el proyecto segun puerto
func Port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"

	}
	return ":" + port
}
