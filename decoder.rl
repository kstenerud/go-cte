package cte

import (
    "fmt"
    "math"
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
    OnListBegin() error
    OnMapBegin() error
    OnMetadataMapBegin() error
    OnContainerEnd() error
    OnBytesBegin() error
    OnStringBegin() error
    OnURIBegin() error
    OnCommentBegin() error
    OnArrayData(bytes []byte) error
    OnArrayEnd() error
    OnMarker(id string) error
    OnReference(id string) error
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
        err = callbacks.OnNil()
        if err != nil {
            fbreak;
        }
    };

    true = "@true" @{
        err = callbacks.OnBool(true)
        if err != nil {
            fbreak;
        }
    };

    false = "@false" @{
        err = callbacks.OnBool(false)
        if err != nil {
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
        this.significandSign = 1
        this.significand = 0
        if err != nil {
            fbreak;
        }
    };

    float_decimal = numeric_first_component '.' fractional_digit_decimal+ exponent_decimal? %{
        err = callbacks.OnDecimalFloat(int64(this.significand) * int64(this.significandSign), (this.exponent+this.exponentAdjust) * this.exponentSign)
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.exponentSign = 1
        this.exponent = 0
        if err != nil {
            fbreak;
        }
    };

    float_hex = integer_hex '.' fractional_digit_hex+ exponent_hex %{
        err = callbacks.OnFloat(float64(this.significandSign) *
                    float64(this.significand) *
                    math.Pow(2.0, float64((this.exponent * this.exponentSign + this.exponentAdjust))))
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.exponentSign = 1
        this.exponent = 0
        if err != nil {
            fbreak;
        }
    };

    float = float_decimal | float_hex;

    inf = significand_sign? "@inf" %{
        err = callbacks.OnFloat(math.Inf(this.significandSign))
        this.significandSign = 1
        if err != nil {
            fbreak;
        }
    };

    nan = "@nan" %{
        err = callbacks.OnFloat(math.NaN())
        if err != nil {
            fbreak;
        }
    };
    snan = "@snan" %{
        // Just map it to regular NaN
        err = callbacks.OnFloat(math.NaN())
        if err != nil {
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
        err = callbacks.OnDate(int(this.significand) * this.significandSign, this.month, this.day)
        this.significandSign = 1
        this.significand = 0
        this.month = 0
        this.day = 0
        if err != nil {
            fbreak;
        }
    };

    time = time_tz_portion %{
        err = callbacks.OnTimeTZ(this.hour,
                this.minute,
                this.second,
                this.subsecond * this.subsecondMultiplier,
                string(this.timezone))
        this.hour = 0
        this.minute = 0
        this.second = 0
        this.subsecond = 0
        this.subsecondMultiplier = 1000000000
        this.timezone = this.timezone[:]
        if err != nil {
            fbreak;
        }
    };

    timestamp = date_portion '/' time_tz_portion %{
        err = callbacks.OnTimestampTZ(int(this.significand) * this.significandSign,
                this.month,
                this.day,
                this.hour,
                this.minute,
                this.second,
                this.subsecond * this.subsecondMultiplier,
                string(this.timezone))
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
        if err != nil {
            fbreak;
        }
    };

    list = '[' @{
        err = callbacks.OnListBegin()
        if err != nil {
            fbreak;
        }
        fcall list_iterate;
    };

    map = '{' @{
        err = callbacks.OnMapBegin()
        if err != nil {
            fbreak;
        }
        fcall map_iterate;
    };

    metadata_map = '(' @{
        err = callbacks.OnMetadataMapBegin()
        if err != nil {
            fbreak;
        }
        fcall metadata_map_iterate;
    };

    comment = "//" @{
        this.arrayStart = fpc + 1
        err = callbacks.OnCommentBegin()
        if err != nil {
            fbreak;
        }
        fcall comment_iterate;
    };

    multiline_comment = "/*" @{
        if this.commentDepth == 0 {
            err = callbacks.OnCommentBegin()
        } else {
            err = callbacks.OnArrayData(this.data[this.arrayStart:fpc+1])
        }
        this.arrayStart = fpc + 1
        this.commentDepth++
        if err != nil {
            fbreak;
        }
        fcall multiline_comment_iterate;
    };

    string = '"' @{
        this.arrayStart = fpc + 1
        err = callbacks.OnStringBegin()
        if err != nil {
            fbreak;
        }
        fcall string_iterate;
    };

    unquoted_string = unquoted_safe_first_char ('"' | unquoted_safe*) >{
        this.arrayStart = fpc - utfCharWidth
    } %{
        if this.data[fpc-1] != '"' {
            err = callbacks.OnStringBegin()
            if err != nil {
                fmt.Printf("Err %v\n", err)
                fbreak;
            }
            err = callbacks.OnArrayData(this.data[this.arrayStart:fpc])
            if err != nil {
                fmt.Printf("Err %v\n", err)
                fbreak;
            }
            err = callbacks.OnArrayEnd()
            if err != nil {
                fmt.Printf("Err %v\n", err)
                fbreak;
            }
        }
    };

    uri = 'u' '"' @{
        this.arrayStart = fpc + 1
        err = callbacks.OnURIBegin()
        if err != nil {
            fbreak;
        }
        fcall uri_iterate;
    };

    binary_hex = 'h' '"' @{
        err = callbacks.OnBytesBegin()
        if err != nil {
            fbreak;
        }
        fcall binary_hex_iterate;
    };

    binary_base64 = 'b' '"' @{
        err = callbacks.OnBytesBegin()
        if err != nil {
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
        err = callbacks.OnArrayData(this.data[this.arrayStart:fpc])
        if err != nil {
            fbreak;
        }
        err = callbacks.OnArrayEnd()
        if err != nil {
            fbreak;
        }
        fret;
    };

    multiline_comment_iterate := any* multiline_comment? "*/" @{
        err = callbacks.OnArrayData(this.data[this.arrayStart:fpc-1])
        if err != nil {
            fbreak;
        }
        this.arrayStart = fpc-1
        this.commentDepth--
        if this.commentDepth == 0 {
            err = callbacks.OnArrayEnd()
            if err != nil {
                fbreak;
            }
        }
        fret;
    };

    string_sequence =  (any - ('"' | '\\') ) |
    ('\\'
        (
            ('\\' @{
                err = this.flushAndAddEscapedCharacter(fpc-1, '\\', callbacks)
                if err != nil {
                    fbreak;
                }
            }) |
            ('n' @{
                err = this.flushAndAddEscapedCharacter(fpc-1, '\n', callbacks)
                if err != nil {
                    fbreak;
                }
            }) |
            ('r' @{
                err = this.flushAndAddEscapedCharacter(fpc-1, '\r', callbacks)
                if err != nil {
                    fbreak;
                }
            }) |
            ('t' @{
                err = this.flushAndAddEscapedCharacter(fpc-1, '\t', callbacks)
                if err != nil {
                    fbreak;
                }
            }) |
            ('"' @{
                err = this.flushAndAddEscapedCharacter(fpc-1, '"', callbacks)
                if err != nil {
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
        err = callbacks.OnArrayData(this.data[this.arrayStart:fpc])
        if err != nil {
            fbreak;
        }
        err = callbacks.OnArrayEnd()
        if err != nil {
            fbreak;
        }
        fret;
    };

    uri_iterate := [ !#-~]+ '"' @{
        err = callbacks.OnArrayData(this.data[this.arrayStart:fpc])
        if err != nil {
            fbreak;
        }
        err = callbacks.OnArrayEnd()
        if err != nil {
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
        err = callbacks.OnArrayData(this.binaryData)
        if err != nil {
            fbreak;
        }
        this.binaryData = this.binaryData[:0]
        err = callbacks.OnArrayEnd()
        if err != nil {
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
        err = callbacks.OnArrayData(this.binaryData)
        if err != nil {
            fbreak;
        }
        this.binaryData = this.binaryData[:0]
        err = callbacks.OnArrayEnd()
        if err != nil {
            fbreak;
        }
        fret;
    };

    list_iterate := ( value (ws value)* )? ws* ']' @{
        err = callbacks.OnContainerEnd()
        if err != nil {
            fbreak;
        }
        fret;
    };

    map_iterate := ( kv_pair (ws kv_pair)* )? ws* '}' @{
        err = callbacks.OnContainerEnd()
        if err != nil {
            fbreak;
        }
        fret;
    };

    metadata_map_iterate := ( kv_pair (ws kv_pair)* )? ws* ')' @{
        err = callbacks.OnContainerEnd()
        if err != nil {
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
    err := callbacks.OnArrayData(this.data[this.arrayStart:index])
    if err != nil {
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
