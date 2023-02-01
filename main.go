package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Produk struct {
	KodeProduk string `json:"kodeproduk"`
	NamaProduk string `json:"namaproduk"`
	Kuantitas int `json:"kuantitas"`
}

type Keranjang struct {
	Produk []Produk
}

var keranjang = &Keranjang{}

func (k *Keranjang) TambahProduk(produk Produk)  {
	for i, p := range k.Produk {
		if p.KodeProduk == produk.KodeProduk {
			k.Produk[i].Kuantitas += produk.Kuantitas
			return
		}
	}
	k.Produk = append(k.Produk, produk)
}

func (k *Keranjang) HapusProduk(kode string)  {
	for i, p := range k.Produk {
		if p.KodeProduk == kode {
			k.Produk = append(k.Produk[:i], k.Produk[i+1:]...)
			return
		}
	}
}

func (k *Keranjang) TampilkanKeranjang(filterNama string, filterKuantitas int) []Produk {
	var hasil []Produk
	for _, p := range k.Produk {
		if filterNama != "" && p.NamaProduk != filterNama {
			continue
		}
		if filterKuantitas > 0 && p.Kuantitas != filterKuantitas {
			continue
		}
		hasil = append(hasil, p)
	}
	return hasil
}

func main() {
	
}