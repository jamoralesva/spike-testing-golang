package main

import "testing"

func TestAdd(t *testing.T) {

	t.Run("empty value", func(t *testing.T) {
		//Arrange
		expected := 0
		//Act
		outcome := Add("")
		//Assert
		if outcome != expected {
			t.Errorf("Test with empty value failed")
		}
	})

	t.Run("one value", func(t *testing.T) {
		//Arrange
		expected := 2
		//Act
		outcome := Add("2")
		//Assert
		if outcome != expected {
			t.Errorf("Test with one value failed")
		}
	})

	t.Run("two values", func(t *testing.T) {
		//Arrange
		expected := 6
		//Act
		outcome := Add("4,2")
		//Assert
		if outcome != expected {
			t.Errorf("Test with one value failed")
		}
	})

	t.Run("unknown values", func(t *testing.T) {
		//Arrange
		expected := 30
		//Act
		outcome := Add("1,6,7,9,3,4")
		//Assert
		if outcome != expected {
			t.Errorf("Test with unknown value failed")
		}
	})
}
