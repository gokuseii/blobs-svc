package requests

import (
	"blobs-svc/resources"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/distributed_lab/ape/apeutil"
)

func TestNewCreateBlobRequest(t *testing.T) {
	data := `
{
	"data": {
		"id": "blah",
		"type": "asset_description",
		"attributes": {
			"value": "foobar"
		},
		"relationships": {
			"owner": {
				"data": {
					"id": "GBSR6JG5AYSAW7HK6EGJFYVIVN54LVGSY3ZLJ6X3IBQZ766EJABCZQTH"
						}
					}
				}
			}
		}
`
	t.Run("valid", func(t *testing.T) {
		r := apeutil.RequestWithURLParams([]byte(data), map[string]string{})
		got, err := NewCreateBlobRequest(r)
		assert.NoError(t, err)
		assert.Equal(t, resources.ResourceType("asset_description"), got.Data.Type)
		assert.Equal(t, "blah", got.Data.ID)
		assert.Equal(t, "foobar", got.Data.Attributes.Value)
		assert.Equal(t, "GBSR6JG5AYSAW7HK6EGJFYVIVN54LVGSY3ZLJ6X3IBQZ766EJABCZQTH", got.Data.Relationships.Owner.Data.ID)
	})
}

func TestNewCreateBlobRequest_NoOwner(t *testing.T) {
	data := `
{
	"data": {
		"id": "blah",
		"type": "asset_description",
		"attributes": {
			"value": "foobar"
		}
		}}
`
	t.Run("invalid", func(t *testing.T) {
		r := apeutil.RequestWithURLParams([]byte(data), map[string]string{})
		_, err := NewCreateBlobRequest(r)
		assert.Error(t, err)
		assert.Equal(t, "/data/relationships/owner/data/id: cannot be blank.", err.Error())
	})
}

func TestNewCreateBlobRequest_WithOwner(t *testing.T) {
	data := `
{
	"data": {
		"id": "blah",
		"type": "asset_description",
		"attributes": {
			"value": "foobar"
		}
		}}
`
	t.Run("valid", func(t *testing.T) {
		r := apeutil.RequestWithURLParams([]byte(data), map[string]string{"address": "GBSR6JG5AYSAW7HK6EGJFYVIVN54LVGSY3ZLJ6X3IBQZ766EJABCZQTH"})
		got, err := NewCreateBlobRequest(r)
		assert.NoError(t, err)
		assert.Equal(t, resources.ResourceType("asset_description"), got.Data.Type)
		assert.Equal(t, "blah", got.Data.ID)
		assert.Equal(t, "foobar", got.Data.Attributes.Value)
		assert.Equal(t, "GBSR6JG5AYSAW7HK6EGJFYVIVN54LVGSY3ZLJ6X3IBQZ766EJABCZQTH", got.Data.Relationships.Owner.Data.ID)
	})
}
