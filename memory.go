package main

import "fmt"

type MemoryManager struct {
	memory *[]int
}

func NewMemory() *MemoryManager {
	m := make([]int, 10)
	return &MemoryManager{&m}
}

func (m MemoryManager) Get(i int) int {
	m.Grow(m.memory, i)
	return (*m.memory)[i]
}

func (m MemoryManager) Put(i int, v int) {
	m.Grow(m.memory, i)
	(*m.memory)[i] = v
}

func (m MemoryManager) Add(i int) {
	m.Grow(m.memory, i)
	(*m.memory)[i]++
	if (*m.memory)[i] > 255 {
		(*m.memory)[i] = 0
	}
}

func (m MemoryManager) Substract(i int) {
	m.Grow(m.memory, i)
	(*m.memory)[i]--
	if (*m.memory)[i] < 0 {
		(*m.memory)[i] = 255
	}
}

func (m MemoryManager) Print(i int) {
	// Memory "Dump"
	m.Grow(m.memory, i)
	for n, v := range *m.memory {
		if n == i {
			fmt.Print("[", v, "] ")
		} else {
			fmt.Print(v, " ")
		}
	}
	fmt.Print("\n")
}

func (m MemoryManager) Grow(a *[]int, i int) {
	// Grow the m.memory array as needed by multiplying it in size
	if i > len(*a)-1 {
		mu := 2
		for {
			if i < len(*a)*mu {
				break
			}
			mu++
		}
		newMemory := make([]int, len(*a)*mu)
		copy(newMemory, *a)
		*a = newMemory
		return
	}
	return
}
