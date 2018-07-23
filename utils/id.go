package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// GeneratorID is
func GeneratorID() string {
	fmt.Println("S")
	// year, month, _ := time.Now().Date()
	timeStr := time.Now().Format("20060102150405")
	year := time.Now().Year()
	month := time.Now().Month() //time.Now().Month().String()
	day := time.Now().Day()
	fmt.Println(year, month)
	fmt.Println(day)
	fmt.Println(timeStr)

	store := strconv.Itoa(rand.Intn(9999))
	o := rand.New(rand.NewSource(time.Now().UnixNano())).Int63()
	s := strconv.FormatInt(o, 10)
	fmt.Println(store)
	fmt.Println(show_substr(s, 2))
	fmt.Println(o)
	numStr := timeStr + store
	fmt.Println(numStr)

	return numStr
}

// GetDaysInYearByThisYear is 年中的第几天
func GetDaysInYearByThisYear() int {
	now := time.Now()
	total := 0
	arr := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	y, month, d := now.Date()
	m := int(month)
	for i := 0; i < m-1; i++ {
		total = total + arr[i]
	}
	if (y%400 == 0 || (y%4 == 0 && y%100 != 0)) && m > 2 {
		total = total + d + 1

	} else {
		total = total + d
	}
	return total
}

// MakeYearDaysRand is 06123xxxxx 生成单号
//sum 最少10位,sum 表示全部单号位数
func MakeYearDaysRand(sum int) string {
	//年
	strs := time.Now().Format("06")
	//一年中的第几天
	days := strconv.Itoa(GetDaysInYearByThisYear())
	count := len(days)
	if count < 3 {
		//重复字符0
		days = strings.Repeat("0", 3-count) + days
	}
	//组合
	strs += days
	//剩余随机数
	sum = sum - 5
	if sum < 1 {
		sum = 5
	}
	//0~9999999的随机数
	// pow := math.Pow(10, float64(sum)) - 1
	//fmt.Println("sum=>", sum)
	//fmt.Println("pow=>", pow)
	result := strconv.Itoa(rand.Intn(99999))
	count = len(result)
	//fmt.Println("result=>", result)
	if count < sum {
		//重复字符0
		result = strings.Repeat("0", sum-count) + result
	}
	//组合
	strs += result
	return strs
}

// showSubstr is
func showSubstr(s string, l int) string {
	if len(s) <= l {
		return s
	}
	ss, sl, rl, rs := "", 0, 0, []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			rl = 1
		} else {
			rl = 2
		}
		if sl+rl > l {
			break
		}
		sl += rl
		ss += string(r)
	}
	return ss
}
