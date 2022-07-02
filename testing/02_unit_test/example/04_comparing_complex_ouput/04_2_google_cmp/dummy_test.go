package _4_2_google_cmp

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

const succeed = "\u2713"
const failed = "\u2717"

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
