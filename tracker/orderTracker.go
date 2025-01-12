package tracker

import (
	"awesomeProject2/src/Pendency_System_Go"
	"awesomeProject2/src/Pendency_System_Go/nodeParser"
	"strconv"
)

type OrderParser struct {
	nodeParser Pendency_System_Go.NodeParser
}

func New() Pendency_System_Go.Tracker {
	o := OrderParser{}
	o.nodeParser = nodeParser.New()
	return &o
}

/*
	StartTracking(orderID int, hierarchy []string)
	GetCount(hierarchy []string) int
	StopTracking(orderID int)
*/

func (o *OrderParser) StartTracking(orderId int, hierarchy []string) {
	o.nodeParser.AddNode(hierarchy, strconv.Itoa(orderId))
}

func (o *OrderParser) GetCount(hierarchy []string) int {
	return o.nodeParser.GetCount(hierarchy)
}

func (o *OrderParser) StopTracking(orderId int) {
	o.nodeParser.DeleteNode(strconv.Itoa(orderId))
	return
}
