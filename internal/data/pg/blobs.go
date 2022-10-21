package pg

import (
	"blobs-svc/internal/data"
	"blobs-svc/internal/types"
	"database/sql"
	"github.com/lib/pq"
	"github.com/pkg/errors"

	"fmt"
	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	blobsTableName    = "blobs"
	blobsPKConstraint = "blobs_pkey"
)

var (
	ErrBlobsConflict = errors.New("blobs primary key conflict")
	blobsSelect      = sq.
				Select("id", "owner_address", "value", "type", "creator_signer_role").
				From(blobsTableName)
)

func NewBlobQ(db *pgdb.DB) data.BlobQ {
	return &BlobQ{
		db:  db.Clone(),
		sql: sq.Select("b.*").From(fmt.Sprintf("%s as b", blobsTableName)),
	}
}

type BlobQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func (q *BlobQ) New() data.BlobQ {
	return NewBlobQ(q.db)
}

func (q *BlobQ) Transaction(fn func(q data.BlobQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *BlobQ) Create(blob *types.Blob) error {
	stmt := sq.Insert(blobsTableName).SetMap(map[string]interface{}{
		"id":                  blob.ID,
		"type":                blob.Type,
		"value":               blob.Value,
		"owner_address":       *blob.OwnerAddress,
		"creator_signer_role": blob.CreatorSignerRole,
	})

	err := q.db.Exec(stmt)
	if err != nil {
		pqerr, ok := errors.Cause(err).(*pq.Error)
		if ok {
			if pqerr.Constraint == blobsPKConstraint {
				return ErrBlobsConflict
			}
		}
	}
	return err
}

func (q *BlobQ) Get() (*types.Blob, error) {
	var result types.Blob
	err := q.db.Get(&result, q.sql)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

func (q *BlobQ) Delete(id string) error {
	stmt := sq.Delete(blobsTableName).Where(sq.Eq{"id": id})
	err := q.db.Exec(stmt)
	return err
}

func (q *BlobQ) Select() ([]types.Blob, error) {
	var result []types.Blob
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *BlobQ) Page(pageParams pgdb.OffsetPageParams) data.BlobQ {
	q.sql = pageParams.ApplyTo(q.sql, "id")
	return q
}

func (q *BlobQ) FilterByID(id string) data.BlobQ {
	q.sql = q.sql.Where(sq.Eq{"b.id": id})
	return q
}

func (q *BlobQ) FilterByOwnerAddress(address types.Address) data.BlobQ {
	q.sql = q.sql.Where(sq.Eq{"b.owner_address": address})
	return q
}
