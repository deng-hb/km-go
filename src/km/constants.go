package km

import "strings"

const (
	KEYBOARD = "k"
	MOUSE    = "m"
	APP      = ""
)

var keyboardMap = map[uint32]string{
	8:   "[BackSpace]",
	9:   "[Tab]",
	12:  "[Clear]",
	13:  "[Enter]",
	17:  "[Ctrl]",
	18:  "[Alt]",
	19:  "[Pause]",
	20:  "[Caps Lock]",
	27:  "[Esc]",
	32:  "[ ]",
	33:  "[Page Up]",
	34:  "[Page Down]",
	35:  "[End]",
	36:  "[Home]",
	37:  "[Left Arrow]",
	38:  "[Up Arrow]",
	39:  "[Right Arrow]",
	40:  "[Down Arrow]",
	44:  "[Print Screen]",
	45:  "[Insert]",
	46:  "[Delete]",
	48:  "0)",
	49:  "1!",
	50:  "2@",
	51:  "3#",
	52:  "4$",
	53:  "5%",
	54:  "6^",
	55:  "7&",
	56:  "8*",
	57:  "9(",
	65:  "A",
	66:  "B",
	67:  "C",
	68:  "D",
	69:  "E",
	70:  "F",
	71:  "G",
	72:  "H",
	73:  "I",
	74:  "J",
	75:  "K",
	76:  "L",
	77:  "M",
	78:  "N",
	79:  "O",
	80:  "P",
	81:  "Q",
	82:  "R",
	83:  "S",
	84:  "T",
	85:  "U",
	86:  "V",
	87:  "W",
	88:  "X",
	89:  "Y",
	90:  "Z",
	91:  "[Win]",
	92:  "[Right Win]",
	93:  "[Menu]",
	96:  "[Num 0]",
	97:  "[Num 1]",
	98:  "[Num 2]",
	99:  "[Num 3]",
	100: "[Num 4]",
	101: "[Num 5]",
	102: "[Num 6]",
	103: "[Num 7]",
	104: "[Num 8]",
	105: "[Num 9]",
	106: "[Num *]",
	107: "[Num +]",
	108: "[Num Enter]",
	109: "[Num -]",
	110: "[Num .]",
	111: "[Num /]",
	112: "[F1]",
	113: "[F2]",
	114: "[F3]",
	115: "[F4]",
	116: "[F5]",
	117: "[F6]",
	118: "[F7]",
	119: "[F8]",
	120: "[F9]",
	121: "[F10]",
	122: "[F11]",
	123: "[F12]",
	144: "[Num Lock]",
	145: "[Scroll Lock]",
	160: "[Shift]",
	161: "[Right Shift]",
	162: "[Ctrl]",
	163: "[Right Ctrl]",
	164: "[Alt]",
	165: "[Right Alt]",
	186: ";:",
	187: "=+",
	188: ",<",
	189: "-_",
	190: ".>",
	191: "/?",
	192: "`~",
	219: "[{",
	220: "\\|",
	221: "]}",
	222: "'\"",
}

func GetKey(code uint32, isCapsLock bool, isShift bool) string {
	key := keyboardMap[code]
	if isShift {
		if len(key) == 1 {
			if isCapsLock {
				key = strings.ToLower(key)
			} else {
				key = strings.ToUpper(key)
			}
		}
		if len(key) == 2 {
			key = key[1:]
		}
	} else {
		if len(key) == 1 {
			if isCapsLock {
				key = strings.ToUpper(key)
			} else {
				key = strings.ToLower(key)
			}
		}
		if len(key) == 2 {
			key = key[0:1]
		}
	}
	return key
}
