package mt

import (
	"fmt"
	"strings"
	"training-golang-corpu/funtion"
	"training-golang-corpu/models"
)

type UstomPst struct {
	models.Peserta
}

func (p UstomPst) FinPst() {
	var psts = funtion.Peserta()

	for k := range psts {
		if strings.ToLower(psts[k].Nama) == strings.ToLower(p.Nama) {
			fmt.Println("ID", k+1)
			fmt.Println("Nama", psts.Nama)
		}
	}
}
