package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	. "github.com/go-ozzo/ozzo-validation"
)

type DeleteBlobRequest struct {
	BlobID string `json:"-"`
}

func NewDeleteBlobRequest(r *http.Request) (DeleteBlobRequest, error) {
	request := DeleteBlobRequest{
		BlobID: chi.URLParam(r, "id"),
	}
	return request, request.Validate()
}

func (r DeleteBlobRequest) Validate() error {
	err := Errors{
		"blobId": Validate(&r.BlobID, Required),
	}
	return err.Filter()
}
