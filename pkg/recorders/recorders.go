package recorders

type Recorder interface {
	RecordLines(lines []string)
}
