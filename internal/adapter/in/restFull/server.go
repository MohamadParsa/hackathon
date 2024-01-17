// Package restfull implements the RESTful server.
// @title     Hackathon API
// @version         1.1
// @description     Hackathon API server.
// @BasePath  beta: http://
package restFull

import (
	"net/http"

	"github.com/MohamadParsa/hackathon/internal/model"
	"github.com/MohamadParsa/hackathon/internal/port"
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
	quickAccess port.QuickAccessApi
	suggestion  port.SuggestionApi
}

// New gets an instance of application layer and returns a new Adapter
func New(quickAccess port.QuickAccessApi, suggestion port.SuggestionApi) *Adapter {
	return &Adapter{
		quickAccess: quickAccess,
		suggestion:  suggestion,
	}
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
	router.GET("/quick-access/", adapter.getAllQuickAccess)
	router.GET("/quick-access/:quickAccessId", adapter.getSpecificQuickAccess) // list quick access
	router.POST("/quick-access/", adapter.addQuickAccess)                      // http status
	router.PATCH("/quick-access/", adapter.modifyQuickAccess)                  // http status
	router.DELETE("/quick-access/:quickAccessId", adapter.deleteQuickAccess)   // http status

	router.GET("/suggestion", adapter.getSuggestionList) // list suggestion

	// router.PUT("/update-file")

}
func (adapter *Adapter) getAllQuickAccess(c *gin.Context) {
	user := getUser(c.Request.Header)
	if user != "" {
		jsonByte, statusCode := adapter.quickAccess.GetQuickAccessList(user)
		writeResponseWithStatusCode(c, jsonByte, statusCode)
	} else {
		writeResponseWithStatusCode(c, nil, http.StatusBadRequest)
	}
}
func (adapter *Adapter) getSpecificQuickAccess(c *gin.Context) {
	user := getUser(c.Request.Header)
	id := c.Param("id")
	if user != "" && id != "" {
		jsonByte, statusCode := adapter.quickAccess.GetSpecificQuickAccess(user, id)
		writeResponseWithStatusCode(c, jsonByte, statusCode)
	} else {
		writeResponseWithStatusCode(c, nil, http.StatusBadRequest)
	}
}
func (adapter *Adapter) addQuickAccess(c *gin.Context) {
	user := getUser(c.Request.Header)
	var quickAccess model.QuickAccess

	err := c.Bind(&quickAccess)
	if user != "" && err == nil {
		quickAccess.UserId = user
		statusCode := adapter.quickAccess.AddQuickAccess(quickAccess)
		writeResponseWithStatusCode(c, nil, statusCode)
	} else {
		writeResponseWithStatusCode(c, nil, http.StatusBadRequest)
	}
}
func (adapter *Adapter) modifyQuickAccess(c *gin.Context) {
	user := getUser(c.Request.Header)
	var quickAccess model.QuickAccess

	err := c.Bind(&quickAccess)
	if user != "" && err == nil {
		quickAccess.UserId = user
		statusCode := adapter.quickAccess.UpdateQuickAccess(quickAccess)
		writeResponseWithStatusCode(c, nil, statusCode)
	} else {
		writeResponseWithStatusCode(c, nil, http.StatusBadRequest)
	}
}
func (adapter *Adapter) deleteQuickAccess(c *gin.Context) {
	user := getUser(c.Request.Header)
	id := c.Param("id")
	if user != "" && id != "" {
		statusCode := adapter.quickAccess.DeleteQuickAccess(user, id)
		writeResponseWithStatusCode(c, nil, statusCode)
	} else {
		writeResponseWithStatusCode(c, nil, http.StatusBadRequest)
	}
}
func (adapter *Adapter) getSuggestionList(c *gin.Context) {
	user := getUser(c.Request.Header)
	if user != "" {
		jsonByte, statusCode := adapter.suggestion.GetSuggestionList(user)
		writeResponseWithStatusCode(c, jsonByte, statusCode)
	} else {
		writeResponseWithStatusCode(c, nil, http.StatusBadRequest)
	}
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
func getUser(header http.Header) string {
	user := header.Get("user")
	if user == "" {
		user = "userX"
	}
	return user
}
