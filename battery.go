package main

import "math"

type Battery struct {
	ID                      int
	status                  string
	columnsList             []Column
	floorRequestButtonsList []FloorRequestButton
}

var floorRequestButtonID = 1
var batt = Battery{1, "Online", []Column{}, []FloorRequestButton{}}

func NewBattery(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn int) *Battery {
	batt.ID = _id
	if _amountOfBasements > 0 {
		batt.createBasementFloorRequestButtons(_amountOfBasements)
		batt.createBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn)
		_amountOfColumns--
	}
	batt.createFloorRequestButtons(_amountOfFloors)
	batt.createColumns(_amountOfColumns, _amountOfFloors, _amountOfElevatorPerColumn)
	return &batt
}

var columnID int = 1

func (b *Battery) createBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn int) {
	var servedFloors []int
	floor := -1
	for i := 0; i < _amountOfBasements; i++ {
		servedFloors = append(servedFloors, floor)
		floor -= 1
	}
	servedFloors = append(servedFloors, 1)
	column := NewColumn(columnID, _amountOfBasements, _amountOfElevatorPerColumn, servedFloors, true)

	b.columnsList = append(b.columnsList, *column)
	columnID += 1
}
func (b *Battery) createColumns(_amountOfColumns, _amountOfFloors, _amountOfElevatorPerColumn int) {
	amountOfFloorsPerColumn := int(math.Ceil(float64(_amountOfFloors / _amountOfColumns)))
	floor := 1
	for i := 0; i < _amountOfColumns; i++ {
		var servedFloors []int
		for n := 0; n < amountOfFloorsPerColumn; n++ {
			if floor <= _amountOfFloors {
				servedFloors = append(servedFloors, floor)
				floor += 1
			}
		}
		servedFloors = append(servedFloors, 1)
		// Console.WriteLine(servedFloors.Count);
		column := NewColumn(columnID, _amountOfFloors, _amountOfElevatorPerColumn, servedFloors, false)
		b.columnsList = append(b.columnsList, *column)
		columnID += 1
		// Console.WriteLine(columnsList.Count);
	}
}

var floorRequestsButtonID int = 1

func (b *Battery) createBasementFloorRequestButtons(_amountOfBasements int) {
	buttonFloor := -1
	for i := 0; i < _amountOfBasements; i++ {
		floorRequestButton := NewFloorRequestButton(floorRequestButtonID, buttonFloor, "down")
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, *floorRequestButton)
		buttonFloor -= 1
		floorRequestButtonID += 1
	}
}
func (b *Battery) createFloorRequestButtons(_amountOfFloors int) {
	buttonFloor := 1
	for i := 0; i < _amountOfFloors; i++ {
		floorRequestButton := NewFloorRequestButton(floorRequestButtonID, buttonFloor, "up")
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, *floorRequestButton)
		buttonFloor += 1
		floorRequestButtonID += 1
	}
}
func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	var returnColumn *Column = nil
	for _, column := range b.columnsList {
		if contains(column.servedFloorsList, _requestedFloor) {
			returnColumn = &column
			return returnColumn
		}
	}
	return returnColumn
}

// Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	var column = batt.findBestColumn(_requestedFloor)
	var elevator = batt.findBestColumn(_requestedFloor).findElevator(1, _direction)

	elevator.addNewRequest(1)
	elevator.move()
	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	return column, elevator
}
