package main
import (
	"fmt"
)

func bubblesort(arr []int)  {
	n := len(arr)
	for i := 0; i<n;i++ {
		for j := 0; j<n-i-1;j++ {
			if(arr[j]>arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main()  {
	fmt.Print("Enter 10 integers : ")
	arr := make([]int,10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	bubblesort(arr)
	fmt.Print("Sorted Array : ")
	fmt.Println(arr)
}