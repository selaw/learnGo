package main

import (
	"fmt"
)

func main() {

	//Struct
	type Hewan struct {
		hewan string
		age   int
	}
	hewan := Hewan{
		"Kucing", 2,
	}
	fmt.Println(hewan)
	/*
		//MAPS
		hewan := map[int]string{
			12020: "Kucing",
			13134: "Sapi",
			14944: "Ayam",
		}
		for id, value := range hewan {
			fmt.Println(id, value)
		}
			//Slice
			hari := [8]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
			Slicehari := hari[:3]
			fmt.Println(Slicehari)
			fmt.Println("Panjangnya ", len(Slicehari))
			fmt.Println("Kapasitasnya ", cap(Slicehari))

				//Array Loop
					nama := [...]string{"Nanda", "Ariel"}

					for i := 0; i < len(nama); i++ {
						fmt.Println("Namanya adalah " + nama[i])

					}
					fmt.Print("\n")
					for _, x := range nama {
						fmt.Printf("Namanya adalah %s ", x)
						fmt.Print("\n")

					} */

}
