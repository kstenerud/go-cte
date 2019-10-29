package cte

import (
    "fmt"
//  "time"
)

type CteDecoderCallbacks interface {
    OnNil() error
  OnBool(value bool) error
//  OnPositiveInt(value uint64) error
//  OnNegativeInt(value uint64) error
//  OnFloat(value float64) error
//  OnDate(value time.Time) error
//  OnTime(value time.Time) error
//  OnTimestamp(value time.Time) error
    OnListBegin() error
  OnOrderedMapBegin() error
  OnUnorderedMapBegin() error
  OnMetadataMapBegin() error
  OnContainerEnd() error
//  OnBytesBegin() error
  OnStringBegin() error
//  OnURIBegin() error
  OnCommentBegin() error
  OnArrayData(bytes []byte) error
  OnArrayEnd() error
}

%%{
    machine cte;
    access this.;

    action on_error
    {
    		err = fmt.Errorf("Parse error pos %v of [%v] (%c)", fpc, string(this.data), fc)
        fbreak;
    }

	ws = [\r\n\t ];
	eol = [\r\n];

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

    comment = "//" any @{
		this.arrayStart = fpc
    		if fc == ' ' {
				this.arrayStart++
		}
        err = callbacks.OnCommentBegin()
        if err != nil {
        		fbreak;
        }
        fcall comment_iterate;
    };

	string = '"' @{
		this.arrayStart = fpc + 1
        err = callbacks.OnStringBegin()
        if err != nil {
        		fbreak;
        }
        fcall string_iterate;
    };


    metadata = metadata_map | comment;
    keyable = true | false | string;
    nonkeyable = nil | list | unordered_map | ordered_map;
    value = ws* (keyable | nonkeyable);
    kv_pair = ws* keyable ws* '=' value;

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
		} $!on_error $/on_error;

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
		} $!on_error $/on_error;

    list_iterate :=
        ( value (ws value)* )?
        ws* ']' @{
            err = callbacks.OnContainerEnd()
            if err != nil {
	        		fbreak;
            }
            fret;
        } $!on_error $/on_error;

    unordered_map_iterate :=
        ( kv_pair (ws kv_pair)* )?
        ws* '}' @{
            err = callbacks.OnContainerEnd()
            if err != nil {
	        		fbreak;
            }
            fret;
        } $!on_error $/on_error;

    ordered_map_iterate :=
        ( kv_pair (ws kv_pair)* )?
        ws* '>' @{
            err = callbacks.OnContainerEnd()
            if err != nil {
	        		fbreak;
            }
            fret;
        } $!on_error $/on_error;

    metadata_map_iterate :=
        ( kv_pair (ws kv_pair)* )?
        ws* ')' @{
            err = callbacks.OnContainerEnd()
            if err != nil {
	        		fbreak;
            }
            fret;
        } $!on_error $/on_error;

    main := value ws* $!on_error @/on_error;
}%%


%% write data nofinal;

type Parser struct {
    cs int // Current Ragel state
    p int // Position: current
    pe int // Position: end of buffer
    eof int // Position: end of file
    ts int // Position: start of token
    te int // Position: end of token
    top int // Index of top of stack
    stack []int
    data []byte
    arrayStart int // Start of the current item of interest
}

func (this *Parser) Init(maxDepth int) {
	this.eof = -1
	this.stack = make([]int, maxDepth)
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
//	if this.ts > 0 {
		// TODO: Read from undeflow buffer
//	}
	this.data = src
    p := 0 // Position: current
    pe := len(this.data) // Position: end of buffer
    eof := pe // Position: end of file
    ts := 0 // Position: start of token
    te := 0 // Position: end of token

	_ = eof
	_ = te
	
    %%{
        write init;
        write exec;
    }%%

	if ts > 0 {
		// TODO: Copy to underflow buffer
		// arrayStart
		// ts doesn't seem to get used?
        // copy(this.underflow, data[ts:pe])
        // p = 0
        // pe = pe - ts
	}
	// TODO
	isComplete = true
	return
}
