package bitmap

import (
	"fmt"
	"testing"
)


func TestAddNum(t *testing.T) {

	bitmap:= NewBitMap()
	bitmap.Add(97128312)
	bitmap.Add(97128315)
	bitmap.Add(97128319)
	bitmap.Add(97128313)
	re:= bitmap.Has(971283122)
	fmt.Println(re)
	//AddNum(19797128312)(
	//AddNum(19797128313)
	//AddNum(19797128315)
	//AddNum(19797128317)
	//AddNum(19797128319)
}
