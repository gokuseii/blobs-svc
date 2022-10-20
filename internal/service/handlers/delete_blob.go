package handlers

import (
	"database/sql"
	"net/http"

	"blobs-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = BlobQ(r).Delete(request.BlobID)
	if err != nil {
		if err == sql.ErrNoRows {
			ape.RenderErr(w, problems.NotFound())
			return
		}
		Log(r).WithError(err).Error("failed to delete blob")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(204)
}
