package nomor3

type Kubus struct {
	Rusuk float64
}

func (k *Kubus) Keliling() float64 {
	r := k.Rusuk
	kelKubus := 12 * r
	return kelKubus
}

func (k *Kubus) Luas() float64 {
	r := k.Rusuk

	luasKubus := r * r * 6
	return luasKubus
}

func (k *Kubus) Volume() float64 {
	r := k.Rusuk

	volKubus := r * r * r
	return volKubus
}
