package main
import "fmt"

func main(){
	const NMAX int =100
	var A[NMAX]int
	var n, i, terbesar, terkecil int

	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
		terbesar = A[i]
		terkecil = A[i]
	}
	for i = 0; i<n; i++{
		if terbesar < A[i]{
			terbesar = A[i]
		}else if terkecil > A[i]{
			terkecil = A[i]
		}
	}
	fmt.Println("Nilai terkecil :", terkecil)
	fmt.Println("Nilai terbesar :", terbesar)

}