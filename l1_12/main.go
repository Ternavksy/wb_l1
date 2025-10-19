package main

import "fmt"

// создаём собственный тип Set для строк
type StringSet struct {
	items map[string]struct{}
}

// конструктор множества
func NewStringSet() *StringSet {
	return &StringSet{
		items: make(map[string]struct{}),
	}
}

// метод добавления элемента
func (s *StringSet) Add(item string) {
	s.items[item] = struct{}{}
}

// метод проверки
func (s *StringSet) Contains(item string) bool {
	_, exists := s.items[item]
	return exists
}

// метод получения всех элементов в виде слайса
func (s *StringSet) Values() []string {
	keys := make([]string, 0, len(s.items))
	for key := range s.items {
		keys = append(keys, key)
	}
	return keys
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	set := NewStringSet()
	for _, word := range words {
		set.Add(word)
	}

	fmt.Println("Множество:", set.Values())
}
