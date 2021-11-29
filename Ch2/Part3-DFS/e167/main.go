// 剪枝; 木棒
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const N = 70

var (
	a              []int
	n, sum, length int
	st             [N]bool
)

func main() {
	defer ot.Flush()

	for n = rnI(); n != 0; n = rnI() {
		for i := 0; i < N; i++ { // 初始化st
			st[i] = false
		}
		a = make([]int, 0) // 初始化a

		sum = 0
		for i := 0; i < n; i++ {
			a = append(a, rnI())
			sum += a[i]
		}
		sort.Slice(a, func(i, j int) bool { // 从大到小枚举
			return a[i] > a[j]
		})

		// 枚举长度
		length = 1
		for {
			// 如果是sum的约数，才有可能枚举成功
			if sum%length == 0 && dfs(0, 0, 0) {
				fmt.Fprintln(ot, length)
				break
			}
			length++
		}
	}
}

func dfs(u, s, start int) bool { // 第u组，第u组已有长度，第u组位置
	if u*length == sum { // 答案可行
		return true
	}

	if s == length {
		return dfs(u+1, 0, 0)
	}

	for i := start; i < n; i++ {
		if st[i] || s+a[i] > length { // 用过了或者超过了 不合法
			continue
		}

		st[i] = true
		if dfs(u, s+a[i], i+1) {
			return true
		}
		st[i] = false

		if s == 0 || s+a[i] == length { // TODO: 还是tm的不理解  <29-11-21, cloud mist> //
			return false
		}

		j := i
		for j < n && a[j] == a[i] { // 如果i不行，和他相等的都不行
			j++
		}
		i = j - 1

	}

	return false
}

/* ======================================================================== */
//
//
//
//						 _____   _   _   ____
//						| ____| | \ | | |  _ \
//						|  _|   |  \| | | | | |
//						| |___  | |\  | | |_| |
//						|_____| |_| \_| |____/
//
//
//
/* ============================PART 1: I/O ================================== */

var (
	ot = bufio.NewWriterSize(os.Stdout, int(1e6))
	in = bufio.NewScanner(os.Stdin)
)

func init()        { in.Split(bufio.ScanWords); in.Buffer(make([]byte, 4096), int(1e9)) }
func rnS() string  { in.Scan(); return in.Text() }
func rnI() int     { i, _ := strconv.Atoi(rnS()); return i }
func rnF() float64 { f, _ := strconv.ParseFloat(rnS(), 64); return f }

func rsI(l, r int) []int {
	t := make([]int, r)
	for i := l; i < r; i++ {
		t[i] = rnI()
	}
	return t
}

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
