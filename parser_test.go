package cte

import (
	"testing"
	// "time"
)

func TestWhitespace(t *testing.T) {
	assertDecoded(t, " nil", NilValue)
	assertDecoded(t, "nil ", NilValue)
	assertDecoded(t, " \tnil \t ", NilValue)
}

// TODO: Test disallowed whitespace

func TestNil(t *testing.T) {
	assertDecoded(t, "nil", NilValue)
}

func TestBool(t *testing.T) {
	assertDecoded(t, "true", true)
	assertDecoded(t, "false", false)
}

func TestList(t *testing.T) {
	assertDecoded(t, "[]", asList())
	assertDecoded(t, "[ ]", asList())
	assertDecoded(t, "[true]", asList(true))
	assertDecoded(t, "[true false]", asList(true, false))
	assertDecoded(t, "  [  true   false ]   ", asList(true, false))
}

func TestUnorderedMap(t *testing.T) {
	assertDecoded(t, "{}", asMap())
	assertDecoded(t, "{ }", asMap())
	assertDecoded(t, "{true = false}", asMap(true, false))
	assertDecoded(t, "{true=false}", asMap(true, false))
	assertDecoded(t, "  {    true   =    false   }  ", asMap(true, false))
	assertDecoded(t, "{true = false false = true}", asMap(true, false, false, true))

	// TODO: Detect this
	// assertDecoded(t, "{true}", asMap())
	// assertDecoded(t, "{true =}", asMap())
}

func TestOrderedMap(t *testing.T) {
	assertDecoded(t, "<>", asMap())
	assertDecoded(t, "< >", asMap())
	assertDecoded(t, "<true = false>", asMap(true, false))
	assertDecoded(t, "<true=false>", asMap(true, false))
	assertDecoded(t, "  <    true   =    false   >  ", asMap(true, false))
	assertDecoded(t, "<true = false false = true>", asMap(true, false, false, true))
}

func TestMixedContainers(t *testing.T) {
	assertDecoded(t, "[<>]", asList(asMap()))
	assertDecoded(t, "[<true = [nil nil]>]", asList(asMap(true, asList(NilValue, NilValue))))
}

// TODO: Wait for metadata to be implemented properly
// func TestMetadataMap(t *testing.T) {
// 	assertDecoded(t, "()", asMap())
// 	assertDecoded(t, "( )", asMap())
// 	assertDecoded(t, "(true = false)", asMap(true, false))
// 	assertDecoded(t, "(true=false)", asMap(true, false))
// 	assertDecoded(t, "  (    true   =    false   )  ", asMap(true, false))
// 	assertDecoded(t, "(true = false false = true)", asMap(true, false, false, true))
// }
