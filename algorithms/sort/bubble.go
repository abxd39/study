package sort

import "log"

//冒泡排序
type Algorithms struct {
}

func (b *Algorithms) Version1(l []int) {
	run := 0
	log.Println("pre", l)
	length := len(l)
	for c := 0; c < len(l); c++ {
		for i, v := range l {
			run++
			//		log.Println("执行次数=", run)
			if i < len(l[:length])-1 {
				j := i + 1
				jv := l[j]
				if v > jv {
					l[j], l[i] = v, jv
				}
			}
		}
	}

	log.Println("sub", l)
}
