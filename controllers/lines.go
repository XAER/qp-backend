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

func SearchLineDataHandler(c *gin.Context) {
	// Handle the reqeust to the server with url /SearchLineData
	// the body use as model models.SearchLineDataRequest

	// Get the request body and parse it to the model
	var req models.SearchLineDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lib.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := http.Post(config.App.C.DATA_SERVER+"/SearchLineData", "application/json", strings.NewReader(`{"languageName":"`+req.LanguageName+`","searchedText":"`+req.SearchedText+`"}`))
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

	parsedRes := models.SearchLineDataResponse{}
	json.Unmarshal([]byte(string(text)), &parsedRes)

	// LOG the request in the Logs table

	// Return the response to the client
	c.JSON(http.StatusOK, parsedRes)

	loggingError := lib.LogRequest(200, "Requested data from line:"+string(req.SearchedText), c.ClientIP(), "OK")
	if loggingError != nil {
		lib.RespondWithError(c, http.StatusInternalServerError, loggingError.Error())
		return
	}

}
