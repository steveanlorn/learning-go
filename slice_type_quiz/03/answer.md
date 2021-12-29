# Question 3 Answer

**Answer**:
```text
Index 0 - 1
Index 1 - 2
Index 2 - 3
Index 3 - 4
Index 4 - 5
```

## Explanation
`range` mechanic in Go will have its own copy of the slice (**value semantic**).
Therefore, any mutation happens inside the value in the range, will not be reflected in
the original slice.

```text
data
[ * ] ---> [1,2,3,4,5]
[ 5 ]       ^
[ 5 ]       |
           range
           [ * ]
           [ 5 ]
           [ 5 ]
```

If we want to apply the mutation into the original slice then we should use **pointer semantic**.
```go
for i := range data {
    data[i]++
}
```
Output:
```text
Index 0 - 2
Index 1 - 3
Index 2 - 4
Index 3 - 5
Index 4 - 6
```