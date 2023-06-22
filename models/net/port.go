package net

type Port struct {
	Number int  `json:"port"`
	Open   bool `json:"open"`
}

func NewPort(number int) *Port {
	return &Port{
		Number: number,
		Open:   false,
	}
}

func (p *Port) OpenPort() {
	p.Open = true
}

func (p *Port) ClosePort() {
	p.Open = false
}
