package selector

import (
	"errors"

	"github.com/RoaringBitmap/roaring"
)

type (
	selector struct {
		Browser map[int]*roaring.Bitmap
		Country map[int]*roaring.Bitmap
		Device  map[int]*roaring.Bitmap
		Os      map[int]*roaring.Bitmap
	}

	SetRequest struct {
		ID              uint32
		BrowserListType int
		BrowserList     []int
		CountryListType int
		CountryList     []int
		DeviceListType  int
		DeviceList      []int
		OsListType      int
		OsList          []int
	}

	GetRequest struct {
		Browser int
		Country int
		Device  int
		Os      int
	}

	response struct {
		ids []uint32
	}
)

var (
	browserCode     = []int{1, 2, 3, 4, 5}
	countryCode     = []int{1, 2, 3, 4, 5}
	deviceCode      = []int{1, 2, 3, 4, 5}
	osCode          = []int{1, 2, 3, 4, 5}
	none            = 0
	allowed         = 1
	blocked         = 2
	listTypeBrowser = 1
	listTypeCountry = 2
	listTypeDevice  = 3
	listTypeOS      = 4
)

func New() *selector {
	s := &selector{
		Browser: make(map[int]*roaring.Bitmap),
		Country: make(map[int]*roaring.Bitmap),
		Device:  make(map[int]*roaring.Bitmap),
		Os:      make(map[int]*roaring.Bitmap),
	}
	s.init()
	return s
}

func (s *selector) init() {
	for v := range browserCode {
		s.Browser[v] = &roaring.Bitmap{}
	}

	for v := range countryCode {
		s.Country[v] = &roaring.Bitmap{}
	}

	for v := range deviceCode {
		s.Device[v] = &roaring.Bitmap{}
	}

	for v := range osCode {
		s.Os[v] = &roaring.Bitmap{}
	}
}

func (s *selector) Set(data SetRequest) error {
	var err error

	if data.ID < 1 {
		err = errors.New("id of item cant be lower than 1")
		return err
	}

	err = s.set(
		listTypeBrowser,
		data.BrowserListType,
		data.BrowserList,
		data.ID,
	)
	err = s.set(
		listTypeCountry,
		data.CountryListType,
		data.CountryList,
		data.ID,
	)
	err = s.set(
		listTypeDevice,
		data.DeviceListType,
		data.DeviceList,
		data.ID,
	)
	err = s.set(
		listTypeOS,
		data.OsListType,
		data.OsList,
		data.ID,
	)

	return err
}

func (s *selector) Get(data GetRequest) ([]uint32, error) {
	res := s.Browser[data.Browser]
	res.And(s.Country[data.Country])
	res.And(s.Device[data.Device])
	res.And(s.Os[data.Os])

	return res.ToArray(), nil
}

func (s *selector) set(
	t int,
	listType int,
	list []int,
	id uint32,
) error {
	var m map[int]*roaring.Bitmap


	switch t {
	case listTypeBrowser:
		m = s.Browser
	case listTypeCountry:
		m = s.Country
	case listTypeDevice:
		m = s.Device
	case listTypeOS:
		m = s.Os
	default:
		return errors.New("undefined type")
	}

	if listType == none {
		for _, v := range m {
			v.Add(id)
		}
	}

	if listType == allowed {
		for _, v := range list {
			m[v].Add(id)
		}
		for k, v := range m {
			if !contains(list, k) {
				v.Remove(id)
			}
		}
	}

	if listType == blocked {
		for _, v := range list {
			m[v].Remove(id)
		}
		for k, v := range m {
			if !contains(list, k) {
				v.Add(id)
			}
		}
	}
	return nil
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
