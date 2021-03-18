//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede UG (haftungsbeschränkt) <info@jj-ideenschmiede.de>
//
// This file is part of uuml.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package uuml

import (
	"strings"
)

// To convert the umlauts in strings
func Convert(text string) string {

	// Create list for umlauts
	list := make(map[string]string)

	// Set umlauts
	list["ß"] = "ss"
	list["Ä"] = "Ae"
	list["ä"] = "ae"
	list["Ö"] = "Oe"
	list["ö"] = "oe"
	list["Ü"] = "Ue"
	list["ü"] = "ue"

	// Check each umlauts in string
	for index, value := range list {

		// Replace umlauts
		text = strings.Replace(text, index, value, -1)

	}

	// Return text without umlauts
	return text

}
