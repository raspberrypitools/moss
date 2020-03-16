package config


func ClockFace(clockNum string) []int{

	switch clockNum{
	case "colon":
		var num = []int{25, 59}
	case "t_1_0":
		var num = []int{0, 1, 2, 31, 33, 34, 36, 65, 67, 68, 69, 70}
	case "t_1_1":
		var num = []int{1, 32, 35, 66, 69}
	case "t_1_2":
		var num = []int{0, 1, 2, 31, 34, 35, 36, 67, 68, 69, 70}
	case "t_1_3":
		var num = []int{0, 1, 2, 31, 34, 35, 36, 65, 68, 69, 70}
	case "t_1_4":
		var num = []int{0, 2, 31, 33, 34, 35, 36, 65, 70}
	case "t_1_5":
		var num = []int{0, 1, 2, 33, 34, 35, 36, 65, 68, 69, 70}
	case "t_1_6":
		var num = []int{0, 1, 2, 33, 34, 35, 36, 65, 67, 68, 69, 70}
	case "t_1_7":
		var num = []int{0, 1, 2, 31, 36, 65, 70}
	case "t_1_8":
		var num = []int{0, 1, 2, 31, 33, 34, 35, 36, 65, 67, 68, 69, 70}
	case "t_1_9":
		var num = []int{0, 1, 2, 31, 33, 34, 35, 36, 65, 68, 69, 70}
	}
	return num
}
