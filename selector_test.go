package selector_test

import (
	"testing"

	selector "github.com/RapidCodeLab/bitmap-index-targeting"
)

func TestSet(t *testing.T) {
	s := selector.New()
	setReq := selector.SetRequest{
		ID: 1,
	}
	err := s.Set(setReq)
	if err != nil {
		t.Fail()
	}

	getReq := selector.GetRequest{
		Browser: 1,
		Country: 1,
		Device:  1,
		Os:      1,
	}
	ids, err := s.Get(getReq)
	if err != nil {
		t.Fail()
	}
	if len(ids) != 1 {
		t.Fail()
	}
	if ids[0] != 1 {
		t.Fail()
	}
}
