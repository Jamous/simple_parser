simple_parser
=============

Simple text parser written in go. Accepts an input file, a list of strings to match, and an output file. 

Usage
-----
-m string File to match values from. (default "match.txt")

-o string File to save results to. (default "output.txt")

-p string File to parse from. (default "input.txt")


Example
-------
input.txt
``` 
This is my string
and my second string
not this one
but this string
or this
```

match.txt
```
string
or
```

output.txt
```
This is my string
and my second string
but this string
or this
````