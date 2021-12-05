package solutions

import "testing"

func TestDayTwoEmptyInput(t *testing.T) {
	var input []string
	expectedA, expectedB := 0, 0
	actualA, actualB, _ := DayTwo(input)
	if expectedA != actualA {
		t.Fatalf("Part A: Expected %v, got %v", expectedA, actualA)
	}
	if expectedB != actualB {
		t.Fatalf("Part B: Expected %v, got %v", expectedB, actualB)
	}
}

func TestDayTwoForwardOnly(t *testing.T) {
	input := []string{"forward 10"}
	expectedA, expectedB := 0, 0
	actualA, actualB, _ := DayTwo(input)
	if expectedA != actualA {
		t.Fatalf("Part A: Expected %v, got %v", expectedA, actualA)
	}
	if expectedB != actualB {
		t.Fatalf("Part B: Expected %v, got %v", expectedB, actualB)
	}
}

func TestDayTwoForwardAndUp(t *testing.T) {
	input := []string{"forward 10", "up 10"}
	expectedA, expectedB := -100, 0
	actualA, actualB, _ := DayTwo(input)
	if expectedA != actualA {
		t.Fatalf("Part A: Expected %v, got %v", expectedA, actualA)
	}
	if expectedB != actualB {
		t.Fatalf("Part B: Expected %v, got %v", expectedB, actualB)
	}
}

func TestDayTwoInvalidDirection(t *testing.T) {
	input := []string{"forward 10", "back 10"}
	expectedA, expectedB := 0, 0
	actualA, actualB, err := DayTwo(input)
	if expectedA != actualA {
		t.Fatalf("Part A: Expected %v, got %v", expectedA, actualA)
	}
	if expectedB != actualB {
		t.Fatalf("Part B: Expected %v, got %v", expectedB, actualB)
	}
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func BenchmarkDayTwo(b *testing.B) {
	input := []string{"forward 10", "up 10", "down 10", "up 5", "forward 1"}
	for n := 0; n < b.N; n++ {
		DayTwo(input)
	}
}