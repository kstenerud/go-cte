package cte

import (
    "fmt"
    "math"
//    "time"
)

type CteDecoderCallbacks interface {
    OnNil() error
    OnBool(value bool) error
    OnPositiveInt(value uint64) error
    OnNegativeInt(value uint64) error
    OnDecimalFloat(significand int64, exponent int) error
    OnFloat(value float64) error
//    OnDate(value time.Time) error
//    OnTime(value time.Time) error
//    OnTimestamp(value time.Time) error
    OnListBegin() error
    OnOrderedMapBegin() error
    OnUnorderedMapBegin() error
    OnMetadataMapBegin() error
    OnContainerEnd() error
//    OnBytesBegin() error
    OnStringBegin() error
    OnURIBegin() error
    OnCommentBegin() error
    OnArrayData(bytes []byte) error
    OnArrayEnd() error
}

%%{
    machine cte;
    access this.;

    ws = [\r\n\t ];

    utf8 = (0xc2..0xdf 0x80..0xbf) |
           (0xe0..0xef 0x80..0xbf 0x80..0xbf) |
           (0xf0..0xf4 0x80..0xbf 0x80..0xbf 0x80..0xbf);

    unquoted_safe_first_char = [A-Za-z_] | utf8;
    unquoted_safe = unquoted_safe_first_char | [0-9.] | '-';
    unquoted_string = unquoted_safe_first_char unquoted_safe*;

    nil = "nil" @{
        err = callbacks.OnNil()
        if err != nil {
            fbreak;
        }
    };

    true = "true" @{
        err = callbacks.OnBool(true)
        if err != nil {
            fbreak;
        }
    };

    false = "false" @{
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
            callbacks.OnPositiveInt(this.significand)
        } else {
            callbacks.OnNegativeInt(this.significand)
        }
        this.significandSign = 1
        this.significand = 0
    };

    float_decimal = numeric_first_component '.' fractional_digit_decimal+ exponent_decimal? %{
        callbacks.OnDecimalFloat(int64(this.significand) * int64(this.significandSign), (this.exponent+this.exponentAdjust) * this.exponentSign)
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.exponentSign = 1
        this.exponent = 0
    };

    float_hex = integer_hex '.' fractional_digit_hex+ exponent_hex %{
        exponent := float64((this.exponent * this.exponentSign + this.exponentAdjust))
        callbacks.OnFloat(float64(this.significandSign) * float64(this.significand) * math.Pow(2.0, exponent))
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.exponentSign = 1
        this.exponent = 0
    };

    float = float_decimal | float_hex;

    list = '[' @{
        err = callbacks.OnListBegin()
        if err != nil {
            fbreak;
        }
        fcall list_iterate;
    };

    unordered_map = '{' @{
        err = callbacks.OnUnorderedMapBegin()
        if err != nil {
            fbreak;
        }
        fcall unordered_map_iterate;
    };

    ordered_map = '<' @{
        err = callbacks.OnOrderedMapBegin()
        if err != nil {
            fbreak;
        }
        fcall ordered_map_iterate;
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
            if err != nil {
                fbreak;
            }
        } else {
            err = callbacks.OnArrayData(this.data[this.arrayStart:fpc+1])
            if err != nil {
                fbreak;
            }
        }
        this.arrayStart = fpc + 1
        this.commentDepth++
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

    uri = 'u' '"' @{
        this.arrayStart = fpc + 1
        err = callbacks.OnURIBegin()
        if err != nil {
            fbreak;
        }
        fcall uri_iterate;
    };


    keyable = true | false | integer | float | string | uri;
    nonkeyable = nil | list | unordered_map | ordered_map;
    object_pre = (ws | metadata_map | comment | multiline_comment)*;
    object_post = (ws | comment | multiline_comment)*;
    value = object_pre (keyable | nonkeyable);
    kv_pair = object_pre keyable object_post '=' value;

    comment_iterate :=
        [^\n]*
        '\n' @{
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

    multiline_comment_iterate :=
        any* multiline_comment? "*/" @{
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

    string_iterate :=
        (
            (any - ('"' | '\\') ) |
            ('\\' (
                ('\\' @{this.flushAndAddEscapedCharacter(fpc-1, '\\', callbacks)}) |
                ('n' @{this.flushAndAddEscapedCharacter(fpc-1, '\n', callbacks)}) |
                ('r' @{this.flushAndAddEscapedCharacter(fpc-1, '\r', callbacks)}) |
                ('t' @{this.flushAndAddEscapedCharacter(fpc-1, '\t', callbacks)}) |
                ('"' @{this.flushAndAddEscapedCharacter(fpc-1, '"', callbacks)}) |
                ([^"nrt\\] @{return false, fmt.Errorf("\\%c: Illegal escape encoding", this.data[fpc])})
                # TODO: hex and unicode
            ))
        )*
        '"' @{
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

    list_iterate :=
        ( value (ws value)* )?
        ws* ']' @{
            err = callbacks.OnContainerEnd()
            if err != nil {
                fbreak;
            }
            fret;
        };

    unordered_map_iterate :=
        ( kv_pair (ws kv_pair)* )?
        ws* '}' @{
            err = callbacks.OnContainerEnd()
            if err != nil {
                fbreak;
            }
            fret;
        };

    ordered_map_iterate :=
        ( kv_pair (ws kv_pair)* )?
        ws* '>' @{
            err = callbacks.OnContainerEnd()
            if err != nil {
                fbreak;
            }
            fret;
        };

    metadata_map_iterate :=
        ( kv_pair (ws kv_pair)* )?
        ws* ')' @{
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


%% write data nofinal;

type Parser struct {
    cs int // Current Ragel state
    ts int // Position: start of token
    te int // Position: end of token
    top int // Index of top of stack
    stack []int
    data []byte
    arrayStart int // Start of the current item of interest
    commentDepth int
    significand uint64
    significandSign int
    exponent int
    exponentSign int
    exponentAdjust int
}

func (this *Parser) Init(maxDepth int) {
    this.stack = make([]int, maxDepth)
    this.significandSign = 1
    this.exponentSign = 1
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
