package handlers

import (
	"encoding/json"
	"fmt"
	"http-service-example/internal/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type AuthenticateInput struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type AuthenticateOutput struct {
	Username string `json:"username"`
	Age int `json:"age"`
}

type Authenticate struct {
	Adapter *database.DataAdapter
}

func (authenticate *Authenticate) Handler() func(*gin.Context) {
	return func(c *gin.Context) {
		var input AuthenticateInput 
		// .Bind will bind fields with 'form' tag
		err := c.Bind(&input)
		if err != nil {
			fmt.Println("[ERROR]", err)
			c.JSON(400, gin.H{
				"status": "bad_request",
				"identity": AuthenticateOutput{},
			})
			return 
		}
		defer func() {
			byteString, _ := json.Marshal(input)
			fmt.Println("[INPUT] /auth:", string(byteString))
		}()

		identity, err := authenticate.Adapter.GetUserByUsernameAndPassword(input.Username, input.Password)
		if err != nil {
			fmt.Println("[ERROR]", err)
			c.JSON(401, gin.H{
				"status": "authenticate_failed",
				"identity": AuthenticateOutput{},
			})
			return 
		}

		c.JSON(200, gin.H{
			"status": "authenticate_success",
			"identity": AuthenticateOutput{
				Username: identity.Username,
				Age: identity.Age,
			},
		})
		return 
	}
}


