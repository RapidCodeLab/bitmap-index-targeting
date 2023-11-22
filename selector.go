package selector

import "github.com/RoaringBitmap/roaring"

type (
	selector struct {
		Browser map[int]*roaring.Bitmap
		Country map[int]*roaring.Bitmap
		Device  map[int]*roaring.Bitmap
		Os      map[int]*roaring.Bitmap
	}

	SetRequest struct {
		id              uint32
		browserListType int
		browserList     []int
		countryListType int
		countryList     []int
		deviceListType  int
		deviceList      []int
		osListType      int
		osList          []int
	}

	GetRequest struct {
		browser int
		country int
		device  int
		os      int
	}

	response struct {
		ids []uint32
	}
)

var (
	browserCode = []int{1, 2, 3, 4, 5}
	countryCode = []int{1, 2, 3, 4, 5}
	deviceCode  = []int{1, 2, 3, 4, 5}
	osCode      = []int{1, 2, 3, 4, 5}
	none        = 0
	allowed     = 1
	blocked     = 2
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
	// if target list not selected,
	// set id enabled in all countries
	if data.browserListType == none {
		for _, v := range s.Browser {
			v.Add(data.id)
		}
	}

	if data.browserListType == allowed {
		for _, v := range data.browserList {
			s.Browser[v].Add(data.id)
		}
	}
	return nil
}

func (s *selector) Get(data GetRequest) ([]uint32, error) {
	return []uint32{}, nil
}
