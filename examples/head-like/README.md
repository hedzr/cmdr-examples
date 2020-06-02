# head-like

The feature `head-like` provides a flag like `--lines' of [`head`(1)](https://www.man7.org/linux/man-pages/man1/head.1.html).

`-n`/`--lines` as we known it can be abbreviated to `-number`, it looks like:

```bash
head -9 some.txt.file   # print the first 9 lines of 'some.txt.file'
```

`head-like` app simulate it with:

```bash
go run ./examples/head-like -n 9
go run ./examples/head-like --lines 9
go run ./examples/head-like -9
go run ./examples/head-like -113
```
