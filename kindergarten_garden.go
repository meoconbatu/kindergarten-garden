package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

const testVersion = 1

type Garden map[string][]string

var plants = map[rune]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	g := make(Garden)
	sChildren := make([]string, len(children))
	copy(sChildren, children)
	sort.Strings(sChildren)
	rows := strings.Split(diagram, "\n")
	err := checkValidDiagram(rows)
	if err != nil {
		return nil, err
	}
	for i := 1; i < len(rows); i++ {
		for j, d := range rows[i] {
			k := j / 2
			if plants[d] == "" {
				return nil, errors.New("Invaid cup codes")
			}
			if len(g[sChildren[k]]) >= 4 {
				return nil, errors.New("Duplicate name")
			}
			g[sChildren[k]] = append(g[sChildren[k]], plants[d])
		}
	}
	return &g, nil
}
func checkValidDiagram(rows []string) error {
	if len(rows) != 3 {
		return errors.New("Wrong diagram format")
	}
	if len(rows[1]) != len(rows[2]) {
		return errors.New("Mismatched rows")
	}
	if len(rows[1])%2 != 0 || len(rows[2])%2 != 0 {
		return errors.New("Odd number of cups")
	}
	return nil
}
func (g *Garden) Plants(child string) (plants []string, err bool) {
	plants = (*g)[child]
	if plants == nil {
		return nil, false
	}
	return plants, true
}
