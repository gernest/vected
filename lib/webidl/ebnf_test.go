package webidl

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/alecthomas/participle/lexer"
	"golang.org/x/exp/ebnf"
)

func TestGrammar(t *testing.T) {
	f, err := os.Open("grammar.ebnf")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	g, err := ebnf.Parse("grammer", f)
	if err != nil {
		t.Fatal(err)
	}
	err = ebnf.Verify(g, "definitions")
	if err != nil {
		t.Fatal(err)
	}
	b, _ := ioutil.ReadAll(f)
	def, err := lexer.EBNF(string(b))
	if err != nil {
		t.Fatal(err)
	}
	for k := range def.Symbols() {
		fmt.Println(k)
	}
	sample := `callback AbortCallback = void ();`
	_, err = lexer.ConsumeAll(def.Lex(strings.NewReader(sample)))
	if err != nil {
		t.Fatal(err)
	}
}
