package main

import (
	"fmt"
	"strconv"
)

var (
	StartTime = 1546300800
	EndTime = 1570675775
)

func GetBetweenString(firstStr string, secondStr string, style... string) string {
	firstInt, _ := strconv.Atoi(firstStr)
	secondInt, _ := strconv.Atoi(secondStr)
	var newStr string
	if len(style) == 0 {
		for i := firstInt; i <= secondInt; i++ {
			newStr = newStr + strconv.Itoa(i)
		}
	} else {
		newStr = firstStr + style[0] + secondStr
	}

	return "[" + newStr + "]"
}

func main()  {
	var regStrList []string
	var breakIndex int
	var sameStr string

	if StartTime != 0 && EndTime != 0 {
		if StartTime < EndTime {
			startTimeStr := strconv.Itoa(StartTime)
			endTimeStr := strconv.Itoa(EndTime)
			// get same part
			for i := 0; i < len(startTimeStr); i++ {
				if string(startTimeStr[i]) != string(endTimeStr[i]) {
					breakIndex = i
					break
				}
			}
			sameStr = startTimeStr[0:breakIndex]
			startTimeStr = startTimeStr[breakIndex:len(startTimeStr)]
			endTimeStr = endTimeStr[breakIndex:len(endTimeStr)]
			var startRegList []string
			var endRegList []string
			for i := len(startTimeStr) - 1; i > -1; i-- {
				if i == len(startTimeStr) - 1 && len(startTimeStr) != 1 {
					startRegList = append(startRegList, sameStr + startTimeStr[0:i] + GetBetweenString(string(startTimeStr[i]), "9"))
				} else if i == 0 {
					zeroToNine := ""
					for j := 0; j < len(startTimeStr[i+1:len(startTimeStr)]); j++ {
						zeroToNine = zeroToNine + GetBetweenString("0", "9")
					}
					intChar, _ := strconv.Atoi(string(startTimeStr[i]))
					endIntChar, _ := strconv.Atoi(string(endTimeStr[i]))
					if intChar == endIntChar - 1 {
					} else {
						startRegList = append(startRegList, sameStr + GetBetweenString(strconv.Itoa(intChar + 1), strconv.Itoa(endIntChar - 1)) + zeroToNine)
					}
					// deal with the bigger timestamp
					for k := 0; k < len(endTimeStr); k++ {
						if k == len(endTimeStr) - 2 {
							nextCharInt, _ := strconv.Atoi(string(endTimeStr[k+1]))
							endRegList = append(endRegList, sameStr + endTimeStr[0:k+1] + GetBetweenString("0", strconv.Itoa(nextCharInt)))
						} else if k == len(endTimeStr) - 1 {
							break
						} else {
							nextCharInt, _ := strconv.Atoi(string(endTimeStr[k+1]))
							if nextCharInt == 0 {
								continue
							}
							zeroToNine := ""
							for m := 0; m < len(endTimeStr) - (len(endTimeStr[0:k+1]) + 1); m++ {
								zeroToNine = zeroToNine + GetBetweenString("0", "9")
							}
							endRegList = append(endRegList, sameStr + endTimeStr[0:k+1] + GetBetweenString("0", strconv.Itoa(nextCharInt - 1)) + zeroToNine)
						}
					}
				} else {
					zeroToNine := ""
					for j := 0; j < len(startTimeStr[i+1:len(startTimeStr)]); j++ {
						zeroToNine = zeroToNine + GetBetweenString("0", "9")
					}
					intChar, _ := strconv.Atoi(string(startTimeStr[i]))
					startRegList = append(startRegList, sameStr + startTimeStr[0:i] + GetBetweenString(strconv.Itoa(intChar + 1), "9") + zeroToNine)
				}
			}
			regStrList = append(startRegList, endRegList...)
			fmt.Println(regStrList)
		} else {
			fmt.Println("Error")
		}
	} else {
		fmt.Println("Error")
	}
}
