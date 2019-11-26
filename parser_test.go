package cte

import (
	"testing"
)

func TestBroken(t *testing.T) {
	assertDecodeFails(t, "$")
}

func TestWhitespace(t *testing.T) {
	assertDecoded(t, " @nil", Nil())
	assertDecoded(t, "@nil ", Nil())
	assertDecoded(t, " \t@nil \t ", Nil())

	assertDecodeFails(t, "\b")
	assertDecodeFails(t, "\f")
	assertDecodeFails(t, "\v")
}

func TestNil(t *testing.T) {
	assertDecoded(t, "@nil", Nil())
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
	assertDecoded(t, "@inf", Inf())
	assertDecoded(t, "-@inf", NInf())
}

func TestNan(t *testing.T) {
	assertDecoded(t, "@nan", Nan())
	assertDecoded(t, "@snan", SNan())
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
	assertDecoded(t, "u\"http://example.com\"", URI("http://example.com"))
	assertDecoded(t, "u\"mailto:me@me.com\"", URI("mailto:me@me.com"))
	assertDecoded(t, "u\"urn:oasis:names:specification:docbook:dtd:xml:4.1.2\"", URI("urn:oasis:names:specification:docbook:dtd:xml:4.1.2"))
}

func TestBinaryHex(t *testing.T) {
	assertDecoded(t, "h\"1f9455\"", Bytes(0x1f, 0x94, 0x55))
}

func TestBase64Hex(t *testing.T) {
	assertBase64Decoded(t, Bytes(0x11, 0xf4, 0x55))
	assertBase64Decoded(t, Bytes(0xff, 0x00, 0x74, 0xa8, 0x1f, 0xff, 0x35, 0x3d, 0x98))
	assertBase64Decoded(t, Bytes(0xff, 0x00, 0x74, 0xa8, 0x1f, 0xff, 0x35, 0x3d))
	assertBase64Decoded(t, Bytes(0xff, 0x00, 0x74, 0xa8, 0x1f, 0xff, 0x35))
	assertBase64Decoded(t, Bytes(0xff, 0x00, 0x74, 0xa8, 0x1f, 0xff))
	assertDecoded(t,  "b\" / w B 0 q B / / \"", Bytes(0xff, 0x00, 0x74, 0xa8, 0x1f, 0xff))
}

func TestDate(t *testing.T) {
	assertDecoded(t, "1000-10-1", Date(1000, 10, 1))
}

func TestTime(t *testing.T) {
	assertDecoded(t, "10:45:01.3014234/Europe/Paris", Time(10, 45, 1, 301423400, "Europe/Paris"))
}

func TestTimestamp(t *testing.T) {
	assertDecoded(t, "2001-1-2/5:08:09.999/America/Los_Angeles", TS(2001, 1, 2, 5, 8, 9, 999000000, "America/Los_Angeles"))
}

func TestList(t *testing.T) {
	assertDecoded(t, "[]", List())
	assertDecoded(t, "[ ]", List())
	assertDecoded(t, "[@true]", List(true))
	assertDecoded(t, "[@true @false]", List(true, false))
	assertDecoded(t, "  [  @true   @false ]   ", List(true, false))
	assertDecoded(t, "[@true @false  \"a string\" ]", List(true, false, "a string"))
}

func TestMap(t *testing.T) {
	assertDecoded(t, "{}", Map())
	assertDecoded(t, "{ }", Map())
	assertDecoded(t, "{@true = @false}", Map(true, false))
	assertDecoded(t, "{@true=@false}", Map(true, false))
	assertDecoded(t, "  {    @true   =    @false   }  ", Map(true, false))
	assertDecoded(t, "{@true = @false @false = @true}", Map(true, false, false, true))
	assertDecoded(t, "{\"true\"=@true \"false\"=@false}", Map("true", true, "false", false))

	assertDecodeFails(t, "{@true}")
	assertDecodeFails(t, "{@true =}")
}

func TestMixedContainers(t *testing.T) {
	assertDecoded(t, "[{}]", List(Map()))
	assertDecoded(t, "[{@true = [@nil @nil]}]", List(Map(true, List(Nil(), Nil()))))
}

func TestMetadata(t *testing.T) {
	assertDecoded(t, "()", Meta())
	assertDecoded(t, "(@true = @false)", Meta(true, false))
	assertDecoded(t, "(@true=@false)", Meta(true, false))
	assertDecoded(t, "  (    @true   =    @false   )  ", Meta(true, false))
	assertDecoded(t, "(@true = @false @false = @true)", Meta(true, false, false, true))
}

func TestComment(t *testing.T) {
	assertDecoded(t, "// This is a comment\n", Comment(" This is a comment"))
	assertDecoded(t, "/* This is a comment */", Comment(" This is a comment "))
	assertDecoded(t, "/* /* This is a comment */ */", Comment(" /* This is a comment */ "))
}
