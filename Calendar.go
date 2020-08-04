package main

import (
	"fmt"
	"time"
)

func main(){

	const tahun int = 2020
	bulan := [12]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli",
						 "Agustus", "September", "Oktober", "November", "Desember"}

	fmt.Printf("\nKalender Tahun %d\n" , tahun)
	for i:=0; i < 12; i++ {
		fmt.Printf("=====\t%s   =====\n", bulan[i] )

		tanggal := Date(tahun, i + 2, 0)
		// fmt.Printf("=====\t%s   =====\n", tanggal.Month())
		fmt.Printf("S   S   R   K   J   S   M\n")

		tanggalSebelum := Date(tahun, i + 1, 0)
		hariSebelum := int(tanggalSebelum.Weekday())

		spasi := ""
		for i := 0; i < hariSebelum; i++{
			spasi = spasi + "    "
		}

		fmt.Printf(spasi)

		for j := 0; j < int(tanggal.Day()); j++ {
			fmt.Printf("%d", j + 1)

			if (j +1) > 9 {
				fmt.Printf("  ")
			} else {
				fmt.Printf("   ")
			}

			hariSebelum++
			if hariSebelum == 7 {
				hariSebelum = 0
				fmt.Printf("\n")
			}

		}
		fmt.Printf("\n\n")


	}

}

func Date(year, month, day int) time.Time  {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}