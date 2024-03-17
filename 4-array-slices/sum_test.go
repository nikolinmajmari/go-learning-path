package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	/*	t.Run("collection of five numbers", func(t *testing.T) {
			numbers := []int{1, 2, 3, 4, 5}

			got := Sum(numbers)
			want := 15

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	*/
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{3, 7})
	want := []int{6, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 3, 5, 7}, []int{2, 4, 6})
		want := []int{15, 10}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 2, 2})
		want := []int{0, 4}
		checkSums(t, got, want)
	})
}
