package hello

import (
	"context"
	"sync"
)

type QueryID uint64
type ResourceManagement struct {
	Priority int32
}
type Spec struct {
	Resources ResourceManagement `json:"resources"`
}
type State int
type Query struct {
	id      QueryID
	spec    Spec
	c       *Controller
	cancel  func()
	stateMu sync.RWMutex
	state   State
	err     error

	parentCtx, currentCtx context.Context
}

func (q *Query) Cancel() {
	q.cancel()
}

func (q *Query) isOK() bool {
	return true
}

type Controller struct {
	shutdownCtx   context.Context
	shutdown      func()
	done          chan struct{}
	queriesMu     sync.RWMutex
	queries       map[QueryID]*Query
	queryDone     chan *Query
	newQueries    chan *Query
	cancelRequest chan QueryID
}

func New() *Controller {
	ctrl := &Controller{
		queries:       make(map[QueryID]*Query),
		queryDone:     make(chan *Query),
		cancelRequest: make(chan QueryID),
		done:          make(chan struct{}),
		newQueries:    make(chan *Query),
	}
	ctrl.shutdownCtx, ctrl.shutdown = context.WithCancel(context.Background())
	go ctrl.run()
	return ctrl
}

func (c *Controller) free(q *Query) {
}
func (c *Controller) run() {
	pq := newPriorityQueue()
	for {
		select {
		case q := <-c.queryDone:
			c.free(q)
			c.queriesMu.Lock()
			delete(c.queries, q.id)
			c.queriesMu.Unlock()
		case q := <-c.newQueries:
			c.queriesMu.Lock()
			c.queries[q.id] = q
			c.queriesMu.Unlock()
		case id := <-c.cancelRequest:
			c.queriesMu.RLock()
			q := c.queries[id]
			c.queriesMu.RUnlock()
			q.Cancel()

		case <-c.shutdownCtx.Done():
			// We have been signaled to shutdown so drain the queues
			// and exit the for loop.
			//c.drain(pq)
			return
		}

		q := pq.Peek()
		if q != nil {
			pop, err := c.processQuery(q)
			if pop {
				pq.Pop()
			}
			if err != nil {
				q.setErr(err)
			}
		}
	}
}
func (q *Query) setErr(err error) {
	q.stateMu.Lock()
	defer q.stateMu.Unlock()

	select {
	case <-q.parentCtx.Done():
	}
}

func (c *Controller) processQuery(q *Query) (pop bool, err error) {
	//consume q
	//...
	//remove query from queue
	pop = true
	return pop, nil
}
