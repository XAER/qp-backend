package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"qp/back/config"
	"qp/back/lib"
	"qp/back/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetStopRealTimeDataHandler(c *gin.Context) {
	// Getting the language name and stop code from the url
	languageName := c.Param("languageName")
	if languageName == "" {
		languageName = "en"
	}
	stopCode := c.Param("stopCode")
	if stopCode == "" || len(stopCode) > 5 || len(stopCode) < 5 {
		lib.RespondWithError(c, http.StatusBadRequest, "Invalid stop code")
		return
	}

	resp, err := http.Post(config.App.C.DATA_SERVER+"/GetRealTimeData", "application/json", strings.NewReader(`{"languageName":"`+languageName+`", "stopCode":"`+stopCode+`"}`))
	if err != nil {
		lib.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		lib.RespondWithError(c, resp.StatusCode, "Error while contacting data server")
		return
	}
	text, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		lib.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	parsedRes := models.GetRealTimeDataResponse{}
	json.Unmarshal([]byte(string(text)), &parsedRes)

	c.JSON(http.StatusOK, parsedRes)

	// Logging request
	loggingError := lib.LogRequest(200, "Requested Real Time Data from stop:"+string(stopCode), c.ClientIP(), "OK")
	if loggingError != nil {
		lib.RespondWithError(c, http.StatusInternalServerError, loggingError.Error())
		return
	}
}
