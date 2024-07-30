package tessera

import (
	"context"
	"sync"
	"time"

	"github.com/globocom/go-buffer"
	"k8s.io/klog/v2"
)

type DeduperStorage interface {
	Set(context.Context, []DedupEntry) error
	Index(context.Context, []byte) (*uint64, error)
}

type Deduper struct {
	ctx     context.Context
	storage DeduperStorage

	mu             sync.Mutex
	numLookups     uint64
	numWrites      uint64
	numCacheDedups uint64
	numDBDedups    uint64
	numPushErrs    uint64

	buf *buffer.Buffer
}

func NewDeduper(ctx context.Context, s DeduperStorage) *Deduper {
	r := &Deduper{
		ctx:     ctx,
		storage: s,
	}

	r.buf = buffer.New(
		buffer.WithSize(64),
		buffer.WithFlushInterval(200*time.Millisecond),
		buffer.WithFlusher(buffer.FlusherFunc(r.flush)),
		buffer.WithPushTimeout(15*time.Second),
	)
	go func(ctx context.Context) {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				r.mu.Lock()
				klog.Infof("DEDUP: # Writes %d, # Lookups %d, # DB hits %v, # buffer Push discards %d", r.numWrites, r.numLookups, r.numDBDedups, r.numPushErrs)
				r.mu.Unlock()
			}
		}
	}(ctx)
	return r
}

type DedupEntry struct {
	ID  []byte
	Idx uint64
}

func (s *Deduper) inc(p *uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	(*p)++
}

func (s *Deduper) Index(ctx context.Context, h []byte) (*uint64, error) {
	s.inc(&s.numLookups)
	r, err := s.storage.Index(ctx, h)
	if r != nil {
		s.inc(&s.numDBDedups)
	}
	return r, err
}

func (s *Deduper) Set(_ context.Context, h []byte, idx uint64) error {
	err := s.buf.Push(DedupEntry{ID: h, Idx: idx})
	if err != nil {
		s.inc(&s.numPushErrs)
		// This means there's pressure flushing dedup writes out, so discard this write.
		if err != buffer.ErrTimeout {
			return err
		}
	}
	return nil
}

func (s *Deduper) flush(items []interface{}) {
	entries := make([]DedupEntry, len(items))
	for i := range items {
		entries[i] = items[i].(DedupEntry)
	}

	ctx, c := context.WithTimeout(s.ctx, 15*time.Second)
	defer c()

	if err := s.storage.Set(ctx, entries); err != nil {
		klog.Infof("Failed to flush dedup entries: %v", err)
		return
	}
	s.mu.Lock()
	s.numWrites += uint64(len(entries))
	s.mu.Unlock()
}
