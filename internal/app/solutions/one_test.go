package solutions

import "testing"

func TestDayOneEmptyInput(t *testing.T) {
	var input []string
	expectedA, expectedB := 0, 0
	actualA, actualB := DayOne(input)
	if expectedA != actualA {
		t.Fatalf("Part A: Expected %v, got %v", expectedA, actualA)
	}
	if expectedB != actualB {
		t.Fatalf("Part B: Expected %v, got %v", expectedB, actualB)
	}
}

func TestDayOneInputLengthOne(t *testing.T) {
	input := []string{"123"}
	expectedA, expectedB := 0, 0
	actualA, actualB := DayOne(input)
	if expectedA != actualA {
		t.Fatalf("Part A: Expected %v, got %v", expectedA, actualA)
	}
	if expectedB != actualB {
		t.Fatalf("Part B: Expected %v, got %v", expectedB, actualB)
	}
}

func TestDayOneInputLengthTwo(t *testing.T) {
	input := []string{"10", "20"}
	expectedA, expectedB := 1, 0
	actualA, actualB := DayOne(input)
	if expectedA != actualA {
		t.Fatalf("Part A: Expected %v, got %v", expectedA, actualA)
	}
	if expectedB != actualB {
		t.Fatalf("Part B: Expected %v, got %v", expectedB, actualB)
	}
}

func TestDayOneInputLengthTwoNoIncrease(t *testing.T) {
	input := []string{"10", "5"}
	expectedA, expectedB := 0, 0
	actualA, actualB := DayOne(input)
	if expectedA != actualA {
		t.Fatalf("Part A: Expected %v, got %v", expectedA, actualA)
	}
	if expectedB != actualB {
		t.Fatalf("Part B: Expected %v, got %v", expectedB, actualB)
	}
}

func TestDayOneInputLengthFive(t *testing.T) {
	input := []string{"10", "20", "15", "5", "50"}
	expectedA, expectedB := 2, 1
	actualA, actualB := DayOne(input)
	if expectedA != actualA {
		t.Fatalf("Part A: Expected %v, got %v", expectedA, actualA)
	}
	if expectedB != actualB {
		t.Fatalf("Part B: Expected %v, got %v", expectedB, actualB)
	}
}

func TestDayOneInputLengthSix(t *testing.T) {
	input := []string{"10", "20", "15", "5", "50", "40"}
	// 45, 40, 70, 95
	expectedA, expectedB := 2, 2
	actualA, actualB := DayOne(input)
	if expectedA != actualA {
		t.Fatalf("Part A: Expected %v, got %v", expectedA, actualA)
	}
	if expectedB != actualB {
		t.Fatalf("Part B: Expected %v, got %v", expectedB, actualB)
	}
}

func BenchmarkDayOne0(b *testing.B) {
	var input []string

	for n :=0; n < b.N; n++ {
		DayOne(input)
	}
}

func BenchmarkDayOne2(b *testing.B) {
	input := []string{"10", "20"}

	for n :=0; n < b.N; n++ {
		DayOne(input)
	}
}

func BenchmarkDayOne5(b *testing.B) {
	input := []string{"10", "20", "15", "5", "50"}

	for n :=0; n < b.N; n++ {
		DayOne(input)
	}
}

func BenchmarkDayOne10(b *testing.B) {
	input := []string{"10", "20", "15", "5", "50", "100", "20", "1", "55", "3000"}
	for n :=0; n < b.N; n++ {
		DayOne(input)
	}
}
