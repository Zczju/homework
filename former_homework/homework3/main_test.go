package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	e := &Elevator{
		maxFloor: 5,
		nowFloor: 3,
	}

	if e.direction != "" {
		t.Fatalf("预期电梯不动，但得到的却是：%s", e.direction)
	}
}

func TestCase2(t *testing.T) {
	e := &Elevator{
		maxFloor: 5,
		nowFloor: 1,
	}
	if err := e.request(3); err != nil {
		t.Fatalf("捕获错误：%v", err)
	}
	if e.direction != "向上" {
		t.Fatalf("预期电梯向上，但得到的却是：%s", e.direction)
	}
	e.elevatorMove()
	if e.nowFloor != 3 {
		t.Fatalf("预期电梯停在3楼，但得到的却是：%d", e.nowFloor)
	}
}

func TestCase3(t *testing.T) {
	e := &Elevator{
		maxFloor: 5,
		nowFloor: 3,
	}
	if err := e.request(4); err != nil {
		t.Fatalf("捕获错误：%v", err)
	}

	if err := e.request(2); err != nil {
		t.Fatalf("捕获错误：%v", err)
	}
	if e.direction != "向上" {
		t.Fatalf("预期电梯向上，但得到的却是：%s", e.direction)
	}
	e.elevatorMove()
	if e.nowFloor != 4 {
		t.Fatalf("预期电梯停在4楼，但得到的却是：%d", e.nowFloor)
	}
	e.elevatorMove()
	if e.nowFloor != 2 {
		t.Fatalf("预期电梯停在2楼，但得到的却是：%d", e.nowFloor)
	}
}
func TestCase4(t *testing.T) {
	e := &Elevator{
		maxFloor: 5,
		nowFloor: 3,
	}
	if err := e.request(4); err != nil {
		t.Fatalf("捕获错误：%v", err)
	}
	if err := e.request(5); err != nil {
		t.Fatalf("捕获错误：%v", err)
	}
	if err := e.request(2); err != nil {
		t.Fatalf("捕获错误：%v", err)
	}
	if e.direction != "向上" {
		t.Fatalf("预期电梯向上，但得到的却是：%s", e.direction)
	}
	e.elevatorMove()
	if e.nowFloor != 4 {
		t.Fatalf("预期电梯停在4楼，但得到的却是：%d", e.nowFloor)
	}
	e.elevatorMove()
	if e.nowFloor != 5 {
		t.Fatalf("预期电梯停在5楼，但得到的却是：%d", e.nowFloor)
	}
	e.elevatorMove()
	if e.nowFloor != 2 {
		t.Fatalf("预期电梯停在2楼，但得到的却是：%d", e.nowFloor)
	}
}
