package writer

import (
	"context"
	"sync"
	"time"
)

type Batch struct {
	Entries [][]byte
}

// SequenceFunc knows how to assign contiguous sequence numbers to the entries in Batch.
// Returns the sequence number of the first entry, or an error.
// Must not return successfully until the assigned sequence numbers are durably stored.
type SequenceFunc func(context.Context, Batch) (uint64, error)

func NewPool(bufferSize int, maxAge time.Duration, s SequenceFunc) *Pool {
	return &Pool{
		current: &batch{
			Done: make(chan struct{}),
		},
		bufferSize: bufferSize,
		seq:        s,
		maxAge:     maxAge,
	}
}

// Pool is a helper for adding entries to a log.
type Pool struct {
	sync.Mutex
	current    *batch
	bufferSize int
	maxAge     time.Duration
	flushTimer *time.Timer

	seq SequenceFunc
}

// Add adds an entry to the tree.
// Returns the assigned sequence number, or an error.
func (p *Pool) Add(e []byte) (uint64, error) {
	p.Lock()
	b := p.current
	// If this is the first entry in a batch, set a flush timer so we attempt to sequence it within maxAge.
	if len(b.Entries) == 0 {
		p.flushTimer = time.AfterFunc(p.maxAge, func() {
			p.Lock()
			defer p.Unlock()
			p.flushWithLock()
		})
	}
	n := b.Add(e)
	// If the batch is full, then attempt to sequence it immediately.
	if n >= p.bufferSize {
		p.flushWithLock()
	}
	p.Unlock()
	<-b.Done
	return b.FirstSeq + uint64(n), b.Err
}

func (p *Pool) flushWithLock() {
	// timer can be nil if a batch was flushed because it because full at about the same time as it hit maxAge.
	// In this case we can just return.
	if p.flushTimer == nil {
		return
	}
	p.flushTimer.Stop()
	p.flushTimer = nil
	b := p.current
	p.current = &batch{
		Done: make(chan struct{}),
	}
	go func() {
		b.FirstSeq, b.Err = p.seq(context.TODO(), Batch{Entries: b.Entries})
		close(b.Done)
	}()
}

type batch struct {
	Entries  [][]byte
	Done     chan struct{}
	FirstSeq uint64
	Err      error
}

func (b *batch) Add(e []byte) int {
	b.Entries = append(b.Entries, e)
	return len(b.Entries)
}
