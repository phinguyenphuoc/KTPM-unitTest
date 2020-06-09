package main

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) string {
	*s = append(*s, str)
	index := len(*s) - 1
	element := (*s)[index]
	return element
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Clear() {
	*s = []string{}
}

func (s *Stack) Peek() string {
	index := len(*s) - 1
	if index >= 0 {
		element := (*s)[index]
		return element
	} else {
		return "-1"
	}
}

func main() {
	var stack Stack

	stack.Push("this")
	stack.Push("is")
	// stack.Clear()
	stack.Push("sparta!!")
	stack.Peek()
	stack.Pop()
	stack.Peek()

	// for len(stack) > 0 {
	// 	x, y := stack.Pop()
	// 	if y == true {
	// fmt.Println(x)
	// 	}
	// }
}
