package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2,3) = %d; men ville ha %d", result, expected)
	}
}

/*
Same test for Add function only with Testify
*/

func TestAddWithTestify(t *testing.T){
	result := Add(2,5)
	assert.Equal(t, 7, result)
}



func TestSubtract(t *testing.T) {
	result := Subtract(10,4)
	expected := 6

	if result != expected {
		t.Errorf("Subtract(10,4) = %d; vill ha %d", result, expected)
	}
}

