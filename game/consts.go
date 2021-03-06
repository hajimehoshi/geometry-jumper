package game

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/uber-go/zap"
)

const (
	DefaultSpeed = 4

	// delay between pattern spawning (milliseconds)
	PatternDelayMillis = 1000

	// 'difficulty' constants for patterns?
	LowDifficulty    = "LOW"
	MediumDifficulty = "MEDIUM"
	HighDifficulty   = "HIGH"

	// Track constants
	UpperTrack = 1
	LowerTrack = 2

	// shape constants
	TriangleType = 1
	CircleType   = 2
	SquareType   = 3

	// triangle movement state constants
	TriangleBeforeSwap = 1
	TriangleDuringSwap = 2
	TriangleAfterSwap  = 3

	// space (in pixels, I guess?) between the two tracks
	UpperTrackYAxis = 150
	LowerTrackYAxis = 250

	// x constants for sides and player
	PlayerX   = 60
	LeftSide  = 50
	RightSide = 375
	// this should be screen width probably, not a constant 400
	TrackLength = 400

	// default angle values IN DEGREES!!!! (go math requires radians but degrees make more sense...)
	DefaultCircleAngleOfDescent   float64 = 35
	DefaultTriangleAngleOfDescent float64 = 45

	JumpHeight    = LowerTrackYAxis - UpperTrackYAxis
	JumpUpSpeed   = 5
	JumpDownSpeed = 3

	StartingUpperSpeedLimit = 4
	MinimumSpeed            = 4

	ScreenWidth  = 400
	ScreenHeight = 400

	// difficulty unlock points
	MediumDifficultyUnlockSeconds = 60
	HighDifficultyUnlockSeconds   = 120
	ColorSwapUnlockSeconds        = 240

	// difficulty lock points
	LowDifficultyLockSeconds    = 150
	MediumDifficultyLockSeconds = 180
)

var (
	Debug = false
	// slice of all shape types
	ShapeTypes = [...]int{TriangleType, CircleType, SquareType}

	// slice of all difficulties
	DifficultyTypes = [...]string{LowDifficulty, MediumDifficulty, HighDifficulty}

	// track mappings so you can use the track ID to get the track's position on the y axis
	TrackMappings = map[int]int{
		UpperTrack: UpperTrackYAxis,
		LowerTrack: LowerTrackYAxis,
	}

	// subsequent track shows us what track comes after the one we're currently on.
	SubsequentTracks = map[int]int{
		UpperTrack: LowerTrack,
		LowerTrack: UpperTrack,
	}

	PersonStandingImage *ebiten.Image
	PersonJumpingImage  *ebiten.Image
	PlatformImage       *ebiten.Image
	SquareImage         *ebiten.Image
	TriangleImage       *ebiten.Image
	CircleImage         *ebiten.Image
	UpperTrackLine      *ebiten.Image
	LowerTrackLine      *ebiten.Image
	TitleImage          *ebiten.Image
	EndImage            *ebiten.Image

	UpperTrackOpts *ebiten.DrawImageOptions
	LowerTrackOpts *ebiten.DrawImageOptions

	ShapeImageMap  map[int][]*ebiten.Image
	HitboxImageMap map[int]*ebiten.Image

	// color mapping constants
	DefaultSquareColorMap   ebiten.ColorM
	DefaultCircleColorMap   ebiten.ColorM
	DefaultTriangleColorMap ebiten.ColorM

	// shape to default color map
	ColorMappings map[int]ebiten.ColorM
	Font          *truetype.Font

	// pattern collection
	GamePatternCollection *PatternCollection
	logger                zap.Logger
)
