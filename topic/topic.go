package topic

import (
	"container/list"
	"github.com/gesius/gokit/rand"
	"runtime"
	"sync"
)

// Summarize returns a list of items meant to summarize the history of the stream so far
// for subscribers joining now.
type Summarize func() []interface{}

// Topic
type Topic struct {
	name      string
	publisher struct {
		sync.Mutex
		src chan<- interface{}
	}
	subscriber struct {
		sum Summarize
		sync.Mutex
		group map[int]*queue //
		n     int
	}
}

func NewTopic(name string, sum Summarize) (t *Topic) {
	src := make(chan interface{})
	tp := &Topic{name: rand.NewRND().RandomString("topic:", name)}
	tp.publisher.src = src
	tp.subscriber.sum = sum
	tp.subscriber.group = make(map[int]*queue)
	go tp.loop(src)
	return tp
}

// Source returns the name of the event source.
func (tp *Topic) Source() string {
	return tp.name
}

// Publish appends a value onto the infinite update stream.
func (tp *Topic) Publish(v interface{}) {
	tp.publisher.Lock()
	defer tp.publisher.Unlock()
	if tp.publisher.src == nil {
		panic("publish after close")
	}
	tp.publisher.src <- v
}

// Close terminates, paradoxically, the infinite update stream.
func (tp *Topic) Close() {
	tp.publisher.Lock()
	defer tp.publisher.Unlock()
	if tp.publisher.src == nil {
		return
	}
	close(tp.publisher.src)
	tp.publisher.src = nil
}

// loop churns messages between the publishing entity, using Publish(),
// and the multiple registered subscriber entities.
func (tp *Topic) loop(src <-chan interface{}) {
	for {
		v, ok := <-src
		if !ok {
			tp.clunk() //if publisher closed channel, release all mmbers and subscriptions
			return
		}
		tp.distribute(v)
	}
}

func (tp *Topic) distribute(v interface{}) {
	tp.subscriber.Lock()
	defer tp.subscriber.Unlock()
	for _, q := range tp.subscriber.group {
		q.distribute(v)
	}
}

func (tp *Topic) clunk() {
	tp.subscriber.Lock()
	defer tp.subscriber.Unlock()
	for _, q := range tp.subscriber.group {
		q.close()
	}
}

// Subscribe creates a new subscription object, whose interface embodies reading from an infinite stream.
// New subscription can join at any time. The input stream of each individual subscription is pre-loaded
// with a sequence of values summarizing all past history. Subsequent values come from the pubish stream.
// Subscriptions are abandoned on garbage-collection.
func (tp *Topic) Subscribe() *Subscription {
	tp.subscriber.Lock()
	defer tp.subscriber.Unlock()
	q := newQueue(tp, tp.subscriber.n)
	tp.subscriber.group[q.id] = q
	tp.subscriber.n++
	// Prefix subscription's input stream with a summary of all history until now
	if tp.subscriber.sum != nil {
		for _, v := range tp.subscriber.sum() {
			q.distribute(v)
		}
	}
	return q.use()
}

// Unsubscribe removes a subscription queue from the member table, only if
// all Subscription handles referring to it have been collected.
func (tp *Topic) Unsubscribe(id int) {
	tp.subscriber.Lock()
	defer tp.subscriber.Unlock()
	q, ok := tp.subscriber.group[id]
	if !ok || q.isBusy() {
		return
	}
	delete(tp.subscriber.group, id)
}

// queue…
type queue struct {
	tp  *Topic
	id  int
	ch1 chan<- interface{} // disribute() => loop()
	ch2 <-chan interface{} // loop() => consume()
	sync.Mutex
	nref   int  // number of references to this queue
	pend   int  // number of buffered messages
	closed bool // true if the source channel has reached EOF
}

func newQueue(tp *Topic, id int) *queue {
	ch1 := make(chan interface{}, 1)
	ch2 := make(chan interface{}, 1)
	q := &queue{
		tp:  tp,
		id:  id,
		ch1: ch1,
		ch2: ch2,
	}
	go q.loop(ch1, ch2)
	return q
}

func (q *queue) addPend(d int) {
	q.Lock()
	defer q.Unlock()
	q.pend += d
}

func (q *queue) setClosed(v bool) {
	q.Lock()
	defer q.Unlock()
	q.closed = v
}

type Stat struct {
	Source  string
	Pending int
	Closed  bool
}

func (q *queue) Peek() Stat {
	q.Lock()
	defer q.Unlock()
	return Stat{
		Source:  q.tp.Source(),
		Pending: q.pend,
		Closed:  q.closed,
	}
}

func (q *queue) close() {
	close(q.ch1)
}

func (q *queue) distribute(v interface{}) {
	q.ch1 <- v
}

// loop churns messages from the main loop onto the internal buffer of this subscription,
// and from there out to the consumer, as requested by calls to Consume.
func (q *queue) loop(ch1 <-chan interface{}, ch2 chan<- interface{}) {
	var l list.List
__preclose:
	for {
		if w := l.Back(); w != nil {
			select {
			case v, ok := <-ch1: // distribute
				if !ok {
					q.setClosed(true)
					break __preclose
				}
				l.PushFront(v)
				q.addPend(1)
			case ch2 <- w.Value: // consume
				l.Remove(w)
				q.addPend(-1)
			}
		} else {
			v, ok := <-ch1
			if !ok {
				q.setClosed(true)
				break __preclose
			}
			l.PushFront(v)
			q.addPend(1)
		}
	}
	// After ch1 has been closed
	for {
		w := l.Back()
		if w == nil {
			close(ch2)
			return
		}
		ch2 <- w.Value
		l.Remove(w)
		q.addPend(-1)
	}
}

func (q *queue) isBusy() bool {
	q.Lock()
	defer q.Unlock()
	return q.nref != 0
}

func (q *queue) recycle() {
	q.Lock()
	defer q.Unlock()
	q.nref--
	if q.nref != 0 {
		return
	}
	go q.tp.Unsubscribe(q.id)
}

// use returns a new subscription, which is effectively a handle for this queue.
func (q *queue) use() *Subscription {
	q.Lock()
	defer q.Unlock()
	q.nref++
	s := &Subscription{q}
	runtime.SetFinalizer(s, func(s2 *Subscription) {
		q.recycle()
	})
	return s
}

func (q *queue) Consume() (v interface{}, ok bool) {
	v, ok = <-q.ch2
	return
}

// Subscription is the user's interface to consuming messages from a topic.
type Subscription struct {
	*queue
}

type Consumer interface {
	Consume() (interface{}, bool)
	Peek() Stat
	Scrub()
}

func (s *Subscription) Scrub() {}

func (s *Subscription) Peek() Stat {
	return s.queue.Peek()
}

func (s *Subscription) Consume() (interface{}, bool) {
	return s.queue.Consume()
}
