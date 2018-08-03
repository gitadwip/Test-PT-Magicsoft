package main

import ("fmt")

func main() {
	number := []int{1, 4, 5, 6, 8, 2}
	fmt.Println("INPUT: Numerical array")
	fmt.Println(number)
	fmt.Println("")
	fmt.Println("OUTPUT: Vertical barcharts")
	sorting(number)
}

func sorting(num []int){
	var max = num[0];
	var n = len(num)
	for i:=1; i<n; i++{
		if max<num[i]{
			max = num[i]
		}
	}
	fmt.Println("STEP VISUALIZATION")
    for i:= 1; i<n; i++{
		j := i;
		for j>0 {
			if num[j-1]> num[j]{
				num[j-1], num[j]=num[j], num[j-1]
			}
			j=j-1
		}
		for k:=max; k>0; k--{
			for l:=0; l<n; l++ {
				if (num[l]+1) > k{
					fmt.Print("! ")
				} else {
					fmt.Print("  ")
				}
			}
		fmt.Println("")
		}
		fmt.Println(num)
	}
}