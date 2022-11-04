package requests

import (
	"encoding/json"
	"net/http"

	"blobs-svc/resources"
	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func NewCreateAccountRequest(r *http.Request) (resources.AccountResponse, error) {
	var request resources.AccountResponse
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}
	return request, ValidateCreateAccountRequest(request)
}

func ValidateCreateAccountRequest(r resources.AccountResponse) error {
	errs := validation.Errors{
		"/data/":                      validation.Validate(r.Data, validation.Required),
		"/data/relationships/signers": validation.Validate(r.Data.Relationships.Signers),
	}

	return errs.Filter()
}
