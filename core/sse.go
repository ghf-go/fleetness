package core

import "fmt"

type Sse struct {
	c *GContent
}

func (s *Sse) Send(data string, event ...string) error {
	if len(event) == 1 {
		_, e := s.c.w.Write([]byte(fmt.Sprintf("event: %s\ndata: %s\n\n", data, event[0])))
		if e == nil {
			s.c.flush()
		}
		return e
	} else {
		_, e := s.c.w.Write([]byte(fmt.Sprintf("data: %s\n\n", data)))
		if e == nil {
			s.c.flush()
		}
		return e
	}
}
