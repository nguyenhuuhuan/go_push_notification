package signal

import "sync"

type topic struct {
	listeners []chan<- struct{}
	mu        *sync.Mutex
}

type signal struct {
	listeners *sync.Map
	topicSize int
}

func (s *signal) Subscribe(id string) (<-chan struct{}, func(), error) {
	topicInf, _ := s.listeners.LoadOrStore(id, &topic{mu: new(sync.Mutex)})
	t := topicInf.(*topic)
	t.mu.Lock()
	defer t.mu.Unlock()
	ch := make(chan struct{}, 1)
	t.listeners = append(t.listeners, ch)
	return ch, func() {
		defer t.mu.Unlock()
		for i := 0; i < len(t.listeners); i++ {
			if t.listeners[i] == ch {
				t.listeners = append(t.listeners[:i], t.listeners[i+1:]...)
			}
		}
	}, nil
}

func (s *signal) Publish(id string) error {
	topicInf, ok := s.listeners.Load(id)
	if !ok {
		return ErrEmpty
	}
	topic := topicInf.(*topic)
	l := len(topic.listeners)
	if l == 0 {
		return ErrEmpty
	}

	for i := 0; i < l; i++ {
		topic.listeners[i] <- struct{}{}
	}
	return nil
}

func NewSignal() Signal {
	return &signal{
		listeners: new(sync.Map),
	}
}
