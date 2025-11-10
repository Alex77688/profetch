package display

import "fmt"

func Render(info []string, logo []string, logo_exists bool) {
	if !logo_exists {
		for _, infoe := range info {
			fmt.Println(infoe);
		}
	}
}