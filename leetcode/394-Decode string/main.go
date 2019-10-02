/*
Given an encoded string, return its decoded string.
The encoding rule is: k[encoded_string], where the encoded_string inside the square
brackets is being repeated exactly k times. Note that k is guaranteed to be a positive
integer.
You may assume that the input string is always valid; No extra white spaces, square
brackets are well-formed, etc.
Furthermore, you may assume that the original data does not contain any digits and that
digits are only for those repeat numbers, k. For example, there won't be input like 3a or
2[4].
 */
package main

type stack struct {
	data []byte
	size int
}
func constructor()*stack{
	return &stack{data:[]byte{}}
}
func (s *stack)isEmpty()bool{
	return s.size == 0
}
func (s *stack)push(e byte){
	s.data =append(s.data,e)
	s.size++
}
func(s *stack)top()byte{
	if s.isEmpty() {
		panic("the stack is empty")
	}
	return s.data[s.size-1]
}
func(s *stack)pop()byte{
	if s.isEmpty() {
		panic("the stack is empty")
	}
	e := s.data[s.size-1]
	s.size--
	s.data = s.data[:s.size]
	return e
}

func decodeString(s string) string {
	kStack := constructor()
	encodedStringStack := constructor()
	for i:=0;i<len(s);{
		if s[i]>='0'&& s[i]<='9'{
			kStack.push(s[i])
		}
		if s[i] == '[' {

		}

	}
}