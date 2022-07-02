package _3_01_counter

import (
	"sync"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestCounter_Up(t *testing.T) {
	testingUpSequential(t)
	testingUpConcurrent(t)
}

func testingUpSequential(t *testing.T) {
	testCases := []struct {
		label     string
		numOfCall int
		want      int
	}{
		{
			label:     "add_1",
			numOfCall: 1,
			want:      1,
		},
		{
			label:     "add_2",
			numOfCall: 2,
			want:      2,
		},
	}

	t.Log("Given the need to add a counter sequentially")
	for i, tc := range testCases {
		t.Logf("\tTest %d:\t%s\n", i, tc.label)
		tf := func(t *testing.T) {
			counter := NewCounter()

			for i := 0; i < tc.numOfCall; i++ {
				counter.Up()
			}

			if got := counter.GetCount(); got != tc.want {
				t.Fatalf("\t%s\tShould get a counter value is %d: %d", failed, tc.want, got)
			}
			t.Logf("\t%s\tShould get a counter value is %d", succeed, tc.want)
		}

		t.Run(tc.label, tf)
	}
}

func testingUpConcurrent(t *testing.T) {
	testCases := []struct {
		label     string
		numOfCall int
		want      int
	}{
		{
			label:     "add_1",
			numOfCall: 1,
			want:      1,
		},
		{
			label:     "add_2",
			numOfCall: 2,
			want:      2,
		},
	}

	t.Log("Given the need to add a counter concurrently")
	for i, tc := range testCases {
		t.Logf("\tTest %d:\t%s\n", i, tc.label)
		tf := func(t *testing.T) {
			counter := NewCounter()

			wg := sync.WaitGroup{}
			wg.Add(tc.numOfCall)

			for i := 0; i < tc.numOfCall; i++ {
				go func() {
					counter.Up()
					wg.Done()
				}()
			}

			wg.Wait()

			if got := counter.GetCount(); got != tc.want {
				t.Fatalf("\t%s\tShould get a counter value is %d: %d", failed, tc.want, got)
			}
			t.Logf("\t%s\tShould get a counter value is %d", succeed, tc.want)
		}

		t.Run(tc.label, tf)
	}
}
