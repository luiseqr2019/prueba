package stack

type Sizer[T comparable] interface {
	Size() int64
}

type Finder[T comparable] interface {
	Find(value T) bool
}

type Stacker[T comparable] interface {
	Sizer[T]
	Finder[T]
	Pop() (T, error)
	Push(value T)
	Top() (T, error)
	Button() (T, error)
}
