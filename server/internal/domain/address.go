package domain

import "fmt"

type Address struct {
	Department   string
	Municipality string
	AddressLine  string
	Latitude     float64
	Longitude    float64
}

func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s, %s", a.AddressLine, a.Municipality, a.Department)
}

func (a Address) HasCoordinates() bool {
	return a.Latitude != 0 && a.Longitude != 0
}

func (a Address) IsComplete() bool {
	return a.Department != "" && a.Municipality != "" && a.AddressLine != ""
}
