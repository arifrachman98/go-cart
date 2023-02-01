package main

import (
	"fmt"
	"strconv"

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
	req := gin.Default()

	req.GET("/keranjang", func(ctx *gin.Context) {
		filterNama := ctx.Query("namaproduk")
		filterKuantitas := ctx.DefaultQuery("count", "0")
		count, _ := strconv.Atoi(filterKuantitas)

		ctx.JSON(200, keranjang.TampilkanKeranjang(filterNama, count))
	})

	req.POST("/keranjang", func(ctx *gin.Context) {
		var produk Produk
		if err := ctx.ShouldBindJSON(&produk); err != nil {
			ctx.JSON(400, gin.H{"error":err.Error()})
			return
		}

		keranjang.TambahProduk(produk)
		ctx.JSON(200, gin.H{"Pesan":"Produk berhasil ditambahkan"})
	})

	req.DELETE("/keranjang/:kodeproduk", func(ctx *gin.Context) {
		kodeP := ctx.Param("kodeproduk")
		
		keranjang.HapusProduk(kodeP)
		ctx.JSON(200, gin.H{"Pesan":"Produk berhasil dihapus"})
	})

	err := req.Run()
	if err != nil {
		fmt.Println(err)
	}

}