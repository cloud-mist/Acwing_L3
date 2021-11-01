// 分组背包变形
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 70
const M = 32010

type p struct {
	v, w int
}

var (
	master [N]p
	slaver [N][]p
	f      [M]int
)

func main() {
	defer ot.Flush()

	m, n := rnI(), rnI()

	for i := 1; i <= n; i++ {
		v, w, q := rnI(), rnI(), rnI()
		if q == 0 {
			master[i] = p{v, v * w}
		} else {
			slaver[q] = append(slaver[q], p{v, v * w})
		}
	}

	for i := 1; i <= n; i++ {
		if master[i].v != 0 {
			for j := m; j >= 0; j-- {

				for k := 0; k < 1<<uint(len(slaver[i])); k++ {
					//  所有物品组情况枚举
					v, w := master[i].v, master[i].w
					for u := 0; u < len(slaver[i]); u++ {
						if (k>>uint(u))&1 == 1 {
							v += slaver[i][u].v
							w += slaver[i][u].w
						}
					}

					if j >= v {
						f[j] = max(f[j], f[j-v]+w)
					}
				}
			}
		}

	}

	fmt.Print(f[m])
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
