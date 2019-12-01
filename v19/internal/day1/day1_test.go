package day1

import "testing"

func TestFuel(t *testing.T) {
	cases := []struct {
		mass, expectedFuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, c := range cases {
		actualFuel := calculateFuelForMass(c.mass)
		if actualFuel != c.expectedFuel {
			t.Errorf("calculateFuelForMass(%d) == %d, expected %d", c.mass, actualFuel, c.expectedFuel)
		}
	}
}

func TestFuelForFuel(t *testing.T) {
	cases := []struct {
		mass, expectedFuel int
	}{
		{2, 0},
		{654, 312},
		{33583, 16763},
	}
	for _, c := range cases {
		actualFuel := calculateFuelForFuel(c.mass)
		if actualFuel != c.expectedFuel {
			t.Errorf("calculateFuelForMass(%d) == %d, expected %d", c.mass, actualFuel, c.expectedFuel)
		}
	}
}

func TestFuelForMassAndFuel(t *testing.T) {
	cases := []struct {
		mass, expectedTotalFuel int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, c := range cases {
		actualTotalFuel := calculateFuelForMassAndFuel(c.mass)
		if actualTotalFuel != c.expectedTotalFuel {
			t.Errorf("calculateFuelForMass(%d) == %d, expected %d", c.mass, actualTotalFuel, c.expectedTotalFuel)
		}
	}
}