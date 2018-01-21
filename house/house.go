// Package house
package house

import "strings"

var (
	a = "This is"
	b = []string{
		"the malt\nthat lay in",
		"the rat\nthat ate",
		"the cat\nthat killed",
		"the dog\nthat worried",
		"the cow with the crumpled horn\nthat tossed",
		"the maiden all forlorn\nthat milked",
		"the man all tattered and torn\nthat kissed",
		"the priest all shaven and shorn\nthat married",
		"the rooster that crowed in the morn\nthat woke",
		"the farmer sowing his corn\nthat kept",
		"the horse and the hound and the horn\nthat belonged to",
	}
	c = "the house that Jack built."
)

// Embed embeds a noun phrase as the object of relative clause with a
// transitive verb.
func Embed(relPhrase, nounPhrase string) string {
	return relPhrase + " " + nounPhrase
}

// Verse generates a verse of a song with relative clauses that have
// a recursive structure.
func Verse(subject string, relPhrases []string, nounPhrase string) string {

	verse := []string{}

	if len(relPhrases) == 0 {
		verse = append(verse, subject+" "+nounPhrase)
		return strings.Join(verse, "")
	} else {
		verse = append(verse, subject+" "+Embed(strings.Join(relPhrases, " "), nounPhrase))
		return strings.Join(verse, "")
	}

}

// Song generates the full text of "The House That Jack Built".
func Song() string {
	var song []string
	for i := 0; i < len(b)+1; i++ {
		song = append(song, Verse(a, flip(b[:i]), c))
	}
	return strings.Join(song, "\n\n")
}

// flip reverses the order of lines in the text of the rhyme
func flip(lines []string) []string {
	l := len(lines)
	var flipped = make([]string, l)
	for k, v := range lines {
		flipped[l-k-1] = v
	}
	return flipped
}
