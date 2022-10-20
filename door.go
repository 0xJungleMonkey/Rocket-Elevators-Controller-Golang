package main

type Door struct {
	ID     int
	status string
}

func NewDoor(_id int) *Door {
	d := Door{_id, "closed"}
	return &d
}
