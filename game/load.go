package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func Load() {
	// This is very fragile. initImages must come first!
	initImages()
	initImageMaps()
}

func initImages() {
	pImage, err := openImage("person-standing.png")
	handleErr(err)

	PersonStandingImage, err = ebiten.NewImageFromImage(pImage, ebiten.FilterNearest)
	handleErr(err)

	pImage2, err := openImage("person-jumping.png")
	handleErr(err)

	PersonJumpingImage, err = ebiten.NewImageFromImage(pImage2, ebiten.FilterNearest)
	handleErr(err)

	sImage, err := openImage("square.png")
	handleErr(err)

	SquareImage, err = ebiten.NewImageFromImage(sImage, ebiten.FilterNearest)
	handleErr(err)

	squareWidth, squareHeight := SquareImage.Size()
	// this is wrong. need to figure out how to do hollow shapes
	SquareBorder, err = ebiten.NewImage(squareWidth, squareHeight, ebiten.FilterNearest)
	SquareBorder.Fill(color.White)

	tImage, err := openImage("triangle.png")
	handleErr(err)

	TriangleImage, err = ebiten.NewImageFromImage(tImage, ebiten.FilterNearest)
	handleErr(err)

	cImage, err := openImage("circle.png")
	handleErr(err)

	CircleImage, err = ebiten.NewImageFromImage(cImage, ebiten.FilterNearest)
	handleErr(err)

	UpperTrackLine, err = ebiten.NewImage(TrackLength, 1, ebiten.FilterNearest)
	UpperTrackLine.Fill(color.White)
	handleErr(err)

	LowerTrackLine, err = ebiten.NewImage(TrackLength, 1, ebiten.FilterNearest)
	LowerTrackLine.Fill(color.White)
	handleErr(err)

	UpperTrackOpts = &ebiten.DrawImageOptions{}
	UpperTrackOpts.GeoM.Translate(0, UpperTrackYAxis)

	LowerTrackOpts = &ebiten.DrawImageOptions{}
	LowerTrackOpts.GeoM.Translate(0, LowerTrackYAxis)
}

func initImageMaps() {
	ShapeImageMap = map[int][]*ebiten.Image{
		TriangleType: []*ebiten.Image{TriangleImage},
		SquareType:   []*ebiten.Image{SquareImage},
		CircleType:   []*ebiten.Image{CircleImage},
	}

	HitboxImageMap = map[int]*ebiten.Image{
		// todo: real values
		TriangleType: SquareBorder,
		SquareType:   SquareBorder,
		CircleType:   SquareBorder,
	}
}