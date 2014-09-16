package pullbarz

import (
  "sort"
  "github.com/tlorens/ansiout"
  "github.com/tlorens/keyboard"
)

var options map[int]string
var keys []int
var OptX[25] int
var OptY[25] int
var curOpt int = 1
var lastOpt int = 1
var count int = 1

func writeOption(s string, selected bool) {
	for i := range s {
		if s[i] > 64 && s[i] < 91 {
			ansiout.Print(11, 0, string(s[i]))
		} else {
			ansiout.Print(7, 0, string(s[i]))
		}
	}
}

func LightBar(X int, Y int) int {
	var ch int = 0
	var choices []byte

	if X == 0 || Y == 0 {
		X, Y = ansiout.CursorXY()
	}

	ansiout.PrintColorXY(11, 0, X, Y, "|")

	for _, k := range keys {
		OptY[count], OptX[count] = ansiout.CursorXY()

		for i := 0; i < len(options[k]); i++ {
			if options[k][i] > 64 &&  options[k][i] < 91 {
				choices = append(choices, options[k][i])
			}
		}
		//ansiout.Print(7, 0, " " + options[k] + " ")
		writeOption(" " + options[k] + " ", false)
		count = count + 1
	}
	ansiout.Print(11,0, "|")

	writeSelect()

	forloop:
		for {
			ch := keyboard.ReadKey()

			switch ch {
				case 96:
					break forloop

				case keyboard.KEY_LF:
					decChoice()
					break;

				case keyboard.KEY_RT:
					incChoice()
					break;

				case 10:
					return int(choices[curOpt-1])
					break;

				default:
					for _, v := range choices {
						if byte(ch-32) == v  || byte(ch-32) == v {
							return int(v)
						}
					}
			}

			writeSelect()
		}

	return ch
}


func incChoice() {
	lastOpt = curOpt
	curOpt = curOpt + 1
	if curOpt > len(options) {
		curOpt = 1
	}
}


func decChoice() {
	lastOpt = curOpt
	curOpt = curOpt - 1
	if curOpt < 1 {
		curOpt = len(options)
	}
}


func writeSelect() {
	ansiout.GotoXY(OptY[lastOpt], OptX[lastOpt])
	writeOption(" " + options[lastOpt] + " ", false)
	ansiout.PrintColorXY(7, 1, OptY[curOpt], OptX[curOpt], " " + options[curOpt] + " ")
}


func setOptions(opts map[int]string) {
	options = opts

	for k := range opts {
		keys = append(keys, k)
	}

	sort.Ints(keys)
}


func Init(opts map[int]string) {
	setOptions(opts)
}
