package model

// R is red
type R uint8

// G is green
type G uint8

// B is blue
type B uint8

// A is alpha
type A uint8

// RGBA is color elements
type RGBA struct {
	R
	G
	B
	A
}

//Colors are a list of rgba color
type Colors []RGBA
