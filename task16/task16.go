package main

import (
	"fmt"
	"math/rand"
)

func qsort(v *[]int) {
	if len(*v) == 0 { // если срез пустой, то выходим из функции
		return
	}

	if len(*v) == 1 { // то же самое, как и в случае нулевого массива
		return
	}

	if len(*v) == 2 { // если срез имеет длину 2, то сортируем его по возрастанию и выходим
		if (*v)[0] > (*v)[1] {
			(*v)[0], (*v)[1] = (*v)[1], (*v)[0]
		}
		return
	}

	pivot := (*v)[len(*v)/2-1] //опорный элемент

	v1 := make([]int, 0) //меньше опорного элемента
	v2 := make([]int, 0) //больше опорного элемента
	v3 := make([]int, 0) //равно опорному элементу

	for i := range *v {
		if (*v)[i] < pivot {
			v1 = append(v1, (*v)[i])
		} else if (*v)[i] > pivot {
			v2 = append(v2, (*v)[i])
		} else {
			v3 = append(v3, (*v)[i])
		}
	}

	qsort(&v1) //вызываем сортировку рекурсивно...
	qsort(&v2)

	//собираем срез образно
	for i := range v1 {
		(*v)[i] = v1[i]
	}

	for i := range v3 {
		(*v)[i+len(v1)] = pivot
	}

	for i := range v2 {
		(*v)[i+len(v1)+len(v3)] = v2[i]
	}
}

const n = 10

func main() {

	sl := make([]int, n)

	for i := range sl {
		sl[i] = rand.Intn(n)
	}

	fmt.Println(sl)

	qsort(&sl)

	fmt.Println(sl)
}
