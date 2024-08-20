package samplepersonality

import (
	"context"
	"net/http"

	"k8s.io/klog/v2"
)

type SamplePersonality struct {
	handlers Handlers
}

type Handlers struct {
	// Read
	Checkpoint,
	Tile,
	EntryBundle,

	// Write
	Add func(http.ResponseWriter, *http.Request)
}

func New(ctx context.Context, handlers Handlers) SamplePersonality {
	return SamplePersonality{
		handlers: handlers,
	}
}

func (p *SamplePersonality) Run(addr string) {
	http.HandleFunc("GET /checkpoint", p.handlers.Checkpoint)
	http.HandleFunc("GET /tile/{level}/{index...}", p.handlers.Tile)
	http.HandleFunc("GET /tile/entries/{index...}", p.handlers.EntryBundle)
	http.HandleFunc("POST /add", p.handlers.Add)

	if err := http.ListenAndServe(addr, http.DefaultServeMux); err != nil {
		klog.Exitf("ListenAndServe: %v", err)
	}
}
