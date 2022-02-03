package algorithms 

import (
	"github.com/sarpdag/boyermoore"
)

func SearchString(src string , subStr string) *int {
	table := boyermoore.CalculateSlideTable(subStr)
	if pos := boyermoore.IndexWithTable(&table, src, subStr); pos > -1 {
		return &pos;
	}
	return nil;
}

func SearchStrings(src []string , subStr string) ([]int  , bool){
	indices := make([]int , 0);
	for i := 0;i < len(src);i++ {
		switch SearchString(src[i] , subStr) {
		case nil:
			continue;
		default:
			indices = append(indices, i);
		}
	}

	if len(indices) == 0 {
		return indices , false;
	} else {
		return indices , true;
	}
}