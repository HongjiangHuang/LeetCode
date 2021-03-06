package easy

import "testing"

func Test_RemoveDuplicates(t *testing.T) {
	arr := []int{0,0,1,1,1,2,2,3,3,4}
	len := RemoveDuplicates(arr)
	if len != 5 {
		t.Error("长度返回错误")
	}

	for k,v := range arr {
		if (k < len && v != arr[k]) {
			t.Error("数组处理错误")
		}
	}

	t.Log("长度返回正常")
}

func Test_maxProfit(t *testing.T) {
	arr := []int{7,1,5,3,6,4}
	profit := MaxProfit(arr)

	if profit != 7 {
		t.Error("解答错误")
	}
}