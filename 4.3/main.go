package main

import "fmt"

/*
问题描述 循环调度算法 名称为name_i的且处理时间为time_i的任务按顺序排成一列，共cpu依次处理，cpu一次处理的一个任务最多时长为q，若q时间内处理任务完毕，则在队列中删除任务
否则，任务放置在队列最后，共下次循环调度处理
 */

const (
	Len = 100
)

type pp struct {
	name string
	t    int
}

var Q [Len]pp

var head, tail, n int

func enqueue(x pp) {
	Q[tail] = x

	//对尾后移一位 超过Len之后，取余数
	tail = (tail + 1) % Len
}

func dequeue() pp {
	rel := Q[head]

	//对头后移一位
	head = (head + 1) % Len
	return rel
}

func main() {

	var works = [5]pp{
		{
			"p1",
			150,
		},
		{
			"p2",
			80,
		},
		{
			"p3",
			200,
		},
		{
			"p4",
			350,
		},
		{
			"p5",
			20,
		},
	}

	for i := 0; i < len(works); i++ {
		Q[i] = works[i]
	}

	//记录经过的时间
	var elaps int = 0

	//时间片
	var q = 100

	f_min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var c int
	n = 5

	head = 0
	tail = n

	var u pp

	for head != tail {
		u = dequeue()
		c = f_min(q, u.t)
		//任务剩余时间
		u.t -= c
		//累计已过时间
		elaps += c
		if (u.t > 0) {
			enqueue(u)
		} else {
			fmt.Println(u.name, "is over at ", elaps)
		}
	}
	return

}
