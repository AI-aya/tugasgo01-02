package main

import "fmt"



func main()  {

	var x int 
	fmt.Println("==========MENU================")
	fmt.Println(" 1. Perhitungan (+,-,/,*,%)")  
	fmt.Println(" 2. Luas Persegi")  
	fmt.Println(" 3. Luas Lingkaran")
	fmt.Println(" 4. Volume Tabung")  
	fmt.Println(" 5. Volume Prisma")    
	fmt.Println(" ")
	fmt.Print("Masukkan Pilihan : ")  
	fmt.Scanf("%d", &x)

	switch x {
	case 1:
		hitung()
		main()
	case 2:
		Luas_Persegi()
		
	case 3:
		Luas_Lingkaran()
		
	case 4:
		Volume_Tabung()
		
	case 5:
		Volume_Prisma()
		
	default:
		fmt.Println("Tidak ada dalam opsi, kembali pilih sesuai opsi")
		
		
	}
	
}



func hitung(){
	var operator string
	var num1 int
	var num2 int

	fmt.Println("=====================")
	fmt.Println("Perhitungan")
	fmt.Print(" ")
	fmt.Scanln()
	fmt.Println("")
	fmt.Print("Masukan angka Pertama: ")
	fmt.Scanln(&num1)
	fmt.Println(" ")
	fmt.Print("Masukan angka Kedua: ")
	fmt.Scanln(&num2)
	fmt.Print("Pilih operator: (+,-,/,*,%) : ")
	fmt.Scanln(&operator)
	output := 0

	switch operator {
	case "+":
		output = num1 + num2
		break
	case "-":
		output = num1 - num2
	case "/":
		output = num1 / num2
	case "*":
		output = num1 * num2
	case "%":
		output = num1 % num2
	default:
		fmt.Println("Operasi perhitungan salah")
		
	}

	fmt.Printf("%d %s %d = %d", num1, operator, num2, output)
	fmt.Println(" ")
}

func Luas_Persegi()  {
	var panjang, lebar, luas int

	fmt.Println("=====================")
	fmt.Print(" ")
	fmt.Scanln()
	fmt.Println("")
	fmt.Print("Masukan panjang persegi: ")
	fmt.Scanln(&panjang)
	
	fmt.Print("Masukan panjang lebar: ")
	fmt.Scanln(&lebar)

	luas = panjang*lebar
	fmt.Printf("Luas Persegi: %d x %d = %d", panjang, lebar, luas)


}

func Luas_Lingkaran()  {
	var r, luaslingkaran float64
	pi := 3.14

	fmt.Println("=====================")
	fmt.Print(" ")
	fmt.Scanln()
	fmt.Println("")
	fmt.Print("Masukan jari-jari: ")
	fmt.Scanln(&r)
	fmt.Println(" ")

	luaslingkaran = pi*r*r
	fmt.Printf("Luas lingkaran: %f x %f x %f = %f ", pi , r, r, luaslingkaran)

}



func Volume_Tabung() {

	var jari, tinggi, volumtabung float64
	pii := 3.14

	fmt.Println("=====================")
	fmt.Print(" ")
	fmt.Scanln()
	fmt.Println("")
	fmt.Print("Masukkan tinggi : ")
	fmt.Scanln(&tinggi)
	fmt.Print("Masukkan jari-jari : ")
	fmt.Scanln(&jari)
	
  
	volumtabung = jari * jari * tinggi * pii
  
	fmt.Printf("Volume balok adalah : %f x %f x %f x %f = %f", jari, jari, tinggi, pii, volumtabung)
  
}

func Volume_Prisma()  {
	var panjangg, lebarr, tinggii, volumprisma float64

	fmt.Println("=====================")
	fmt.Scanln()
	fmt.Print("Masukan panjang: ")
	fmt.Scanln(&panjangg)
	fmt.Print("Masukan lebar: ")
	fmt.Scanln(&lebarr)
	fmt.Print("Masukan tinggi: ")
	fmt.Scanln(&tinggii)

	volumprisma = 0.5 * panjangg * lebarr * tinggii
	
	fmt.Printf("Volume prisma adalah : 0.5 x %f x %f x %f = %f", panjangg, lebarr, tinggii, volumprisma)

	
}