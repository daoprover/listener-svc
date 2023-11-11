package requests

import (
	"encoding/json"
	"github.com/daoprover/listener-svc/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"net/http"
)

type CreateInfoOrderRequest struct {
	Data resources.CreateInfoOrder
}

func NewCreateInfoOrderRequest(r *http.Request) (*CreateInfoOrderRequest, error) {
	request := new(CreateInfoOrderRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, errors.Wrap(err, "failed to decode raw request ")
	}

	if err := validateTemplateData(request.Data); err != nil {
		return nil, errors.Wrap(err, "failed to validate data")

	}

	return request, nil
}

func validateTemplateData(template resources.CreateInfoOrder) error {
	return MergeErrors(validation.Errors{
		"/attributes/name": validation.Validate(template.Attributes.Name,
			validation.Required),
		"/attributes/link": validation.Validate(template.Attributes.Link,
			validation.Required),
	}).Filter()
}
