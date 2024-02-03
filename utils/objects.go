package utils

type Placeholder [5]string

var zero = Placeholder{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var one = Placeholder{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var two = Placeholder{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var three = Placeholder{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var four = Placeholder{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var five = Placeholder{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var six = Placeholder{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var seven = Placeholder{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var eight = Placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var nine = Placeholder{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

var Separators = Placeholder{
	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}

var Digits = [...]Placeholder{
	zero, one, two, three, four, five, six, seven, eight, nine,
}

var first = Placeholder{
	"   ",
	"   ",
	"   ",
	"   ",
	"   ",
}

var second = Placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"█ █",
}

var third = Placeholder{
	"█  ",
	"█  ",
	"█  ",
	"█  ",
	"███",
}

var fourth = Placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"█ █",
}

var fifth = Placeholder{
	"██ ",
	"█ █",
	"██ ",
	"█ █",
	"█ █",
}

var sixth = Placeholder{
	"█ █",
	"███",
	"█ █",
	"█ █",
	"█ █",
}

var seventh = Placeholder{
	" █ ",
	" █ ",
	" █ ",
	"   ",
	" █ ",
}

var last = Placeholder{
	"   ",
	"   ",
	"   ",
	"   ",
	"   ",
}

var Alarm = [...]Placeholder{
	first, second, third, fourth, fifth, sixth, seventh, last,
}
