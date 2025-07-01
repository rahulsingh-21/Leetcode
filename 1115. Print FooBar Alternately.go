package main

type FooBar struct {
	n  int
	ch chan int
}

func NewFooBar(n int) *FooBar {
	ch := make(chan int)
	return &FooBar{n, ch}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.ch <- 1
		<-fb.ch
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		// printBar() outputs "bar". Do not change or remove this line.
		<-fb.ch
		printBar()
		fb.ch <- 1
	}
}
