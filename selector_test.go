package selector_test

import (
	"math/rand"
	"testing"
	"time"

	selector "github.com/RapidCodeLab/bitmap-index-targeting"
)

func TestWithNoTargerts(t *testing.T) {
	s := selector.New()
	setReq := selector.SetRequest{
		ID: 11,
	}
	err := s.Set(setReq)
	if err != nil {
		t.Fail()
	}

	getReq := selector.GetRequest{
		Browser: 1,
		Country: 2,
		Device:  3,
		Os:      4,
	}
	ids, err := s.Get(getReq)
	if err != nil {
		t.Fail()
	}
	if len(ids) != 1 {
		t.Fail()
	}
	if ids[0] != 11 {
		t.Fail()
	}
}

func TestWithBrowserAllowed(t *testing.T) {
	s := selector.New()

	testCases := []struct {
		setReq         []selector.SetRequest
		getReq         selector.GetRequest
		expectedResLen int
	}{
		{
			setReq: []selector.SetRequest{
				{
					ID:              11,
					BrowserListType: selector.Allowed,
					BrowserList:     []int{1},
				},
			},
			getReq: selector.GetRequest{
				Browser: 1,
				Country: 2,
				Device:  3,
				Os:      4,
			},
			expectedResLen: 1,
		},
		{
			setReq: []selector.SetRequest{
				{
					ID:              11,
					BrowserListType: selector.Allowed,
					BrowserList:     []int{1},
				},
			},
			getReq: selector.GetRequest{
				Browser: 2,
				Country: 2,
				Device:  3,
				Os:      4,
			},
			expectedResLen: 0,
		},
		{
			setReq: []selector.SetRequest{
				{
					ID:              11,
					BrowserListType: selector.Allowed,
					BrowserList:     []int{1},
				},
				{
					ID: 12,
				},
			},
			getReq: selector.GetRequest{
				Browser: 1,
				Country: 2,
				Device:  3,
				Os:      4,
			},
			expectedResLen: 2,
		},
	}

	for _, tc := range testCases {

		for _, req := range tc.setReq {
			err := s.Set(req)
			if err != nil {
				t.Fail()
			}
		}

		ids, err := s.Get(tc.getReq)
		if err != nil {
			t.Fail()
		}

		if len(ids) != tc.expectedResLen {
			t.Fail()
		}
	}
}

func TestWithBrowserBlocked(t *testing.T) {
	s := selector.New()

	testCases := []struct {
		setReq         []selector.SetRequest
		getReq         selector.GetRequest
		expectedResLen int
	}{
		{
			setReq: []selector.SetRequest{
				{
					ID:              11,
					BrowserListType: selector.Blocked,
					BrowserList:     []int{1},
				},
			},
			getReq: selector.GetRequest{
				Browser: 1,
				Country: 2,
				Device:  3,
				Os:      4,
			},
			expectedResLen: 0,
		},
		{
			setReq: []selector.SetRequest{
				{
					ID:              11,
					BrowserListType: selector.Blocked,
					BrowserList:     []int{1},
				},
			},
			getReq: selector.GetRequest{
				Browser: 2,
				Country: 2,
				Device:  3,
				Os:      4,
			},
			expectedResLen: 1,
		},
		{
			setReq: []selector.SetRequest{
				{
					ID:              11,
					BrowserListType: selector.Blocked,
					BrowserList:     []int{1},
				},
				{
					ID: 12,
				},
			},
			getReq: selector.GetRequest{
				Browser: 1,
				Country: 2,
				Device:  3,
				Os:      4,
			},
			expectedResLen: 1,
		},
	}

	for _, tc := range testCases {

		for _, req := range tc.setReq {
			err := s.Set(req)
			if err != nil {
				t.Fail()
			}
		}

		ids, err := s.Get(tc.getReq)
		if err != nil {
			t.Fail()
		}

		if len(ids) != tc.expectedResLen {
			t.Fail()
		}
	}
}

func BenchmarkSet(b *testing.B) {
	s := selector.New()
	itemsAmount := 99999
	testSet := []selector.SetRequest{}

	for i := 0; i < itemsAmount; i++ {
		testSet = append(testSet, selector.SetRequest{
			ID:              uint32(i),
			BrowserListType: 1,
			BrowserList:     []int{1, 2, 3, 4, 5},
			CountryListType: 2,
			CountryList:     []int{1, 3},
			DeviceListType:  1,
			DeviceList:      []int{1},
			OsListType:      2,
			OsList:          []int{4, 5},
		})
	}

	b.ResetTimer()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		s.Set(testSet[r.Intn(itemsAmount)])
	}
}

func BenchmarkGet(b *testing.B) {
	s := selector.New()
	itemsAmount := 99999

	for i := 0; i < itemsAmount; i++ {
		s.Set(selector.SetRequest{
			ID:              uint32(i),
			BrowserListType: 1,
			BrowserList:     []int{1, 2, 3, 4, 5},
			CountryListType: 2,
			CountryList:     []int{1, 3},
			DeviceListType:  1,
			DeviceList:      []int{1},
			OsListType:      2,
			OsList:          []int{4, 5},
		})
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		req := selector.GetRequest{
			Browser: 1,
			Country: 2,
			Device:  3,
			Os:      5,
		}
		s.Get(req)
	}
}
