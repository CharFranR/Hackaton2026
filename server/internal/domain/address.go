package domain

type Adress struct {
	Departament  string
	Municipality string
	AddressLine  string
	Latitude     float32
	Longitude    float32
}

func (a Adress) FullAddress() bool {
	return a.Departament != "" && a.Municipality != "" && a.AddressLine != "" && a.Latitude > 0 && a.Longitude > 0
}
