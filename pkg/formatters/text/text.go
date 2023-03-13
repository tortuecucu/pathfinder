package text

import (
	"fmt"
	"strconv"

	"github.com/tortuecucu/pathfinder/pkg/core"
)

type TextFormatter struct {
}

func (f TextFormatter) Format(exe *core.FactCollection) []string {
	var r []string

	for key, fact := range exe.Facts {
		var valueString string

		switch fact.Value.(type) {
		case int:
			valueString = strconv.Itoa(fact.Value.(int))
		case float64:
			valueString = strconv.FormatFloat(fact.Value.(float64), 'E', -1, 64)
		case string:
			valueString = fmt.Sprintf("%v", fact.Value)
		case error:
			valueString = fmt.Sprintf("%v", fact.Value) //TODO: improve it
		default:
			valueString = fmt.Sprintf("%v", fact.Value) //TODO: handle error
		}

		r = append(r, key+": '"+valueString+"'")

	}

	return r
}
