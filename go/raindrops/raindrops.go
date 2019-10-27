package raindrops

import (
	"fmt"
)

func main() {

}

func Convert(input int) string {
	str := ""
	if (input % 3 == 0) {
		str += "Pling"
	}
	if (input % 5 == 0) {
		str += "Plang"
	}
	if (input % 7 == 0) {
		str += "Plong"
	}
	if (str == ""){
		str = fmt.Sprintf("%d", input)
	}
	return str
}