package requests

import (
	"blobs-svc/internal/types"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
	"net/http"
)

type GetBlobsListRequest struct {
	pgdb.OffsetPageParams
	FilterOwnerAddress *types.Address `filter:"owner_address"`
}

func NewGetBlobsListRequest(r *http.Request) (GetBlobsListRequest, error) {
	request := GetBlobsListRequest{}
	err := urlval.Decode(r.URL.Query(), &request)
	return request, err
}
