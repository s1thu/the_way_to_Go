package main

import (
	"fmt"
	"strconv"
	"strings"
)

const EDUCATIVE string = "educative"

func IdentifyPrefixPostfix(userID, email string) bool {
	// Implement this function
	if strings.HasPrefix(userID, "UID") && strings.HasSuffix(email, "educative.io") {
		return true
	}
	return false
}

func ContainsEducative(email string) bool {
	// Implement this function
	return strings.Contains(email, EDUCATIVE)
}

func MaskUserName(email string) string {
	// Implement this function
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email // return original if email format is invalid
	}
	username := parts[0]
	domain := parts[1]

	masked := string(username[0]) + strings.Repeat("*", len(username)-2) + string(username[len(username)-1])
	return masked + "@" + domain
}

func IndexOfAtSymbol(email string) int {
	// Implement this function
	return strings.LastIndex(email, "@")
}

func TrimAndSplitUserID(userID string) string {
	// Implement this function
	splitString := strings.Split(userID, "-")
	return splitString[1]
}

func ConvertStringToInt(str string) int {
	// Implement this function
	var returnValue int
	returnValue, _ = strconv.Atoi(str)
	return returnValue
}

func strings_functions() {
	// Test your functions here
	fmt.Println(IdentifyPrefixPostfix("UID-0123", "evangeline@educative.io")) // true
	fmt.Println(ContainsEducative("evangeline@educative.io"))                 // true
	fmt.Println(MaskUserName("evangeline@educative.io"))                      // e******e@educative.io
	fmt.Println(IndexOfAtSymbol("evangeline@educative.io"))                   // 10
	fmt.Println(TrimAndSplitUserID("UID-0123"))                               // 0123
	fmt.Println(ConvertStringToInt("123"))                                    // 123
}
