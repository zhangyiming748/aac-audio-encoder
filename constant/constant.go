package constant

type Param struct {
	Root string
	Vbr  string
}

func (p *Param) GetRoot() string {
	return p.Root
}
func (p *Param) SetRoot(s string) {
	p.Root = s
}
func (p *Param) GetVbr() string {
	return p.Vbr
}
func (p *Param) SetVbr(s string) {
	p.Vbr = s
}
