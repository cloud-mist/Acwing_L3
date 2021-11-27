// TAG：外部搜索，搜索顺序。
// 单词接龙
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 25

var (
	n    int
	word []string
	g    [N][N]int
	used [N]int
	ans  int
)

func main() {
	defer ot.Flush()

	// input
	n = rnI()
	for i := 0; i < n; i++ {
		word = append(word, rnS())
	}
	start := rnS()

	// 预处理 每种接法的重合的部分值
	for i, a := range word {
		for j, b := range word {
			la, lb := len(a), len(b)

			for k := 1; k < min(la, lb); k++ {
				// 从小到大枚举，因为是为了让龙更长,如果有重合就记录且break
				if a[la-k:] == b[:k] {
					g[i][j] = k
					break
				}
			}
		}
	}

	// 从能开始的单词dfs
	for i, v := range word {
		if string(v[0]) == start {
			dfs(v, i)
		}
	}
	fmt.Println(ans)
}

func dfs(dragon string, last int) {
	ans = max(len(dragon), ans) // 每次取最大值

	used[last]++ // 使用了1次

	for i := 0; i < n; i++ {
		// 如果说以字符串i为结尾的值不为0,说明可以接，并且i的使用次数要小于2.
		// 怎么接，前面是dragon,后面是，以非公共部分开始
		if g[last][i] != 0 && used[i] < 2 {
			dfs(dragon+word[i][g[last][i]:], i)
		}
	}

	used[last]-- // 回复现场
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
