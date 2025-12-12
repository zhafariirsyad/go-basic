package main

import "fmt"

type BangunDatar interface{
	HitungLuas() float64
}

type Persegi struct{
	Sisi float64
}

type PersegiPanjang struct{
	Panjang float64
	Lebar   float64
}

func (persegi Persegi) HitungLuas() float64{
	luas := persegi.Sisi * persegi.Sisi

	return luas
}

func (persegiPanjang PersegiPanjang) HitungLuas() float64{
	luas := persegiPanjang.Panjang * persegiPanjang.Lebar

	return luas
}

func main(){
	// var luasPersegi, luasPersegiPanjang BangunDatar
	var bangunDatarPersegi Persegi
	var bangunDatarPersegiPanjang PersegiPanjang

	bangunDatarPersegi = Persegi{Sisi: 29.5}
	fmt.Println("Hitung Luas Bangun Datar Persegi:", HitungLuasBangunDatar(bangunDatarPersegi))

	bangunDatarPersegiPanjang = PersegiPanjang{Panjang: 10.77, Lebar: 59.43}
	fmt.Println("Luas Persegi Panjang:", HitungLuasBangunDatar(bangunDatarPersegiPanjang))
}

func HitungLuasBangunDatar(bangunDatar BangunDatar) float64{
	return bangunDatar.HitungLuas()
}
