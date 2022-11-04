package handlers

import (
	"blobs-svc/internal/data"
	"blobs-svc/internal/service/requests"
	"blobs-svc/resources"
	"encoding/json"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/connectors/submit"
	"gitlab.com/tokend/go/xdrbuild"
	"net/http"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	log := Log(r)

	request, err := requests.NewCreateAccountRequest(r)
	if err != nil {
		log.WithError(err).Error("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	signers, err := getSigners(request)
	if err != nil {
		log.WithError(err).Error("troubles with signers")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	signersData := make([]xdrbuild.SignerData, 0, len(signers))

	for _, s := range signers {
		signersData = append(signersData, xdrbuild.SignerData{
			PublicKey: s.SignerID,
			RoleID:    s.RoleID,
			Weight:    s.Weight,
			Identity:  s.Identity,
			Details:   Details{},
		})
	}

	tx := Submitter(r)
	keys := Keys(r)

	env, _, err := tx.CreateTransaction(keys.Source, keys.Signer, &xdrbuild.CreateAccount{
		Destination: request.Data.ID,
		RoleID:      1,
		Signers:     signersData,
	})
	if err != nil {
		log.WithError(err).Error("failed to create transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	resp, err := tx.Submit(r.Context(), env, false)
	if err != nil {
		if serr, ok := err.(submit.TxFailure); ok {
			log.WithError(err).Error("tx failure")
			w.WriteHeader(http.StatusBadRequest)
			ape.Render(w, serr.Response)
			return
		}

		log.WithError(err).Error("failed to submit transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resp)
}

func getSigners(request resources.AccountResponse) ([]data.AccountSigner, error) {
	var signers []data.AccountSigner
	for _, signerKey := range request.Data.Relationships.Signers.Data {
		signer := request.Included.MustSigner(signerKey)
		if signer == nil {
			return nil, validation.Errors{"/included": errors.New("missed signer include")}
		}
		signers = append(signers, data.AccountSigner{
			SignerID: signerKey.ID,
			RoleID:   uint64(signer.Attributes.RoleId),
			Weight:   uint32(signer.Attributes.Weight),
			Identity: uint32(signer.Attributes.Identity),
		})
	}
	return signers, nil
}

type Details map[string]interface{}

func (d Details) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}(d))
}
