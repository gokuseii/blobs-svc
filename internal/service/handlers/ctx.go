package handlers

import (
	"blobs-svc/internal/data"
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	blobQCtxKey
)

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
