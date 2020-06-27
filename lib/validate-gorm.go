package lib

import "log"

type gorm struct {
	err []string
}

func (g *gorm) validate(err error) []string {
	log.Printf("%#v", err)

	return g.err
}
