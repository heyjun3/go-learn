package sequencing

type Frame struct {
	pixles []int
}

func NewFrame(size int) *Frame {
	return &Frame{
		pixles: make([]int, size),
	}
}

func (f *Frame) Clear() {
	for i, _ := range f.pixles {
		f.pixles[i] = 0
	}
}

func (f *Frame) Draw(x int) {
	f.pixles[x] = 1
}

func (f Frame) GetPixels() []int {
	return f.pixles
}
