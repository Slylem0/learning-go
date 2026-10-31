package main

import (
	"fmt"
)

type Clock struct {
	hour    int
	minutes int
}

func New(h, m int) Clock {
	minTotales := h*60 + m
	if h < 0 || m < 0 {
	}
	mTotales := minTotales % 1440

	htotal := mTotales / 60
	mtotal := mTotales % 60

	return Clock{htotal, mtotal}
}

func main() {
	fmt.Println(New(2, 50))
}
