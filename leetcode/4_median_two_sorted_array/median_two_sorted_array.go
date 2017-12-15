package median_two_sorted_array


func MedianTwoSortedArray(l1, l2 []int) float32 {
	switch v :=(len(l1) + len(l2)) % 2; v {
	case 1:
		i, r := FindIndexTwoSortedArray(l1, l2, (len(l1) + len(l2) + 1) / 2)
		if i  == 1 {
			return float32(l1[r])
		} else {
			return float32(l2[r])
		}
	default:
		ret := float32(0.0)
		i, r := FindIndexTwoSortedArray(l1, l2, (len(l1) + len(l2)) / 2)
		if i == 1 {
			ret = float32(l1[r])
		} else {
			ret = float32(l2[r])
		}
		i, r = FindIndexTwoSortedArray(l1, l2, ((len(l1) + len(l2)) / 2) + 1)
		if i == 1 {
			return  (ret + float32(l1[r])) / 2
		} else {
			return (ret + float32(l2[r])) / 2
		}
	}

}

func FindIndexTwoSortedArray(l1, l2 []int, pos int) (int, int) {
	if len(l1)==0 {
		if pos > len(l2) {
			panic("pos > len(l2)")
		}
		return 2, pos - 1
	}
	if len(l2)==0 {
		if pos > len(l1) {
			panic("pos > len(l1)")
		}
		return 1, pos - 1
	}

	mid := len(l1) / 2
	ret := findPos2(l2, l1[mid])
	v := mid + ret + 1
	switch {
	case v <  pos:
		i, p := FindIndexTwoSortedArray(l1[mid+1 : len(l1)], l2[ret : len(l2)], pos - mid - ret -1)
		if i == 1 {
			return i, p + mid
		} else {
			return i, p + ret
		}
	case v > pos:
		return FindIndexTwoSortedArray(l1[0: mid], l2[0: ret], pos)
	default:
		return 1, mid
	}
}

//func FindPosTwoSortedArray(l1, l2 []int, pos int) int, int {
//	start1, start2, end1, end2 := 0, 0, len(l1) - 1, len(l2) - 1
//	pos1, pos2 := 0, 0
//	ret1 : = 0
//	for {
//		if start1 <= end1 {
//			pos1 = (start1 + end1) / 2
//			pos2  = findPos(l2, start2, end2, l1[pos1])
//			ret1 = 0
//		} else {
//			pos2 = (start2 + end2) / 2
//			pos2 = findPos(l2, start2, end2, l2[pos2])
//			ret1 = 1
//		}
//		if pos1 + pos2 + 1 == pos {
//			if ret1 {
//				return ret1, pos2
//			} else {
//				return ret1, pos1
//			}
//		} else if pos1 + pos2 + 1 < pos {
//			if ret1 {
//				start1 = pos1 + 1
//				start2 = pos2
//			} else {
//				start1 = pos1
//				start2 = pos2 +1
//			}
//		} else {
//			if ret1 {	
//				end1 = pos1 - 1
//				end2 = pos2
//			} else {
//				end1 = pos1
//				ende = pos2 - 1
//			}
//		}
//	}
//}


func findPos(l []int, start int, end int, value int) int {
	var mid int = -1

	for start <= end {
		mid = (start + end) / 2
		if l[mid] <= value {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}

func findPos2(l []int, value int) int {
	start, end := 0, len(l) - 1
	var mid int = 0

	for start <= end {
		mid = (start + end) / 2
		if l[mid] <= value {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}
