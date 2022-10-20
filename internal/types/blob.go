package types

import (
	"fmt"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Blob struct {
	ID                string
	Type              BlobType
	Value             string
	OwnerAddress      *Address `db:"owner_address"`
	CreatorSignerRole *uint64  `db:"creator_signer_role"`
}

func GetBlobType(v string) (b BlobType, err error) {
	err = b.UnmarshalJSON([]byte(fmt.Sprintf(`"%s"`, v)))
	if err != nil {
		return b, errors.Wrap(err, "failed to unmarshal blob type")
	}
	return b, nil
}
