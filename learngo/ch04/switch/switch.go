package main

import "fmt"

func main() {
	/*
	switch var1 {
	    case val2:
	        ...
	    case val3:
	        ...
	    default:
	        ...
	}
	 */
	//中文的星期几， 输出对应的英文
	day := "星期一"
	switch day {
	case "星期一":
		fmt.Println("Monday")
	case "星期二":
		fmt.Println("Tuesday")
	case "星期三":
		fmt.Println("Wednesday")
	case "星期四":
		fmt.Println("Thursday")
	default:
		fmt.Println("UnKnown")
	}

	score := 95
	switch {
	case score <60:
		fmt.Println("E")
	case score >=60 && score < 70:
		fmt.Println("D")
	case score >=70 && score < 80:
		fmt.Println("C")
	case score >=80 && score < 90:
		fmt.Println("B")
	case score >=90 && score <= 100:
		fmt.Println("A")
	}

	score = 70
	switch score{
	case 60, 70, 80:
		fmt.Println("Wow")
	default:
		fmt.Println("default")
	}
}
