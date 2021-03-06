package g

// one Go per goroutine (goroutine not safe)
type Go struct {
	ChanCb    chan func()
	pendingGo int
}

func New(l int) *Go {
	g := new(Go)
	g.ChanCb = make(chan func(), l)
	return g
}

func (g *Go) Go(f func(), cb func()) {
	g.pendingGo++

	go func() {
		f()
		g.ChanCb <- cb
	}()
}

func (g *Go) Cb(cb func()) {
	if cb != nil {
		cb()
	}

	g.pendingGo--
}

func (g *Go) Close() {
	for g.pendingGo > 0 {
		g.Cb(<-g.ChanCb)
	}
}
