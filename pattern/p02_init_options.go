package main

import "fmt"

type Player struct {
	name   string
	gender string
	age    int
}

type Option func(*Player)

// 思想就是用一些列函数来修改player的各个属性
func NewPlayer(options ...Option) *Player {
	p := &Player{
		name:   "defaultname",
		gender: "defaultgender",
		age:    0,
	}

	for _, option := range options {
		option(p)
	}
	return p
}

func OptionName(name string) Option {
	return func(p *Player) {
		p.name = name
	}
}

func OptionGender(gender string) Option {
	return func(p *Player) {
		p.gender = gender
	}
}

func main() {
	p := NewPlayer(OptionName("jack"), OptionGender("male"))
	fmt.Println(p)
}
