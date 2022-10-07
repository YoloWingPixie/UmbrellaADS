package iads

import (
	"umbrella/internal/network"
)

type IADS struct {
	Name    string
	Network network.Network
}

func Run() {

}

func NewIads(name string) *IADS {
	var iads *IADS
	iads.Name = name

	return iads
}
