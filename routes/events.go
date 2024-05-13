package routes

import (
	"net/http"
	"strconv"
	"udemy/restapi/models"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldnt fetch data, try again later"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "couldnt parse eventId"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldnt fetch event."})
		return
	}

	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {
	// token := context.Request.Header.Get("Authorization")

	// if token == "" {
	// 	context.JSON(http.StatusUnauthorized, gin.H{
	// 		"message": "Not Authorized",
	// 	})
	// 	return
	// }

	// userId, err := utils.VerifyToken(token)
	// if err != nil {
	// 	context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
	// 	return
	// }
	////VERIFY TOKEN should be in Middleware

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "couln't parse data"})
		return
	}

	//dummy value
	// event.ID = 1

	userId := context.GetInt64("userId") //didapat dari proses verify token
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldnt create event, try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})

}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt parse event id"})
		return
	}

	// check id exist or not
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt Fetch the event"})
		return
	}

	// CHECK only user with specific id that can update or delete events
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to Update event!"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt parse request data"})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Update event successfully"})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt parse event id"})
		return
	}

	// check id exist or not
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt Fetch the event"})
		return
	}
	// CHECK only user with specific id that can update or delete events
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to Delete event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Delete event successfully"})

}
