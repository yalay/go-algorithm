package main

import (
	"flag"
	"fmt"
	"strings"
)

// 考题反序检测调整算法，支持任意最大数，不限于10
/*
   请用阿拉伯数字填空,使下面十句话都正确.
   0在这十句话里出现的次数是_____；
   1在这十句话里出现的次数是_____；
   2在这十句话里出现的次数是_____；
   3在这十句话里出现的次数是_____；
   4在这十句话里出现的次数是_____；
   5在这十句话里出现的次数是_____；
   6在这十句话里出现的次数是_____；
   7在这十句话里出现的次数是_____；
   8在这十句话里出现的次数是_____；
   9在这十句话里出现的次数是_____；
*/
const (
	maxRetryTimes = 200
)

var (
	num int = 10
)

type BuffData []int

func init() {
	flag.IntVar(&num, "num", 10, "num:0-100")
	flag.Parse()
}

func NewBuff() BuffData {
	buff := make(BuffData, num)
	for i := 0; i < num; i++ {
		buff[i] = 1
	}
	return buff
}

func (buff BuffData) String() string {
	result := ""
	for i, value := range buff {
		result += fmt.Sprintf("%d-%d\n", i, value)
	}
	return result
}

func (buff BuffData) Adjust() {
	for i := num - 1; i >= 0; i-- {
		curIndex := fmt.Sprintf("%d", i)
		buff[i] = strings.Count(buff.String(), curIndex)
	}
}

func (buff BuffData) IsOver() bool {
	curString := buff.String()
	for i, value := range buff {
		curIndex := fmt.Sprintf("%d", i)
		if strings.Count(curString, curIndex) == value {
			continue
		}
		return false
	}
	return true
}

func main() {
	buffData := NewBuff()
	for i := 0; i < maxRetryTimes; i++ {
		if !buffData.IsOver() {
			buffData.Adjust()
		} else {
			fmt.Printf("result:\n%s\n", buffData.String())
			return
		}
	}
	fmt.Printf("no result\n")
}
