/**
 * long long time
 * @author denghb
 * @since 2020-09-10
 */
var KEYBOARD = {
    k104: {
        "[Esc]": { x: 38, y: 36},

        "[F1]": { x: 144, y: 36},
        "[F2]": { x: 198, y: 36},
        "[F3]": { x: 254, y: 36},
        "[F4]": { x: 308, y: 36},

        "[F5]": { x: 388, y: 36},
        "[F6]": { x: 442, y: 36},
        "[F7]": { x: 496, y: 36},
        "[F8]": { x: 550, y: 36},

        "[F9]": { x: 634, y: 36},
        "[F10]": { x: 688, y: 36},
        "[F11]": { x: 742, y: 36},
        "[F12]": { x: 796, y: 36},

        "[Print Screen]": { x: 860, y: 38},
        "[Scroll Lock]": { x: 914, y: 36},
        "[Pause]": { x: 968, y: 36},

        "`~": { x: 38, y: 116},
        "`": { x: 38, y: 116},
        "~": { x: 38, y: 116},
        "1!": { x: 91, y: 116},
        "1": { x: 91, y: 116},
        "!": { x: 91, y: 116},
        "2@": { x: 145, y: 116},
        "2": { x: 145, y: 116},
        "@": { x: 145, y: 116},
        "3#": { x: 199, y: 116},
        "3": { x: 199, y: 116},
        "#": { x: 199, y: 116},
        "4$": { x: 253, y: 116},
        "4": { x: 253, y: 116},
        "$": { x: 253, y: 116},
        "5%": { x: 307, y: 116},
        "5": { x: 307, y: 116},
        "%": { x: 307, y: 116},
        "6^": { x: 361, y: 116},
        "6": { x: 361, y: 116},
        "^": { x: 361, y: 116},
        "7&": { x: 415, y: 116},
        "7": { x: 415, y: 116},
        "&": { x: 415, y: 116},
        "8*": { x: 469, y: 116},
        "8": { x: 469, y: 116},
        "*": { x: 469, y: 116},
        "9(": { x: 523, y: 116},
        "9": { x: 523, y: 116},
        "(": { x: 523, y: 116},
        "0)": { x: 577, y: 116},
        "0": { x: 577, y: 116},
        ")": { x: 577, y: 116},
        "-_": { x: 631, y: 116},
        "-": { x: 631, y: 116},
        "_": { x: 631, y: 116},
        "=+": { x: 685, y: 116},
        "=": { x: 685, y: 116},
        "+": { x: 685, y: 116},
        "[BackSpace]": { x: 766, y: 116},

        "[Insert]": { x: 860, y: 116},
        "[Home]": { x: 914, y: 116},
        "[Page Up]": { x: 968, y: 116},

        "[Tab]": { x: 54, y: 170},
        "Q": { x: 117, y: 170},
        "W": { x: 171, y: 170},
        "E": { x: 225, y: 170},
        "R": { x: 279, y: 170},
        "T": { x: 333, y: 170},
        "Y": { x: 387, y: 170},
        "U": { x: 441, y: 170},
        "I": { x: 495, y: 170},
        "O": { x: 549, y: 170},
        "P": { x: 603, y: 170},
        "[{": { x: 657, y: 170},
        "[": { x: 657, y: 170},
        "{": { x: 657, y: 170},
        "]}": { x: 711, y: 170},
        "]": { x: 711, y: 170},
        "}": { x: 711, y: 170},
        "\\": { x: 780, y: 170},
        "|": { x: 780, y: 170},
        "[Delete]": { x: 860, y: 170},
        "[End]": { x: 914, y: 170},
        "[Page Down]": { x: 968, y: 170},

        "[Caps Lock]": { x: 56, y: 224},
        "A": { x: 131, y: 224},
        "S": { x: 185, y: 224},
        "D": { x: 239, y: 224},
        "F": { x: 293, y: 224},
        "G": { x: 347, y: 224},
        "H": { x: 401, y: 224},
        "J": { x: 455, y: 224},
        "K": { x: 509, y: 224},
        "L": { x: 563, y: 224},
        ";:": { x: 617, y: 224},
        ";": { x: 617, y: 224},
        ":": { x: 617, y: 224},
        "'\"": { x: 671, y: 224},
        "'": { x: 671, y: 224},
        "\"": { x: 671, y: 224},
        "[Enter]": { x: 760, y: 224},

        "[Shift]": { x: 74, y: 276},
        "Z": { x: 160, y: 276},
        "X": { x: 214, y: 276},
        "C": { x: 268, y: 276},
        "V": { x: 322, y: 276},
        "B": { x: 376, y: 276},
        "N": { x: 430, y: 276},
        "M": { x: 484, y: 276},
        ",<": { x: 538, y: 276},
        ",": { x: 538, y: 276},
        "<": { x: 538, y: 276},
        ".>": { x: 592, y: 276},
        ".": { x: 592, y: 276},
        ">": { x: 592, y: 276},
        "/?": { x: 646, y: 276},
        "/": { x: 646, y: 276},
        "?": { x: 646, y: 276},
        "[Right Shift]": { x: 746, y: 276},

        "[Up Arrow]": { x: 914, y: 276},

        "[Ctrl]": { x: 44, y: 332},
        "[Win]": { x: 114, y: 332},
        "[Alt]": { x: 180, y: 332},
        "[ ]": { x: 390, y: 332},
        "[Right Alt]": { x: 582, y: 332},
        "[Right Win]": { x: 654, y: 332},
        "[Menu]": { x: 722, y: 332},
        "[Right Ctrl]": { x: 786, y: 332},

        "[Left Arrow]": { x: 860, y: 332},
        "[Down Arrow]": { x: 914, y: 332},
        "[Right Arrow]": { x: 968, y: 332},

        "[Num Lock]": { x: 1036, y: 116},
        "[Num /]": { x: 1090, y: 116},
        "[Num *]": { x: 1144, y: 116},
        "[Num -]": { x: 1198, y: 116},

        "[Num 7]": { x: 1036, y: 170},
        "[Num 8]": { x: 1090, y: 170},
        "[Num 9]": { x: 1144, y: 170},
        "[Num +]": { x: 1198, y: 200},

        "[Num 4]": { x: 1036, y: 224},
        "[Num 5]": { x: 1090, y: 224},
        "[Num 6]": { x: 1144, y: 224},


        "[Num 1]": { x: 1036, y: 276},
        "[Num 2]": { x: 1090, y: 276},
        "[Num 3]": { x: 1144, y: 276},
        "[Num Enter]": { x: 1198, y: 310},
        
        "[Num 0]": { x: 1064, y: 332},
        "[Num .]": { x: 1144, y: 332},
    },
}
