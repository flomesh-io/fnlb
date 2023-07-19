package xlbnet

import (
	"sync"
	"time"
)

// DpEbpfH - context container
type DpEbpfH struct {
	ticker  *time.Ticker
	tDone   chan bool
	ctBcast chan bool
	tbN     uint
	CtSync  bool
	RssEn   bool
	ToMapCh chan interface{}
	ToFinCh chan int
	mtx     sync.RWMutex
	ctMap   map[string]*DpCtInfo
}
