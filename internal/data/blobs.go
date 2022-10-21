package data

import (
	. "blobs-svc/internal/types"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type BlobQ interface {
	New() BlobQ
	Transaction(fn func(BlobQ) error) error
	Create(blob *Blob) error
	Get() (*Blob, error)
	Delete(id string) error
	Select() ([]Blob, error)
	Page(pageParams pgdb.OffsetPageParams) BlobQ
	FilterByID(id string) BlobQ
	FilterByOwnerAddress(address Address) BlobQ
}
