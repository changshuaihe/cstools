package cstools

type Thread struct {
	count int      `json:"count"`
	ch    chan int `json:"ch"`
}

var Threads map[string]Thread

func (t *Thread) Init(chCount int) {
	t.ch = make(chan int, chCount)
}

func (t *Thread) Start() {
	if t.ch == nil {
		t.Init(1)
	}
	t.ch <- 1
}

func (t *Thread) End() {
	<-t.ch
}
