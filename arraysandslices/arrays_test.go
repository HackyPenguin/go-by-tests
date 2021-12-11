package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {

	assertTrue := func(t testing.TB, got int, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d wanted %d given %v ", got, want, numbers)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertTrue(t, got, want, numbers)
	})

}

func TestSumAll(t *testing.T) {
	t.Run("test multiple slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func SumAllTails() {

}
