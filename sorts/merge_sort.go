//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

func merge(a, b []int) (R []int) {
	l, r := 0, 0

        for l < len(a) && r < len(b) {
                if a[l] < b[r] {
                        R = append(R, a[l])
                        l++
                } else {
                        R = append(R, b[r])
                        r++
                }
        }

        R = append(R, b[r:]...)
        R = append(R, a[l:]...)

        return R

}

// func merge(a []int, b []int) []int {
	
// 	r := make([]int, len(a)+len(b))
// 	i, j := 0, 0

// 	for i < len(a) && j < len(b) {

// 		if a[i] <= b[j] {
// 			r[i+j] = a[i]
// 			i++
// 		} else {
// 			r[i+j] = b[j]
// 			j++
// 		}

// 	}

// 	for i < len(a) {
// 		r[i+j] = a[i]
// 		i++
// 	}
// 	for j < len(b) {
// 		r[i+j] = b[j]
// 		j++
// 	}

// 	return r
// }

//Mergesort Perform mergesort on a slice of ints
func Mergesort(items []int) []int {

	if len(items) < 2 {
		return items

	}

	var middle = len(items) / 2
	var a = Mergesort(items[:middle])
	var b = Mergesort(items[middle:])
	return merge(a, b)

}
