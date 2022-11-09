package modsort

import (
	"sort"
)

func ModuloSort(arr []string) (res []string) {
	a := [3][]string{}
	for _, v := range arr {
		a[len(v)%3] = append(a[len(v)%3], v)
	}
	for i := range a {
		sort.Strings(a[i])
		res = append(res, a[i]...)
	}
	return
}
