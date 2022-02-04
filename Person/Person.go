package Person

import "errors"

type Person struct {
	name string
}

func (p *Person) GetName() (string, error) {
	if p.name == "" {
		return "", errors.New("name is not set")
	}
	return p.name, nil
}

func (p *Person) SetName(s string) {
	p.name = s
}
