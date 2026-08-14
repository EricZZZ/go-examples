package main

func main() {
	n := 10
	println(read(&n))
	println(newInt())
}

func read(p *int)(v int) {
	v = *p
	return
}

func newInt()(p *int) {
	var n int
	return &n
}
