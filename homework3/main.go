package main

import "time"

var requestFloors []int

type Elevator struct {
	floor     int
	nowFloor  int
	direction string
}

func (e *Elevator) getDirection(requestFloors []int) string {
	if len(requestFloors) == 0 {
		e.direction = "不动"
		return e.direction
	} else {
		if requestFloors[0] > e.nowFloor {
			e.direction = "向上"
			return e.direction
		} else if requestFloors[0] < e.nowFloor {
			e.direction = "向下"
			return e.direction
		}

	}

	return ""
}

func (e *Elevator) elevatorMove(direction string, nowFloor int) {
	if direction == "向上" {
		for {
			time.Sleep(1 * time.Second)
			nowFloor++
			if nowFloor == requestFloors[0] {
				requestFloors = requestFloors[1:]
				return
			}
		}

	}
	if direction == "向下" {
		for {
			time.Sleep(1 * time.Second)
			nowFloor--
			if nowFloor == requestFloors[0] {
				requestFloors = requestFloors[1:]
				return
			}
		}

	}
	return
}
