# Question 4 Answer

**Answer**:
```text
1
2
3
4
5
Data: [1 2]
```

## Explanation
`range` mechanic in Go will have its own copy of the slice (**value semantic**).
Therefore, any mutation of the original slice inside the range, 
will not reflect the ongoing range iteration.

We should use pointer semantic if we want the new slice will be reflected in the `print`
