package api

import (
	"github.com/alioth-center/foreseen/app/entity"
	"github.com/alioth-center/foreseen/app/service"
	"github.com/alioth-center/infrastructure/network/http"
)

var NotifyApi NotifyApiImpl

type NotifyApiImpl struct {
	srv *service.NotifyService
}

func (impl NotifyApiImpl) NotifyLark() http.Chain[*entity.LarkNotifyRequest, *entity.BaseResponse] {
	return http.NewChain[*entity.LarkNotifyRequest, *entity.BaseResponse](service.CheckToken[*entity.LarkNotifyRequest], impl.srv.NotifyLark)
}
