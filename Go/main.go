package main

import (
	"fmt"
	"sort"
	// "strconv"
	// "strings"
	"time"
	//"regexp"
	"unicode/utf8"
	//"testing/quick"
	"regexp"
	"strings"
	//"io"
	"strconv"
)

func main(){
	paths := [][]string{
		{"A","Z"},
		// {"D","B"},
		// {"C", "A"},
	}
	fmt.Println(destCity(paths))
}

func destCity(paths [][]string) string {
	//var citySquence []string
	var sliceRight [] string
	var sliceLeft [] string
	
	for _, i := range paths{
		sliceRight = append(sliceRight, i[0])
		sliceLeft = append(sliceLeft, i[1])
	}

	rightText := strings.Join(sliceRight, "")
	
	for _, i := range sliceLeft{
		if strings.Count(rightText, i) == 0{
			return i
		}
	}

	return ""
}


func arrayRankTransform(arr []int) {
	ans := []int{}
	
	sortedNum := make([]int, len(arr))
	copy(sortedNum, arr)
	sort.Ints(sortedNum)

	rankMap := make(map[int]int)
	for i, num := range sortedNum{
		if i > 0 && sortedNum[i] == sortedNum[i - 1]{
			rankMap[num] = i - 1
			continue
		}

		rankMap[num] = i + 1
	}

	for _, i := range arr{
		ans = append(ans, rankMap[i])
	}

	fmt.Println(ans)
}

func truncateSentence(s string, k int) string {
	sList := strings.Split(s, " ")
	ansList := []string{}

	for i := 0; k > i; i++{
		ansList = append(ansList, sList[i])
	}

	ans := strings.Join(ansList, " ")

	return ans
}

func maxSubArray(nums []int) int {
	pos := 0
	length := 0
	tmpNum := nums[pos]
	ans := 0

	for true{
		if length == len(nums) - 1 && pos == len(nums) - 1{
			break
		}

		if nums[pos] > 0{
			
		}

		if length == len(nums) - 1{
			pos += 1
			length = 0
			tmpNum = nums[pos]
			continue
		}

		if nums[pos] == nums[length]{
			length += 1
			continue
		}

		tmpNum += nums[length]
		if tmpNum > ans{
			ans = tmpNum
			fmt.Println(ans)
		}

		length += 1
	}
	return ans
}


func canMake(arr []int) bool {
	type firstNum struct{
		Num int
	}

	n := firstNum{
		Num: arr[0],
	}

	flagNum := n.Num - arr[1]
	for _, i := range arr[2:]{
		if n.Num / i == flagNum {
			
		}
	}

	return true
}

func commonFactors(a int, b int) int {
	type maxNum struct{
		Num int
	}

	ans := []int{}
	p := maxNum{
		Num: 0,
	}

	if a > b{
		p.Num = a
	} else {
		p.Num = b
	}

	for i := 1; p.Num > i ; i++{
		if a % i == 0 && b % i == 0{
			ans = append(ans, i)
		}
	}
	
	return len(ans)
}


func merge(word1 string, word2 string) string {
	stringSlice := []string{}
	result := []string{}
	word1Length := utf8.RuneCountInString(word1)
	word2Length := utf8.RuneCountInString(word2)

	if word1Length == word2Length{
		word1Slice := strings.Split(word1, "")
		v := 0
		for i := 0; word1Length * 2 > i ; i++{
			if i % 2 != 0{
				stringSlice = append(stringSlice, "")
				continue
			}
			stringSlice = append(stringSlice, word1Slice[v])
			v += 1
		}
		v = 0
		word2Slice := strings.Split(word2, "")
		for _, i := range stringSlice{
			if i == ""{
				result = append(result, word2Slice[v])
				v += 1
				continue
			}
			result = append(result, i)
		}
		return strings.Join(result, "")
	}

	if word1Length > word2Length{
		word1Slice := strings.Split(word1, "")
		v := 0
		for i := 0; word1Length * 2 > i ; i++{
			if i % 2 != 0{
				stringSlice = append(stringSlice, "")
				continue
			}
			stringSlice = append(stringSlice, word1Slice[v])
			v += 1
		}

		v = 0
		word2Slice := strings.Split(word2, "")
		for _, i := range stringSlice{
			if i == "" && len(word2Slice) > v{
				result = append(result, word2Slice[v])
				v += 1
				continue
			} else {
				result = append(result, "")
			}
			result = append(result, i)
		}
		return strings.Join(result, "")
	}

	if word2Length > word1Length{
		word2Slice := strings.Split(word1, "")
		v := 0
		for i := 0; word2Length * 2 > i ; i++{
			if i % 2 == 0{
				stringSlice = append(stringSlice, "")
				continue
			}
			stringSlice = append(stringSlice, word2Slice[v])
			v += 1
		}

		v = 0
		word1Slice := strings.Split(word1, "")
		for _, i := range stringSlice{
			if i == "" && len(word1Slice) > v{
				result = append(result, word1Slice[v])
				v += 1
				continue
			} else {
				result = append(result, "")
			}
			result = append(result, i)
		}
		return strings.Join(result, "")
	}
	return strings.Join(result, "")
}


func majorty(nums []int) int {
	numMap := make(map[int]int)
	ans := 0
	for _, num := range nums{
		numMap[num]++
	}
	for key, value := range numMap{
		if value >= ans{
			ans = key
		}
	}
	fmt.Println(ans)
	return ans
}

func removeDuplicates(s string) string{
	result := []string{}
	slice := strings.Split(s, "")
	mapForRemoveDup := make(map[string]int)
	for _, word := range slice{
		mapForRemoveDup[word]++
	}

	for word, num := range mapForRemoveDup{
		if num == 1 {
			result = append(result, word)
		} else {
			result = append(result, word, word)
		}
	}
	return strings.Join(result, "")
}

func removeChar(s []string, c string, n int) []string {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			count++
			if count <= n {
				s = append(s[:i], s[i+1:]...)
				i--
			} else {
				break
			}
		}
	}
	return s
}

func longestPalindrome(s string) int {
	var result int
	listS := strings.Split(s, "")
	wordMap := make(map[string]int)
	for _, word := range listS{
		wordMap[word]++
	}
	for word , num := range wordMap{
		if num % 2 == 0 {
			result += num
			delete(wordMap, word)
		}
	}
	nums := getValues(wordMap)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for _ , num := range nums{
		if num % 2 != 0 && result % 2 == 0{
			result += num
		}
	}
	return result
}

func getValues(m map[string]int) []int {
    values := []int{}
    for _, v := range m {
        values = append(values, v)
    }
    return values
}


func findOccrrences(text string, first string, second string) []string{
	ans := []string{}
	textSlice := strings.Split(text, " ")
	fmt.Println(textSlice)
	for i := 0; len(textSlice) > i ; i++{
		if i == len(textSlice) - 2{
			break
		} 
		if textSlice[i] == first && textSlice[i + 1] == second{
			ans = append(ans, textSlice[i + 2])
		}
	} 
	return ans
}

func maximum69Number(num int) int {
	str := strconv.FormatInt(int64(num), 10)
	slice := strings.Split(str, "")

	result := num
	tmpSlice := make([]string, len(slice))
	for i := 0; i < len(slice); i++{
		copy(tmpSlice, slice) 
		if tmpSlice[i] == "6"{
			tmpSlice[i] = "9"
		}
		if tmpSlice[i] == "9"{
			tmpSlice[i] = "6"
		}
		str := strings.Join(tmpSlice, "")
		tmp, _ := strconv.Atoi(str)
		
		fmt.Println(tmpSlice)

		if tmp > result{
			num = tmp
		}
	} 
	return result
}

func dayOfYear(date string) int {
	regex, _ := regexp.Compile("[^a-zA-Z0-9]+")
	dates := regex.ReplaceAllString(date, " ")
	dateSlice := strings.Split(dates, " ")
	year, month, day := dateSlice[0], time.Month(date[1]), dateSlice[2]
	intYear, _ := strconv.Atoi(year)
	intDate, _ := strconv.Atoi(day)
	t := time.Date(intYear, month, intDate, 0, 0, 0, 0, time.UTC)
	_, daysInYear := t.ISOWeek()

	return daysInYear
}

func sumZero(n int) []int {
	ans := []int{}
	positive := 1
	negative := -1
	if n % 2 != 0{
		ans = append(ans, 0)
	}
	for n > len(ans) {
			ans = append(ans, positive)
			ans = append(ans, negative)
			positive += 1
			negative -= 1
		}
	return ans
} 



func uncommonFromSentences(s1 string, s2 string) []string {
    str1 := strings.Split(s1, " ")
    str2 := strings.Split(s2, " ")
    wordSet1, wordSet2 := make(map[string]int), make(map[string]int)
    for _, word := range str1 {
        wordSet1[word]++
    }

    for _, word := range str2 {
        wordSet2[word]++
    }

    var res []string
    for word, fre := range wordSet1 {
        if _, ok := wordSet2[word]; !ok && fre == 1 {
            res = append(res, word)
        }
    }

    for word, fre := range wordSet2 {
        if _, ok := wordSet1[word]; !ok && fre == 1 {
            res = append(res, word)
        }
    }

    return res
}




func mostCommonWord(paragraph string, banned []string) string {
	regex, _ := regexp.Compile("[^a-zA-Z0-9]+")
	para := regex.ReplaceAllString(paragraph, " ")
	par := strings.ToLower(para)
	paraList := strings.Split(par, " ")
	

	wordMap := make(map[string]bool)
	for _, w := range banned{
		wordMap[w] = true
	}

	var filteredWords []string
	for _, w := range paraList{
		if _, exists := wordMap[w]; !exists{
			filteredWords = append(filteredWords, w)
		}
	}

	mapW := make(map[string]int)
	for _, word := range filteredWords{
		mapW[word]++
	}

	var result string
	maxCount := 0
	for word, count := range mapW{
		if count > maxCount{
			result = word
			maxCount = count
		}
	}
	return result
}