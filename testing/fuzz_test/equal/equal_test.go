package equal

import "testing"

const succeed = "\u2713"
const failed = "\u2717"

func TestSliceByte(t *testing.T) {
	testCases := []struct {
		label string
		a     []byte
		b     []byte
		want  bool
	}{
		{
			label: "equal",
			a:     []byte{1, 2, 3},
			b:     []byte{1, 2, 3},
			want:  true,
		},
		{
			label: "notEqual",
			a:     []byte{1, 2, 3},
			b:     []byte{3, 2, 1},
			want:  false,
		},
	}

	t.Log("Given a need to test equality of byte slice")
	for i, tc := range testCases {
		t.Logf("\tTest %d:\t%s\n", i, tc.label)

		tf := func(t *testing.T) {
			if got := SliceByte(tc.a, tc.b); got != tc.want {
				t.Fatalf("\t%s\tShould get an equality %v, got %v", failed, tc.want, got)
			}
			t.Logf("\t%s\tShould get an equality %v", succeed, tc.want)
		}

		t.Run(tc.label, tf)
	}
}

func FuzzSlice(f *testing.F) {
	ff := func(t *testing.T, a []byte, b []byte) {
		SliceByte(a, b)
	}
	f.Fuzz(ff)
}
