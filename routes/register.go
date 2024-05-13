package routes

import (
	"net/http"
	"strconv"
	"udemy/restapi/models"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt parse event id"})
		return
	}

	//CHECK ID EXIST or NOT
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt fetch event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt register user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered!"})

}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, _ := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event

	event.ID = eventId

	err := event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt cancel registration of user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled!"})

}
