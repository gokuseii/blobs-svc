package data

import (
	. "blobs-svc/internal/types"
	"gitlab.com/distributed_lab/kit/pgdb"
)

//
//type BlobsQ interface {
//	New() BlobsQ
//
//	Get() (*Blob, error)
//	Select() ([]Blob, error)
//
//	Transaction(fn func(q BlobsQ) error) error
//
//	Insert(data Blob) (Blob, error)
//
//	Page(pageParams pgdb.OffsetPageParams) BlobsQ
//
//	FilterByID(id ...int64) BlobsQ
//}

//go:generate mockery -case underscore -name Blobs
type BlobQ interface {
	New() BlobQ
	Transaction(fn func(BlobQ) error) error
	Create(blob *Blob) error
	Get(id string) (*Blob, error)
	Delete(id string) error
	Select() ([]Blob, error)
	Page(pageParams pgdb.OffsetPageParams) BlobQ
}
