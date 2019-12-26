/**
 * @Author: wangyingwen
 * @Description:插入排序算法 对于少量元素的排序 是一种有效的算法
 * @File:  insertion
 * @Version: 1.0.0
 * @Date: 2019-12-19 15:48
 */
// https://github.com/abxd39/study/doc/images/insertion.jpeg
package sort

import (
	"log"
)

func (a *Algorithms) Insertion(l []int) {
	//打牌的比喻很好。
	log.Println(l)
	for j := 2; j < len(l); j++ {
		key := l[j]
		i := j - 1
		for i >= 0 && l[i] > key {
			l[i+1] = l[i]
			i = i - 1
		}
		l[i+1] = key
		log.Printf("set %d %v", j, l)
	}
	log.Println("insertion ", l)
}
