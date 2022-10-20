package main

import "math"

type Column struct {
	ID               int
	status           string
	servedFloorsList []int
	elevatorsList    []Elevator
	callButtonList   []CallButton
	isBasement       bool
}

func NewColumn(_id, _amountOfFloors, _amountOfElevators int, _servedFloors []int, _isBasement bool) *Column {
	var c = Column{1, "Online", []int{}, []Elevator{}, []CallButton{}, true}

	c.ID = _id
	c.servedFloorsList = _servedFloors
	c.createElevators(_amountOfFloors, _amountOfElevators)
	c.createCallButtons(_amountOfFloors, _isBasement)
	return &c
}

func (c *Column) createCallButtons(_amountOfFloors int, _isBasement bool) {
	if _isBasement {
		callButtonID := 1
		buttonFloor := -1
		for i := 0; i < _amountOfFloors; i++ {
			callButton := NewCallButton(callButtonID, buttonFloor, "up")
			c.callButtonList = append(c.callButtonList, *callButton)
			buttonFloor++
			callButtonID++
		}
	} else {
		buttonFloor := 1
		callButtonID := 1
		for i := 0; i < _amountOfFloors; i++ {

			callButton := NewCallButton(callButtonID, buttonFloor, "down")
			c.callButtonList = append(c.callButtonList, *callButton)

			buttonFloor++
			callButtonID++
		}
	}
}
func (c *Column) createElevators(_amountOfFloors int, _amountOfElevators int) {
	elevatorID := 1
	for i := 0; i < _amountOfElevators; i++ {
		elevator := NewElevator(elevatorID, "idle", _amountOfFloors, []int{1})
		c.elevatorsList = append(c.elevatorsList, *elevator)
		elevatorID++
		// Console.WriteLine(elevatorsList.Count);
	}

}

// Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(userPosition int, direction string) *Elevator {
	elevator := c.findElevator(userPosition, direction)
	elevator.addNewRequest(userPosition)
	elevator.move()
	return elevator
}

func (c *Column) findElevator(_requestedFloor int, _requestedDirection string) *Elevator {
	BestElevatorInformations := []int{0, 6, 1000000}

	if _requestedFloor == 1 {
		for i := 0; i < len(c.elevatorsList); i++ {
			if c.elevatorsList[i].currentFloor == 1 && c.elevatorsList[i].status == "stopped" {
				BestElevatorInformations[0], BestElevatorInformations[1], BestElevatorInformations[2] = checkIfElevatorIsBetter(1, &c.elevatorsList[i], BestElevatorInformations[1], BestElevatorInformations[2], BestElevatorInformations[0], _requestedFloor)
			} else if c.elevatorsList[i].currentFloor == 1 && c.elevatorsList[i].status == "idle" {
				BestElevatorInformations[0], BestElevatorInformations[1], BestElevatorInformations[2] = checkIfElevatorIsBetter(2, &c.elevatorsList[i], BestElevatorInformations[1], BestElevatorInformations[2], BestElevatorInformations[0], _requestedFloor)
			} else if 1 > c.elevatorsList[i].currentFloor && c.elevatorsList[i].direction == "up" {
				BestElevatorInformations[0], BestElevatorInformations[1], BestElevatorInformations[2] = checkIfElevatorIsBetter(3, &c.elevatorsList[i], BestElevatorInformations[1], BestElevatorInformations[2], BestElevatorInformations[0], _requestedFloor)
			} else if 1 < c.elevatorsList[i].currentFloor && c.elevatorsList[i].direction == "down" { // Console.WriteLine("mewo");
				BestElevatorInformations[0], BestElevatorInformations[1], BestElevatorInformations[2] = checkIfElevatorIsBetter(3, &c.elevatorsList[i], BestElevatorInformations[1], BestElevatorInformations[2], BestElevatorInformations[0], _requestedFloor)
			} else if c.elevatorsList[i].status == "idle" {
				BestElevatorInformations[0], BestElevatorInformations[1], BestElevatorInformations[2] = checkIfElevatorIsBetter(4, &c.elevatorsList[i], BestElevatorInformations[1], BestElevatorInformations[2], BestElevatorInformations[0], _requestedFloor)
			} else {
				// Console.WriteLine("meow2");
				BestElevatorInformations[0], BestElevatorInformations[1], BestElevatorInformations[2] = checkIfElevatorIsBetter(5, &c.elevatorsList[i], BestElevatorInformations[1], BestElevatorInformations[2], BestElevatorInformations[0], _requestedFloor)
			}

			// Console.WriteLine("Wow" + bestElevator.ID);
			// Console.WriteLine(bestScore);
			// Console.WriteLine(referenceGap);
		}
	} else {
		for i := 0; i < len(c.elevatorsList); i++ {
			if _requestedFloor == c.elevatorsList[i].currentFloor && c.elevatorsList[i].status == "stopped" && _requestedDirection == c.elevatorsList[i].direction {
				BestElevatorInformations[0], BestElevatorInformations[1], BestElevatorInformations[2] = checkIfElevatorIsBetter(1, &c.elevatorsList[i], BestElevatorInformations[1], BestElevatorInformations[2], BestElevatorInformations[0], _requestedFloor)
			} else if _requestedFloor > c.elevatorsList[i].currentFloor && c.elevatorsList[i].direction == "up" && _requestedDirection == "up" {
				BestElevatorInformations[0], BestElevatorInformations[1], BestElevatorInformations[2] = checkIfElevatorIsBetter(2, &c.elevatorsList[i], BestElevatorInformations[1], BestElevatorInformations[2], BestElevatorInformations[0], _requestedFloor)
			} else if _requestedFloor < c.elevatorsList[i].currentFloor && c.elevatorsList[i].direction == "down" && _requestedDirection == "down" {
				BestElevatorInformations[0], BestElevatorInformations[1], BestElevatorInformations[2] = checkIfElevatorIsBetter(2, &c.elevatorsList[i], BestElevatorInformations[1], BestElevatorInformations[2], BestElevatorInformations[0], _requestedFloor)
			} else if c.elevatorsList[i].status == "idle" {
				BestElevatorInformations[0], BestElevatorInformations[1], BestElevatorInformations[2] = checkIfElevatorIsBetter(4, &c.elevatorsList[i], BestElevatorInformations[1], BestElevatorInformations[2], BestElevatorInformations[0], _requestedFloor)
			} else {
				BestElevatorInformations[0], BestElevatorInformations[1], BestElevatorInformations[2] = checkIfElevatorIsBetter(5, &c.elevatorsList[i], BestElevatorInformations[1], BestElevatorInformations[2], BestElevatorInformations[0], _requestedFloor)
			}

		}
	}

	return &c.elevatorsList[BestElevatorInformations[0]-1]
}
func checkIfElevatorIsBetter(scoreToCheck int, newElevator *Elevator, bestScore int, referenceGap int, bestElevator int, floor int) (int, int, int) {
	if scoreToCheck < bestScore {
		bestScore = scoreToCheck
		bestElevator = newElevator.ID
		referenceGap = int(math.Abs(float64(newElevator.currentFloor - floor)))
		// Console.WriteLine(referenceGap);
	} else if bestScore == scoreToCheck {
		gap := int(math.Abs(float64(newElevator.currentFloor - floor)))
		if referenceGap > gap {
			bestElevator = newElevator.ID
			referenceGap = gap
		}
	}

	return bestElevator, bestScore, referenceGap
}
