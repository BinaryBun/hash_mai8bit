## **About**
Basic hash-map[string]string on 8bit Pearson's hash-function.


## **Example**
<table>
  <tr>
    <td> <b>Code</b> </td> <td> <b>Output</b> </td>
  </tr>
  <tr>
  <td>

```go
// ...
// some code
func example() {
  data := Map{}
  data.insert("hello", "1")
  data.insert("world", "1")
  data.insert("hello", "2")
  data.insert("!", "3")

  fmt.Println(data.get_value("world"))
  fmt.Println(data.keys())
}
```
  </td>
  <td>

  ```bash
1
["hello" "world" "!"]
```
</td>
</tr>
</table>
