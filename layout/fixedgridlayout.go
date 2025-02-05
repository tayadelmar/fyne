package layout

import (
	"math"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

// Declare conformity with Layout interface
var _ fyne.Layout = (*fixedGridLayout)(nil)

type fixedGridLayout struct {
	CellSize fyne.Size
	colCount int
	rowCount int
}

// Layout is called to pack all child objects into a specified size.
// For a FixedGridLayout this will attempt to lay all the child objects in a row
// and wrap to a new row if the size is not large enough.
func (g *fixedGridLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	g.colCount = 1
	g.rowCount = 1

	if size.Width > g.CellSize.Width {
		g.colCount = int(math.Floor(float64(size.Width) / float64(g.CellSize.Width+theme.Padding())))
	}

	x, y := 0, 0
	for i, child := range objects {
		child.Move(fyne.NewPos(x, y))
		child.Resize(g.CellSize)

		if (i+1)%g.colCount == 0 {
			x = 0
			y += g.CellSize.Height + theme.Padding()
			g.rowCount++
		} else {
			x += g.CellSize.Width + theme.Padding()
		}
	}
}

// MinSize finds the smallest size that satisfies all the child objects.
// For a FixedGridLayout this is simply the specified cellsize as a single column
// layout has no padding. The returned size does not take into account the number
// of columns as this layout re-flows dynamically.
func (g *fixedGridLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize((g.CellSize.Width),
		((g.CellSize.Height * g.rowCount) + ((g.rowCount - 1) * theme.Padding())))
}

// NewFixedGridLayout returns a new FixedGridLayout instance
func NewFixedGridLayout(size fyne.Size) fyne.Layout {
	return &fixedGridLayout{size, 1, 1}
}
