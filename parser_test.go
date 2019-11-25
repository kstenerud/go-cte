package cte

import (
	"math"
	"testing"
)

func TestBroken(t *testing.T) {
	assertDecodeFails(t, "$")
}

func TestWhitespace(t *testing.T) {
	assertDecoded(t, " @nil", NilValue)
	assertDecoded(t, "@nil ", NilValue)
	assertDecoded(t, " \t@nil \t ", NilValue)

	assertDecodeFails(t, "\b")
	assertDecodeFails(t, "\f")
	assertDecodeFails(t, "\v")
}

func TestNil(t *testing.T) {
	assertDecoded(t, "@nil", NilValue)
}

func TestBool(t *testing.T) {
	assertDecoded(t, "@true", true)
	assertDecoded(t, "@false", false)
}

func TestInteger(t *testing.T) {
	assertDecoded(t, "0", 0)
	assertDecoded(t, "1", 1)
	assertDecoded(t, "-1", -1)
	assertDecoded(t, "934757932785", 934757932785)
	assertDecoded(t, "-844452230", -844452230)
}

func TestIntegerBinary(t *testing.T) {
	assertDecoded(t, "0b0", 0)
	assertDecoded(t, "0b1", 1)
	assertDecoded(t, "0b100111110100", 0x9f4)
}

func TestIntegerOctal(t *testing.T) {
	assertDecoded(t, "0o0", 0)
	assertDecoded(t, "0o1", 1)
	assertDecoded(t, "0o755", 0x1ed)
}

func TestIntegerHex(t *testing.T) {
	assertDecoded(t, "0x0", 0)
	assertDecoded(t, "0x1", 1)
	assertDecoded(t, "0x123456789abcdef", 0x123456789abcdef)
}

func TestFloat(t *testing.T) {
	assertDecoded(t, "0.1", 0.1)
	assertDecoded(t, "1.0", 1.0)
	assertDecoded(t, "51.34", 51.34)
	assertDecoded(t, "-51.34", -51.34)
	assertDecoded(t, "0510.03040", 510.0304)
	assertDecoded(t, "4.5104e-20", 4.5104e-20)
	assertDecoded(t, "4.5104e20", 4.5104e20)
	assertDecoded(t, "4.5104e+20", 4.5104e+20)
	assertDecoded(t, "-4.5104e+20", -4.5104e+20)
	assertDecoded(t, "-1.49465123566589", -1.49465123566589)
}

func TestFloatHex(t *testing.T) {
	assertDecoded(t, "0x1.0p0", 0x1.0p0)
	assertDecoded(t, "0x0.1p0", 0x0.1p0)
	assertDecoded(t, "0x1.f9a773c1p0", 0x1.f9a773c1p0)
	assertDecoded(t, "-0xf.9a05p41", -0xf.9a05p41)
	assertDecoded(t, "-0xf.9a05p+41", -0xf.9a05p+41)
	assertDecoded(t, "-0xf.9a05p-41", -0xf.9a05p-41)
	assertDecoded(t, "0x1.0p-1", 0x1.0p-1)
}

func TestInf(t *testing.T) {
	assertDecoded(t, "@inf", math.Inf(1))
	assertDecoded(t, "-@inf", math.Inf(-1))
}

func TestNan(t *testing.T) {
	assertDecoded(t, "@nan", math.NaN())
	// Note: snan is converted to regular nan
	assertDecoded(t, "@snan", math.NaN())
}

func TestString(t *testing.T) {
	assertDecoded(t, "\"This is a string\"", "This is a string")
	assertDecoded(t, "\"String\\twith\\nnewlines\\rcrs\\\" and \\\\ quotes\"", "String\twith\nnewlines\rcrs\" and \\ quotes")
}

func TestUnquotedString(t *testing.T) {
	assertDecoded(t, "a", "a")
	assertDecoded(t, "b", "b")
	assertDecoded(t, "u", "u")
	assertDecoded(t, "abcd", "abcd")
	assertDecoded(t, "_a_b_c_d_", "_a_b_c_d_")
	assertDecoded(t, "number_8", "number_8")
	assertDecoded(t, "飲み物", "飲み物")
}

func TestURI(t *testing.T) {
	assertDecoded(t, "u\"http://example.com\"", asURL("http://example.com"))
	assertDecoded(t, "u\"mailto:me@me.com\"", asURL("mailto:me@me.com"))
	assertDecoded(t, "u\"urn:oasis:names:specification:docbook:dtd:xml:4.1.2\"", asURL("urn:oasis:names:specification:docbook:dtd:xml:4.1.2"))
}

func TestBinaryHex(t *testing.T) {
	assertDecoded(t, "h\"1f9455\"", []byte{0x1f, 0x94, 0x55})
}

func TestBase64Hex(t *testing.T) {
	assertBase64Decoded(t, []byte{0x11, 0xf4, 0x55})
	assertBase64Decoded(t, []byte{0xff, 0x00, 0x74, 0xa8, 0x1f, 0xff, 0x35, 0x3d, 0x98})
	assertBase64Decoded(t, []byte{0xff, 0x00, 0x74, 0xa8, 0x1f, 0xff, 0x35, 0x3d})
	assertBase64Decoded(t, []byte{0xff, 0x00, 0x74, 0xa8, 0x1f, 0xff, 0x35})
	assertBase64Decoded(t, []byte{0xff, 0x00, 0x74, 0xa8, 0x1f, 0xff})
	assertDecoded(t,  "b\" / w B 0 q B / / \"", []byte{0xff, 0x00, 0x74, 0xa8, 0x1f, 0xff})
}

func TestDate(t *testing.T) {
	assertDecoded(t, "1000-10-1", asDate(1000, 10, 1))
}

func TestTime(t *testing.T) {
	assertDecoded(t, "10:45:01.3014234/Europe/Paris", asTime(10, 45, 1, 301423400, "Europe/Paris"))
}

func TestTimestamp(t *testing.T) {
	assertDecoded(t, "2001-1-2/5:08:09.999/America/Los_Angeles", asTimestamp(2001, 1, 2, 5, 8, 9, 999000000, "America/Los_Angeles"))
}

func TestList(t *testing.T) {
	assertDecoded(t, "[]", asList())
	assertDecoded(t, "[ ]", asList())
	assertDecoded(t, "[@true]", asList(true))
	assertDecoded(t, "[@true @false]", asList(true, false))
	assertDecoded(t, "  [  @true   @false ]   ", asList(true, false))
	assertDecoded(t, "[@true @false  \"a string\" ]", asList(true, false, "a string"))
}

func TestMap(t *testing.T) {
	assertDecoded(t, "{}", asMap())
	assertDecoded(t, "{ }", asMap())
	assertDecoded(t, "{@true = @false}", asMap(true, false))
	assertDecoded(t, "{@true=@false}", asMap(true, false))
	assertDecoded(t, "  {    @true   =    @false   }  ", asMap(true, false))
	assertDecoded(t, "{@true = @false @false = @true}", asMap(true, false, false, true))
	assertDecoded(t, "{\"true\"=@true \"false\"=@false}", asMap("true", true, "false", false))

	assertDecodeFails(t, "{@true}")
	assertDecodeFails(t, "{@true =}")
}

func TestMixedContainers(t *testing.T) {
	assertDecoded(t, "[{}]", asList(asMap()))
	assertDecoded(t, "[{@true = [@nil @nil]}]", asList(asMap(true, asList(NilValue, NilValue))))
}

func TestMetadata(t *testing.T) {
	assertDecoded(t, "()", asMetadataMap())
	assertDecoded(t, "(@true = @false)", asMetadataMap(true, false))
	assertDecoded(t, "(@true=@false)", asMetadataMap(true, false))
	assertDecoded(t, "  (    @true   =    @false   )  ", asMetadataMap(true, false))
	assertDecoded(t, "(@true = @false @false = @true)", asMetadataMap(true, false, false, true))
}

func TestComment(t *testing.T) {
	assertDecoded(t, "// This is a comment\n", " This is a comment")
	assertDecoded(t, "/* This is a comment */", " This is a comment ")
	assertDecoded(t, "/* /* This is a comment */ */", " /* This is a comment */ ")
}
