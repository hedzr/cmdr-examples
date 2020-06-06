# actions

demo app for PreAction, Action and PostAction:

```bash
go run ./examples/actions soundex quick fox jumps

# enable the debugger attached mode:
go run -tags=delve ./examples/actions/ snd quick fox jumps

# complex case:
go run -tags=delve ./examples/actions soundex -f64 7.32 -f 9.9 quick -u 72 fox -c64 2.718+5.71i jumps
```
