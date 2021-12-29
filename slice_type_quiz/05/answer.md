# Question 5 Answer

**Answer**:
```text
[0] 0xc00007e000 - apple
[1] 0xc00007e010 - banana
[2] 0xc00007e020 - coconut
[3] 0xc00007e030 - grape
[4] 0xc00007e040 - avocado
-----------------
[0] 0xc00007e020 - coconut
[1] 0xc00007e030 - grape
[2] 0xc00007e040 - avocado
```

## Explanation
On line 16, we create a new slice `fruits2`, the value will be:
```text
[orange,grape]
```
```text
 f1         0  1  2  3  4 5 6 7
[ * ] ---> [a][b][o][g][m][][][]
[ 5 ]             ^
[ 8 ]             |
                  f2
                 [ * ]
                 [ 2 ]
                 [ 6 ]
```

On line 17, we replace first element of `fruits2` with `coconut`.
The value of `fruits2` will be:
```text
[coconut,grape]
```
Because of `fruits2` has the same reference with `fruit1` therefore the second element
value of `fruit1` also change to `coconut`

On line 18, we append a new string to `fruits2`.
The value of `fruits2` will be:
```text
[coconut,grape,avocado]
```
Because of `fruits2` has the same reference with `fruit1` therefore the appended
string will also be appended to `fruit1`.