package web

import (
	"github.com/kataras/iris/v12"
	"resk/infra"
	"resk/infra/base"
	service "resk/services"
)

func init() {
	infra.RegisterApi(new(EnvelopeApi))
}

type EnvelopeApi struct {
	service service.RedEnvelopeService
}

func (e *EnvelopeApi) Init() {
	e.service = service.GetRedEnvelopeService()
	groupRoute := base.Iris().Party("/v1/envelope")
	groupRoute.Post("/sendout", e.sendOutHandler)
}
func (e *EnvelopeApi) sendOutHandler(context iris.Context) {
	dto := service.RedEnvelopeSendingDTO{}
	err := context.ReadJSON(&dto)
	var res base.Res
	if err != nil {
		res.Code = base.ResCodeRequestParamsError
		res.Message = err.Error()
		context.JSON(&res)
		return
	}
	activity, err := e.service.SendOut(dto)
	if err != nil {
		res.Code = base.ResCodeInnerServerError
		res.Message = err.Error()
		context.JSON(&res)
		return
	}
	res.Code = base.ResCodeOk
	res.Date = activity
	context.JSON(&res)
	return
}
