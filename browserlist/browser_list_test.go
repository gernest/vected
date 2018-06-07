package browserlist

import (
	"testing"
)

type stats struct {
	name    string
	version version
	usage   float64
	match   bool
}

var sample = []stats{
	{name: "and_chr", version: "64", usage: 30.4129},
	{name: "and_ff", version: "57", usage: 0.25047},
	{name: "and_qq", version: "1.2", usage: 0},
	{name: "and_uc", version: "11.8", usage: 7.50853},
	{name: "android", version: "2.1", usage: 0},
	{name: "android", version: "2.2", usage: 0},
	{name: "android", version: "2.3", usage: 0},
	{name: "android", version: "3", usage: 0},
	{name: "android", version: "4", usage: 0},
	{name: "android", version: "4.1", usage: 0.067427},
	{name: "android", version: "4.2-4.3", usage: 0.205212},
	{name: "android", version: "4.4", usage: 0.578992},
	{name: "android", version: "4.4.3-4.4.4", usage: 0.35619},
	{name: "android", version: "62", usage: 0},
	{name: "baidu", version: "7.12", usage: 0},
	{name: "bb", version: "7", usage: 0.0133584},
	{name: "bb", version: "10", usage: 0.0534336},
	{name: "chrome", version: "4", usage: 0.004706},
	{name: "chrome", version: "5", usage: 0.004879},
	{name: "chrome", version: "6", usage: 0.004879},
	{name: "chrome", version: "7", usage: 0.005591},
	{name: "chrome", version: "8", usage: 0.005591},
	{name: "chrome", version: "9", usage: 0.005591},
	{name: "chrome", version: "10", usage: 0.004534},
	{name: "chrome", version: "11", usage: 0.008866},
	{name: "chrome", version: "12", usage: 0.004283},
	{name: "chrome", version: "13", usage: 0.004879},
	{name: "chrome", version: "14", usage: 0.004706},
	{name: "chrome", version: "15", usage: 0.009154},
	{name: "chrome", version: "16", usage: 0.004393},
	{name: "chrome", version: "17", usage: 0.004393},
	{name: "chrome", version: "18", usage: 0.017732},
	{name: "chrome", version: "19", usage: 0.004418},
	{name: "chrome", version: "20", usage: 0.004393},
	{name: "chrome", version: "21", usage: 0.004433},
	{name: "chrome", version: "22", usage: 0.017732},
	{name: "chrome", version: "23", usage: 0.008786},
	{name: "chrome", version: "24", usage: 0.013299},
	{name: "chrome", version: "25", usage: 0.008866},
	{name: "chrome", version: "26", usage: 0.008866},
	{name: "chrome", version: "27", usage: 0.008866},
	{name: "chrome", version: "28", usage: 0.008866},
	{name: "chrome", version: "29", usage: 0.363506},
	{name: "chrome", version: "30", usage: 0.017732},
	{name: "chrome", version: "31", usage: 0.026598},
	{name: "chrome", version: "32", usage: 0.008866},
	{name: "chrome", version: "33", usage: 0.013299},
	{name: "chrome", version: "34", usage: 0.022165},
	{name: "chrome", version: "35", usage: 0.013299},
	{name: "chrome", version: "36", usage: 0.017732},
	{name: "chrome", version: "37", usage: 0.017732},
	{name: "chrome", version: "38", usage: 0.039897},
	{name: "chrome", version: "39", usage: 0.017732},
	{name: "chrome", version: "40", usage: 0.017732},
	{name: "chrome", version: "41", usage: 0.022165},
	{name: "chrome", version: "42", usage: 0.022165},
	{name: "chrome", version: "43", usage: 0.062062},
	{name: "chrome", version: "44", usage: 0.013299},
	{name: "chrome", version: "45", usage: 0.035464},
	{name: "chrome", version: "46", usage: 0.022165},
	{name: "chrome", version: "47", usage: 0.04433},
	{name: "chrome", version: "48", usage: 0.124124},
	{name: "chrome", version: "49", usage: 0.820105},
	{name: "chrome", version: "50", usage: 0.026598},
	{name: "chrome", version: "51", usage: 0.053196},
	{name: "chrome", version: "52", usage: 0.053196},
	{name: "chrome", version: "53", usage: 0.026598},
	{name: "chrome", version: "54", usage: 0.084227},
	{name: "chrome", version: "55", usage: 0.359073},
	{name: "chrome", version: "56", usage: 0.115258},
	{name: "chrome", version: "57", usage: 0.097526},
	{name: "chrome", version: "58", usage: 0.164021},
	{name: "chrome", version: "59", usage: 0.115258},
	{name: "chrome", version: "60", usage: 0.159588},
	{name: "chrome", version: "61", usage: 0.172887},
	{name: "chrome", version: "62", usage: 0.203918},
	{name: "chrome", version: "63", usage: 1.29887},
	{name: "chrome", version: "64", usage: 15.2229},
	{name: "chrome", version: "65", usage: 8.28528},
	{name: "chrome", version: "66", usage: 0.053196},
	{name: "chrome", version: "67", usage: 0.022165},
	{name: "chrome", version: "68", usage: 0},
	{name: "edge", version: "12", usage: 0.026598},
	{name: "edge", version: "13", usage: 0.035464},
	{name: "edge", version: "14", usage: 0.097526},
	{name: "edge", version: "15", usage: 0.128557},
	{name: "edge", version: "16", usage: 1.57371},
	{name: "edge", version: "17", usage: 0.022165},
	{name: "edge", version: "18", usage: 0},
	{name: "firefox", version: "2", usage: 0.004433},
	{name: "firefox", version: "3", usage: 0.004433},
	{name: "firefox", version: "3.5", usage: 0.008786},
	{name: "firefox", version: "3.6", usage: 0.008866},
	{name: "firefox", version: "4", usage: 0.013299},
	{name: "firefox", version: "5", usage: 0.004879},
	{name: "firefox", version: "6", usage: 0.020136},
	{name: "firefox", version: "7", usage: 0.005725},
	{name: "firefox", version: "8", usage: 0.008866},
	{name: "firefox", version: "9", usage: 0.00533},
	{name: "firefox", version: "10", usage: 0.004283},
	{name: "firefox", version: "11", usage: 0.004433},
	{name: "firefox", version: "12", usage: 0.004418},
	{name: "firefox", version: "13", usage: 0.004486},
	{name: "firefox", version: "14", usage: 0.00453},
	{name: "firefox", version: "15", usage: 0.004433},
	{name: "firefox", version: "16", usage: 0.004433},
	{name: "firefox", version: "17", usage: 0.004349},
	{name: "firefox", version: "18", usage: 0.004393},
	{name: "firefox", version: "19", usage: 0.004443},
	{name: "firefox", version: "20", usage: 0.004283},
	{name: "firefox", version: "21", usage: 0.004418},
	{name: "firefox", version: "22", usage: 0.004393},
	{name: "firefox", version: "23", usage: 0.004433},
	{name: "firefox", version: "24", usage: 0.008786},
	{name: "firefox", version: "25", usage: 0.008836},
	{name: "firefox", version: "26", usage: 0.004393},
	{name: "firefox", version: "27", usage: 0.004393},
	{name: "firefox", version: "28", usage: 0.004418},
	{name: "firefox", version: "29", usage: 0.008866},
	{name: "firefox", version: "30", usage: 0.004433},
	{name: "firefox", version: "31", usage: 0.008866},
	{name: "firefox", version: "32", usage: 0.004433},
	{name: "firefox", version: "33", usage: 0.004433},
	{name: "firefox", version: "34", usage: 0.008866},
	{name: "firefox", version: "35", usage: 0.008866},
	{name: "firefox", version: "36", usage: 0.013299},
	{name: "firefox", version: "37", usage: 0.008866},
	{name: "firefox", version: "38", usage: 0.039897},
	{name: "firefox", version: "39", usage: 0.008866},
	{name: "firefox", version: "40", usage: 0.013299},
	{name: "firefox", version: "41", usage: 0.013299},
	{name: "firefox", version: "42", usage: 0.008866},
	{name: "firefox", version: "43", usage: 0.031031},
	{name: "firefox", version: "44", usage: 0.075361},
	{name: "firefox", version: "45", usage: 0.035464},
	{name: "firefox", version: "46", usage: 0.008866},
	{name: "firefox", version: "47", usage: 0.053196},
	{name: "firefox", version: "48", usage: 0.137423},
	{name: "firefox", version: "49", usage: 0.022165},
	{name: "firefox", version: "50", usage: 0.035464},
	{name: "firefox", version: "51", usage: 0.062062},
	{name: "firefox", version: "52", usage: 0.4433},
	{name: "firefox", version: "53", usage: 0.026598},
	{name: "firefox", version: "54", usage: 0.035464},
	{name: "firefox", version: "55", usage: 0.04433},
	{name: "firefox", version: "56", usage: 0.119691},
	{name: "firefox", version: "57", usage: 0.106392},
	{name: "firefox", version: "58", usage: 1.98155},
	{name: "firefox", version: "59", usage: 1.57371},
	{name: "firefox", version: "60", usage: 0.048763},
	{name: "firefox", version: "61", usage: 0},
	{name: "ie", version: "5.5", usage: 0.009298},
	{name: "ie", version: "6", usage: 0.00911931},
	{name: "ie", version: "7", usage: 0.013679},
	{name: "ie", version: "8", usage: 0.173267},
	{name: "ie", version: "9", usage: 0.118551},
	{name: "ie", version: "10", usage: 0.113991},
	{name: "ie", version: "11", usage: 2.76315},
	{name: "ie_mob", version: "10", usage: 0.033396},
	{name: "ie_mob", version: "11", usage: 0.178112},
	{name: "ios_saf", version: "3.2", usage: 0.00107165},
	{name: "ios_saf", version: "4.0-4.1", usage: 0},
	{name: "ios_saf", version: "4.2-4.3", usage: 0.00107165},
	{name: "ios_saf", version: "5.0-5.1", usage: 0.0192897},
	{name: "ios_saf", version: "6.0-6.1", usage: 0.0150031},
	{name: "ios_saf", version: "7.0-7.1", usage: 0.0525108},
	{name: "ios_saf", version: "8", usage: 0.0150031},
	{name: "ios_saf", version: "8.1-8.4", usage: 0.079302},
	{name: "ios_saf", version: "9.0-9.2", usage: 0.0600123},
	{name: "ios_saf", version: "9.3", usage: 0.429731},
	{name: "ios_saf", version: "10.0-10.2", usage: 0.428659},
	{name: "ios_saf", version: "10.3", usage: 0.993418},
	{name: "ios_saf", version: "11.0-11.2", usage: 8.61391},
	{name: "ios_saf", version: "11.3", usage: 0},
	{name: "op_mini", version: "all", usage: 2.66111},
	{name: "op_mob", version: "10", usage: 0},
	{name: "op_mob", version: "11", usage: 0},
	{name: "op_mob", version: "11.1", usage: 0},
	{name: "op_mob", version: "11.5", usage: 0},
	{name: "op_mob", version: "12", usage: 0},
	{name: "op_mob", version: "12.1", usage: 0},
	{name: "op_mob", version: "37", usage: 0.0111391},
	{name: "opera", version: "9", usage: 0.0082},
	{name: "opera", version: "9.5-9.6", usage: 0.00685},
	{name: "opera", version: "10.0-10.1", usage: 0},
	{name: "opera", version: "10.5", usage: 0.008392},
	{name: "opera", version: "10.6", usage: 0.004706},
	{name: "opera", version: "11", usage: 0.016581},
	{name: "opera", version: "11.1", usage: 0.006229},
	{name: "opera", version: "11.5", usage: 0.004879},
	{name: "opera", version: "11.6", usage: 0.008786},
	{name: "opera", version: "12", usage: 0.008786},
	{name: "opera", version: "12.1", usage: 0.04433},
	{name: "opera", version: "15", usage: 0.00685},
	{name: "opera", version: "16", usage: 0.00685},
	{name: "opera", version: "17", usage: 0.00685},
	{name: "opera", version: "18", usage: 0.005014},
	{name: "opera", version: "19", usage: 0.006015},
	{name: "opera", version: "20", usage: 0.004879},
	{name: "opera", version: "21", usage: 0.006597},
	{name: "opera", version: "22", usage: 0.006597},
	{name: "opera", version: "23", usage: 0.013434},
	{name: "opera", version: "24", usage: 0.006702},
	{name: "opera", version: "25", usage: 0.006015},
	{name: "opera", version: "26", usage: 0.005595},
	{name: "opera", version: "27", usage: 0.004393},
	{name: "opera", version: "28", usage: 0.008698},
	{name: "opera", version: "29", usage: 0.004879},
	{name: "opera", version: "30", usage: 0.004879},
	{name: "opera", version: "31", usage: 0.004433},
	{name: "opera", version: "32", usage: 0.005152},
	{name: "opera", version: "33", usage: 0.005014},
	{name: "opera", version: "34", usage: 0.009758},
	{name: "opera", version: "35", usage: 0.004879},
	{name: "opera", version: "36", usage: 0.039897},
	{name: "opera", version: "37", usage: 0.004283},
	{name: "opera", version: "38", usage: 0.004367},
	{name: "opera", version: "39", usage: 0.004534},
	{name: "opera", version: "40", usage: 0.004367},
	{name: "opera", version: "41", usage: 0.004227},
	{name: "opera", version: "42", usage: 0.004418},
	{name: "opera", version: "43", usage: 0.008668},
	{name: "opera", version: "44", usage: 0.004227},
	{name: "opera", version: "45", usage: 0.008866},
	{name: "opera", version: "46", usage: 0.004433},
	{name: "opera", version: "47", usage: 0.004433},
	{name: "opera", version: "48", usage: 0.004433},
	{name: "opera", version: "49", usage: 0.008866},
	{name: "opera", version: "50", usage: 0.048763},
	{name: "opera", version: "51", usage: 0.740311},
	{name: "opera", version: "52", usage: 0.026598},
	{name: "safari", version: "3.1", usage: 0},
	{name: "safari", version: "3.2", usage: 0.008692},
	{name: "safari", version: "4", usage: 0},
	{name: "safari", version: "5", usage: 0.013299},
	{name: "safari", version: "5.1", usage: 0.070928},
	{name: "safari", version: "6", usage: 0.004349},
	{name: "safari", version: "6.1", usage: 0.013299},
	{name: "safari", version: "7", usage: 0.008866},
	{name: "safari", version: "7.1", usage: 0.004283},
	{name: "safari", version: "8", usage: 0.031031},
	{name: "safari", version: "9", usage: 0.031031},
	{name: "safari", version: "9.1", usage: 0.164021},
	{name: "safari", version: "10", usage: 0.097526},
	{name: "safari", version: "10.1", usage: 0.319176},
	{name: "safari", version: "11", usage: 1.56928},
	{name: "safari", version: "11.1", usage: 0.022165},
	{name: "safari", version: "TP", usage: 0},
	{name: "samsung", version: "4", usage: 0.943211},
	{name: "samsung", version: "5", usage: 0.228029},
	{name: "samsung", version: "6.2", usage: 1.71022},
}

func TestQueryNormal(t *testing.T) {
	t.Run("> %", func(ts *testing.T) {
		s := []stats{
			{name: "samsung", version: "4", usage: 0.943211, match: true},
			{name: "samsung", version: "5", usage: 0.228029, match: false},
			{name: "samsung", version: "6.2", usage: 1.71022, match: true},
		}
		q := query("> 50%")
		for _, v := range s {
			g := q(v.name, v.version, v.usage)
			if g != v.match {
				t.Errorf("expected %v got %v", v.match, g)
			}
		}
	})
	t.Run(">= %", func(ts *testing.T) {
		s := []stats{
			{name: "samsung", version: "4", usage: 0.943211, match: true},
			{name: "samsung", version: "5", usage: 0.228029, match: false},
			{name: "samsung", version: "5.4", usage: 0.5, match: true},
			{name: "samsung", version: "6.2", usage: 1.71022, match: true},
		}
		q := query(">= 50%")
		for _, v := range s {
			g := q(v.name, v.version, v.usage)
			if g != v.match {
				t.Errorf("expected %v got %v :%v", v.match, g, v.usage)
			}
		}
	})
	t.Run("< %", func(ts *testing.T) {
		s := []stats{
			{name: "samsung", version: "4", usage: 0.943211, match: false},
			{name: "samsung", version: "5", usage: 0.228029, match: true},
			{name: "samsung", version: "5.4", usage: 0.5, match: false},
			{name: "samsung", version: "6.2", usage: 1.71022, match: false},
		}
		q := query("< 50%")
		for _, v := range s {
			g := q(v.name, v.version, v.usage)
			if g != v.match {
				t.Errorf("expected %v got %v :%v", v.match, g, v.usage)
			}
		}
	})
	t.Run("< %", func(ts *testing.T) {
		s := []stats{
			{name: "samsung", version: "4", usage: 0.943211, match: false},
			{name: "samsung", version: "5", usage: 0.228029, match: true},
			{name: "samsung", version: "5.4", usage: 0.5, match: true},
			{name: "samsung", version: "6.2", usage: 1.71022, match: false},
		}
		q := query("<= 50%")
		for _, v := range s {
			g := q(v.name, v.version, v.usage)
			if g != v.match {
				t.Errorf("expected %v got %v :%v", v.match, g, v.usage)
			}
		}
	})
	t.Run("cover %", func(ts *testing.T) {
		s := []stats{
			{name: "samsung", version: "4", usage: 0.943211, match: true},
			{name: "samsung", version: "5", usage: 0.228029, match: false},
			{name: "samsung", version: "5.4", usage: 0.5, match: true},
			{name: "samsung", version: "6.2", usage: 1.71022, match: true},
		}
		q := query("cover 50%")
		for _, v := range s {
			g := q(v.name, v.version, v.usage)
			if g != v.match {
				t.Errorf("expected %v got %v :%v", v.match, g, v.usage)
			}
		}
	})
}
