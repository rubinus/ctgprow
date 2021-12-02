//找到数组中的中位数 6 / 3 = 2
//[10,12,15,20,30,31]

package codecov

import (
	"fmt"
	"sort"
	"strings"
)

const (
	a = iota
	b
	c
	d
)

type st struct {
	A string
}

func ReverseWords2(s string) string {
	source := strings.Fields(s)
	fmt.Println(source)
	an := len(source)
	am := an / 2
	for i := 0; i < am; i++ {
		source[i], source[an-1-i] = source[an-1-i], source[i]
	}
	return strings.Join(source, " ")
}
func ReverseWords(s string) string {
	source := strings.Split(s, " ")
	an := len(source)
	am := an / 2
	for i := 0; i < am; i++ {
		left := reverse(source[i])
		right := reverse(source[an-1-i])
		source[i], source[an-1-i] = right, left
	}

	return strings.Join(source, " ")
}

func reverse(source string) string {
	b := []rune(source) //中文
	n := len(b)
	mid := n / 2
	for i := 0; i < mid; i++ {
		b[i], b[n-1-i] = b[n-1-i], b[i]
	}
	return string(b)
}

func Append(ss []int) {
	ss = append(ss, 5)
}

func boyun(arr []int) int {
	sum := 0
	l := len(arr)
	if l == 0 {
		return 0
	}
	if l < 2 {
		return arr[0]
	}
	for _, v := range arr {
		sum += v
	}
	mid := sum / l
	fmt.Printf("sum=%d, len=%d,mid=%d", sum, l, mid)

	//排序
	sort.Ints(arr)

	var result int
	for i := 0; i < l/2; i++ {
		//头，尾对撞，快速查找

		//处理左半部分
		left := arr[i]
		right := arr[i+1]
		result_left := search(left, mid, right)

		//右半部分
		right_left := arr[l-2]
		right_right := arr[l-1]
		result_right := search(right_left, mid, right_right)

		if result_left < result_right {
			result = result_left
		} else {
			result = result_right
		}

	}

	return result
}

func search(left, mid, right int) (result int) {
	if left < mid && mid < right {
		//[10,12,15,20,30,31]
		ll := result - left  //1
		rr := right - result //2
		//找到
		if ll < rr {
			return left
		} else {
			return right
		}
	}
	return
}

func maxChildArr(arr []int, count int) (int, []int) {

	// l := len(arr)
	// mid := l / 2
	if count <= 0 {
		return 0, nil
	}

	max, begin, k := 1, 0, 1
	for idx, _ := range arr {
		if idx > 0 && arr[idx] > arr[idx-1] {
			k++
		} else {
			k = 1
		}
		if k > max {
			max = k

			begin = idx - max + 1
		}
	}
	return max, arr[begin : begin+max]
}
