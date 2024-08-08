package service

import (
	"encoding/json"

	"github.com/alioth-center/foreseen/app/entity"
	"github.com/alioth-center/infrastructure/logger"
	"github.com/alioth-center/infrastructure/network/http"
	"github.com/alioth-center/infrastructure/thirdparty/lark"
)

func NewCallbackService() *CallbackService {
	return &CallbackService{client: lark.NewClient(lark.Config{AppID: entity.GlobalConfig.AppID, AppSecret: entity.GlobalConfig.AppSecret})}
}

type CallbackService struct {
	client lark.Client
}

func (srv *CallbackService) Callback(ctx http.Context[any, json.RawMessage]) {
	log.Info(logger.NewFields(ctx).WithMessage("callback received").WithData(ctx.Request()))

	payload, _ := json.Marshal(ctx.Request())

	_ = srv.client.SendMarkdownMessage(
		ctx,
		lark.Receiver{
			Type:     lark.LarkReceiverIdTypeChatID,
			Receiver: entity.GlobalConfig.DefaultChatID,
		},
		"Webhook Received",
		string(payload),
		lark.LarkMarkdownMessageThemeBlue,
	)

	ctx.SetStatusCode(200)
	ctx.SetResponse([]byte("{}"))
}
