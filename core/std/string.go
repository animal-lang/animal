package std

import (
	"fmt"
)

// pelt(value, times): repeat string value times times
func AnimalPelt(args []interface{}) interface{} {
	if len(args) != 2 {
		return fmt.Errorf("pelt expects 2 arguments (value, times)")
	}
	n := args[0]
	times, ok := args[1].(float64)
	if !ok {
		return fmt.Errorf("pelt expects (value, times)")
	}

	str := fmt.Sprintf("%v", n)
	repeated := ""
	for i := 0; i < int(times); i++ {
		repeated += str
	}
	return repeated
}

// nuzzle(a, b): merge two lists or two strings
func AnimalNuzzle(args []interface{}) interface{} {
	if len(args) != 2 {
		return fmt.Errorf("nuzzle expects 2 arguments")
	}

	switch a := args[0].(type) {
	case []interface{}:
		if b, ok := args[1].([]interface{}); ok {
			return append(a, b...)
		}
	case string:
		if b, ok := args[1].(string); ok {
			return a + b
		}
	}

	return fmt.Errorf("nuzzle expects (list,list) or (string,string)")
}

// squirrel(string, delimiter): split string by delimiter into a list
func AnimalSquirrel(args []interface{}) interface{} {
	if len(args) != 2 {
		return fmt.Errorf("squirrel expects 2 arguments (string, delimiter)")
	}
	str, ok := args[0].(string)
	if !ok {
		return fmt.Errorf("squirrel expects first argument as string")
	}
	delimiter, ok := args[1].(string)
	if !ok {
		return fmt.Errorf("squirrel expects second argument as string")
	}

	var result []interface{}
	start := 0
	for i := 0; i <= len(str)-len(delimiter); {
		if str[i:i+len(delimiter)] == delimiter {
			result = append(result, str[start:i])
			start = i + len(delimiter)
			i += len(delimiter)
		} else {
			i++
		}
	}
	result = append(result, str[start:])

	return result
}

// parrot(list of strings): join list of strings into a single string
func AnimalParrot(args []interface{}) interface{} {
	if len(args) != 1 {
		return fmt.Errorf("parrot expects 1 argument (list of strings)")
	}
	list, ok := args[0].([]interface{})
	if !ok {
		return fmt.Errorf("parrot expects a list")
	}

	result := ""
	for _, v := range list {
		str, ok := v.(string)
		if !ok {
			return fmt.Errorf("parrot expects a list of strings")
		}
		result += str
	}
	return result
}

// rat(string): uppercase string
func AnimalRat(args []interface{}) interface{} {
	if len(args) != 1 {
		return fmt.Errorf("rat expects 1 argument (string)")
	}
	str, ok := args[0].(string)
	if !ok {
		return fmt.Errorf("rat expects a string")
	}
	
	result := ""
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			result += string(char - ('a' - 'A'))
		} else {
			result += string(char)
		}
	}
	return result
}

// mole(string): lowercase string
func AnimalMole(args []interface{}) interface{} {
	if len(args) != 1 {
		return fmt.Errorf("mole expects 1 argument (string)")
	}
	str, ok := args[0].(string)
	if !ok {
		return fmt.Errorf("mole expects a string")
	}
	
	result := ""
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			result += string(char + ('a' - 'A'))
		} else {
			result += string(char)
		}
	}
	return result
}

// snipe(string, symbol = " "): trim symbol from both ends of string
func AnimalSnipe(args []interface{}) interface{} {
	if len(args) < 1 || len(args) > 2 {
		return fmt.Errorf("snipe expects 1 or 2 arguments (string, symbol = ' ')")
	}
	str, ok := args[0].(string)
	if !ok {
		return fmt.Errorf("snipe expects first argument as string")
	}
	symbol := " "
	if len(args) == 2 {
		symbol, ok = args[1].(string)
		if !ok || len(symbol) != 1 {
			return fmt.Errorf("snipe expects second argument as a single character string")
		}
	}

	start := 0
	end := len(str) - 1

	for start <= end && string(str[start]) == symbol {
		start++
	}
	for end >= start && string(str[end]) == symbol {
		end--
	}

	if start > end {
		return ""
	}
	return str[start : end+1]
}
