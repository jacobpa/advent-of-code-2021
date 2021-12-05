package solutions

import "testing"

func TestDayTwoEmptyInput(t *testing.T) {
	var input []string
	expected := 0
	actual, _ := DayTwo(input)
	if expected != actual {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestDayTwoForwardOnly(t *testing.T) {
	input := []string{"forward 10"}
	expected := 0
	actual, _ := DayTwo(input)
	if expected != actual {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestDayTwoForwardAndUp(t *testing.T) {
	input := []string{"forward 10", "up 10"}
	expected := -100
	actual, _ := DayTwo(input)
	if expected != actual {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestDayTwoInvalidDirection(t *testing.T) {
	input := []string{"forward 10", "back 10"}
	expected := 0
	actual, err := DayTwo(input)
	if expected != actual {
		t.Fatalf("Expected %v, got %v", expected, actual)
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