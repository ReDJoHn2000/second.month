package function

import (
	"fmt"
)

func SumTarget(nums [5]int, target int){

	var flag bool = false

    result:=make([]int,2)

    for i:=0; i < len(nums)-1; i++{
        for j:=i+1; j < len(nums); j++{
            if nums[i]+nums[j] == target {
                result[0] = i
                result[1] = j
                flag = true;
            }
        }
    }
      
    if(!flag){
        fmt.Println("not possible") 
    } else {
      fmt.Println(result)
   }
}
