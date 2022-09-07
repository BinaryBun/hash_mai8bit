package map8bit

import ("fmt"
        "math/rand"
        "time"
        )

var table[256]int = func()[256]int {
  var data [256]int
  rand.Seed(time.Now().Unix())
  for index, _ := range(data) { data[index] = index }
  for index, _ := range(data) {
    rv := rand.Intn(256)
    data[index], data[rv] = data[rv], data[index]
  }
  return data
}()

type Node struct {
  key   string
  value string
  next  *Node
}

type List struct {
  head 	*Node
}

type Map struct {
  map_list  [256]*List
  len       int64
}

func (m *Map) insert(key, val string) {
  if m.map_list[hash8(key)] == nil {
    m.map_list[hash8(key)] = &List{}
  }
  m.map_list[hash8(key)].add(key, val, m)
}

func (m *Map) get_value(key string) string {
  if m.map_list[hash8(key)] == nil { panic("Map is nil!!") }

  current := m.map_list[hash8(key)].head
  if current != nil && current.key == key { return current.value }
  for current.next != nil {
    if current.next.key == key { return  current.next.value }
    current = current.next
  }
  panic("Key is missing!!")
}

func (m *Map) keys() []string {
  var answer [] string
  for _, data := range(m.map_list) {
    if data != nil { answer = append(answer, data.print()...)}
  }

  return answer
}

func (l *List) add(key, val string, m *Map) {
  new_data := &Node{value: val, key: key}
  if l.head == nil {
    l.head = new_data
    m.len++
  } else {
    current := l.head
    for current.next != nil {
      if key == current.key {
        current.value = val
        return
      }
      current = current.next
    }
    if key == current.key {
      current.value = val
      return
    }
    current.next = new_data
    m.len++
  }
}

func (l *List) print() []string {
  var result []string
  current := l.head
  if current != nil {
    result = append(result, current.key)
  }
  for current.next != nil {
    result = append(result, current.next.key)
    current = current.next
  }

  return result
}

func hash8(message string) int {
  // person hash
  // using global table
  var hash int = len(message)%256

  for _, value := range(message) {
    hash = table[(hash + int(value))%256]
  }

  return hash
}
