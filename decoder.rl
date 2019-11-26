package cte

import (
    "fmt"
    "math"
)

type ContainerType int
const (
	ContainerTypeList = ContainerType(iota)
	ContainerTypeMap
	ContainerTypeMetadataMap
)

type ArrayType int
const (
	ArrayTypeBinary = ArrayType(iota)
	ArrayTypeString
	ArrayTypeURI
	ArrayTypeComment
)

type CteDecoderCallbacks interface {
    OnNil() error
    OnBool(value bool) error
    OnPositiveInt(value uint64) error
    OnNegativeInt(value uint64) error
    OnDecimalFloat(significand int64, exponent int) error
    OnFloat(value float64) error
    OnDate(year, month, day int) error
    OnTimeTZ(hour, minute, second, nanosecond int, tz string) error
    OnTimeLoc(hour, minute, second, nanosecond int, latitude, longitude float32) error
    OnTimestampTZ(year, month, day, hour, minute, second, nanosecond int, tz string) error
    OnTimestampLoc(year, month, day, hour, minute, second, nanosecond int, latitude, longitude float32) error
    OnContainerBegin(containerType ContainerType) error
    OnContainerEnd() error
    OnArrayBegin(arrayType ArrayType) error
    OnArrayData(bytes []byte) error
    OnArrayEnd() error
}

%%{
    machine cte;
    access this.;

    ws = [\r\n\t ];

    utf8 = (0xc2..0xdf 0x80..0xbf) @{utfCharWidth = 2} |
           (0xe0..0xef 0x80..0xbf 0x80..0xbf) @{utfCharWidth = 3} |
           (0xf0..0xf4 0x80..0xbf 0x80..0xbf 0x80..0xbf) @{utfCharWidth = 4};

    unquoted_safe_first_char = ([A-Za-z_] @{utfCharWidth = 1}) | utf8;
    unquoted_safe = unquoted_safe_first_char | [0-9];

    nil = "@nil" @{
        if err = callbacks.OnNil(); err != nil {
            fbreak;
        }
    };

    true = "@true" @{
        if err = callbacks.OnBool(true); err != nil {
            fbreak;
        }
    };

    false = "@false" @{
        if err = callbacks.OnBool(false); err != nil {
            fbreak;
        }
    };

    leading_zeroes = 0*;

    significand_sign = '-' @{
        this.significandSign = -1
    };

    exponent_sign = '+' | ('-' @{
        this.exponentSign = -1
    });

    whole_digit_decimal = [0-9] @{
        this.significand = this.significand * 10 + uint64(fc - '0')
    };

    whole_digit_hex = [0-9] @{
        this.significand = (this.significand << 4) | uint64(fc - '0')
    } | [a-f] @{
        this.significand = (this.significand << 4) | uint64(fc - 'a' + 10)
    };

    fractional_digit_decimal = [0-9] @{
        this.significand = this.significand * 10 + uint64(fc - '0')
        this.exponentAdjust--
    };

    fractional_digit_hex = [0-9] @{
        this.significand = (this.significand << 4) | uint64(fc - '0')
        this.exponentAdjust -= 4
    } | [a-f] @{
        this.significand = (this.significand << 4) | uint64(fc - 'a' + 10)
        this.exponentAdjust -= 4
    };

    exponent_digit = [0-9] @{
        this.exponent = this.exponent * 10 + int(fc - '0')
    };

    exponent_decimal = 'e' exponent_sign? exponent_digit+;
    exponent_hex = 'p' exponent_sign? exponent_digit+;

    numeric_first_component = significand_sign? leading_zeroes whole_digit_decimal+;

    integer_binary = significand_sign? "0b" [0-1]+ @{
        this.significand = (this.significand << 1) | uint64(fc - '0')
    };

    integer_octal = significand_sign? "0o" [0-7]+ @{
        this.significand = (this.significand << 3) | uint64(fc - '0')
    };

    integer_hex = significand_sign? "0x" whole_digit_hex+;

    integer = (numeric_first_component | integer_binary | integer_octal | integer_hex) %{
        if this.significandSign >= 0 {
            err = callbacks.OnPositiveInt(this.significand)
        } else {
            err = callbacks.OnNegativeInt(this.significand)
        }
        if err != nil {
            fbreak;
        }
        this.significandSign = 1
        this.significand = 0
    };

    float_decimal = numeric_first_component '.' fractional_digit_decimal+ exponent_decimal? %{
        if err = callbacks.OnDecimalFloat(int64(this.significand) * int64(this.significandSign),
        				(this.exponent+this.exponentAdjust) * this.exponentSign); err != nil {
            fbreak;
        }
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.exponentSign = 1
        this.exponent = 0
    };

    float_hex = integer_hex '.' fractional_digit_hex+ exponent_hex %{
        if err = callbacks.OnFloat(float64(this.significandSign) *
                    float64(this.significand) *
                    math.Pow(2.0, float64((this.exponent * this.exponentSign + this.exponentAdjust)))); err != nil {
            fbreak;
        }
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.exponentSign = 1
        this.exponent = 0
    };

    float = float_decimal | float_hex;

    inf = significand_sign? "@inf" %{
        if err = callbacks.OnFloat(math.Inf(this.significandSign)); err != nil {
            fbreak;
        }
        this.significandSign = 1
    };

    nan = "@nan" %{
        if err = callbacks.OnFloat(math.NaN()); err != nil {
            fbreak;
        }
    };
    snan = "@snan" %{
        // Just map it to regular NaN
        if err = callbacks.OnFloat(math.NaN()); err != nil {
            fbreak;
        }
    };

    month = [0-9]{1,2} ${
        this.month = this.month * 10 + int(fc - '0')
    };

    day = [0-9]{1,2} ${
        this.day = this.day * 10 + int(fc - '0')
    };

    hour = [0-9]{1,2} ${
        this.hour = this.hour * 10 + int(fc - '0')
    };

    minute = [0-9]{2} ${
        this.minute = this.minute * 10 + int(fc - '0')
    };

    second = [0-9]{2} ${
        this.second = this.second * 10 + int(fc - '0')
    };

    subsecond = [0-9]{1,9} ${
        this.subsecond = this.subsecond * 10 + int(fc - '0')
        this.subsecondMultiplier /= 10
    };

    timezone = [a-zA-Z0-9+_/\-]+ @{
        this.timezone = append(this.timezone, fc)
    };

    date_portion = numeric_first_component '-' month '-' day;
    time_tz_portion = hour ':' minute ':' second ('.' subsecond)? ('/' timezone)?;

    date = date_portion %{
        if err = callbacks.OnDate(int(this.significand) * this.significandSign, this.month, this.day); err != nil {
            fbreak;
        }
        this.significandSign = 1
        this.significand = 0
        this.month = 0
        this.day = 0
    };

    time = time_tz_portion %{
        if err = callbacks.OnTimeTZ(this.hour,
                this.minute,
                this.second,
                this.subsecond * this.subsecondMultiplier,
                string(this.timezone)); err != nil {
            fbreak;
        }
        this.hour = 0
        this.minute = 0
        this.second = 0
        this.subsecond = 0
        this.subsecondMultiplier = 1000000000
        this.timezone = this.timezone[:]
    };

    timestamp = date_portion '/' time_tz_portion %{
        if err = callbacks.OnTimestampTZ(int(this.significand) * this.significandSign,
                this.month,
                this.day,
                this.hour,
                this.minute,
                this.second,
                this.subsecond * this.subsecondMultiplier,
                string(this.timezone)); err != nil {
            fbreak;
        }
        this.significandSign = 1
        this.significand = 0
        this.month = 0
        this.day = 0
        this.hour = 0
        this.minute = 0
        this.second = 0
        this.subsecond = 0
        this.subsecondMultiplier = 1000000000
        this.timezone = this.timezone[:]
    };

    list = '[' @{
        if err = callbacks.OnContainerBegin(ContainerTypeList); err != nil {
            fbreak;
        }
        fcall list_iterate;
    };

    map = '{' @{
        if err = callbacks.OnContainerBegin(ContainerTypeMap); err != nil {
            fbreak;
        }
        fcall map_iterate;
    };

    metadata_map = '(' @{
        if err = callbacks.OnContainerBegin(ContainerTypeMetadataMap); err != nil {
            fbreak;
        }
        fcall metadata_map_iterate;
    };

    comment = "//" @{
        this.arrayStart = fpc + 1
        if err = callbacks.OnArrayBegin(ArrayTypeComment); err != nil {
            fbreak;
        }
        fcall comment_iterate;
    };

    multiline_comment = "/*" @{
        if this.commentDepth == 0 {
            err = callbacks.OnArrayBegin(ArrayTypeComment)
        } else {
            err = callbacks.OnArrayData(this.data[this.arrayStart:fpc+1])
        }
        if err != nil {
            fbreak;
        }
        this.arrayStart = fpc + 1
        this.commentDepth++
        fcall multiline_comment_iterate;
    };

    string = '"' @{
        this.arrayStart = fpc + 1
        if err = callbacks.OnArrayBegin(ArrayTypeString); err != nil {
            fbreak;
        }
        fcall string_iterate;
    };

    unquoted_string = unquoted_safe_first_char ('"' | unquoted_safe*) >{
        this.arrayStart = fpc - utfCharWidth
    } %{
        if this.data[fpc-1] != '"' {
            if err = callbacks.OnArrayBegin(ArrayTypeString); err != nil {
                fmt.Printf("Err %v\n", err)
                fbreak;
            }
            if err = callbacks.OnArrayData(this.data[this.arrayStart:fpc]); err != nil {
                fmt.Printf("Err %v\n", err)
                fbreak;
            }
            if err = callbacks.OnArrayEnd(); err != nil {
                fmt.Printf("Err %v\n", err)
                fbreak;
            }
        }
    };

    uri = 'u' '"' @{
        this.arrayStart = fpc + 1
        if err = callbacks.OnArrayBegin(ArrayTypeURI); err != nil {
            fbreak;
        }
        fcall uri_iterate;
    };

    binary_hex = 'h' '"' @{
        if err = callbacks.OnArrayBegin(ArrayTypeBinary); err != nil {
            fbreak;
        }
        fcall binary_hex_iterate;
    };

    binary_base64 = 'b' '"' @{
        if err = callbacks.OnArrayBegin(ArrayTypeBinary); err != nil {
            fbreak;
        }
        fcall binary_base64_iterate;
    };


    keyable = true | false | integer | float | inf | string | unquoted_string | uri | binary_hex | binary_base64 | date | time | timestamp;
    nonkeyable = nil | list | map | nan | snan;
    object_pre = (ws | metadata_map | comment | multiline_comment)*;
    object_post = (ws | comment | multiline_comment)*;
    value = object_pre (keyable | nonkeyable);
    kv_pair = object_pre keyable object_post '=' value;

    comment_iterate := [^\n]* '\n' @{
        if err = callbacks.OnArrayData(this.data[this.arrayStart:fpc]); err != nil {
            fbreak;
        }
        if err = callbacks.OnArrayEnd(); err != nil {
            fbreak;
        }
        fret;
    };

    multiline_comment_iterate := any* multiline_comment? "*/" @{
        if err = callbacks.OnArrayData(this.data[this.arrayStart:fpc-1]); err != nil {
            fbreak;
        }
        this.arrayStart = fpc-1
        this.commentDepth--
        if this.commentDepth == 0 {
            if err = callbacks.OnArrayEnd(); err != nil {
                fbreak;
            }
        }
        fret;
    };

    string_sequence =  (any - ('"' | '\\') ) |
    ('\\'
        (
            ('\\' @{
                if err = this.flushAndAddEscapedCharacter(fpc-1, '\\', callbacks); err != nil {
                    fbreak;
                }
            }) |
            ('n' @{
                if err = this.flushAndAddEscapedCharacter(fpc-1, '\n', callbacks); err != nil {
                    fbreak;
                }
            }) |
            ('r' @{
                if err = this.flushAndAddEscapedCharacter(fpc-1, '\r', callbacks); err != nil {
                    fbreak;
                }
            }) |
            ('t' @{
                if err = this.flushAndAddEscapedCharacter(fpc-1, '\t', callbacks); err != nil {
                    fbreak;
                }
            }) |
            ('"' @{
                if err = this.flushAndAddEscapedCharacter(fpc-1, '"', callbacks); err != nil {
                    fbreak;
                }
            }) |
            ([^"nrt\\] @{
                return false, fmt.Errorf("\\%c: Illegal escape encoding", this.data[fpc])
            })
            # TODO: hex and unicode
        )
    );

    string_iterate := string_sequence* '"' @{
        if err = callbacks.OnArrayData(this.data[this.arrayStart:fpc]); err != nil {
            fbreak;
        }
        if err = callbacks.OnArrayEnd(); err != nil {
            fbreak;
        }
        fret;
    };

    uri_iterate := [ !#-~]+ '"' @{
        if err = callbacks.OnArrayData(this.data[this.arrayStart:fpc]); err != nil {
            fbreak;
        }
        if err = callbacks.OnArrayEnd(); err != nil {
            fbreak;
        }
        fret;
    };

    hex_09_hi = [0-9] @{
        this.binaryNext = (fc - '0') << 4
    };
    hex_af_hi = [a-f] @{
        this.binaryNext = (fc - 'a' + 10) << 4
    };
    hex_09_lo = [0-9] @{
        this.binaryNext |= fc - '0'
    };
    hex_af_lo = [a-f] @{
        this.binaryNext |= fc - 'a' + 10
    };
    binary_hex_sequence = ws* (hex_09_hi | hex_af_hi) ws* (hex_09_lo | hex_af_lo) @{
        this.binaryData = append(this.binaryData, this.binaryNext)
    };

    binary_hex_iterate := binary_hex_sequence* '"' @{
        if err = callbacks.OnArrayData(this.binaryData); err != nil {
            fbreak;
        }
        this.binaryData = this.binaryData[:0]
        if err = callbacks.OnArrayEnd(); err != nil {
            fbreak;
        }
        fret;
    };

    base64_AZ = [A-Z] @{
        this.binaryAccumulator = (this.binaryAccumulator << 6) | uint(fc - 'A')
    };
    base64_az = [a-z] @{
        this.binaryAccumulator = (this.binaryAccumulator << 6) | uint(fc - 'a' + 26)
    };
    base64_09 = [0-9] @{
        this.binaryAccumulator = (this.binaryAccumulator << 6) | uint(fc - '0' + 52)
    };
    base64_plus = '+' @{
        this.binaryAccumulator = (this.binaryAccumulator << 6) | 62
    };
    base64_slash = '/' @{
        this.binaryAccumulator = (this.binaryAccumulator << 6) | 63
    };

    base64_digit = ws* (base64_AZ | base64_az | base64_09 | base64_plus | base64_slash) @{
        this.base64Digits++
        if this.base64Digits == 4 {
            this.binaryData = append(this.binaryData, byte(this.binaryAccumulator >> 16))
            this.binaryData = append(this.binaryData, byte(this.binaryAccumulator >> 8))
            this.binaryData = append(this.binaryData, byte(this.binaryAccumulator))
            this.binaryAccumulator = 0
            this.base64Digits = 0
        }
    };

    base64_sequence = base64_digit* ws* '"' @{
        switch this.base64Digits {
            case 0:
                break
            case 1:
                // TODO: Invalid
            case 2:
                this.binaryData = append(this.binaryData, byte(this.binaryAccumulator >> 4))
            case 3:
                this.binaryData = append(this.binaryData, byte(this.binaryAccumulator >> 10))
                this.binaryData = append(this.binaryData, byte(this.binaryAccumulator >> 2))
        }
        this.binaryAccumulator = 0
        this.base64Digits = 0
    };

    binary_base64_iterate := base64_sequence @{
        if err = callbacks.OnArrayData(this.binaryData); err != nil {
            fbreak;
        }
        this.binaryData = this.binaryData[:0]
        if err = callbacks.OnArrayEnd(); err != nil {
            fbreak;
        }
        fret;
    };

    list_iterate := ( value (ws value)* )? ws* ']' @{
        if err = callbacks.OnContainerEnd(); err != nil {
            fbreak;
        }
        fret;
    };

    map_iterate := ( kv_pair (ws kv_pair)* )? ws* '}' @{
        if err = callbacks.OnContainerEnd(); err != nil {
            fbreak;
        }
        fret;
    };

    metadata_map_iterate := ( kv_pair (ws kv_pair)* )? ws* ')' @{
        if err = callbacks.OnContainerEnd(); err != nil {
            fbreak;
        }
        fret;
    };

    main := value ws* @/{
        err = fmt.Errorf("Incomplete document")
    };
}%%


%% write data;

type Parser struct {
    cs int // Current Ragel state
    ts int // Position: start of token
    te int // Position: end of token
    top int // Index of top of stack
    stack []int
    data []byte
    arrayStart int // Start of the current item of interest
    binaryData []byte
    binaryNext byte
    binaryAccumulator uint
    base64Digits int
    commentDepth int
    significand uint64
    significandSign int
    exponent int
    exponentSign int
    exponentAdjust int
    month int
    day int
    hour int
    minute int
    second int
    subsecond int
    subsecondMultiplier int
    timezone []byte
}

func (this *Parser) Init(maxDepth int) {
    this.stack = make([]int, maxDepth)
    this.significandSign = 1
    this.exponentSign = 1
    this.subsecondMultiplier = 1000000000
    this.timezone = make([]byte, 0, 40)
    this.binaryData = make([]byte, 0, 500)
}

func NewParser(maxDepth int) *Parser {
    this := new(Parser)
    this.Init(maxDepth)
    return this
}

func (this *Parser) flushByteArray(index int, callbacks CteDecoderCallbacks) error {
    if err := callbacks.OnArrayData(this.data[this.arrayStart:index]); err != nil {
        return err
    }
    this.arrayStart = index
    return nil
}

func (this *Parser) flushAndAddEscapedCharacter(index int, escapedCharacter byte, callbacks CteDecoderCallbacks) error {
    this.data[index] = escapedCharacter
    if err := this.flushByteArray(index+1, callbacks); err != nil {
        return err
    }
    // Get past escape initiator and escape char
    this.arrayStart = index + 2
    return nil
}

func (this *Parser) Parse(src []byte, callbacks CteDecoderCallbacks) (isComplete bool, err error) {
//  if this.ts > 0 {
        // TODO: Read from undeflow buffer
//  }
    this.data = src
    p := 0 // Position: current
    pe := len(this.data) // Position: end of buffer
    // TODO: Change to -1 and check for end of file
    eof := pe // Position: end of file
    utfCharWidth := 0

    _ = eof
    
    %%{
        write init;
        write exec;
    }%%

    if this.ts > 0 {
        // TODO: Copy to underflow buffer
        // arrayStart
        // ts doesn't seem to get used?
        // copy(this.underflow, data[ts:pe])
        // p = 0
        // pe = pe - ts
    }
    // TODO
    if this.cs == cte_error {
        err = fmt.Errorf("Parse error at %v", p)
    }
//    isComplete = this.cs == cte_parse_first;
    // TODO: Maybe there's no way to detect completion?
    isComplete = true
    return
}
