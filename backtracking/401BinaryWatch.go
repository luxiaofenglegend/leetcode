package backtracking

import (
	"fmt"
	"strconv"
)

/*
A binary watch has 4 LEDs on the top which represent the hours (0-11), and the 6 LEDs on the bottom represent the minutes (0-59).

Given a non-negative integer n which represents the number of LEDs that are currently on, return all possible times the watch could represent.

Input: n = 1
Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]

The hour must not contain a leading zero, for example "01:00" is not valid, it should be "1:00".
The minute must be consist of two digits and may contain a leading zero, for example "10:2" is not valid, it should be "10:02".
*/
type clock struct {
	hour   int
	minute int
}

func readBinaryWatch(num int) []string {
	bitLen := 4 + 6
	c := clock{0, 0}
	clocks := innerReadBinaryWatch(num, 0, bitLen, c)
	return clocks2Strings(clocks)
}

func innerReadBinaryWatch(num int, index int, bitLen int, c clock) []clock {
	result := make([]clock, 0)
	if num == 0 {
		result = append(result, c)
		return result
	}

	leftBits := bitLen - index
	if leftBits < num {
		return nil
	}
	// vars
	watchRepresent := []int{1, 2, 4, 8, 16, 32, 1, 2, 4, 8}
	minuteMaxIndex := 5 + 1

	if index < minuteMaxIndex {
		curMinute := c.minute + watchRepresent[index]
		if curMinute < 60 {
			c.minute = curMinute
			addedClocks := innerReadBinaryWatch(num-1, index+1, bitLen, c)
			c.minute -= watchRepresent[index]
			result = appendClocks(result, addedClocks)
		}
		// not add
		notAddedClocks := innerReadBinaryWatch(num, index+1, bitLen, c)
		result = appendClocks(result, notAddedClocks)
	} else {
		curHour := c.hour + watchRepresent[index]
		if curHour < 12 {
			c.hour = curHour
			addedClocks := innerReadBinaryWatch(num-1, index+1, bitLen, c)
			c.hour -= watchRepresent[index]
			result = appendClocks(result, addedClocks)
		}
		// not add
		notAddedClocks := innerReadBinaryWatch(num, index+1, bitLen, c)
		result = appendClocks(result, notAddedClocks)
	}

	return result
}

func appendClocks(result []clock, arr []clock) []clock {
	for _, v := range arr {
		result = append(result, v)
	}
	return result
}

func clocks2Strings(clocks []clock) []string {
	result := make([]string, 0)
	for _, v := range clocks {
		result = append(result, clock2String(v))
	}
	return result
}

func clock2String(c clock) string {
	strHour := strconv.Itoa(c.hour)
	strMinute := strconv.Itoa(c.minute)
	if c.minute < 10 {
		strMinute = "0" + strMinute
	}
	return strHour + ":" + strMinute
}

func Test401() {
	myPrint := func(strs []string) {
		for _, v := range strs {
			fmt.Println(v)
		}
	}
	strs := readBinaryWatch(1)
	myPrint(strs)
}
