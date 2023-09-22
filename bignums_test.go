package bignums

import (
	"math/big"
	"testing"
)

func TestBigIntChain_Add(t *testing.T) {
	chain := NewBigIntChain(10).Add(20)
	val, err := chain.Value()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := big.NewInt(30)
	if val.Cmp(expected) != 0 {
		t.Fatalf("Expected %v, got %v", expected, val)
	}
}

func TestBigIntChain_Subtract(t *testing.T) {
	chain := NewBigIntChain(30).Subtract(20)
	val, err := chain.Value()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := big.NewInt(10)
	if val.Cmp(expected) != 0 {
		t.Fatalf("Expected %v, got %v", expected, val)
	}
}

func TestBigIntChain_Multiply(t *testing.T) {
	chain := NewBigIntChain(5).Multiply(4)
	val, err := chain.Value()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := big.NewInt(20)
	if val.Cmp(expected) != 0 {
		t.Fatalf("Expected %v, got %v", expected, val)
	}
}

func TestBigIntChain_Divide(t *testing.T) {
	chain := NewBigIntChain(20).Divide(4)
	val, err := chain.Value()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := big.NewInt(5)
	if val.Cmp(expected) != 0 {
		t.Fatalf("Expected %v, got %v", expected, val)
	}
}

func TestBigIntChain_DivideByZero(t *testing.T) {
	chain := NewBigIntChain(20).Divide(0)
	_, err := chain.Value()

	if err == nil || err.Error() != "division by zero" {
		t.Fatalf("Expected division by zero error, got %v", err)
	}
}

func TestBigIntChain_Brackets(t *testing.T) {
	chain := NewBigIntChain(10).Begin().Add(10).End().Multiply(2)
	val, err := chain.Value()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := big.NewInt(40)
	if val.Cmp(expected) != 0 {
		t.Fatalf("Expected %v, got %v", expected, val)
	}
}

func TestBigIntChain_MismatchedBrackets(t *testing.T) {
	chain := NewBigIntChain(10).Begin().Add(10)
	_, err := chain.Value()

	if err == nil || err.Error() != "mismatched brackets" {
		t.Fatalf("Expected mismatched brackets error, got %v", err)
	}
}

func TestBigFloatChain_BasicOperations(t *testing.T) {
	chain := NewBigFloatChain(10.5).Add(20.5).Subtract(10).Multiply(2).Divide(2)
	val, err := chain.Value()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := big.NewFloat(21)
	if val.Cmp(expected) != 0 {
		t.Fatalf("Expected %v, got %v", expected, val)
	}
}

func TestConvertToBigInt(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected *big.Int
		err      bool
	}{
		{input: "0x10", expected: big.NewInt(16), err: false},
		{input: 10, expected: big.NewInt(10), err: false},
		{input: "invalid", expected: nil, err: true},
	}

	for _, test := range tests {
		val, err := convertToBigInt(test.input)
		if (err != nil) != test.err {
			t.Fatalf("Expected error: %v, got: %v", test.err, err)
		}
		if err == nil && val.Cmp(test.expected) != 0 {
			t.Fatalf("Expected %v, got %v", test.expected, val)
		}
	}
}

func TestBigFloatChain_Add(t *testing.T) {
	chain := NewBigFloatChain(10.5).Add(20.5)
	val, err := chain.Value()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := big.NewFloat(31)
	if val.Cmp(expected) != 0 {
		t.Fatalf("Expected %v, got %v", expected, val)
	}
}

func TestBigFloatChain_Subtract(t *testing.T) {
	chain := NewBigFloatChain(30.5).Subtract(20.5)
	val, err := chain.Value()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := big.NewFloat(10)
	if val.Cmp(expected) != 0 {
		t.Fatalf("Expected %v, got %v", expected, val)
	}
}

func TestBigFloatChain_Multiply(t *testing.T) {
	chain := NewBigFloatChain(5.5).Multiply(4)
	val, err := chain.Value()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := big.NewFloat(22)
	if val.Cmp(expected) != 0 {
		t.Fatalf("Expected %v, got %v", expected, val)
	}
}

func TestBigFloatChain_Divide(t *testing.T) {
	chain := NewBigFloatChain(22).Divide(4)
	val, err := chain.Value()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := big.NewFloat(5.5)
	if val.Cmp(expected) != 0 {
		t.Fatalf("Expected %v, got %v", expected, val)
	}
}

func TestBigFloatChain_DivideByZero(t *testing.T) {
	chain := NewBigFloatChain(20).Divide(0)
	_, err := chain.Value()

	if err == nil || err.Error() != "division by zero" {
		t.Fatalf("Expected division by zero error, got %v", err)
	}
}

func TestBigFloatChain_Brackets(t *testing.T) {
	chain := NewBigFloatChain(10).Begin().Add(10).End().Multiply(2)
	val, err := chain.Value()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := big.NewFloat(40)
	if val.Cmp(expected) != 0 {
		t.Fatalf("Expected %v, got %v", expected, val)
	}
}

func TestBigFloatChain_MismatchedBrackets(t *testing.T) {
	chain := NewBigFloatChain(10).Begin().Add(10)
	_, err := chain.Value()

	if err == nil || err.Error() != "mismatched brackets" {
		t.Fatalf("Expected mismatched brackets error, got %v", err)
	}
}

func TestConvertToBigFloat(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected *big.Float
		err      bool
	}{
		{input: "10.5", expected: big.NewFloat(10.5), err: false},
		{input: 10, expected: big.NewFloat(10), err: false},
		{input: "invalid", expected: nil, err: true},
	}

	for _, test := range tests {
		val, err := convertToBigFloat(test.input)
		if (err != nil) != test.err {
			t.Fatalf("Expected error: %v, got: %v", test.err, err)
		}
		if err == nil && val.Cmp(test.expected) != 0 {
			t.Fatalf("Expected %v, got %v", test.expected, val)
		}
	}
}
