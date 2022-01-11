package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	e := &Elevator{
		floor:    5,
		nowFloor: 3,
	}
	requestFloors = []int{}
	direction := e.getDirection(requestFloors) // 0表示无人请求电梯
	if direction != "不动" {
		t.Fatalf("预期电梯不动，但得到的却是：%s", direction)
	}
}

func TestCase2(t *testing.T) {
	e := &Elevator{
		floor:    5,
		nowFloor: 1,
	}
	requestFloors = append(requestFloors, 3)
	direction := e.getDirection(requestFloors)
	if direction != "向上" {
		t.Fatalf("预期电梯向上，但得到的却是：%s", direction)
	}
	e.elevatorMove(e.direction, e.nowFloor)
	if e.nowFloor != 3 {
		t.Fatalf("预期电梯停在三楼，但得到的却是：%d", e.nowFloor)
	}

}

func TestCase3(t *testing.T) {
	e := &Elevator{
		floor:    5,
		nowFloor: 3,
	}
	requestFloors = append(requestFloors, 4, 2)
	direction := e.getDirection(requestFloors)
	if direction != "向上" {
		t.Fatalf("预期电梯向上，但得到的却是：%s", direction)
	}
	e.elevatorMove(e.direction, e.nowFloor)
	if e.nowFloor != 4 {
		t.Fatalf("预期电梯停在四楼，但得到的却是：%d", e.nowFloor)
	}

	direction = e.getDirection(requestFloors)
	if direction != "向下" {
		t.Fatalf("预期电梯向下，但得到的却是：%s", direction)
	}
	e.elevatorMove(e.direction, e.nowFloor)
	if e.nowFloor != 2 {
		t.Fatalf("预期电梯停在四楼，但得到的却是：%d", e.nowFloor)
	}

}
func TestCase4(t *testing.T) {
	e := &Elevator{
		floor:    5,
		nowFloor: 3,
	}
	requestFloors = append(requestFloors, 4, 5, 2)
	direction := e.getDirection(requestFloors)
	if direction != "向上" {
		t.Fatalf("预期电梯向上，但得到的却是：%s", direction)
	}
	e.elevatorMove(e.direction, e.nowFloor)

}
