package etc

//Output 1,2,3,4,5,6,7 ....N
func GetFibo1(N int) []int {
	Fibo1 := make([]int, N)

	a := 0

	for i := 0; i < N; i++ {
		a++
		Fibo1[i] += a
	}
	return Fibo1
}

//Output 1,3,5,7,9 .... N
func GetGanjil(N int) []int {
	Ganjil := make([]int, N)
	a := 1

	for i := 0; i < N; i++ {
		Ganjil[i] = a
		a += 2
	}
	return Ganjil
}

func GetPrima(N int) []int {
	Prima := make([]int, N)

	k := 0
	for i := 2; i <= 100; i++ {

		isPrime := true

		for j := 2; j < i; j++ {

			if i%j == 0 {

				isPrime = false
			}
		}

		if isPrime == true {

			Prima[k] = i
			k++
		}
	}
	return Prima
}

/*
func GetPrima(N int) []int {
	var Prima = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 163, 167, 173, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263}
	return Prima
}
*/
//Output 1,2,4,6,8,10 ...... N
func GetGenap(N int) []int {
	Genap := make([]int, N)

	a := 0

	for i := 0; i < N; i++ {
		if a <= 1 {
			a++
			Genap[i] = a
		} else {
			a += 2
			Genap[i] = a
		}
	}
	return Genap
}

//Output 1,1,2,3,5,8,13 ... N
func GetFibo2(N int) []int {
	Fibo2 := make([]int, N)

	var a, b, c int = 1, 1, 0

	for i := 0; i < N; i++ {
		Fibo2[i] = a
		c = a + b
		a = b
		b = c
	}
	return Fibo2
}

//Output 1,1,1,3,5,9 ..... N
func GetFibo3(N int) []int {
	Fibo3 := make([]int, N)

	var a, b, c, d int = 1, 1, 1, 1

	for i := 0; i < N; i++ {
		Fibo3[i] = a
		d = a + b + c
		a = b
		b = c
		c = d
	}
	return Fibo3
}

//Output 1,3,5,7,9 ..... N
func GetSegi3(N int) []int {
	Segi3 := make([]int, N)

	a := 1

	for i := 0; i < N; i++ {
		Segi3[i] = a
		a += 2
	}
	return Segi3
}

//Output 1,3,6,10, ..... N
func GetSegi4(N int) []int {
	Segi4 := make([]int, N+1)

	a := 1

	for i := 2; i <= N+2; i++ {
		Segi4[i-2] += a
		a += i
	}
	return Segi4
}

func GetAbjadB(N int) []string {
	AbjadB := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	return AbjadB
}
