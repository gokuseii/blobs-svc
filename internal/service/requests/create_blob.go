package requests

import (
	"blobs-svc/internal/types"
	"encoding/base32"
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"

	"blobs-svc/docs/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/hash"
)

type CreateBlobRequest struct {
	Data resources.Blob
}

func NewCreateBlobRequest(r *http.Request) (CreateBlobRequest, error) {
	var request CreateBlobRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *CreateBlobRequest) validate() error {
	return validation.Errors{
		"/data/type":                        validation.Validate(&r.Data.Type, validation.Required),
		"/data/attributes/value":            validation.Validate(&r.Data.Attributes.Value, validation.Required),
		"/data/relationships/owner/data/id": validation.Validate(&r.Data.Relationships.Owner.Data.ID, validation.Required),
	}.Filter()
}

func Blob(r resources.Blob, allowedRole uint64) (*types.Blob, error) {
	blob, err := types.GetBlobType(string(r.Type))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create blob")
	}
	msg := fmt.Sprintf("%s%d%s", r.Relationships.Owner.Data.ID, blob, r.Attributes.Value)
	hash := hash.Hash([]byte(msg))
	owner := types.Address(r.Relationships.Owner.Data.ID)

	return &types.Blob{
		ID:                base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(hash[:]),
		Type:              blob,
		Value:             r.Attributes.Value,
		CreatorSignerRole: &allowedRole,
		OwnerAddress:      &owner,
	}, nil
}
