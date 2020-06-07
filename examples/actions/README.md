# actions

demo app for PreAction, Action and PostAction:

```bash
go run ./examples/actions soundex quick fox jumps

# enable the debugger attached mode:
go run -tags=delve ./examples/actions/ snd quick fox jumps
```

The expected result:

```
...
[ACTION]     0. quick => q222
[ACTION]     1. fox => f12
[ACTION]     2. jumps => j2512
...
```




#### Flags test

```bash
# complex case:
go run -tags=delve ./examples/actions soundex -f64 7.32 -f 9.9 quick -u 72 fox -c64 2.718+5.71i jumps -i 073 -u 065 -f64 3e+7
```

The expected result:

```
...
         [--bool] false,
         [--int] 59,
         [--int64] 81,
         [--uint] 53,
         [--uint64] 4,
         [--float32] 9.9,
         [--float64] 3e+07,
         [--complex64] (2.718+5.71i),
         [--complex128] (3.14+9i),
...
```



#### Combining Short Flags

```bash
TRACE=1 go run -tags=delve ./examples/actions soundex -f64 7.32 -f 9.9 quick -u 72 fox -c64 2.718+5.71i jumps -i 073 -i64 81 -u 065 -nwsm -f64 3e+7
```

The expected result:

```
...
         [--single] true,
         [--double] false,
         [--norway] true,
         [--mongo] true,
...
```

