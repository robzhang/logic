package main

import (
	"fmt"
	"time"
)

type MatchFunc = func([]int, int) bool

func main() {
	matchFunc := []MatchFunc{
		nil,
		match02,
		match03,
		match04,
		match05,
		match06,
		match07,
		match08,
		match09,
		match10,
	}

	answer1 := make([]int, len(matchFunc))
	answer2 := make([]int, len(matchFunc))

	//使用2种不同的方法穷举所有解

	t1 := time.Now()
	solve(answer1, 0, matchFunc)
	t2 := time.Now()
	fmt.Printf("Solve done, takes [%f] seconds!\n", t2.Sub(t1).Seconds())

	t1 = time.Now()
	solve2(answer2, matchFunc)
	t2 = time.Now()
	fmt.Printf("Solve done, takes [%f] seconds!\n", t2.Sub(t1).Seconds())

	return
}

const (
	A = iota
	B
	C
	D
)

/* 递归穷举所有可能解，对每个解使用match判断是否复合每道题的逻辑描述 */
func solve(answer []int, n int, matchFunc []MatchFunc) {
	for i := A; i <= D; i++ {
		answer[n] = i
		if n < len(answer)-1 {
			solve(answer, n+1, matchFunc)
		} else {
			if match(answer, matchFunc) {
				printAnswer(answer)
			}
		}
	}
}

/* 循环穷举所有可能解，对每个解使用match判断是否复合每道题的逻辑描述 */
func solve2(answer []int, matchFunc []MatchFunc) {
	n := len(answer)
	/* 每道题有4种选择：ABCD；一共n道题，所以一共有4的 n 次方种可能解 */
	total := 1 << (uint(n) * 2)
	for i := 0; i < total; i++ {
		k := 0
		j := i
		for j != 0 {
			answer[k] = (j & 0x3) //每道题有【A，B，C，D】四种选择，用【0，1，2，3】表示
			k++
			j = (j >> 2)
		}

		if match(answer, matchFunc) {
			printAnswer(answer)
		}
	}
}

func match(answer []int, matchFunc []MatchFunc) bool {
	for i, fun := range matchFunc {
		if fun != nil {
			if !fun(answer, answer[i]) {
				return false
			}
		}
	}

	return true
}

func printAnswer(answer []int) {
	asc := [...]rune{'A', 'B', 'C', 'D'}
	fmt.Println("Answer:")
	for i, n := range answer {
		fmt.Printf("%d. %c\n", i, asc[n])
	}
}

func match02(answer []int, k int) bool {
	options := [...]int{C, D, A, B}
	return answer[4] == options[k]
}

func match03(answer []int, k int) bool {
	var options3 = [...]int{3 - 1, 6 - 1, 2 - 1, 4 - 1}
	var nk [4]int

	nk[answer[options3[0]]]++
	nk[answer[options3[1]]]++
	nk[answer[options3[2]]]++
	nk[answer[options3[3]]]++

	return nk[answer[options3[uint(k+1)&0x3]]] == 3 && nk[answer[options3[k]]] == 1
}

func match04(answer []int, k int) bool {
	//1 5, 2 7, 1 9, 6 10
	options := [4][2]int{{1 - 1, 5 - 1}, {2 - 1, 7 - 1}, {1 - 1, 9 - 1}, {6 - 1, 10 - 1}}

	if answer[options[k][0]] != answer[options[k][1]] {
		return false
	}

	for i := 1; i < 4; i++ {
		j := (k + i) % 4
		if answer[options[j][0]] == answer[options[j][1]] {
			return false
		}
	}

	return true
}

func match05(answer []int, k int) bool {
	options := [...]int{8 - 1, 4 - 1, 9 - 1, 7 - 1}

	if k != answer[options[k]] {
		return false
	}

	for i := 1; i < 4; i++ {
		j := (k + i) % 4
		if k == answer[options[j]] {
			return false
		}
	}

	return true
}

func match06(answer []int, k int) bool {
	options := [4][2]int{{2 - 1, 4 - 1}, {1 - 1, 6 - 1}, {3 - 1, 10 - 1}, {5 - 1, 9 - 1}}

	k8 := answer[7]

	if answer[options[k][0]] != k8 || answer[options[k][1]] != k8 {
		return false
	}

	for i := 1; i < 4; i++ {
		j := (k + i) % 4
		if answer[options[j][0]] == k8 && answer[options[j][1]] == k8 {
			return false
		}
	}

	return true
}

func match07(answer []int, k int) bool {
	options := [...]int{C, B, A, D}
	var nk [4]int

	for _, v := range answer {
		nk[v]++
	}

	n := nk[options[k]]
	for _, v := range nk {
		if n > v {
			return false
		}
	}

	return true
}

func match08(answer []int, k int) bool {
	options := [...]int{7 - 1, 5 - 1, 2 - 1, 10 - 1}
	k1 := answer[0]

	if k1+1 == answer[options[k]] || k1-1 == answer[options[k]] {
		return false
	}

	for i := 1; i < 4; i++ {
		j := (k + i) % 4
		if k1+1 != answer[options[j]] && k1-1 != answer[options[j]] {
			return false
		}
	}

	return true
}

func match09(answer []int, k int) bool {
	//16, x5
	k1 := answer[0]
	k6 := answer[5]

	b16 := (k1 == k6)

	k5 := answer[4]

	//6 10 2 9
	options := [...]int{6 - 1, 10 - 1, 2 - 1, 9 - 1}

	if b16 == (k5 == answer[options[k]]) {
		return false
	}

	for i := 1; i < 4; i++ {
		j := (k + i) % 4
		b5x := (answer[options[j]] == k5)
		if b5x != b16 {
			return false
		}
	}

	return true
}

func match10(answer []int, k int) bool {

	var nk [4]int
	for _, v := range answer {
		nk[v]++
	}

	min, max := nk[0], nk[1]
	if min > max {
		min, max = max, min
	}

	for i := 2; i < 4; i++ {
		if nk[i] > max {
			max = nk[i]
		} else if nk[i] < min {
			min = nk[i]
		}
	}

	options := [...]int{3, 2, 4, 1}
	if max-min != options[k] {
		return false
	}

	return true
}
