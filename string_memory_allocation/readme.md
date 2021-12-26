# String Memory Allocation In Go

Computer understands only two things 0 and 1 (binary)
and programing language job to show this binary data based on their data type. 

For example, if we have binary `01101100`, it represents 108 integers if the data type is an integer.
What if binary `01101100` has a data type string?
The programming language will try to map this numeric data into character **code-points**.

## ASCII & Unicode
ASCII defines 128 characters, identified by the code-points 0–127. 
It covers English letters, Latin numbers, and a few other characters.
https://en.wikipedia.org/wiki/ASCII#Control_code_chart

Unicode, which is a superset of ASCII, defines 1,114,112 code-points.
https://en.wikipedia.org/wiki/Unicode

Go as a modern programming language want to support characters in Unicode.
Go uses UTF-8 encoding to represent the Unicode characters. 

## UTF-8 Encoding
UTF stands for Unicode Transformation Format. 
The ‘8’ means it uses 8-bit blocks to represent a character.

UTF-8 is just a method to represent a binary number to send any symbol with a minimum number of chunks. 
In UTF-8, 1 chunk means 8 bits of data. 
Send `$` symbol, the software will convert this number to binary 100100 and 
send this binary in one chunk that’s 8 bits `00100100`.

See `01_symbol_binary_convertion`

- 1 byte: 0xxxxxx
- 2 bytes: 110xxxxx 10xxxxxx
- 3 bytes: 1110xxxx 10xxxxxx 10xxxxxx
- 4 bytes: 11110xxx 10 XXXXXX 10xxxxxx 10xxxxxx

## String in Go
See `02_string_loop`
A string is a sequence of bytes (or uint8)
ASCII characters are encoded with one byte, while other code points use more.

## Runes
See `03_runes`
When we do iteration over string it reads byte by byte and 
on a special character, loop read it 3 bytes as 3 different characters 
not a single one that’s why they printed 3 different characters (2 characters are non-visible). 
To avoid this issue, Go has a concept of `rune`.

The rune type is an alias for `int32` and is used to emphasize that an integer represents a code point.
Why `int32`? Because in UTF-8, all Unicode code points will be encoded using one to four bytes.

## Resource:
- https://perennialsky.medium.com/how-string-works-in-golang-7ac4d797164b
- https://dev.to/luispa/what-the-heck-is-a-rune-in-golang-5bl5