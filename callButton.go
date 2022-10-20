package main

// Button on a floor or basement to go back to lobby
type CallButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

func NewCallButton(_ID int, _floor int, _direction string) *CallButton {
	cb := CallButton{_ID, "idle", _floor, _direction}
	return &cb
}
