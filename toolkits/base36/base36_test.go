package base36

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"strings"
	"testing"
)

var raw []uint64 = []uint64{0, 50, 100, 999, 1000, 1111, 5959, 99999,
	123456789, 5481594952936519619, math.MaxInt64 / 2048, math.MaxInt64 / 512,
	math.MaxInt64, math.MaxUint64}
var num []string = []string{"JNU5AE"}

var encoded []string = []string{"", "1E", "2S", "RR", "RS", "UV", "4LJ", "255R",
	"21I3V9", "15N9Z8L3AU4EB", "18CE53UN18F", "4XDKKFEK4XR",
	"1Y2P0IJ32E8E7", "3W5E11264SGSF"}

func TestDecodeToBytes(t *testing.T) {
	const (
		NFC      = "0"
		IPT      = "0"
		BLE      = "1"
		WIFI     = "0"
		Reserved = "0000"
		version  = "000"
		cat      = 7
	)
	catb := fmt.Sprintf("%08b", cat)
	setupcodeb := fmt.Sprintf("%027b", 51808582)
	//binx := version + Reserved   + catb  + WIFI + BLE + IPT + NFC + setupcodeb
	fmt.Println(len(setupcodeb))
	bin := version + Reserved + catb + WIFI + BLE + IPT + NFC + setupcodeb
	fmt.Println(bin)
	ui, err := strconv.ParseUint(bin, 2, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	xxx := Encode(ui)

	fmt.Printf("%09s\n", xxx)
	for i, _ := range num {
		out := Decode(num[i])
		fmt.Println(ui)
		fmt.Println(out)
		//assert.Equal(t, encoded[i], encoded(num[i]))
	}
}
func TestEncode(t *testing.T) {

	for i, v := range raw {
		assert.Equal(t, encoded[i], Encode(v))
	}
}

func TestDecode(t *testing.T) {

	for i, v := range encoded {
		assert.Equal(t, raw[i], Decode(v))
		assert.Equal(t, raw[i], Decode(strings.ToLower(v)))
	}
}

func BenchmarkEncode(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Encode(5481594952936519619)
	}
}

func TestGenUUID(t *testing.T) {
	x := uuid.NewV4()
	//x.SetVariant()
	fmt.Println(x.String())
}

func BenchmarkDecode(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Decode("1Y2P0IJ32E8E7")
	}
}

func TestTwo(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	result := twoSum(nums, target)
	fmt.Println(result)
}
func TestAddTwoNumbers(t *testing.T) {
	//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	//Output: 7 -> 0 -> 8
	//Explanation: 342 + 465 = 807.
	//a := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	a := []int{5}
	//b := []int{5, 6, 4}
	b := []int{5}
	var l1, l2 *ListNode
	for i := len(a) - 1; i >= 0; i-- {
		x := &ListNode{a[i], nil}
		if l1 == nil {
			l1 = x
		} else {
			o := l1
			x.Next = o
			l1 = x
		}
	}
	for i := len(b) - 1; i >= 0; i-- {
		x := &ListNode{b[i], nil}
		if l2 == nil {
			l2 = x
		} else {
			o := l2
			x.Next = o
			l2 = x
		}
	}

	result := addTwoNumbers(l1, l2)
	fmt.Println(result)
}

func TestAmountConvert(t *testing.T) {
	a := AmountConvert(112312.12, true)
	fmt.Println(a)
}
