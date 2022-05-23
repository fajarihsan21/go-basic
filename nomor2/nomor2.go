package nomor2

import "fmt"

type DeretBilangan struct {
	N int
}

func (d *DeretBilangan) Prima() int {
	fmt.Println("Prima: ")
	var i, j, a int

	for i = 1; i <= d.N; i++ {
		a = 0
		for j = 1; j <= i; j++ {
			if i%j == 0 {
				a++

			}
		}

		if a == 2 {
			fmt.Println(i)
		}
		if i%d.N == 0 {
			return i
		}
	}
	return i
}

func (d *DeretBilangan) Ganjil() int {
	fmt.Println("Ganjil: ")
	i := 0
	for ; i <= d.N; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
	return i
}

func (d *DeretBilangan) Genap() int {
	fmt.Println("Genap: ")
	i := 0
	for ; i <= d.N; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	return i
}

func (d *DeretBilangan) Fibonacci() int {
	fmt.Println("Fibonacci: ")
	bil1 := 0
	bil2 := 1
	bil3 := bil2
	fmt.Println(bil2)
	for {
		bil3 = bil2
		bil2 = bil1 + bil2
		if bil2 >= d.N {
			break
		}
		bil1 = bil3
		fmt.Println(bil2)
	}
	return bil2
}

/* Output

Prima:
2
3
5
7
11
13
17
19
23
29
Ganjil:
1
3
5
7
9
11
13
15
17
19
21
23
25
27
29
Genap:
0
2
4
6
8
10
12
14
16
18
20
22
24
26
28
30
Fibonacci:
1
1
2
3
5
8
13
21

		**/
