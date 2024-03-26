package wakling

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
				City string
				Age  int
			}{"Nick", "Tirane", 23},
			[]string{"Nick", "Tirane"},
		},
		{
			"nested fields",
			&Person{
				"Nick",
				[]Profile{
					{31, "Tirane"},
					{31, "Durres"},
				},
			},
			[]string{"Nick", "Tirane", "Durres"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "papa"},
				{22, "Apa"},
			}, []string{"papa", "Apa"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"papa": "papa",
			"apa":  "apa",
			"aa":   "a",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "papa")
		assertContains(t, got, "apa")
		assertContains(t, got, "a")
	})

	t.Run("with channels", func(t *testing.T) {
		channel1 := make(chan Profile)

		go func() {
			channel1 <- Profile{33, "a"}
			channel1 <- Profile{44, "b"}
			close(channel1)
		}()
		var got []string
		want := []string{"a", "b"}
		walk(channel1, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		function := func() (Profile, Profile) {
			return Profile{22, "a"}, Profile{3, "b"}
		}

		var got []string
		want := []string{"a", "b"}
		walk(function, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, hystack []string, neddle string) {
	t.Helper()
	contains := false
	for _, x := range hystack {
		if x == neddle {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it did'nt", hystack, neddle)
	}
}
