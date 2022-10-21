package handlers

import (
	"blobs-svc/internal/service/requests"
	"net/http"

	"blobs-svc/docs/resources"
	"blobs-svc/internal/types"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	blob, err := BlobQ(r).FilterByID(request.BlobID).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get blob")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if blob == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	response := resources.BlobResponse{
		Data: NewBlob(blob),
	}

	ape.Render(w, &response)
}

func NewBlob(blob *types.Blob) resources.Blob {
	b := resources.Blob{
		Key: resources.Key{
			ID:   blob.ID,
			Type: resources.ResourceType(blob.Type.String()),
		},
		Attributes: resources.BlobAttributes{
			Value: blob.Value,
		},
	}
	return b
}
