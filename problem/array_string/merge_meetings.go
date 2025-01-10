
/*
Problem:
- Given a list of unsorted, independent meetings, returns a list of a merged
  one.

Example:
- Input: []meeting{{1, 2}, {2, 3}, {4, 5}}
  Output: []meeting{{1, 3}, {4, 5}}
- Input: []meeting{{1, 5}, {2, 3}}
  Output: []meeting{{1, 5}}

Approach:
- Sort the list in ascending order so that meetings that might need to be
  merged are next to each other.
- Can merge two meetings together if the first one's end time is greater or
  or equal than the second one's start time.

Solution:
- Sort the list in ascending order.
- Create a new list of merged meetings and consider the first meeting in the
  original list to be the last merged one.
- Iterate through the original list and verify if the last merged meeting's
  end time is greater or equal than the current meeting's start time.
- If it is true, merge them using the last merged meeting's start time
  and the larger one's end time.

Cost:
- O(nlogn) time, O(n) space.
- Because we sort all meeting first, the runtime is O(nlogn). We create a new
  list of merged meeting times, so the space cost is O(n).
*/

package arraystring


func MergeMeetingsMain() {

}