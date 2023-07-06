package main

func main() {
}


type GetList[T Stringer] func() (T, error)

type Stringer interface{}

type Product struct {
	GetList[Str]
}

func NewProduct() Product{
	return Product{
		GetList: func() (Str, error) { return Str{}, nil},
	}
}

type Str struct {}