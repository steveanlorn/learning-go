# Question 1 Answer

**Answer**:
```
len 0, cap 0, address 0x0, true
len 0, cap 0, address 0x11a7c50, false
```

## Explanation
The difference between slice `a` and `b` is 
that slice `a` has no reference to any slice, but 
`b` has a reference to something.

Slice `a` is an example of a zero declaration of a slice.  
Slice `b` is an example of an empty literal declaration of a slice.

```
Zero declaration:
[nil] // reference of underlaying array
[ 0 ] // len
[ 0 ] // cap

Empty literal declaration:
[ * ] ---> pointer to an empty struct type
[ 0 ] // len
[ 0 ] // cap
```

Example of an empty struct type
```
var es struct{}
```

Why do we need two different declarations?
1. There will be a requirement where the slice should be nil (does not exist).
2. There will be a requirement where the slice should be empty.