package api

import (
	"encoding/json"
	"github.com/alioth-center/foreseen/app/service"
	"github.com/alioth-center/infrastructure/network/http"
)

var CallbackApi CallbackApiImpl

type CallbackApiImpl struct {
	srv *service.CallbackService
}

func (impl CallbackApiImpl) Callback() http.Chain[any, json.RawMessage] {
	return http.NewChain[any, json.RawMessage](impl.srv.Callback)
}
