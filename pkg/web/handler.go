package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/nityanandagohain/rtms/pkg/pubsub"
)

type Handler struct {
	*chi.Mux
	PubSubService *pubsub.Service
}

func NewHandler(redisAddress string, redisPassword string) *Handler {
	h := Handler{
		Mux:           chi.NewMux(),
		PubSubService: pubsub.New(redisAddress, redisPassword),
	}

	h.Use(middleware.Logger)

	h.Post("/subscribe/{name}", h.Subscribe())
	h.Post("/publish/{name}", h.Publish())

	return &h
}

func (h *Handler) Subscribe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)

		name := chi.URLParam(r, "name")

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		notify := r.Context().Done()

		data := make(chan []byte)

		h.PubSubService.Subscribe(r.Context(), name, data)
		for {
			select {
			case <-notify: // break condition i.e client closed the connection
				return
			case msg := <-data: // stream the messages and keep the connecion open
				fmt.Fprintf(w, "data: %s\n\n", msg)

				flusher.Flush()
			}
		}
	}
}

func (h *Handler) Publish() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Query().Get("message")

		name := chi.URLParam(r, "name")

		err := h.PubSubService.Publish(r.Context(), name, message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
