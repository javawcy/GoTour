package main

import (
	"fmt"
	"runtime"
)

func main() {
	sum := 0

	for i := 0; i < 100; i++ {
		sum += 1
	}
	fmt.Println(sum)

	//初始语句和后置语句是可选的 - while
	for sum < 1000 {
		sum += 1
	}
	fmt.Println(sum)

	//无限循环
	/* 	for {
		fmt.Println("HAHA")
	} */

	if sum < 9999 {
		fmt.Println("YES")
	}

	//可以在if中定义变量
	if a := 10000; sum < a {
		fmt.Println("YEZ")
	}

	//swich 默认break,fallthrough需要指定
	switch os := runtime.GOOS; os {
	case "drawin":
		fmt.Println("MAC")
	case "linux":
		fmt.Println("LINUX")
	default:
		fmt.Println("WIN")
	}
	//switch 可以不指定条件类似 if-elseif-elseif-else

}
