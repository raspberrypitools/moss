package config


import (
	"sort"
	"github.com/wangbokun/go/log"
)


func ClockFaces(displacement int, clockNum string) []int{

	var t_0 = []int{0, 1, 2, 31, -1, 33, 34, -1, 36, 65, -1, 67, 68, 69, 70}
	var t_1 = []int{-1, 1, -1, -1, 32, -1, -1, 35, -1, -1, 66, -1, -1, 69, -1}
	var t_2 = []int{0, 1, 2, -1, -1, 31, 34, 35, 36, 67, -1, -1, 68, 69, 70}
	var t_3 = []int{0, 1, 2, -1, -1,  31, 34, 35, 36, -1, -1, 65, 68, 69, 70}
	var t_4 = []int{0, -1, 2, 31, 33, -1, 34, 35, 36, -1, -1, 65, -1, -1, 70}
	var t_5 = []int{0, 1, 2, 33, -1, -1, 34, 35, 36, -1, -1, 65, 68, 69, 70}
	var t_6 = []int{0, 1, 2, 33, -1, -1, 34, 35, 36, 65, 67, -1, 68, 69, 70}
	var t_7 = []int{0, 1, 2, -1, -1, 31, -1, -1, 36, -1, -1, 65, -1, -1, 70}
	var t_8 = []int{0, 1, 2, 31, 33, -1, 34, 35, 36, 65, 67, -1, 68, 69, 70}
	var t_9 = []int{0, 1, 2, 31, 33, -1, 34, 35, 36, -1, -1, 65, 68, 69, 70}

	switch clockNum{
	case "t_0":
		return transform(displacement, t_0)
	case "t_1":
		return transform(displacement, t_1)
	case "t_2":
		return transform(displacement, t_2)
	case "t_3":
		return transform(displacement, t_3)
	case "t_4":
		return transform(displacement, t_4)
	case "t_5":
		return transform(displacement, t_5)
	case "t_6":
		return transform(displacement, t_6)
	case "t_7":
		return transform(displacement, t_7)
	case "t_8":
		return transform(displacement, t_8)
	case "t_9":
		return  transform(displacement, t_9)
	default:
		return  nil
	}
}



func transform(displacement int, n []int) []int{
	switch displacement{
	// 0 
	case 1:
		return disNum(0,n)
	// 4 
	case 2:
		return disNum(4,n)
	// 10 
	case 3:
		return disNum(10,n)
	// 14
	case 4:
		return disNum(14,n)
	}
	return nil
}


func disNum(d int,n []int) []int{

	log.Error("%d ======= %#v",d,n)
	if d == 0 {
		sort.Ints(DeleteSlice(n)[:])
		return  n
	}

	var ctx []int

	for k,v := range  n{
		if v != -1{
			switch  k{
			case  3:
				ctx = append(ctx,v - d)
			case  4:
				ctx = append(ctx,v - d)
			case  5:
				ctx = append(ctx,v - d)
			case  9:
				ctx = append(ctx,v - d)
			case  10:
				ctx = append(ctx,v - d)
			case  11:
				ctx = append(ctx,v - d)
			default:
				ctx = append(ctx,v + d)
			}
		}
	}

	sort.Ints(DeleteSlice(ctx)[:])
	return  ctx
}


func DeleteSlice(a []int) []int{
	for i := 0; i < len(a); i++ {
		if a[i] == -1 {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	return a
}
