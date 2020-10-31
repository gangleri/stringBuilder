# Playing with Builder vs Concat vs append

Primarily interested in concat vs builder but included append to see 
how bad I can make this.

`go test -bench=.`

| Test | Ran |           Time|
| :--- | ---:| -------------:|
|Build | 674773|   1758 ns/op|
|Concat|  10995| 110091 ns/op|
|Append|   5721| 210602 ns/op|


