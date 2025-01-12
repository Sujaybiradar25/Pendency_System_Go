package Pendency_System_Go

type Tracker interface {
	StartTracking(orderID int, hierarchy []string)
	GetCount(hierarchy []string) int
	StopTracking(orderID int)
}
