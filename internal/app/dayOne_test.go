package app

import "testing"

func TestDayOneEmptyInput(t *testing.T) {
	var input []string
	expected := 0
	actual := DayOne(input)
	if expected != actual {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestDayOneInputLengthOne(t *testing.T) {
	input := []string{"123"}
	expected := 0
	actual := DayOne(input)
	if expected != actual {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestDayOneInputLengthTwo(t *testing.T) {
	input := []string{"10", "20"}
	expected := 1
	actual := DayOne(input)
	if expected != actual {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestDayOneInputLengthTwoNoIncrease(t *testing.T) {
	input := []string{"10", "5"}
	expected := 0
	actual := DayOne(input)
	if expected != actual {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestDayOneInputLengthFive(t *testing.T) {
	input := []string{"10", "20", "15", "5", "50"}
	expected := 2
	actual := DayOne(input)
	if expected != actual {
		t.Fatalf("Expected %v, got %v", expected, actual)
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
