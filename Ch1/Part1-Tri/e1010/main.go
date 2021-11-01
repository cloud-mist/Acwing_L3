// 导弹拦截
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 1010

var (
	a, f, g [N]int
)

func main() {
	defer ot.Flush()

	// input
	t := strings.Split(rlS(), " ")
	for i := range t {
		a[i+1], _ = strconv.Atoi(t[i])
	}
	n := len(t)

	// p1
	res1 := 0
	for i := 1; i <= n; i++ {
		f[i] = 1
		for j := 1; j < i; j++ {
			if a[j] >= a[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		res1 = max(res1, f[i])
	}

	// p2
	res2 := 0
	for i := 1; i <= n; i++ {
		// 找到要替换的尾
		k := 0
		for k < res2 && g[k] < a[i] {
			k++
		}
		g[k] = a[i]

		// 说明要添加新坑
		if k >= res2 {
			res2++
		}
	}

	// output
	fmt.Printf("%d\n%d", res1, res2)

}

/* ======================================================================== */
//
//
//
//							 _____   _   _   ____
//							| ____| | \ | | |  _ \
//							|  _|   |  \| | | | | |
//							| |___  | |\  | | |_| |
//							|_____| |_| \_| |____/
//
//
//
/* ============================PART1: I/O ================================== */

var (
	ot = bufio.NewWriterSize(os.Stdout, int(1e6))
	in = bufio.NewScanner(os.Stdin)
)

func init()       { in.Buffer(make([]byte, 4096), int(1e9)) }
func rlS() string { in.Scan(); return in.Text() }

/* ===========================PART 2: Math Func ============================  */
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func memset(a []int, v int) {
	if len(a) == 0 {
		return
	}
	a[0] = v
	for bp := 1; bp < len(a); bp *= 2 {
		copy(a[bp:], a[:bp])
	}
}
