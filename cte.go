
//line cte.rl:1
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


//line cte.rl:212




//line cte.go:39
var _cte_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 8, 1, 9, 1, 10, 1, 11, 
	1, 12, 1, 13, 1, 14, 1, 15, 
	1, 16, 1, 17, 1, 18, 1, 19, 
	1, 20, 
}

var _cte_key_offsets []byte = []byte{
	0, 0, 12, 13, 13, 14, 15, 16, 
	17, 18, 19, 20, 21, 22, 23, 25, 
	30, 43, 48, 49, 49, 50, 51, 52, 
	53, 54, 55, 56, 57, 58, 66, 71, 
	83, 88, 89, 89, 90, 91, 92, 93, 
	94, 95, 96, 97, 98, 99, 100, 101, 
	102, 103, 104, 105, 113, 118, 130, 135, 
	136, 136, 137, 138, 139, 140, 141, 142, 
	143, 144, 145, 146, 147, 148, 149, 150, 
	151, 152, 160, 165, 177, 182, 183, 183, 
	184, 185, 186, 187, 188, 189, 190, 191, 
	192, 193, 194, 195, 196, 197, 198, 199, 
	203, 203, 203, 203, 203, 203, 
}

var _cte_trans_keys []byte = []byte{
	13, 32, 34, 47, 60, 91, 102, 110, 
	116, 123, 9, 10, 47, 97, 108, 115, 
	101, 105, 108, 114, 117, 101, 10, 34, 
	92, 34, 92, 110, 114, 116, 13, 32, 
	34, 47, 60, 91, 93, 102, 110, 116, 
	123, 9, 10, 13, 32, 93, 9, 10, 
	47, 97, 108, 115, 101, 105, 108, 114, 
	117, 101, 13, 32, 34, 102, 116, 125, 
	9, 10, 13, 32, 61, 9, 10, 13, 
	32, 34, 47, 60, 91, 102, 110, 116, 
	123, 9, 10, 13, 32, 125, 9, 10, 
	47, 97, 108, 115, 101, 105, 108, 114, 
	117, 101, 97, 108, 115, 101, 114, 117, 
	101, 13, 32, 34, 62, 102, 116, 9, 
	10, 13, 32, 61, 9, 10, 13, 32, 
	34, 47, 60, 91, 102, 110, 116, 123, 
	9, 10, 13, 32, 62, 9, 10, 47, 
	97, 108, 115, 101, 105, 108, 114, 117, 
	101, 97, 108, 115, 101, 114, 117, 101, 
	13, 32, 34, 41, 102, 116, 9, 10, 
	13, 32, 61, 9, 10, 13, 32, 34, 
	47, 60, 91, 102, 110, 116, 123, 9, 
	10, 13, 32, 41, 9, 10, 47, 97, 
	108, 115, 101, 105, 108, 114, 117, 101, 
	97, 108, 115, 101, 114, 117, 101, 13, 
	32, 9, 10, 
}

var _cte_single_lengths []byte = []byte{
	0, 10, 1, 0, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 2, 5, 
	11, 3, 1, 0, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 6, 3, 10, 
	3, 1, 0, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 6, 3, 10, 3, 1, 
	0, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 6, 3, 10, 3, 1, 0, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 2, 
	0, 0, 0, 0, 0, 0, 
}

var _cte_range_lengths []byte = []byte{
	0, 1, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	1, 1, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 1, 1, 1, 
	1, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 1, 1, 1, 1, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 1, 1, 1, 1, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 1, 
	0, 0, 0, 0, 0, 0, 
}

var _cte_index_offsets []int16 = []int16{
	0, 0, 12, 14, 15, 17, 19, 21, 
	23, 25, 27, 29, 31, 33, 35, 38, 
	44, 57, 62, 64, 65, 67, 69, 71, 
	73, 75, 77, 79, 81, 83, 91, 96, 
	108, 113, 115, 116, 118, 120, 122, 124, 
	126, 128, 130, 132, 134, 136, 138, 140, 
	142, 144, 146, 148, 156, 161, 173, 178, 
	180, 181, 183, 185, 187, 189, 191, 193, 
	195, 197, 199, 201, 203, 205, 207, 209, 
	211, 213, 221, 226, 238, 243, 245, 246, 
	248, 250, 252, 254, 256, 258, 260, 262, 
	264, 266, 268, 270, 272, 274, 276, 278, 
	282, 283, 284, 285, 286, 287, 
}

var _cte_trans_targs []byte = []byte{
	1, 1, 95, 2, 95, 95, 4, 8, 
	10, 95, 1, 0, 3, 0, 95, 5, 
	0, 6, 0, 7, 0, 95, 0, 9, 
	0, 95, 0, 11, 0, 12, 0, 95, 
	0, 96, 13, 97, 15, 14, 14, 14, 
	14, 14, 14, 14, 16, 16, 17, 18, 
	17, 17, 98, 20, 24, 26, 17, 16, 
	0, 16, 16, 98, 16, 0, 19, 0, 
	17, 21, 0, 22, 0, 23, 0, 17, 
	0, 25, 0, 17, 0, 27, 0, 28, 
	0, 17, 0, 29, 29, 30, 44, 48, 
	99, 29, 0, 30, 30, 31, 30, 0, 
	31, 31, 32, 33, 32, 32, 35, 39, 
	41, 32, 31, 0, 29, 29, 99, 29, 
	0, 34, 0, 32, 36, 0, 37, 0, 
	38, 0, 32, 0, 40, 0, 32, 0, 
	42, 0, 43, 0, 32, 0, 45, 0, 
	46, 0, 47, 0, 30, 0, 49, 0, 
	50, 0, 30, 0, 51, 51, 52, 100, 
	66, 70, 51, 0, 52, 52, 53, 52, 
	0, 53, 53, 54, 55, 54, 54, 57, 
	61, 63, 54, 53, 0, 51, 51, 100, 
	51, 0, 56, 0, 54, 58, 0, 59, 
	0, 60, 0, 54, 0, 62, 0, 54, 
	0, 64, 0, 65, 0, 54, 0, 67, 
	0, 68, 0, 69, 0, 52, 0, 71, 
	0, 72, 0, 52, 0, 73, 73, 74, 
	101, 88, 92, 73, 0, 74, 74, 75, 
	74, 0, 75, 75, 76, 77, 76, 76, 
	79, 83, 85, 76, 75, 0, 73, 73, 
	101, 73, 0, 78, 0, 76, 80, 0, 
	81, 0, 82, 0, 76, 0, 84, 0, 
	76, 0, 86, 0, 87, 0, 76, 0, 
	89, 0, 90, 0, 91, 0, 74, 0, 
	93, 0, 94, 0, 74, 0, 95, 95, 
	95, 0, 0, 0, 0, 0, 0, 0, 
	
}

var _cte_trans_actions []byte = []byte{
	0, 0, 17, 0, 13, 9, 0, 0, 
	0, 11, 0, 0, 0, 0, 15, 0, 
	0, 0, 0, 0, 0, 7, 0, 0, 
	0, 3, 0, 0, 0, 0, 0, 5, 
	0, 19, 0, 33, 0, 0, 29, 21, 
	23, 25, 27, 31, 0, 0, 17, 0, 
	13, 9, 35, 0, 0, 0, 11, 0, 
	1, 0, 0, 35, 0, 1, 0, 0, 
	15, 0, 0, 0, 0, 0, 0, 7, 
	0, 0, 0, 3, 0, 0, 0, 0, 
	0, 5, 0, 0, 0, 17, 0, 0, 
	37, 0, 1, 0, 0, 0, 0, 0, 
	0, 0, 17, 0, 13, 9, 0, 0, 
	0, 11, 0, 0, 0, 0, 37, 0, 
	1, 0, 0, 15, 0, 0, 0, 0, 
	0, 0, 7, 0, 0, 0, 3, 0, 
	0, 0, 0, 0, 5, 0, 0, 0, 
	0, 0, 0, 0, 7, 0, 0, 0, 
	0, 0, 5, 0, 0, 0, 17, 39, 
	0, 0, 0, 1, 0, 0, 0, 0, 
	0, 0, 0, 17, 0, 13, 9, 0, 
	0, 0, 11, 0, 0, 0, 0, 39, 
	0, 1, 0, 0, 15, 0, 0, 0, 
	0, 0, 0, 7, 0, 0, 0, 3, 
	0, 0, 0, 0, 0, 5, 0, 0, 
	0, 0, 0, 0, 0, 7, 0, 0, 
	0, 0, 0, 5, 0, 0, 0, 17, 
	41, 0, 0, 0, 1, 0, 0, 0, 
	0, 0, 0, 0, 17, 0, 13, 9, 
	0, 0, 0, 11, 0, 0, 0, 0, 
	41, 0, 1, 0, 0, 15, 0, 0, 
	0, 0, 0, 0, 7, 0, 0, 0, 
	3, 0, 0, 0, 0, 0, 5, 0, 
	0, 0, 0, 0, 0, 0, 7, 0, 
	0, 0, 0, 0, 5, 0, 0, 0, 
	0, 1, 1, 1, 1, 1, 1, 1, 
	
}

var _cte_eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 1, 1, 0, 
	1, 1, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 1, 0, 0, 
	1, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 1, 0, 0, 1, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 1, 0, 0, 1, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	1, 1, 1, 1, 1, 1, 
}

const cte_start int = 1
const cte_error int = 0

const cte_en_comment_iterate int = 13
const cte_en_string_iterate int = 14
const cte_en_list_iterate int = 16
const cte_en_unordered_map_iterate int = 29
const cte_en_ordered_map_iterate int = 51
const cte_en_metadata_map_iterate int = 73
const cte_en_main int = 1


//line cte.rl:216

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
	
    
//line cte.go:310
	{
	 this.cs = cte_start
	 this.top = 0
	}

//line cte.go:316
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if  this.cs == 0 {
		goto _out
	}
_resume:
	_keys = int(_cte_key_offsets[ this.cs])
	_trans = int(_cte_index_offsets[ this.cs])

	_klen = int(_cte_single_lengths[ this.cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case  this.data[p] < _cte_trans_keys[_mid]:
				_upper = _mid - 1
			case  this.data[p] > _cte_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_cte_range_lengths[ this.cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case  this.data[p] < _cte_trans_keys[_mid]:
				_upper = _mid - 2
			case  this.data[p] > _cte_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	 this.cs = int(_cte_trans_targs[_trans])

	if _cte_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_cte_trans_actions[_trans])
	_nacts = uint(_cte_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _cte_actions[_acts-1] {
		case 0:
//line cte.rl:35

    		err = fmt.Errorf("Parse error pos %v of [%v] (%c)", p, string(this.data),  this.data[p])
        p++; goto _out

    
		case 1:
//line cte.rl:51

        err = callbacks.OnNil()
        if err != nil {
        		p++; goto _out

        }
    
		case 2:
//line cte.rl:58

        err = callbacks.OnBool(true)
        if err != nil {
        		p++; goto _out

        }
    
		case 3:
//line cte.rl:65

        err = callbacks.OnBool(false)
        if err != nil {
        		p++; goto _out

        }
    
		case 4:
//line cte.rl:72

        err = callbacks.OnListBegin()
        if err != nil {
        		p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 16; goto _again

    
		case 5:
//line cte.rl:80

        err = callbacks.OnUnorderedMapBegin()
        if err != nil {
        		p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 29; goto _again

    
		case 6:
//line cte.rl:88

        err = callbacks.OnOrderedMapBegin()
        if err != nil {
        		p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 51; goto _again

    
		case 7:
//line cte.rl:104

		this.arrayStart = p
    		if  this.data[p] == ' ' {
				this.arrayStart+=4
		}
        err = callbacks.OnCommentBegin()
        if err != nil {
        		p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 13; goto _again

    
		case 8:
//line cte.rl:116

		this.arrayStart = p + 1
        err = callbacks.OnStringBegin()
        if err != nil {
        		p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 14; goto _again

    
		case 9:
//line cte.rl:134

            err = callbacks.OnArrayData(this.data[this.arrayStart:p])
            if err != nil {
                    p++; goto _out

            }
            err = callbacks.OnArrayEnd()
            if err != nil {
                    p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

		
		case 10:
//line cte.rl:150
this.flushAndAddEscapedCharacter(p-1, '\\', callbacks)
		case 11:
//line cte.rl:151
this.flushAndAddEscapedCharacter(p-1, '\n', callbacks)
		case 12:
//line cte.rl:152
this.flushAndAddEscapedCharacter(p-1, '\r', callbacks)
		case 13:
//line cte.rl:153
this.flushAndAddEscapedCharacter(p-1, '\t', callbacks)
		case 14:
//line cte.rl:154
this.flushAndAddEscapedCharacter(p-1, '"', callbacks)
		case 15:
//line cte.rl:155
return false, fmt.Errorf("\\%c: Illegal escape encoding", this.data[p])
		case 16:
//line cte.rl:159

            err = callbacks.OnArrayData(this.data[this.arrayStart:p])
            if err != nil {
                    p++; goto _out

            }
            err = callbacks.OnArrayEnd()
            if err != nil {
                    p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

		
		case 17:
//line cte.rl:173

            err = callbacks.OnContainerEnd()
            if err != nil {
	        		p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 18:
//line cte.rl:183

            err = callbacks.OnContainerEnd()
            if err != nil {
	        		p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 19:
//line cte.rl:193

            err = callbacks.OnContainerEnd()
            if err != nil {
	        		p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 20:
//line cte.rl:203

            err = callbacks.OnContainerEnd()
            if err != nil {
	        		p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
//line cte.go:588
		}
	}

_again:
	if  this.cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		__acts := _cte_eof_actions[ this.cs]
		__nacts := uint(_cte_actions[__acts]); __acts++
		for ; __nacts > 0; __nacts-- {
			__acts++
			switch _cte_actions[__acts-1] {
			case 0:
//line cte.rl:35

    		err = fmt.Errorf("Parse error pos %v of [%v] (%c)", p, string(this.data),  this.data[p])
        p++; goto _out

    
//line cte.go:614
			}
		}
	}

	_out: {}
	}

//line cte.rl:277


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
