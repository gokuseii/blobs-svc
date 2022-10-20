package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	. "github.com/go-ozzo/ozzo-validation"
)

type GetBlobRequest struct {
	BlobID string `json:"-"`
}

func NewGetBlobRequest(r *http.Request) (GetBlobRequest, error) {
	request := GetBlobRequest{
		BlobID: chi.URLParam(r, "id"),
	}
	return request, request.Validate()
}

func (r GetBlobRequest) Validate() error {
	err := Errors{
		"blobId": Validate(&r.BlobID, Required),
	}
	return err.Filter()
}
