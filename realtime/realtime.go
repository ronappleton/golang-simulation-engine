package realtime

// This is a helper package to convert real time to simulation time.
// The simulation time is based on the Gregorian calendar, but the
// units below are used to convert real time to simulation time.
// The microseconds to the right of each constant is the number of
// microseconds required to run a game year in real time (the constant name).
// For example, OneHour is the number of microseconds in one hour of real time.

const (
	OneHour          = 8760988000
	TwoHours         = 4380494000
	ThreeHours       = 2920329333
	FourHours        = 2190247000
	FiveHours        = 1752197600
	SixHours         = 1460988000
	SevenHours       = 1251412571
	EightHours       = 1095049500
	NineHours        = 973325333
	TenHours         = 876098800
	ElevenHours      = 796355454
	TwelveHours      = 730049333
	ThirteenHours    = 673841538
	FourteenHours    = 625712857
	FifteenHours     = 583999200
	SixteenHours     = 547499600
	SeventeenHours   = 515292588
	EighteenHours    = 486665555
	NineteenHours    = 461049578
	TwentyHours      = 438049400
	TwentyOneHours   = 417141142
	TwentyTwoHours   = 398181090
	TwentyThreeHours = 380868913
	TwentyFourHours  = 365098800
)
