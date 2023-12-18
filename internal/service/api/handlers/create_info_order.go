package handlers

import (
	"github.com/daoprover/listener-svc/internal/service/api/requests"
	"github.com/daoprover/listener-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

func CreateInfoOrder(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateInfoOrderRequest(r)
	if err != nil {
		Log(r).Error(errors.Wrap(err, "failed  to get request"))
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	masterRunner := MasterRunner(r)
	orderID, err := masterRunner.AddToQuery(request.Data.Attributes.Name, request.Data.Attributes.Link, request.Data.Attributes.TimeFrom, request.Data.Attributes.TimeTo)
	if err != nil {
		Log(r).Error(errors.Wrap(err, "failed  to  init order"))
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	w.WriteHeader(http.StatusCreated)
	ape.Render(w, CreateInfoOrderResponse(*orderID))

}

func CreateInfoOrderResponse(orderID string) resources.OrderResponse {
	return resources.OrderResponse{
		Data: resources.Order{
			Attributes: resources.OrderAttributes{
				Id: orderID,
			},
		},
	}
}
