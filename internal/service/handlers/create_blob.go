package handlers

import (
	"blobs-svc/internal/data"
	"blobs-svc/internal/data/pg"
	"blobs-svc/resources"
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"blobs-svc/internal/service/requests"
)

func CreateBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateBlobRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	blob, err := requests.Blob(request.Data, 0)
	if err != nil {
		Log(r).WithError(err).Warn("invalid blob type")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{"/data/type": errors.New("invalid blob type")})...)
		return
	}

	err = BlobQ(r).Transaction(func(blobQ data.BlobQ) error {
		err := blobQ.Create(blob)
		if err != nil {
			return errors.Wrap(err, "failed to create blob")
		}
		return nil
	})

	if err != nil {
		if errors.Cause(err) == pg.ErrBlobsConflict {
			Log(r).WithError(err).Error("conflict")
			ape.RenderErr(w, problems.Conflict())
			return
		}

		Log(r).WithError(err).Error("failed to save blob")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := resources.BlobResponse{
		Data: NewBlob(blob),
	}

	w.WriteHeader(201)
	ape.Render(w, &response)
}
