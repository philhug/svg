package svg

import mt "github.com/rustyoz/Mtransform"

// added by jueewo (c) jueewo.com

// Text is an SVG XML text element
type Text struct {
	ID    string `xml:"id,attr"`
	Style string `xml:"style,attr"`
	X     string `xml:"x,attr"`
	Y     string `xml:"y,attr"`

	transform mt.Transform
	group     *Group
}

// ParseDrawingInstructions implements the DrawingInstructionParser
// interface
func (r *Text) ParseDrawingInstructions() (chan *DrawingInstruction, chan error) {
	draw := make(chan *DrawingInstruction)
	errs := make(chan error)

	defer close(draw)
	defer close(errs)

	return draw, errs
}
