package cte

import (
	"testing"
	// "time"
)

func TestBroken(t *testing.T) {
	assertDecodeFails(t, "$")
}

func TestWhitespace(t *testing.T) {
	assertDecoded(t, " nil", NilValue)
	assertDecoded(t, "nil ", NilValue)
	assertDecoded(t, " \tnil \t ", NilValue)

	assertDecodeFails(t, "\b")
	assertDecodeFails(t, "\f")
	assertDecodeFails(t, "\v")
}

func TestNil(t *testing.T) {
	assertDecoded(t, "nil", NilValue)
}

func TestBool(t *testing.T) {
	assertDecoded(t, "true", true)
	assertDecoded(t, "false", false)
}

func TestInteger(t *testing.T) {
	assertDecoded(t, "0", 0)
	assertDecoded(t, "1", 1)
	assertDecoded(t, "-1", -1)
	assertDecoded(t, "934757932785", 934757932785)
	assertDecoded(t, "-844452230", -844452230)
}

func TestString(t *testing.T) {
	assertDecoded(t, "\"This is a string\"", "This is a string")
	assertDecoded(t, "\"String\\twith\\nnewlines\\rcrs\\\" and \\\\ quotes\"", "String\twith\nnewlines\rcrs\" and \\ quotes")
}

func TestList(t *testing.T) {
	assertDecoded(t, "[]", asList())
	assertDecoded(t, "[ ]", asList())
	assertDecoded(t, "[true]", asList(true))
	assertDecoded(t, "[true false]", asList(true, false))
	assertDecoded(t, "  [  true   false ]   ", asList(true, false))
	assertDecoded(t, "[true false  \"a string\" ]", asList(true, false, "a string"))
}

func TestUnorderedMap(t *testing.T) {
	assertDecoded(t, "{}", asMap())
	assertDecoded(t, "{ }", asMap())
	assertDecoded(t, "{true = false}", asMap(true, false))
	assertDecoded(t, "{true=false}", asMap(true, false))
	assertDecoded(t, "  {    true   =    false   }  ", asMap(true, false))
	assertDecoded(t, "{true = false false = true}", asMap(true, false, false, true))
	assertDecoded(t, "{\"true\"=true \"false\"=false}", asMap("true", true, "false", false))

	assertDecodeFails(t, "{true}")
	assertDecodeFails(t, "{true =}")
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

func TestMetadataMap(t *testing.T) {
	assertDecoded(t, "()", asMap())
	assertDecoded(t, "( )", asMap())
	assertDecoded(t, "(true = false)", asMap(true, false))
	assertDecoded(t, "(true=false)", asMap(true, false))
	assertDecoded(t, "  (    true   =    false   )  ", asMap(true, false))
	assertDecoded(t, "(true = false false = true)", asMap(true, false, false, true))
}

func TestComment(t *testing.T) {
	assertDecoded(t, "// This is a comment\n", " This is a comment")
	assertDecoded(t, "/* This is a comment */", " This is a comment ")
	assertDecoded(t, "/* /* This is a comment */ */", " /* This is a comment */ ")
}
