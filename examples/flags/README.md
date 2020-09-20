# flags


#### Flags test

```bash
# complex case:
go run -tags=delve ./examples/flags f -f64 7.32 -f 9.9 quick -u 72 fox -c64 2.718+5.71i jumps -i 073 -u 065 -f64 3e+7
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

##### I

```bash
TRACE=1 go run -tags=delve ./examples/flags f -f64 7.32 -f 9.9 quick -u 72 fox -c64 2.718+5.71i jumps -i 073 -i64 81 -u 065 -nwsm -f64 3e+7
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


##### II

```bash
TRACE=1 go run -tags=delve ./examples/flags f -sv"t80k" -sv=zjfk -sv fjksdl -tdv 6s
```

The expected result:

```
...
[--string-value] fjksdl,
[--time-duration-value] 6s,
...
```


##### III: Array

```bash
TRACE=1 go run -tags=delve ./examples/flags f -ssv"aa" -ssv bb -ssv cc,dd,ee
```

The expected result:

```
...
[--string-slice-value] [aa bb cc dd ee],
...
```

For int slice:

```bash
TRACE=1 go run -tags=delve ./examples/flags f -isv11 -isv 22 -isv 33,44,55
```

The expected result:

```
...
[--int-slice-value] [11 22 33 44 55],
...
```



## The New Features for `cmdr`

### Pluggable Extensions

Copy `<project-dir>/ci/local/share/flags` as `$HOME/.local/share/flags`, and run:

```bash
cp -R ./ci/local/share/flags $HOME/.local/share/
go run ./examples/flags
```

> 1. For your final deploy, extensions (shell scripts, elf binary, ...) should be put into `$HOME/.local/share/<appname>/ext/`, with linux executable permissions (`chmod +x`.
>
> 2. For debugging in developing, just put the extensions into `<project-dir>/ci/local/share/<appname>/ext` and `cmdr` will find out it.

And **three extensions** will be plugged into Sub-commands group **`Extensions`**:

![image-20200920111517653](https://i.loli.net/2020/09/20/x5D2ZoFlM7HcOnK.png)

And of course it can be invoked:

![image-20200920111706101](https://i.loli.net/2020/09/20/6qwftmyeWcJ4gbG.png)

