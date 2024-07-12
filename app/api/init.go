package api

import "github.com/alioth-center/foreseen/app/service"

func init() {
	NotifyApi = NotifyApiImpl{srv: service.NewNotifyService()}
	CallbackApi = CallbackApiImpl{srv: service.NewCallbackService()}
}
