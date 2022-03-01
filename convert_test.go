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
	"fmt"
	"testing"
)

// To test the convert function
func TestConvert(t *testing.T) {
	text := "Hallöchen, mein Name ist Fränklin Meißter!"
	fmt.Println(Convert(text))
}
