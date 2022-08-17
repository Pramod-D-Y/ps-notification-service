package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/roppenlabs/ps-notification-service/internal/model"
)

func CreateNotification(ctx *gin.Context) {
	var notification model.Notification

	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	publishError := PublishNotification(notification)

	if publishError != nil {
		panic(publishError)
	}

	ctx.JSON(http.StatusCreated, notification)
}

func PublishNotification(notification model.Notification) error {
	posturl := []string{
		"https://hooks.slack.com/services/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",
		"https://hooks.slack.com/services/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"}
	reqBody := map[string]string{

		"text": fmt.Sprintf("<@U03PLJQRNET> Sent request for Access"),
	}
	bodyJson, _ := json.Marshal(reqBody)

	for index := 0; index < len(posturl); index++ {
		r, err := http.NewRequest("POST", posturl[index], bytes.NewBuffer(bodyJson))
		if err != nil {
			panic(err)
		}

		r.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(r)
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()
	}
	return nil
}
