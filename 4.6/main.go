package main

import (
	"reflect"
	"unsafe"
	"fmt"
)

const (
	Len = 1000
)

type stack struct {
	top    int
	values []interface{}
}

func newStack() *stack {
	return &stack{
		top:    -1,
		values: make([]interface{}, 0, Len),
	}
}

func (s *stack) push(key interface{}) {
	s.top++
	s.values = append(s.values, key)
}

func (s *stack) pop() {
	s.values = s.values[:s.top]
	s.top--
	return
}

func (s *stack) topValue() interface{} {
	return s.values[s.top]
}

func (s *stack) size() int {
	return s.top + 1
}

func byteTostring(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

//first
type pair struct {
	first  int
	second int
}

func main() {
	order := `\\///\_/\/\\\\/_/\\///__\\\_\\/_\/_/\`

	sum := 0
	s1 := newStack()
	s2 := newStack()

	for k, v := range order {
		v_str := string([]byte{byte(v)})

		if v_str == `\` {
			//压栈
			s1.push(k)
		} else if (v_str == `/` && s1.size() > 0) {
			//取出栈顶元素
			j := s1.topValue().(int)
			//出栈
			s1.pop()
			//总面积相加
			sum += (k - j)

			//记录一段\/的面积
			a := k - j

			//把s2中满足是一块区域的存储面积加上
			for s2.size() > 0 && s2.topValue().(*pair).first > j {
				a += s2.topValue().(*pair).second
				s2.pop()
			}
			s2.push(&pair{j, a})
		}
	}

	s3 := newStack()
	for s2.size() > 0 {
		s3.push(s2.topValue().(*pair).second)
		s2.pop()
	}

	fmt.Println("sum ", sum)
	for i := s3.top; i > -1; i-- {
		fmt.Print(s3.values[i])
		fmt.Print("  ")
	}

}
