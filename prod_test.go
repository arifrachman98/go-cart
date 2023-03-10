package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTambahProduk(t *testing.T) {
	keranjang := &Keranjang{}
	produk := Produk{
		KodeProduk: "A001",
		NamaProduk: "Kaca Mata",
		Kuantitas:  100,
	}

	keranjang.TambahProduk(produk)

	assert.Len(t, keranjang.Prod, 1)
	assert.Equal(t, produk, keranjang.Prod[0])

	produk2 := Produk{
		KodeProduk: "B001",
		NamaProduk: "Sepatu Renang",
		Kuantitas:  102,
	}
	keranjang.TambahProduk(produk2)

	assert.Len(t, keranjang.Prod, 2)
	assert.Equal(t, "B001", keranjang.Prod[1].KodeProduk)
	assert.Equal(t, "Sepatu Renang", keranjang.Prod[1].NamaProduk)
	assert.Equal(t, 102, keranjang.Prod[1].Kuantitas)
}

func TestHapusProduk(t *testing.T) {
	keranjang := &Keranjang{
		Prod: []Produk{
			{
				KodeProduk: "A001",
				NamaProduk: "Sepatu Tali",
				Kuantitas: 10,
			},
			{
				KodeProduk: "A002",
				NamaProduk: "Sepatu Balet",
				Kuantitas: 50,
			},
		},
	}
	
	keranjang.HapusProduk("A001")
	assert.Len(t, keranjang.Prod, 1)
	assert.Equal(t, "A002", keranjang.Prod[0].KodeProduk)
	assert.Equal(t, "Sepatu Balet", keranjang.Prod[0].NamaProduk)
	assert.Equal(t, 50, keranjang.Prod[0].Kuantitas)
}

func TestTampilkanProduk(t *testing.T) {
	keranjang := &Keranjang{
		Prod: []Produk{
			{
				KodeProduk: "A001",
				NamaProduk: "Kunci Gantung",
				Kuantitas: 200,
			},
			{
				KodeProduk: "A002",
				NamaProduk: "Sepatu Bola",
				Kuantitas: 250,
			},
			{
				KodeProduk: "A003",
				NamaProduk: "Lampu Bola",
				Kuantitas: 25,
			},
		},
	}

	produk := keranjang.TampilkanKeranjang("", 0)
	assert.Len(t, produk, 3)
	assert.Equal(t, "A001", produk[0].KodeProduk)
	assert.Equal(t, "Kunci Gantung", produk[0].NamaProduk)
	assert.Equal(t, 200, produk[0].Kuantitas)
	assert.Equal(t, "A002", produk[1].KodeProduk)
	assert.Equal(t, "Sepatu Bola", produk[1].NamaProduk)
	assert.Equal(t, 250, produk[1].Kuantitas)
	assert.Equal(t, "A003", produk[2].KodeProduk)
	assert.Equal(t, "Lampu Bola", produk[2].NamaProduk)
	assert.Equal(t, 25, produk[2].Kuantitas)
}