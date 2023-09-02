package controllers

import (
	"context"
	"mime"
	"zapmeow/models"
	"zapmeow/services"
	"zapmeow/utils"

	"github.com/gin-gonic/gin"
	"github.com/vincent-petithory/dataurl"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

type sendAudioMessageController struct {
	wppService     services.WppService
	messageService services.MessageService
}

func NewSendAudioMessageController(
	wppService services.WppService,
	messageService services.MessageService,
) *sendAudioMessageController {
	return &sendAudioMessageController{
		wppService:     wppService,
		messageService: messageService,
	}
}

func (a *sendAudioMessageController) Handler(c *gin.Context) {
	type Body struct {
		Phone  string
		Base64 string
	}

	var body Body
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.RespondBadRequest(c, "Body data is invalid")
		return
	}

	jid, ok := utils.MakeJID(body.Phone)
	if !ok {
		utils.RespondBadRequest(c, "Invalid phone")
		return
	}
	instanceId := c.Param("instanceId")

	mimitype, err := utils.GetMimeTypeFromDataURI(body.Base64)
	if err != nil {
		utils.RespondInternalServerError(c, err.Error())
		return
	}

	instance, err := a.wppService.GetAuthenticatedInstance(instanceId)
	if err != nil {
		utils.RespondInternalServerError(c, err.Error())
		return
	}

	audioURL, err := dataurl.DecodeString(body.Base64)
	if err != nil {
		utils.RespondInternalServerError(c, err.Error())
		return
	}

	uploaded, err := instance.Upload(
		context.Background(),
		audioURL.Data,
		whatsmeow.MediaAudio,
	)
	if err != nil {
		utils.RespondInternalServerError(c, err.Error())
		return
	}

	msg := &waProto.Message{
		AudioMessage: &waProto.AudioMessage{
			Ptt:           proto.Bool(true),
			Url:           proto.String(uploaded.URL),
			DirectPath:    proto.String(uploaded.DirectPath),
			MediaKey:      uploaded.MediaKey,
			Mimetype:      proto.String(mimitype),
			FileEncSha256: uploaded.FileEncSHA256,
			FileSha256:    uploaded.FileSHA256,
			FileLength:    proto.Uint64(uint64(len(audioURL.Data))),
		},
	}

	resp, err := instance.SendMessage(context.Background(), jid, msg)
	if err != nil {
		utils.RespondInternalServerError(c, err.Error())
		return
	}

	dir, err := utils.MakeUserDirectory(instance.Store.ID.User)
	exts, _ := mime.ExtensionsByType(mimitype)
	path, err := utils.SaveMedia(
		audioURL.Data,
		dir,
		resp.ID,
		exts[0],
	)
	if err != nil {
		utils.RespondInternalServerError(c, err.Error())
		return
	}

	message := models.Message{
		ChatJID:   jid.User,
		SenderJID: instance.Store.ID.User,
		MeJID:     instance.Store.ID.User,
		MediaType: "audio",
		MediaPath: path,
		Timestamp: resp.Timestamp,
		FromMe:    true,
		MessageID: resp.ID,
	}

	err = a.messageService.CreateMessage(&message)
	if err != nil {
		utils.RespondInternalServerError(c, err.Error())
		return
	}

	utils.RespondWithSuccess(c, gin.H{
		"Message": a.messageService.ToJSON(message),
	})
}