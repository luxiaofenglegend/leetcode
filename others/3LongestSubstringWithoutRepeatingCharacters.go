package others

import (
	"bytes"
	"fmt"
)

/*
Given a string, find the length of the longest substring without repeating characters.
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

*/

func lengthOfLongestSubstring(s string) int {
	//const a = int32('a')
	longest := ""
	for idx, c := range s {
		buf := bytes.Buffer{}
		buf.WriteByte(byte(c))
		//used := make([]bool, 26)
		used := make(map[int32]bool)
		used[c] = true
		for _, ch := range s[idx+1:] {
			_, ok := used[ch]
			if ok {
				break
			} else {
				used[ch] = true
				buf.WriteByte(byte(ch))
			}
		}
		// if longer
		tmpStr := buf.String()
		if len(tmpStr) > len(longest) {
			longest = tmpStr
		}
	}
	return len(longest)
}

func Test3() {
	testCase1 := "pwwkew"
	testCase2 := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(testCase1))
	fmt.Println(lengthOfLongestSubstring(testCase2))
}
