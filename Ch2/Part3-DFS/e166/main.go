// 剪枝 sudoku
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const (
	N = 9
	M = 1 << N
)

var (
	str      []byte
	row, col [N]int
	cell     [3][3]int
	mNum     map[int]int
	one      [M]int
)

func main() {
	defer ot.Flush()

	// 预处理
	mNum = make(map[int]int)
	for i := 0; i < N; i++ { // 映射1所在位置和数字的关系
		mNum[1<<i] = i
	}
	for i := 0; i < M; i++ { // 记录每个数字1的个数
		one[i] = getoneNum(i)
	}

	for str = []byte(rnS()); !bytes.Equal(str, []byte("end")); str = []byte(rnS()) {
		ini()
		cnt := 0

		k := 0
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if str[k] != '.' {
					t := str[k] - '1'
					draw(i, j, int(t), true)
				} else {
					cnt++
				}
				k++
			}
		}

		dfs(cnt)

		fmt.Fprintf(ot, "%s\n", str)
	}
}

func dfs(cnt int) bool {
	// 如果cnt没了，说明得到了答案。因为答案只有一个
	if cnt == 0 {
		return true
	}

	minv := 10
	var x, y int
	// 获取最少方案的空格，开始dfs
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if str[i*N+j] == '.' {
				// 如果是空，get到此位置的合法方案，如果这个空格的合法方案数（1的个数）比minv要少，就将其记录。
				state := get(i, j)
				if one[state] < minv {
					minv = one[state]
					x = i
					y = j
				}
			}
		}
	}

	state := get(x, y)
	for i := state; i != 0; i -= lowbit(i) {
		t := mNum[lowbit(i)]
		draw(x, y, t, true)
		if dfs(cnt - 1) {
			return true
		}
		draw(x, y, t, false)
	}

	return false
}

func draw(x, y, t int, isSet bool) {
	// 如果是填数，就填，如果是清空，就清空
	if isSet {
		str[x*N+y] = '1' + byte(t)
	} else {
		str[x*N+y] = '.'
	}

	// 然后需要加/减设置数,因为1表示用过，所以-v将其置为0,表示这个数填过，而加将其变为1，表示没填过
	v := 1 << t
	if !isSet {
		v = -v
	}

	row[x] -= v
	col[y] -= v
	cell[x/3][y/3] -= v
}

// 返回的数的1的位置代表可以填的合法方案
func get(x, y int) int {
	return row[x] & col[y] & cell[x/3][y/3]
}

// lowbit
func lowbit(x int) int {
	return x & (-x)
}

// 初始化为全为1,即都可以填
func ini() {
	for i := 0; i < N; i++ {
		row[i] = M - 1
		col[i] = M - 1
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cell[i][j] = M - 1
		}
	}
}

// 得到某个数1的个数
func getoneNum(x int) (ans int) {
	for x != 0 {
		x -= lowbit(x)
		ans++
	}
	return
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
