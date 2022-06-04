package __1_comparing_type

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"net/http"
	"reflect"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestDetermineEquality(t *testing.T) {
	testCases := []struct {
		name string
		f    func(t *testing.T)
	}{
		{
			name: "integer",
			f: func(t *testing.T) {
				a := 1
				b := 2

				if a != b {
					t.Errorf("\t%s\t%d is not equal with %d", failed, a, b)
				}
			},
		},
		{
			name: "string",
			f: func(t *testing.T) {
				a := "Steve"
				b := "steve"

				if a != b {
					t.Errorf("\t%s\t%s is not equal with %s", failed, a, b)
				}
			},
		},
		{
			name: "emptySlice",
			f: func(t *testing.T) {
				a := []int{}
				var b []int

				if !reflect.DeepEqual(a, b) {
					t.Errorf("\t%s\t%#v is not equal with %#v", failed, a, b)
				} else {
					t.Logf("\t%s\t%#v is equal with %#v", succeed, a, b)
				}
			},
		},
		{
			name: "unorderedSlice",
			f: func(t *testing.T) {
				a := map[string]int{"a": 1, "b": 2}
				b := map[string]int{"b": 2, "a": 1}

				if !reflect.DeepEqual(a, b) {
					t.Errorf("\t%s\t%v is not equal with %v", failed, a, b)
				} else {
					t.Logf("\t%s\t%v is equal with %v", succeed, a, b)
				}
			},
		},
		{
			name: "unorderedMap",
			f: func(t *testing.T) {
				a := []int{1, 2, 3}
				b := []int{1, 3, 2}

				if !reflect.DeepEqual(a, b) {
					t.Errorf("\t%s\t%v is not equal with %v", failed, a, b)
				} else {
					t.Logf("\t%s\t%v is equal with %v", succeed, a, b)
				}
			},
		},
		{
			name: "anonymousStruct",
			f: func(t *testing.T) {
				a := struct {
					ID int
				}{
					ID: 1,
				}

				b := struct {
					ID int
				}{
					ID: 1,
				}

				if !reflect.DeepEqual(a, b) {
					t.Errorf("\t%s\t%+v is not equal with %+v", failed, a, b)
				} else {
					t.Logf("\t%s\t%+v is equal with %+v", succeed, a, b)
				}
			},
		},
		{
			name: "struct",
			f: func(t *testing.T) {
				type one struct {
					ID int
				}

				type two struct {
					ID int
				}

				a := one{1}
				b := two{ID: 1}

				if !reflect.DeepEqual(a, b) {
					t.Errorf("\t%s\t%+v is not equal with %+v", failed, a, b)
				} else {
					t.Logf("\t%s\t%+v is equal with %+v", succeed, a, b)
				}
			},
		},
		{
			name: "error",
			f: func(t *testing.T) {
				a := context.Canceled
				b := fmt.Errorf("context canceled")

				t.Log("\tWith '==' operator")
				if a != b {
					t.Errorf("\t\t%s\t'%+v' is not equal with '%+v'", failed, a, b)
				} else {
					t.Logf("\t\t%s\t'%+v' is equal with '%+v'", succeed, a, b)
				}

				t.Log("\tWith 'errors.Is'")
				if errors.Is(a, b) {
					t.Errorf("\t\t%s\t'%+v' is not equal with '%+v'", failed, a, b)
				} else {
					t.Logf("\t\t%s\t'%+v' is equal with '%+v'", succeed, a, b)
				}

				t.Log("\tWith 'reflect.DeepEqual'")
				if !reflect.DeepEqual(a, b) {
					t.Errorf("\t\t%s\t'%+v' is not equal with '%+v'", failed, a, b)
				} else {
					t.Logf("\t\t%s\t'%+v' is equal with '%+v'", succeed, a, b)
				}
			},
		},
		{
			name: "interface",
			f: func(t *testing.T) {
				var a interface {
					Count() error
				}

				var b interface {
					Count()
				}

				//t.Log("\tWith '==' operator")
				//if a != b {
				//	t.Errorf("\t\t%s\t'%+v' is not equal with '%+v'", failed, a, b)
				//} else {
				//	t.Logf("\t\t%s\t'%+v' is equal with '%+v'", succeed, a, b)
				//}

				t.Log("\tWith 'reflect.DeepEqual'")
				if !reflect.DeepEqual(a, b) {
					t.Errorf("\t\t%s\t'%+v' is not equal with '%+v'", failed, a, b)
				} else {
					t.Logf("\t\t%s\t'%+v' is equal with '%+v'", succeed, a, b)
				}
			},
		},
		{
			name: "pointer",
			f: func(t *testing.T) {
				a := &http.Server{Addr: ":7070"}
				b := &http.Server{Addr: ":7070"}

				t.Log("\tWith '==' operator")
				if a != b {
					t.Errorf("\t\t%s\t'%+v' is not equal with '%+v'", failed, a, b)
				} else {
					t.Logf("\t\t%s\t'%+v' is equal with '%+v'", succeed, a, b)
				}

				t.Log("\tWith 'reflect.DeepEqual'")
				if !reflect.DeepEqual(a, b) {
					t.Errorf("\t\t%s\t'%+v' is not equal with '%+v'", failed, a, b)
				} else {
					t.Logf("\t\t%s\t'%+v' is equal with '%+v'", succeed, a, b)
				}
			},
		},
	}

	for i, tc := range testCases {
		tf := func(t *testing.T) {
			t.Logf("Test %d:\t%s\n", i, tc.name)
			tc.f(t)
		}
		t.Run(tc.name, tf)
	}
}

func TestCompareDiff(t *testing.T) {
	a := &http.Server{Addr: ":7070"}
	b := &http.Server{Addr: ":7071"}

	t.Log("With 'reflect.DeepEqual'")
	if !reflect.DeepEqual(a, b) {
		t.Errorf("\t%s\t'%+v' is not equal with '%+v'", failed, a, b)
	} else {
		t.Logf("\t%s\t'%+v' is equal with '%+v'", succeed, a, b)
	}

	t.Log("With 'cmp'")
	if diff := cmp.Diff(a, b, cmpopts.IgnoreUnexported(http.Server{})); diff != "" {
		t.Logf("\t%s\ta & b are not equal:", failed)
		t.Errorf(diff)
	} else {
		t.Logf("\t%s\ta & b are equal:", succeed)
	}
}
