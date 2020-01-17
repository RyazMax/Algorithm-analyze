package sorts

func min(a, b int) (res int) {
	res = a
	if b < a {
		res = b
	}
	return
}

// InsertSort sortes int array with insertion sort
func InsertSort(a []int) {
	for i := 0; i < len(a); i++ {
		toInsert := a[i]
		j := i
		for ; j > 0 && toInsert < a[j-1]; j-- {
			a[j] = a[j-1]
		}
		a[j] = toInsert
	}
}

func merge(dst, src []int, lb, le, rb, re int) {
	for i := lb; lb < le || rb < re; i++ {
		if lb >= le {
			dst[i] = src[rb]
			rb++
		} else if rb >= re {
			dst[i] = src[lb]
			lb++
		} else if src[rb] < src[lb] {
			dst[i] = src[rb]
			rb++
		} else {
			dst[i] = src[lb]
			lb++
		}
	}
}

// MergeSort sortes int array with non-recursive merge sort
func MergeSort(a []int) {
	tmp := make([]int, len(a))
	for k := 1; k < len(a); k <<= 1 {
		for i := 0; i < len(a)-k+1; i += k << 1 {
			merge(tmp, a, i, i+k, i+k, min(i+k<<1, len(a)))
		}
		copy(a, tmp)
	}
}

// RadixSort sortes int array of ONLY POSITIVE numbers with LSD method
func RadixSort(a []int) {
	digits := make([]int, 4096) // 65536
	anses := make([]int, len(a))
	mask := 0xFFF // FFFF
	shift := uint(12)
	count := uint(0)

	for mask > 0 {
		// Обнуляем количество чисел
		for i := range digits {
			digits[i] = 0
		}
		// Подсчитываем число чисел
		for _, val := range a {
			digits[val&mask>>(count*shift)]++
		}
		// Пересчитываем границы диапозона
		for i := range digits {
			if i > 0 {
				digits[i] += digits[i-1]
			}
		}
		// Упорядочивем значения
		for i := range a {
			anses[digits[a[len(a)-i-1]&mask>>(count*shift)]-1] = a[len(a)-i-1]
			digits[a[len(a)-i-1]&mask>>(count*shift)]--
		}
		a, anses = anses, a
		mask <<= shift
		count++
	}
}
