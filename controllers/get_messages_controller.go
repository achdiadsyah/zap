package controllers

import (
	"zapmeow/services"
	"zapmeow/utils"

	"github.com/gin-gonic/gin"
)

type getMessagesController struct {
	wppService     services.WppService
	messageService services.MessageService
}

func NewGetMessagesController(
	wppService services.WppService,
	messageService services.MessageService,
) *getMessagesController {
	return &getMessagesController{
		wppService:     wppService,
		messageService: messageService,
	}
}

func (m *getMessagesController) Handler(c *gin.Context) {
	type Body struct {
		Phone string
	}

	var body Body
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.RespondBadRequest(c, "Body data is invalid")
		return
	}
	instanceID := c.Param("instanceId")

	instance, err := m.wppService.GetAuthenticatedInstance(instanceID)
	if err != nil {
		utils.RespondInternalServerError(c, err.Error())
		return
	}

	messages, err := m.messageService.GetChatMessages(
		instance.Store.ID.User,
		body.Phone,
	)
	if err != nil {
		utils.RespondInternalServerError(c, err.Error())
		return
	}

	var data = []gin.H{}
	for _, message := range *messages {
		data = append(data, m.messageService.ToJSON(message))
	}

	utils.RespondWithSuccess(c, gin.H{
		"Messages": data,
	})
}