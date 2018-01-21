// Package tournament tallies the results of a small football competition.
package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

const testVersion = 4

// board tallies the score
type board []team

// team is the score for a particular team
type team struct {
	name           string
	mp, w, l, d, p int
}

// Tally tallies the results of a small football competition
func Tally(r io.Reader, w io.Writer) error {

	// let's create a score board
	b := board{
		team{"Allegoric Alaskians", 0, 0, 0, 0, 0},
		team{"Blithering Badgers", 0, 0, 0, 0, 0},
		team{"Courageous Californians", 0, 0, 0, 0, 0},
		team{"Devastating Donkeys", 0, 0, 0, 0, 0},
	}

	// scan the input
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		l := sc.Text()

		// skip over empty  and commented lines
		if l == "" || strings.HasPrefix(l, "#") {
			continue
		}

		// validate line
		line := strings.Split(l, ";")
		if len(line) != 3 {
			return errors.New("burp: wrong number of fields")
		}

		// parse a line (single game)
		err := b.parseLine(line)
		if err != nil {
			return err
		}

	}

	// let's work on the output
	// stream header first
	header := "Team                           | MP |  W |  D |  L |  P\n"
	io.WriteString(w, header)

	// sort results
	sort.Slice(b, func(i, j int) bool {
		return b[i].p > b[j].p
	})

	// stream / print results
	for _, v := range b {
		output := fmt.Sprintf("%-31s| %2d | %2d | %2d | %2d | %2d\n",
			v.name,
			v.mp,
			v.w,
			v.d,
			v.l,
			v.p,
		)
		io.WriteString(w, output)
	}

	return nil
}

// attribute integers to game values (points, match played, etc)
func (b board) parseLine(line []string) error {
	var home, away int
	for i, v := range b {
		if v.name == line[0] {
			home = i
		}
		if v.name == line[1] {
			away = i
		}
	}

	switch line[2] {
	case "win":
		b[home].mp++
		b[home].w++
		b[home].p += 3
		b[away].mp++
		b[away].l++
		return nil
	case "loss":
		b[home].mp++
		b[away].mp++
		b[home].l++
		b[away].w++
		b[away].p += 3
		return nil
	case "draw":
		b[home].mp++
		b[away].mp++
		b[home].d++
		b[away].d++
		b[home].p++
		b[away].p++
		return nil
	default:
		return errors.New("burp: unknown game outcome")
	}

}
