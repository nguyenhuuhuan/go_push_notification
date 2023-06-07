package bolbol

import (
	"push_notification/signal"
	"push_notification/storage"
	"time"
)

type Bolbol struct {
	Storage        storage.Storage
	Signal         signal.Signal
	defaultTimeout time.Duration
}

func NewBolbol(str storage.Storage, sig signal.Signal) *Bolbol {
	return &Bolbol{
		Storage:        str,
		Signal:         sig,
		defaultTimeout: 2 * time.Minute,
	}
}

func Build() *Bolbol {
	str := storage.NewMemoryWithChannel(100)
	sig := signal.NewSignal()
	return NewBolbol(str, sig)
}
