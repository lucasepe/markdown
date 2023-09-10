package markdown

import (
	"fmt"
	"testing"
)

func TestImageSize(t *testing.T) {
	data := `![foo bar](/path/to/image.jpg "=200x300" )`

	render := New(
		HTML(false),
		Tables(true),
		Linkify(true),
		Typographer(true),
		XHTMLOutput(true),
	)

	tokens := render.Parse([]byte(data))
	inline, ok := tokens[1].(*Inline)
	if !ok {
		t.Fatalf("expected *markdown.Inline, got: %T", tokens[1])
	}
	tok := inline.Children[0].(*Image)

	w, h := parseImageSize(tok.Title)
	fmt.Printf("w: %d, h: %d\n", w, h)

	//got := render.RenderToString([]byte(data))
	//fmt.Println(got)
}
