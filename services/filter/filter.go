package filter

import (
	"log"
	"paf/shared"
	"strings"
)

var stopWords = []string{"bird-watching", "ailurophobia", "mango"}

type Filter struct {
	Next shared.Service
}

func (c *Filter) Execute(p *shared.Process) {

	skip := false
	for _, word := range stopWords {
		if strings.Contains(p.Message.Message, word) {
			log.Printf("filtered because message contains: '%s'", word)
			skip = true
		}
	}

	if !skip {
		log.Printf("Passed checks!")
		if c.Next != nil {
			c.Next.Execute(p)
		}
	}
}

func (c *Filter) SetNext(next shared.Service) {
	c.Next = next
}
