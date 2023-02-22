package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func sortedKeys(mp map[int]int) []int {
	a := make([]int, 0)
	for k := range mp {
		a = append(a, k)
	}
	sort.Ints(a)
	return a
}

func primeFactorize(n int) []int {
	a := make([]int, 0)
	for n%2 == 0 {
		a = append(a, 2)
		n /= 2
	}
	f := 3
	for f*f <= n {
		if n%f == 0 {
			a = append(a, f)
			n /= f
		} else {
			f += 2
		}
	}
	if n != 1 {
		a = append(a, n)
	}
	return a
}

func count(slc []int) map[int]int {
	c := make(map[int]int)
	for _, n := range slc {
		c[n]++
	}
	return c
}

func show(num int, mp map[int]int) {
	fmt.Printf("%v = ", num)
	for i, k := range sortedKeys(mp) {
		if i == 0 {
			fmt.Printf("%v^%v", k, mp[k])
		} else {
			fmt.Printf(" * %v^%v", k, mp[k])
		}
	}
	fmt.Println()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: The argument must be integer")
		os.Exit(1)
	}
	arg := os.Args[1]
	num, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ans := primeFactorize(num)
	//fmt.Println(ans)
	counted := count(ans)
	//fmt.Println(counted)
	show(num, counted)
}
