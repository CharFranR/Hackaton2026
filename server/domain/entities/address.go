package domain

import "fmt"

type Address struct {
	Department   string
	Municipality string
	AddressLine  string
	Latitude     float64
	Longitude    float64
}

// Builder

func NewAddress(department, municipality, addressLine string) (*Address, error) {
	if department == "" {
		return nil, ErrDepartmentRequired
	}
	if municipality == "" {
		return nil, ErrMunicipalityRequired
	}
	if addressLine == "" {
		return nil, ErrAddressLineRequired
	}

	return &Address{
		Department:   department,
		Municipality: municipality,
		AddressLine:  addressLine,
	}, nil
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
