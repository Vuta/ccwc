# Coding Challenges - wc

Simply using Go's `bufio` package to buffer IO calls.

- With default flags
```
$ ./ccwc test.txt
    7145    58164    342190    test.txt
$ wc test.txt
    7145   58164  342190 test.txt
```

- With -c flag
```
$ ./ccwc -c test.txt
    342190    test.txt
$ wc -c test.txt
    342190 test.txt
```

- With -l flag
```
$ ./ccwc -l test.txt
    7145    test.txt
$ wc -l test.txt
    7145 test.txt
```

- With -w flag
```
$ ./ccwc -w test.txt
    58164    test.txt
$ wc -w test.txt
    58164 test.txt
```

- With -m flag
```
$ ./ccwc -m test.txt
    339292    test.txt
$ wc -m test.txt
    339292 test.txt
```
