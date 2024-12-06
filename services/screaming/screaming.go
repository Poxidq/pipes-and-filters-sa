package screaming

import (
	"log"
	"paf/shared"
	"strings"
)

type Screaming struct {
	Next shared.Service
}

func (c *Screaming) Execute(p *shared.Process) {

	p.Message.Message = strings.ToUpper(p.Message.Message)
	log.Printf("made uppercase: %s", p.Message.Message)

	if c.Next != nil {
		c.Next.Execute(p)
	}
}

func (c *Screaming) SetNext(next shared.Service) {
	c.Next = next
}
