package handlers

import (
	"blobs-svc/internal/service/helpers"
	"blobs-svc/internal/types"
	"net/http"

	"blobs-svc/docs/resources"

	"blobs-svc/internal/data"
	"blobs-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetBlobsList(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobsListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	blobsQ := BlobQ(r)
	applyFilters(blobsQ, request)
	blobs, err := blobsQ.Select()
	if err != nil {
		Log(r).WithError(err).Error("failed to get blob list")
		ape.Render(w, problems.InternalError())
		return
	}

	blobIds := make([]string, len(blobs))
	for i, item := range blobs {
		blobIds[i] = item.ID
	}

	response := resources.BlobListResponse{
		Data:  newBlobsList(blobs),
		Links: helpers.GetOffsetLinks(r, request.OffsetPageParams),
	}

	ape.Render(w, response)
}

func applyFilters(q data.BlobQ, request requests.GetBlobsListRequest) {
	q.Page(request.OffsetPageParams)
}

func newBlobsList(blobs []types.Blob) []resources.Blob {
	result := make([]resources.Blob, len(blobs))
	for i, blob := range blobs {
		result[i] = NewBlob(&blob)
	}
	return result
}
