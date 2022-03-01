//********************************************************************************************************************//
//
// This file is part of uuml.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package uuml

import (
	"strings"
)

// list defines all possible german umlauts
var list = map[string]string{
	"ß": "ss",
	"Ä": "Ae",
	"ä": "ae",
	"Ö": "Oe",
	"ö": "oe",
	"Ü": "Ue",
	"ü": "ue",
}

// Convert is to convert the umlauts in strings
// The function check each umlaut and replace it with the correct written out spelling
func Convert(text string) string {

	for index, value := range list {
		text = strings.Replace(text, index, value, -1)
	}

	return text

}
