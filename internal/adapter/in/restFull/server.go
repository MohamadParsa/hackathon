// Package restfull implements the RESTful server.
// @title     Hackathon API
// @version         1.1
// @description     Hackathon API server.
// @BasePath  beta: http://
package restFull

import (
	"fmt"
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
	router.MaxMultipartMemory = 8 << 20
	router.Use(CORSMiddleware())
	setHealthMethod(router)

	routerV1 := router.Group("/v1")
	adapter.setAPIMethodsV1(routerV1)

	log.Printf("api server listening at %v \n", port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Errorf("error on server listening: %v", err)
	}
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func (adapter *Adapter) setAPIMethodsV1(router *gin.RouterGroup) {
	router.GET("/quick-access/", adapter.getAllQuickAccess)
	router.GET("/quick-access/:quickAccessId", adapter.getSpecificQuickAccess) // list quick access
	router.POST("/quick-access/", adapter.addQuickAccess)                      // http status
	router.PATCH("/quick-access/", adapter.modifyQuickAccess)                  // http status
	router.DELETE("/quick-access/:quickAccessId", adapter.deleteQuickAccess)   // http status
	router.GET("/purchase-history/", adapter.purcahseHistory)                  // http status
	router.GET("/purchase-history/:serviceType", adapter.purcahseHistory)      // http status
	router.GET("/suggestion", adapter.getSuggestionList)                       // list suggestion

	router.POST("/update-file", adapter.uploadFile)

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
	id := c.Param("quickAccessId")
	fmt.Println(id, user)
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
	id := c.Param("quickAccessId")
	if user != "" && id != "" {
		statusCode := adapter.quickAccess.DeleteQuickAccess(user, id)
		writeResponseWithStatusCode(c, nil, statusCode)
	} else {
		writeResponseWithStatusCode(c, nil, http.StatusBadRequest)
	}
}
func (adapter *Adapter) purcahseHistory(c *gin.Context) {
	user := getUser(c.Request.Header)
	serviceType := c.Param("serviceType")
	if user != "" {
		jsonByte, statusCode := adapter.quickAccess.PurcahseHistory(user, serviceType)
		writeResponseWithStatusCode(c, jsonByte, statusCode)
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
func (adapter *Adapter) uploadFile(c *gin.Context) {
	file, err := c.FormFile("uploadFile")
	if err != nil {
		log.Error("1", err)
		writeResponseWithStatusCode(c, nil, http.StatusBadRequest)
		return
	}
	fileContent, err := file.Open()
	if err == nil {
		jsonByte, statusCode := adapter.quickAccess.UploadFile(fileContent, file.Filename)
		writeResponseWithStatusCode(c, jsonByte, statusCode)
	} else {
		log.Error("2", err)

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
