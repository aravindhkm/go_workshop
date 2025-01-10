package shufflearray

import "fmt"


type Solution struct {
	orig []int
    copy []int
}

func Constructor(nums []int) Solution {
    copyArr:=make([]int,len(nums))
    copy(copyArr,nums)
    return Solution{
        orig: nums,
        copy: copyArr,
    }
}

// func (this *Solution) Reset() []int {
//     return this.orig
// }

// func (this *Solution) Shuffle() []int {
//     rand.Shuffle(len(this.orig),func(i,j int){
//          this.copy[i],this.copy[j]=this.copy[j],this.copy[i]
//     })
    
//     return this.copy
// }

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func ShuffleExampleOne() {
	arr := []int{4,5,6,8}

	data := Constructor(arr)

	fmt.Println("data", data)
}