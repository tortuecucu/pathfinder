package text

import "fmt"

type TextRecorder struct{}

func (t TextRecorder) RecordLines(lines *[]string) {
	fmt.Println(lines)
}
