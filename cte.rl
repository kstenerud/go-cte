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
//  OnBytesBegin(byteCount uint64) error
//  OnStringBegin(byteCount uint64) error
//  OnURIBegin(byteCount uint64) error
//  OnCommentBegin(byteCount uint64) error
//  OnArrayData(bytes []byte) error
}

%%{
    machine cte;
    access this.;

    action on_error
    {
    		err = fmt.Errorf("Parse Error")
        fbreak;
    }

	ws = [\r\n\t ];

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

    comment = "//" [^\n]* '\n' @{
    };

    metadata = metadata_map | comment;
    keyable = true | false;
    nonkeyable = nil | list | unordered_map | ordered_map;
    value = ws* (keyable | nonkeyable);
    kv_pair = ws* keyable ws* '=' value;


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
        // copy(this.underflow, data[ts:pe])
        // p = 0
        // pe = pe - ts
	}
	// TODO
	isComplete = true
	return
}
