package base36

import (
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"
)

var Reader io.Reader

const dash byte = '-'

// Read is a helper function that calls Reader.Read using io.ReadFull.
// On return, n == len(b) if and only if err == nil.
func Read(b []byte) (n int, err error) {
	return io.ReadFull(Reader, b)
}

type UUID [16]byte

// SetVersion sets version bits.
func (u *UUID) SetVersion(v byte) {
	u[6] = (u[6] & 0x0f) | (v << 4)
}

// SetVariant sets variant bits as described in RFC 4122.
func (u *UUID) SetVariant() {
	u[8] = (u[8] & 0xbf) | 0x80
}

func (u UUID) String() string {
	buf := make([]byte, 36)

	hex.Encode(buf[0:8], u[0:4])
	buf[8] = dash
	hex.Encode(buf[9:13], u[4:6])
	buf[13] = dash
	hex.Encode(buf[14:18], u[6:8])
	buf[18] = dash
	hex.Encode(buf[19:23], u[8:10])
	buf[23] = dash
	hex.Encode(buf[24:], u[10:])

	return string(buf)
}

func NewV4() UUID {
	u := UUID{}
	safeRandom(u[:])
	u.SetVersion(4)
	u.SetVariant()

	return u
}

func safeRandom(dest []byte) {
	fmt.Print(hex.Dump(dest))
	if _, err := Read(dest); err != nil {
		panic(err)
	}
	fmt.Print(hex.Dump(dest))
}

func twoSum(nums []int, target int) []int {
	fmt.Println(len(nums))
	for i := 0; i < len(nums)-1; i++ {
		for j := len(nums) - 1; j > i; j-- {
			fmt.Println(i, j)
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	fmt.Println("sfasf")
	for i := 0; i < len(nums)-1; i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] > nums[j] {
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			}
		}
	}
	l := 0
	r := len(nums) - 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[l]+nums[r] > target {
			r--
		}
		if nums[l]+nums[r] < target {
			l++
		}
	}
	return []int{l, r}
}
func AmountConvert(p_money float64, p_round bool) string {
	i := len(strconv.Itoa(112121))
	fmt.Println(i)
	var NumberUpper = []string{"壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖", "零"}
	var Unit = []string{"分", "角", "圆", "拾", "佰", "仟", "万", "拾", "佰", "仟", "亿", "拾", "佰", "仟"}
	var regex = [][]string{
		{"零拾", "零"}, {"零佰", "零"}, {"零仟", "零"}, {"零零零", "零"}, {"零零", "零"},
		{"零角零分", "整"}, {"零分", "整"}, {"零角", "零"}, {"零亿零万零元", "亿元"},
		{"亿零万零元", "亿元"}, {"零亿零万", "亿"}, {"零万零元", "万元"}, {"万零元", "万元"},
		{"零亿", "亿"}, {"零万", "万"}, {"拾零圆", "拾元"}, {"零圆", "元"}, {"零零", "零"}}
	str, DigitUpper, Unit_Len, round := "", "", 0, 0
	if (p_money == 0) {
		return "零"
	}
	if (p_money < 0) {
		str = "负";
		p_money = math.Abs(p_money)
	}
	if (p_round) {
		round = 2
	} else {
		round = 1
	}
	Digit_byte := []byte(strconv.FormatFloat(p_money, 'f', round+1, 64)) //注意币种四舍五入
	fmt.Println(Digit_byte)
	Unit_Len = len(Digit_byte) - round

	for _, v := range (Digit_byte) {
		if (Unit_Len >= 1 && v != 46) {
			s, _ := strconv.ParseInt(string(v), 10, 0)
			fmt.Println(s)
			if (s != 0) {
				DigitUpper = NumberUpper[s-1]

			} else {
				DigitUpper = "零"
			}
			str = str + DigitUpper + Unit[Unit_Len-1]
			Unit_Len = Unit_Len - 1
		}
	}
	for i, _ := range (regex) {
		reg := regexp.MustCompile(regex[i][0])
		str = reg.ReplaceAllString(str, regex[i][1])
	}
	if (string(str[0:3]) == "元") {
		str = string(str[3:len(str)])
	}
	if (string(str[0:3]) == "零") {
		str = string(str[3:len(str)])
	}
	return str
}
