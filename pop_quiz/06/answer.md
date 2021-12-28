# Question 6 Answer

**Answer**:
```text
User: 0, likes: {0}
User: 1, likes: {1}
User: 2, likes: {0}
User: 3, likes: {0}
```

## Explanation
On line 15, we create a variable with value pointer to element 1 of slice `user`.
```text
  u         0   1   2  
[ * ] ---> [u1][u2][u3]
[ 3 ]            ^
[ 3 ]            |
                 su
                [ * ]
```

On line 16, we increment the likes. So current user value will be:
```text
User: 0, likes: {0}
User: 1, likes: {1}
User: 2, likes: {0}
```

On line 22, we have a new user, so we append it to the slice of user.
Since the capacity already full, then `append` will allocate a new array.
Now slice `user` reference to a new array.

The problem is that the `shareUser` pointer has not been updated, therefore
a new like will not be reflected in the new `user` slice.
```text
  u         0   1   2   4   5   6
[ * ] ---> [u1][u2][u3][  ][  ][  ] <--- a new slice 
[ 4 ]           
[ 6 ]                
```

## Tips
This kind of bug is hard to find. We may not see it during review, we may not be able
to capture it in the test either.

My tips, if we found an append during the code review, we may want to slow down
our review and double-check the append usage there.