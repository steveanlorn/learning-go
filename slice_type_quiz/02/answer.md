# Question 2 Answer

**Answer**:
```text
Value: [0 0 0 0 0], len 5, cap 5
Value: [0 0 0 0 0 1], len 6, cap 10
Value: [0 0 0 0 0 1 2], len 7, cap 10
Value: [0 0 0 0 0 1 2 3], len 8, cap 10
Value: [0 0 0 0 0 1 2 3 4], len 9, cap 10
Value: [0 0 0 0 0 1 2 3 4 5], len 10, cap 10
```

## Explanation
If the capacity and length of the slice are equal, then `append`
mechanic will create a new slice with double capacity (until 1k elements, more than that will grow 25%) from the
original slice.

## Optimization
With that in mind, we can optimize the code:

1. if we know exact maximum number of expected length of slice,
   we can keep using `data := make([]int, 5)`, but remove the `append` and replace with:
   ```go
   data[record] = record
   ```
   Output:
   ```text
   Value: [0 0 0 0 0], len 5, cap 5
   Value: [1 0 0 0 0], len 5, cap 5
   Value: [1 2 0 0 0], len 5, cap 5
   Value: [1 2 3 0 0], len 5, cap 5
   Value: [1 2 3 4 0], len 5, cap 5
   Value: [1 2 3 4 5], len 5, cap 5
   ```
2. By modifying the slice initialization by set the length & capacity explicitly to zero. 
Use this approach if we do not know the expected maximum capacity.
    ```go
    data := make([]int, 0)
    ```
    Output:
    ```text
    Value: [], len 0, cap 0
    Value: [1], len 1, cap 1
    Value: [1 2], len 2, cap 2
    Value: [1 2 3], len 3, cap 4
    Value: [1 2 3 4], len 4, cap 4
    Value: [1 2 3 4 5], len 5, cap 8
    ```

2. by modifying the slice initialization by set the length explicitly to zero
    ```go
    data := make([]int, 0, 5)
    ```
   Output:
    ```text
    Value: [], len 0, cap 5
    Value: [1], len 1, cap 5
    Value: [1 2], len 2, cap 5
    Value: [1 2 3], len 3, cap 5
    Value: [1 2 3 4], len 4, cap 5
    Value: [1 2 3 4 5], len 5, cap 5
    ```

