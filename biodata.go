package main

import (
	"fmt"
	"os"
	"training-golang-corpu/models"
	"training-golang-corpu/mt"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: main <nama>")
		return

		nama := args[1]
		peserta := mt.UstomPst{
			Peserta: models.Peserta{
				Nama: nama,
			},
		}
		peserta.FinPst()
	}
}
