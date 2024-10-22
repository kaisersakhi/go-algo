## Insertion Sort

Insertion sort is a simple sorting algorithm that builds the final sorted array (or list) one item at a time by comparisons. It is much less efficient on large lists than more advanced algorithms such as Quick-Sort, Heap-Sort, or Merge-Sort. However insertion sort provides several advantages:

- _Simple implementation:_ Jon Bentley shows a version that is three lines in C-like pseudo-code, and five lines when optimized.

- More efficient in practice than most other simple quadratic algorithms such as selection sort or bubble sort.
- _Adaptive:_ efficient for data sets that are already substantially sorted: the time complexity is O(kn) when each element in the input is no more than _k_ places away from its sorted position.
- _Stable:_ does not change the relative order of elements with equal keys
- _In-Place:_ only requires a constant amount _O(1)_ of additional memory space.
- _Online_ can sort a list as it receives it.

## Algorithm

Insertion sort iterates, consuming one input element each repetition, and grows a sorted output list. At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list, and inserts it there. It repeats until no input element remain.

```go
func InsertionSort(list: []int){
    for i := 0; i < len(list); i++ {
        j := i
        for j > 0 && list[j - 1] > list[j] {
            temp := list[j]
            list[j] = list[j-1]
            list[j - 1] = temp

            j--
        }
    }
}
```
