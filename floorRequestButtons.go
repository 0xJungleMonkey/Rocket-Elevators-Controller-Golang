package main

var floorRequestsButton = 1
var floor int

// FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

func NewFloorRequestButton(_id int, _floor int, _direction string) *FloorRequestButton {
	fb := FloorRequestButton{_id, "idle", _floor, _direction}
	return &fb
}
