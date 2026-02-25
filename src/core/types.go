package core

type Cell struct {
	X, Y int
}

const (
	Empty = iota
	Conductor
	ElectronHead
	ElectronTail
)
