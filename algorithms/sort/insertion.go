/**
 * @Author: wangyingwen
 * @Description:插入排序算法 对于少量元素的排序 是一种有效的算法
 * @File:  insertion
 * @Version: 1.0.0
 * @Date: 2019-12-19 15:48
 */
// https://github.com/abxd39/study/doc/images/insertion.jpeg
package sort

import "log"

func (b *Algorithms) Insertion(l []int) {
	//打牌的比喻很好。
	for i := 2; i < len(l); i++ {
		key := l[i]
		j := i - 1
		for j > 0 && l[j] > key {
			l[j+1] = l[j]
			j = j - 1
		}
		l[i] = key
	}
	log.Println("insertion ", l)
}
