package cte

import (
	"math"
	"testing"
)


func TestBroken(t *testing.T) {
	assertDecodeFails(t, "$")
}

func TestWhitespace(t *testing.T) {
	assertDecoded(t, " @nil", NilValue, 0)
	assertDecoded(t, "@nil ", NilValue, 0)
	assertDecoded(t, " \t@nil \t ", NilValue, 0)

	assertDecodeFails(t, "\b")
	assertDecodeFails(t, "\f")
	assertDecodeFails(t, "\v")
}

func TestNil(t *testing.T) {
	assertDecoded(t, "@nil", NilValue, 0)
}

func TestBool(t *testing.T) {
	assertDecoded(t, "@true", true, 0)
	assertDecoded(t, "@false", false, 0)
}

func TestInteger(t *testing.T) {
	assertDecoded(t, "0", 0, 0)
	assertDecoded(t, "1", 1, 0)
	assertDecoded(t, "-1", -1, 0)
	assertDecoded(t, "934757932785", 934757932785, 0)
	assertDecoded(t, "-844452230", -844452230, 0)
}

func TestIntegerBinary(t *testing.T) {
	assertDecoded(t, "0b0", 0, 0)
	assertDecoded(t, "0b1", 1, 0)
	assertDecoded(t, "0b100111110100", 0x9f4, 0)
}

func TestIntegerOctal(t *testing.T) {
	assertDecoded(t, "0o0", 0, 0)
	assertDecoded(t, "0o1", 1, 0)
	assertDecoded(t, "0o755", 0x1ed, 0)
}

func TestIntegerHex(t *testing.T) {
	assertDecoded(t, "0x0", 0, 0)
	assertDecoded(t, "0x1", 1, 0)
	assertDecoded(t, "0x123456789abcdef", 0x123456789abcdef, 0)
}

func TestFloat(t *testing.T) {
	assertDecoded(t, "0.1", 0.1, 0)
	assertDecoded(t, "1.0", 1.0, 0)
	assertDecoded(t, "51.34", 51.34, 0)
	assertDecoded(t, "-51.34", -51.34, 0)
	assertDecoded(t, "0510.03040", 510.0304, 0.00000001)
	assertDecoded(t, "4.5104e-20", 4.5104e-20, 0.000000001)
	assertDecoded(t, "4.5104e20", 4.5104e20, 0.000000001)
	assertDecoded(t, "4.5104e+20", 4.5104e+20, 0.000000001)
	assertDecoded(t, "-4.5104e+20", -4.5104e+20, 0.000000001)
	assertDecoded(t, "-1.49465123566589", -1.49465123566589, 0.0000000000000001)
}

func TestFloatHex(t *testing.T) {
	assertDecoded(t, "0x1.0p0", 0x1.0p0, 0)
	assertDecoded(t, "0x0.1p0", 0x0.1p0, 0)
	assertDecoded(t, "0x1.f9a773c1p0", 0x1.f9a773c1p0, 0)
	assertDecoded(t, "-0xf.9a05p41", -0xf.9a05p41, 0)
	assertDecoded(t, "-0xf.9a05p+41", -0xf.9a05p+41, 0)
	assertDecoded(t, "-0xf.9a05p-41", -0xf.9a05p-41, 0)
	assertDecoded(t, "0x1.0p-1", 0x1.0p-1, 0)
}

func TestInf(t *testing.T) {
	assertDecoded(t, "@inf", math.Inf(1), 0)
	assertDecoded(t, "-@inf", math.Inf(-1), 0)
}

func TestNan(t *testing.T) {
	assertDecoded(t, "@nan", math.NaN(), 0)
	// Note: snan is converted to regular nan
	assertDecoded(t, "@snan", math.NaN(), 0)
}

func TestString(t *testing.T) {
	assertDecoded(t, "\"This is a string\"", "This is a string", 0)
	assertDecoded(t, "\"String\\twith\\nnewlines\\rcrs\\\" and \\\\ quotes\"", "String\twith\nnewlines\rcrs\" and \\ quotes", 0)
}

func TestUnquotedString(t *testing.T) {
	assertDecoded(t, "a", "a", 0)
	assertDecoded(t, "b", "b", 0)
	assertDecoded(t, "u", "u", 0)
	assertDecoded(t, "abcd", "abcd", 0)
	assertDecoded(t, "_a_b_c_d_", "_a_b_c_d_", 0)
	assertDecoded(t, "number_8", "number_8", 0)
	// assertDecoded(t, "飲み物", "飲み物", 0)
}

func TestURI(t *testing.T) {
	assertDecoded(t, "u\"http://example.com\"", newURL("http://example.com"), 0)
	assertDecoded(t, "u\"mailto:me@me.com\"", newURL("mailto:me@me.com"), 0)
	assertDecoded(t, "u\"urn:oasis:names:specification:docbook:dtd:xml:4.1.2\"", newURL("urn:oasis:names:specification:docbook:dtd:xml:4.1.2"), 0)
}

func TestBinaryHex(t *testing.T) {
	assertDecoded(t, "h\"1f9455\"", []byte{0x1f, 0x94, 0x55}, 0)
}

func TestDate(t *testing.T) {
	assertDecoded(t, "1000-10-1", newDate(1000, 10, 1), 0)
}

func TestTime(t *testing.T) {
	assertDecoded(t, "10:45:01.3014234/Europe/Paris", newTimeTZ(10, 45, 1, 301423400, "Europe/Paris"), 0)
}

func TestTimestamp(t *testing.T) {
	assertDecoded(t, "2001-1-2/5:08:09.999/America/Los_Angeles", newTimestampTZ(2001, 1, 2, 5, 8, 9, 999000000, "America/Los_Angeles"), 0)
}

func TestList(t *testing.T) {
	assertDecoded(t, "[]", asList(), 0)
	assertDecoded(t, "[ ]", asList(), 0)
	assertDecoded(t, "[@true]", asList(true), 0)
	assertDecoded(t, "[@true @false]", asList(true, false), 0)
	assertDecoded(t, "  [  @true   @false ]   ", asList(true, false), 0)
	assertDecoded(t, "[@true @false  \"a string\" ]", asList(true, false, "a string"), 0)
}

func TestUnorderedMap(t *testing.T) {
	assertDecoded(t, "{}", asMap(), 0)
	assertDecoded(t, "{ }", asMap(), 0)
	assertDecoded(t, "{@true = @false}", asMap(true, false), 0)
	assertDecoded(t, "{@true=@false}", asMap(true, false), 0)
	assertDecoded(t, "  {    @true   =    @false   }  ", asMap(true, false), 0)
	assertDecoded(t, "{@true = @false @false = @true}", asMap(true, false, false, true), 0)
	assertDecoded(t, "{\"true\"=@true \"false\"=@false}", asMap("true", true, "false", false), 0)

	assertDecodeFails(t, "{@true}")
	assertDecodeFails(t, "{@true =}")
}

func TestOrderedMap(t *testing.T) {
	assertDecoded(t, "<>", asMap(), 0)
	assertDecoded(t, "< >", asMap(), 0)
	assertDecoded(t, "<@true = @false>", asMap(true, false), 0)
	assertDecoded(t, "<@true=@false>", asMap(true, false), 0)
	assertDecoded(t, "  <    @true   =    @false   >  ", asMap(true, false), 0)
	assertDecoded(t, "<@true = @false @false = @true>", asMap(true, false, false, true), 0)
}

func TestMixedContainers(t *testing.T) {
	assertDecoded(t, "[<>]", asList(asMap()), 0)
	assertDecoded(t, "[<@true = [@nil @nil]>]", asList(asMap(true, asList(NilValue, NilValue))), 0)
}

func TestMetadataMap(t *testing.T) {
	assertDecoded(t, "()", asMap(), 0)
	assertDecoded(t, "( )", asMap(), 0)
	assertDecoded(t, "(@true = @false)", asMap(true, false), 0)
	assertDecoded(t, "(@true=@false)", asMap(true, false), 0)
	assertDecoded(t, "  (    @true   =    @false   )  ", asMap(true, false), 0)
	assertDecoded(t, "(@true = @false @false = @true)", asMap(true, false, false, true), 0)
}

func TestComment(t *testing.T) {
	assertDecoded(t, "// This is a comment\n", " This is a comment", 0)
	assertDecoded(t, "/* This is a comment */", " This is a comment ", 0)
	assertDecoded(t, "/* /* This is a comment */ */", " /* This is a comment */ ", 0)
}
