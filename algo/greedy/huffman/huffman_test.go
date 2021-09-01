package huffman

import (
	"encoding/hex"
	"fmt"
	strconv "strconv"
	"testing"
)

func TestNewHuffman(t *testing.T) {
	//var str = "BBC News is an operational business division of the British Broadcasting Corporation responsible for the gathering and broadcasting of news and current affairs."
	var str = "operaaaaaaaaaaaaatioooooooooooonalbusiness"
	var hf = NewHuffman(str)
	//hf.tree.For2(hf.tree.Height,64)
	m := make(map[int32]string)
	//hf.tree.Print()
	ret := hf.tree.Show(m, "")
	fmt.Println(ret)
	new_str := ""
	for _, c := range str {
		new_str += ret[c]
	}
	fmt.Println(new_str)
	x := bitString(new_str)
	fmt.Println(len(x))
	fmt.Println(len(x.AsByteSlice()))
	fmt.Printf("%b",x.AsByteSlice())
}


type bitString string

func (b bitString) AsByteSlice() []byte {
	var out []byte
	var str string

	for i := len(b); i > 0; i -= 8 {
		if i-8 < 0 {
			str = string(b[0:i])
		} else {
			str = string(b[i-8 : i])
		}
		v, err := strconv.ParseUint(str, 2, 8)
		if err != nil {
			panic(err)
		}
		out = append([]byte{byte(v)}, out...)
	}
	return out
}

func (b bitString) AsHexSlice() []string {
	var out []string
	byteSlice := b.AsByteSlice()
	for _, b := range byteSlice {
		out = append(out, "0x" + hex.EncodeToString([]byte{b}))
	}
	return out
}


func TestGetHeight(t *testing.T) {
	x := "sdfasfd"
	fmt.Println(x[7:])
}
func TestGetMinTree(t *testing.T) {
	x := []byte("00010101011")
	fmt.Println(len(x))
}
