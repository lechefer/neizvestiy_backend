package entity

type Image struct {
	Sizes Sizes
}

type Sizes struct {
	M ImageSize
	X ImageSize
	O ImageSize
}

type ImageSize struct {
	Size Size
	Url  string
}

type Size string

const (
	X Size = "x"
	M Size = "m"
	O Size = "o"
)
