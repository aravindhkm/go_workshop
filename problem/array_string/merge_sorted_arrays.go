/*
Problem:
- Merge two sorted arrays.

Example:
- Input: []int{1, 3, 5}, []int{2, 4, 6}
  Output: []int{1, 2, 3, 4, 5, 6}
- Input: []int{1, 3, 5}, []int{2, 4, 6, 7}
  Output: []int{1, 2, 3, 4, 5, 6, 7},

Approach:
- Since these arrays are sorted, can use two pointers approach to iterate
  through both of them and append the smaller value to a new merged list at
  each step.

Solution:
- Have two pointers start at the beginning of these two arrays.
- While both of them does not reach the end, compare two current values
  at each step and append the smaller one two a new merged list.
- Move the two pointers up accordingly as values get merged in.
- In the case where one of these pointers reach the end first and the
  other one is still in the middle of the array, simply add the rest of
  its values to the merged list since they are all sorted and guaranteed
  to be in ascending order.

Cost:
- O(n) time, O(n) space.
*/

package arraystring

import "fmt"


// mergeSortedArrays merges two sorted arrays and returns a single sorted array
func ExpectedSortedArrays(arr1, arr2 []int){
  i, j := 0, 0
  merged := []int{}

  for i < len(arr1) && j < len(arr2) {
      if arr1[i] < arr2[j] {
          merged = append(merged, arr1[i])
          i++
      } else {
          merged = append(merged, arr2[j])
          j++
      }
  }

  fmt.Println(i)
  // Add any remaining elements from arr1
  if i < len(arr1) {
      merged = append(merged, arr1[i:]...)
  }

  fmt.Println(j)
  // Add any remaining elements from arr2
  if j < len(arr2) {
      merged = append(merged, arr2[j:]...)
  }

  fmt.Println(merged)
}


func MyMergeSortedArrays(data ...[]int) {
  result := []int{}

  least := 0

  for {
    for i:=0; i<len(data); i++ {
      for _, value := range data[i] {
          if value == least + 1  {
            result = append(result, value)
          }
      }
    }

    curr := result[len(result) - 1]
    if least == curr {
      break
    }

    least = curr
  }

	fmt.Println(result)
}

func MergeSortedArraysTest() {
  arr_1 := []int{1, 3, 5}
  arr_2 := []int{2, 4, 6}
  arr_3 := []int{1, 3, 5}
  arr_4 := []int{2, 4, 6, 7}

	// MyMergeSortedArrays(arr_1,arr_2)
  // MyMergeSortedArrays(arr_3,arr_4)

  ExpectedSortedArrays(arr_1,arr_2)
  ExpectedSortedArrays(arr_3,arr_4)

}
