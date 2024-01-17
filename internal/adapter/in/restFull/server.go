// Package restfull implements the RESTful server.
// @title     Hackathon API
// @version         1.1
// @description     Hackathon API server.
// @BasePath  beta: http://
package restFull

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Failure struct {
	Result string `json:"result"`
}
type GeneralResponse struct {
	Result string `json:"result"`
}

// Adapter struct uses to connecting to application layer and call application methods.
type Adapter struct {
}

// New gets an instance of application layer and returns a new Adapter
func New() *Adapter {
	return &Adapter{}
}

func (adapter *Adapter) Serve(port string) {
	router := gin.New()

	setHealthMethod(router)

	routerV1 := router.Group("/v1")
	adapter.setAPIMethodsV1(routerV1)

	log.Printf("api server listening at %v \n", port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Errorf("error on server listening: %v", err)
	}
}

func (adapter *Adapter) setAPIMethodsV1(router *gin.RouterGroup) {

}

func writeResponseWithStatusCode(c *gin.Context, jsonByteResult []byte, statusCode int) {

	if statusCode != 200 {
		c.JSON(statusCode, gin.H{"result": http.StatusText(statusCode)})
		return
	}
	c.Data(statusCode, "application/json", jsonByteResult)
}

func setHealthMethod(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
}
