<table>
  <tr>
    <td> **Code** </td> <td> **Output** </td>
  </tr>
  <tr>
  <td>

  ^ Extra blank line above!
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
  V Extra blank line below!

  </td>
  <td>

  ^ Extra blank line above!
  ```bash
1
[hello world !]
```
V Extra blank line below!

</td>
</tr>
</table>
