// Package twelve outputs a song
package twelve

const testVersion = 1

type thing struct {
	ordinal string
	present string
}

var twelveThings = map[int]thing{
	1:  {"first", "a Partridge"},
	2:  {"second", "two Turtle Doves"},
	3:  {"third", "three French Hens"},
	4:  {"fourth", "four Calling Birds"},
	5:  {"fifth", "five Gold Rings"},
	6:  {"sixth", "six Geese-a-Laying"},
	7:  {"seventh", "seven Swans-a-Swimming"},
	8:  {"eighth", "eight Maids-a-Milking"},
	9:  {"ninth", "nine Ladies Dancing"},
	10: {"tenth", "ten Lords-a-Leaping"},
	11: {"eleventh", "eleven Pipers Piping"},
	12: {"twelfth", "twelve Drummers Drumming"},
}

func getPresent(i int) string {
	var p string
	for j := i; j > 0; j-- {
		if j == 1 && i != 1 {
			p += "and " + twelveThings[j].present
			return p
		}
		if i == 1 {
			return twelveThings[j].present
		}

		p += twelveThings[j].present + ", "
	}
	return p
}

// Verse outputs the line at the point of i
func Verse(i int) string {
	return "On the " + twelveThings[i].ordinal +
		" day of Christmas my true love gave to me, " +
		getPresent(i) +
		" in a Pear Tree."

}

// Song generates the whole song, lines 1 to 12
func Song() string {
	var song string
	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}
	return song
}
