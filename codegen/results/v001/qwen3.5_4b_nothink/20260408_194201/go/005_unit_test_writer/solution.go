package main

import (
	"fmt"
	"testing"
)

func calculateDiscount(price float64, quantity int) float64 {
	if quantity < 10 {
		return price * float64(quantity)
	} else if quantity < 50 {
		return price * float64(quantity) * 0.9
	}
	return price * float64(quantity) * 0.75
}

func TestDiscount_Threshold1_BelowTen(t *testing.T) {
	price := 10.0
	quantity := 9
	expected := 90.0
	result := calculateDiscount(price, quantity)
	
	if result != expected {
		t.Errorf("Quantity below 10: expected %v, got %v", expected, result)
	}
}

func TestDiscount_Threshold2_AtTen(t *testing.T) {
	price := 10.0
	quantity := 10
	expected := 90.0
	result := calculateDiscount(price, quantity)
	
	if result != expected {
		t.Errorf("Quantity at 10: expected %v, got %v", expected, result)
	}
}

func TestDiscount_Threshold3_BelowFifty(t *testing.T) {
	price := 10.0
	quantity := 49
	expected := 441.0
	result := calculateDiscount(price, quantity)
	
	if result != expected {
		t.Errorf("Quantity below 50: expected %v, got %v", expected, result)
	}
}

func TestDiscount_Threshold4_AtFifty(t *testing.T) {
	price := 10.0
	quantity := 50
	expected := 375.0
	result := calculateDiscount(price, quantity)
	
	if result != expected {
		t.Errorf("Quantity at 50: expected %v, got %v", expected, result)
	}
}

func main() {
	t := &testing.T{}
	
	tests := []struct {
		name     string
		fn       func(t *testing.T)
	}{
		{"Quantity below 10", TestDiscount_Threshold1_BelowTen},
		{"Quantity at 10", TestDiscount_Threshold2_AtTen},
		{"Quantity below 50", TestDiscount_Threshold3_BelowFifty},
		{"Quantity at 50", TestDiscount_Threshold4_AtFifty},
	}

	for _, tt := range tests {
		tt.fn(t)
		if t.Failed() {
			fmt.Printf("FAIL: %s\n", tt.name)
			// Print expected and got values directly if assertion failed
			// This part is simulated since we can't access t.Failed() context directly without a wrapper
		} else {
			fmt.Printf("PASS: %s\n", tt.name)
		}
	}
}