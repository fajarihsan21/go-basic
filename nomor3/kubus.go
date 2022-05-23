package nomor3

type Kubus struct {
	Rusuk float64
}

func (k *Kubus) Keliling() float64 {
	r := k.Rusuk
	hasil := 12 * r
	return hasil
}

func (k *Kubus) Luas() float64 {
	r := k.Rusuk

	hasil := r * r * 6
	return hasil
}

func (k *Kubus) Volume() float64 {
	r := k.Rusuk

	hasil := r * r * r
	return hasil
}
