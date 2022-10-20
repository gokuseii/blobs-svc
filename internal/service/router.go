package service

import (
	"blobs-svc/internal/config"
	"blobs-svc/internal/data/pg"
	"blobs-svc/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxBlobQ(pg.NewBlobQ(cfg.DB())),
		),
	)

	r.Route("/blobs", func(r chi.Router) {
		r.Post("/", handlers.CreateBlob)
		r.Get("/", handlers.GetBlobsList)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetBlob)
			r.Delete("/", handlers.DeleteBlob)
		})
	})

	return r
}
