package handlers

import (
	"blobs-svc/internal/data"
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/connectors/keyer"
	"gitlab.com/tokend/connectors/submit"
	"net/http"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	blobQCtxKey
	keysCtxKey
	submitterCtxKey
)

func CtxSubmitter(v *submit.Submitter) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, submitterCtxKey, v)
	}
}

func Submitter(r *http.Request) *submit.Submitter {
	return r.Context().Value(submitterCtxKey).(*submit.Submitter)
}

func CtxKeys(v keyer.Keys) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, keysCtxKey, v)
	}
}

func Keys(r *http.Request) keyer.Keys {
	return r.Context().Value(keysCtxKey).(keyer.Keys)
}

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxBlobQ(entry data.BlobQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, blobQCtxKey, entry)
	}
}

func BlobQ(r *http.Request) data.BlobQ {
	return r.Context().Value(blobQCtxKey).(data.BlobQ).New()
}
