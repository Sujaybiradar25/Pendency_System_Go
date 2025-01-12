package Pendency_System_Go

type NodeParser interface {
	AddNode(path []string, value string)
	DeleteNode(value string)
	GetCount(path []string) int
}
