package globals

import (
	"sync"

	"github.com/gesius/gokit/topic"
)

// Registry for global accessible objects, options, functions and subscriptions
type Globals struct {
	r *registry
}

var (
	r = &registry{
		objects:  make(map[string]*interface{}),
		commands: make(map[string]func() interface{}),
		options:  make(map[string]func() interface{}),
		topics:   make(map[string]*topic.Topic),
	}
)

// Globals returns an Object given name. Returns nil if the Object does not exist.
func (g *Globals) Object(name string) (*interface{}, bool) {
	if g == nil {
		return nil, false
	}
	return g.r.object(name)
}

// Globals returns an Command given name. Returns nil if the Command does not exist.
func (g *Globals) Commmand(name string) (func() interface{}, bool) {
	if g == nil {
		return nil, false
	}
	return g.r.command(name)
}

// Globals returns an Command given name. Returns nil if the Command does not exist.
func (g *Globals) Option(name string) (func() interface{}, bool) {
	if g == nil {
		return nil, false
	}
	return g.r.option(name)
}

// Globals returns an Command given name. Returns nil if the Command does not exist.
func (g *Globals) Topic(name string) (*topic.Topic, bool) {
	if g == nil {
		return nil, false
	}
	return g.r.topic(name)
}

type registry struct {
	sync.Mutex // protects entries from concurrent mutation
	objects    map[string]*interface{}
	commands   map[string]func() interface{}
	options    map[string]func() interface{}
	topics     map[string]*topic.Topic
}

func (r *registry) addobject(key string, o *interface{}) {
	r.Lock()
	defer r.Unlock()
	r.objects[key] = o
}

func (r *registry) addcommand(key string, f func() interface{}) {
	r.Lock()
	defer r.Unlock()
	r.commands[key] = f
}

func (r *registry) addoption(key string, f func() interface{}) {
	r.Lock()
	defer r.Unlock()
	r.options[key] = f
}

func (r *registry) addtopic(key string, t *topic.Topic) {
	r.Lock()
	defer r.Unlock()
	r.topics[key] = t
}

func (r *registry) object(key string) (*interface{}, bool) {
	r.Lock()
	defer r.Unlock()
	f, ok := r.objects[key]
	return f, ok
}

func (r *registry) command(key string) (func() interface{}, bool) {
	r.Lock()
	defer r.Unlock()
	f, ok := r.commands[key]
	return f, ok
}

func (r *registry) option(key string) (func() interface{}, bool) {
	r.Lock()
	defer r.Unlock()
	f, ok := r.options[key]
	return f, ok
}

func (r *registry) topic(key string) (*topic.Topic, bool) {
	r.Lock()
	defer r.Unlock()
	f, ok := r.topics[key]
	return f, ok
}

func (r *registry) keys() (k []string) {
	r.Lock()
	defer r.Unlock()
	for e := range r.objects {
		k = append(k, e)
	}

	for e := range r.commands {
		k = append(k, e)
	}

	for e := range r.options {
		k = append(k, e)
	}

	for e := range r.topics {
		k = append(k, e)
	}
	return
}
