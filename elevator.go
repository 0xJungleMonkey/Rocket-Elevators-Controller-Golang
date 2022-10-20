package main

import "sort"

var elevatorID = 1

type Elevator struct {
	ID                    int
	status                string
	currentFloor          int
	completedRequestsList []int
	floorRequestsList     []int
	direction             string
}

func NewElevator(_ID int, _status string, _currentFloor int, _floorRequestsList []int) *Elevator {
	e := Elevator{_ID, "idle", _currentFloor, []int{}, _floorRequestsList, "stopped"}
	return &e
}

func (e *Elevator) move() {
	for len(e.floorRequestsList) > 0 {
		e.sortFloorList()
		destination := e.floorRequestsList[0]
		e.status = "moving"
		if e.currentFloor < destination {
			e.direction = "up"
			for e.currentFloor < destination {
				e.currentFloor++
			}
		} else if e.currentFloor > destination {
			e.direction = "down"
			for e.currentFloor > destination {
				e.currentFloor--
			}
		}
		e.status = "stopped"
		e.completedRequestsList = append(e.completedRequestsList, destination)
		e.floorRequestsList = e.floorRequestsList[1:]
	}
	e.status = "idle"
}

func (e *Elevator) sortFloorList() {
	if e.direction == "up" {
		sort.Sort(sort.IntSlice(e.floorRequestsList))
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(e.floorRequestsList)))
	}
}
func (e *Elevator) addNewRequest(_requestedFloor int) {
	if contains(e.floorRequestsList, _requestedFloor) == false {
		e.floorRequestsList = append(e.floorRequestsList, _requestedFloor)
		// Console.WriteLine(floorRequestsList[0] + "AddNewRequest");
	}
	if e.currentFloor < _requestedFloor {
		e.direction = "up"
	}
	if e.currentFloor > _requestedFloor {
		e.direction = "down"
	}
}

// func contains(s []int, e int) bool {
// 	for _, a := range s {
// 		if a == e {
// 			return true
// 		}
// 	}
// 	return false
// }
