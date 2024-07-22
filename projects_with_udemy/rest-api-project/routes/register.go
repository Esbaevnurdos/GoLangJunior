package routes

import (
	"net/http"
	"strconv"

	"example.com/project/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
    eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch event."})
		return
	}

    err = event.Register(userId)
		if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message":"Registired"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
    eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	    err = event.Register(userId)
		if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not cancel registration."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message":"Canceled"})
}