
//line cte.rl:1
package cte

import (
    "fmt"
//    "time"
)

type CteDecoderCallbacks interface {
    OnNil() error
    OnBool(value bool) error
//    OnPositiveInt(value uint64) error
//    OnNegativeInt(value uint64) error
//    OnFloat(value float64) error
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
//    OnURIBegin() error
    OnCommentBegin() error
    OnArrayData(bytes []byte) error
    OnArrayEnd() error
}


//line cte.rl:244




//line cte.go:39
var _cte_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 8, 1, 9, 1, 10, 1, 11, 
	1, 12, 1, 13, 1, 14, 1, 15, 
	1, 16, 1, 17, 1, 18, 1, 19, 
	1, 20, 1, 21, 1, 22, 
}

var _cte_key_offsets []int16 = []int16{
	0, 0, 13, 15, 16, 17, 18, 19, 
	20, 21, 22, 23, 24, 25, 27, 29, 
	31, 33, 38, 52, 57, 70, 72, 73, 
	74, 75, 76, 77, 78, 79, 80, 81, 
	91, 97, 99, 112, 117, 119, 120, 121, 
	122, 123, 124, 125, 126, 127, 128, 137, 
	139, 140, 141, 142, 143, 144, 145, 146, 
	156, 162, 164, 177, 182, 184, 185, 186, 
	187, 188, 189, 190, 191, 192, 193, 202, 
	204, 205, 206, 207, 208, 209, 210, 211, 
	221, 227, 229, 242, 247, 249, 250, 251, 
	252, 253, 254, 255, 256, 257, 258, 267, 
	269, 270, 271, 272, 273, 274, 275, 276, 
	280, 280, 282, 282, 282, 282, 282, 
}

var _cte_trans_keys []byte = []byte{
	13, 32, 34, 40, 47, 60, 91, 102, 
	110, 116, 123, 9, 10, 42, 47, 97, 
	108, 115, 101, 105, 108, 114, 117, 101, 
	10, 42, 47, 42, 47, 42, 47, 34, 
	92, 34, 92, 110, 114, 116, 13, 32, 
	34, 40, 47, 60, 91, 93, 102, 110, 
	116, 123, 9, 10, 13, 32, 93, 9, 
	10, 13, 32, 34, 40, 47, 60, 91, 
	102, 110, 116, 123, 9, 10, 42, 47, 
	97, 108, 115, 101, 105, 108, 114, 117, 
	101, 13, 32, 34, 40, 47, 102, 116, 
	125, 9, 10, 13, 32, 47, 61, 9, 
	10, 42, 47, 13, 32, 34, 40, 47, 
	60, 91, 102, 110, 116, 123, 9, 10, 
	13, 32, 125, 9, 10, 42, 47, 97, 
	108, 115, 101, 105, 108, 114, 117, 101, 
	13, 32, 34, 40, 47, 102, 116, 9, 
	10, 42, 47, 97, 108, 115, 101, 114, 
	117, 101, 13, 32, 34, 40, 47, 62, 
	102, 116, 9, 10, 13, 32, 47, 61, 
	9, 10, 42, 47, 13, 32, 34, 40, 
	47, 60, 91, 102, 110, 116, 123, 9, 
	10, 13, 32, 62, 9, 10, 42, 47, 
	97, 108, 115, 101, 105, 108, 114, 117, 
	101, 13, 32, 34, 40, 47, 102, 116, 
	9, 10, 42, 47, 97, 108, 115, 101, 
	114, 117, 101, 13, 32, 34, 40, 41, 
	47, 102, 116, 9, 10, 13, 32, 47, 
	61, 9, 10, 42, 47, 13, 32, 34, 
	40, 47, 60, 91, 102, 110, 116, 123, 
	9, 10, 13, 32, 41, 9, 10, 42, 
	47, 97, 108, 115, 101, 105, 108, 114, 
	117, 101, 13, 32, 34, 40, 47, 102, 
	116, 9, 10, 42, 47, 97, 108, 115, 
	101, 114, 117, 101, 13, 32, 9, 10, 
	42, 47, 
}

var _cte_single_lengths []byte = []byte{
	0, 11, 2, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 2, 2, 2, 
	2, 5, 12, 3, 11, 2, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 8, 
	4, 2, 11, 3, 2, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 7, 2, 
	1, 1, 1, 1, 1, 1, 1, 8, 
	4, 2, 11, 3, 2, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 7, 2, 
	1, 1, 1, 1, 1, 1, 1, 8, 
	4, 2, 11, 3, 2, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 7, 2, 
	1, 1, 1, 1, 1, 1, 1, 2, 
	0, 2, 0, 0, 0, 0, 0, 
}

var _cte_range_lengths []byte = []byte{
	0, 1, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 1, 1, 1, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 1, 
	1, 0, 1, 1, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 1, 0, 
	0, 0, 0, 0, 0, 0, 0, 1, 
	1, 0, 1, 1, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 1, 0, 
	0, 0, 0, 0, 0, 0, 0, 1, 
	1, 0, 1, 1, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 1, 0, 
	0, 0, 0, 0, 0, 0, 0, 1, 
	0, 0, 0, 0, 0, 0, 0, 
}

var _cte_index_offsets []int16 = []int16{
	0, 0, 13, 16, 18, 20, 22, 24, 
	26, 28, 30, 32, 34, 36, 39, 42, 
	45, 48, 54, 68, 73, 86, 89, 91, 
	93, 95, 97, 99, 101, 103, 105, 107, 
	117, 123, 126, 139, 144, 147, 149, 151, 
	153, 155, 157, 159, 161, 163, 165, 174, 
	177, 179, 181, 183, 185, 187, 189, 191, 
	201, 207, 210, 223, 228, 231, 233, 235, 
	237, 239, 241, 243, 245, 247, 249, 258, 
	261, 263, 265, 267, 269, 271, 273, 275, 
	285, 291, 294, 307, 312, 315, 317, 319, 
	321, 323, 325, 327, 329, 331, 333, 342, 
	345, 347, 349, 351, 353, 355, 357, 359, 
	363, 364, 367, 368, 369, 370, 371, 
}

var _cte_trans_targs []byte = []byte{
	1, 1, 103, 1, 2, 103, 103, 3, 
	7, 9, 103, 1, 0, 1, 1, 0, 
	4, 0, 5, 0, 6, 0, 103, 0, 
	8, 0, 103, 0, 10, 0, 11, 0, 
	103, 0, 104, 12, 14, 15, 13, 14, 
	105, 13, 14, 15, 13, 106, 17, 16, 
	16, 16, 16, 16, 16, 16, 18, 18, 
	19, 20, 21, 19, 19, 107, 22, 26, 
	28, 19, 18, 0, 18, 18, 107, 18, 
	0, 20, 20, 19, 20, 21, 19, 19, 
	22, 26, 28, 19, 20, 0, 20, 20, 
	0, 23, 0, 24, 0, 25, 0, 19, 
	0, 27, 0, 19, 0, 29, 0, 30, 
	0, 19, 0, 31, 31, 32, 46, 47, 
	48, 52, 108, 31, 0, 32, 32, 33, 
	34, 32, 0, 32, 32, 0, 34, 34, 
	35, 34, 36, 35, 35, 37, 41, 43, 
	35, 34, 0, 31, 31, 108, 31, 0, 
	34, 34, 0, 38, 0, 39, 0, 40, 
	0, 35, 0, 42, 0, 35, 0, 44, 
	0, 45, 0, 35, 0, 46, 46, 32, 
	46, 47, 48, 52, 46, 0, 46, 46, 
	0, 49, 0, 50, 0, 51, 0, 32, 
	0, 53, 0, 54, 0, 32, 0, 55, 
	55, 56, 70, 71, 109, 72, 76, 55, 
	0, 56, 56, 57, 58, 56, 0, 56, 
	56, 0, 58, 58, 59, 58, 60, 59, 
	59, 61, 65, 67, 59, 58, 0, 55, 
	55, 109, 55, 0, 58, 58, 0, 62, 
	0, 63, 0, 64, 0, 59, 0, 66, 
	0, 59, 0, 68, 0, 69, 0, 59, 
	0, 70, 70, 56, 70, 71, 72, 76, 
	70, 0, 70, 70, 0, 73, 0, 74, 
	0, 75, 0, 56, 0, 77, 0, 78, 
	0, 56, 0, 79, 79, 80, 94, 110, 
	95, 96, 100, 79, 0, 80, 80, 81, 
	82, 80, 0, 80, 80, 0, 82, 82, 
	83, 82, 84, 83, 83, 85, 89, 91, 
	83, 82, 0, 79, 79, 110, 79, 0, 
	82, 82, 0, 86, 0, 87, 0, 88, 
	0, 83, 0, 90, 0, 83, 0, 92, 
	0, 93, 0, 83, 0, 94, 94, 80, 
	94, 95, 96, 100, 94, 0, 94, 94, 
	0, 97, 0, 98, 0, 99, 0, 80, 
	0, 101, 0, 102, 0, 80, 0, 103, 
	103, 103, 0, 0, 14, 15, 13, 0, 
	0, 0, 0, 0, 
}

var _cte_trans_actions []byte = []byte{
	0, 0, 19, 13, 0, 11, 7, 0, 
	0, 0, 9, 0, 0, 17, 15, 0, 
	0, 0, 0, 0, 0, 0, 5, 0, 
	0, 0, 1, 0, 0, 0, 0, 0, 
	3, 0, 21, 0, 0, 0, 0, 0, 
	23, 0, 17, 0, 0, 37, 0, 0, 
	33, 25, 27, 29, 31, 35, 0, 0, 
	19, 13, 0, 11, 7, 39, 0, 0, 
	0, 9, 0, 0, 0, 0, 39, 0, 
	0, 0, 0, 19, 13, 0, 11, 7, 
	0, 0, 0, 9, 0, 0, 17, 15, 
	0, 0, 0, 0, 0, 0, 0, 5, 
	0, 0, 0, 1, 0, 0, 0, 0, 
	0, 3, 0, 0, 0, 19, 13, 0, 
	0, 0, 41, 0, 0, 0, 0, 0, 
	0, 0, 0, 17, 15, 0, 0, 0, 
	19, 13, 0, 11, 7, 0, 0, 0, 
	9, 0, 0, 0, 0, 41, 0, 0, 
	17, 15, 0, 0, 0, 0, 0, 0, 
	0, 5, 0, 0, 0, 1, 0, 0, 
	0, 0, 0, 3, 0, 0, 0, 19, 
	13, 0, 0, 0, 0, 0, 17, 15, 
	0, 0, 0, 0, 0, 0, 0, 5, 
	0, 0, 0, 0, 0, 3, 0, 0, 
	0, 19, 13, 0, 43, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 17, 
	15, 0, 0, 0, 19, 13, 0, 11, 
	7, 0, 0, 0, 9, 0, 0, 0, 
	0, 43, 0, 0, 17, 15, 0, 0, 
	0, 0, 0, 0, 0, 5, 0, 0, 
	0, 1, 0, 0, 0, 0, 0, 3, 
	0, 0, 0, 19, 13, 0, 0, 0, 
	0, 0, 17, 15, 0, 0, 0, 0, 
	0, 0, 0, 5, 0, 0, 0, 0, 
	0, 3, 0, 0, 0, 19, 13, 45, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 17, 15, 0, 0, 0, 
	19, 13, 0, 11, 7, 0, 0, 0, 
	9, 0, 0, 0, 0, 45, 0, 0, 
	17, 15, 0, 0, 0, 0, 0, 0, 
	0, 5, 0, 0, 0, 1, 0, 0, 
	0, 0, 0, 3, 0, 0, 0, 19, 
	13, 0, 0, 0, 0, 0, 17, 15, 
	0, 0, 0, 0, 0, 0, 0, 5, 
	0, 0, 0, 0, 0, 3, 0, 0, 
	0, 0, 0, 0, 17, 0, 0, 0, 
	0, 0, 0, 0, 
}

const cte_start int = 1
const cte_error int = 0

const cte_en_comment_iterate int = 12
const cte_en_multiline_comment_iterate int = 13
const cte_en_string_iterate int = 16
const cte_en_list_iterate int = 18
const cte_en_unordered_map_iterate int = 31
const cte_en_ordered_map_iterate int = 55
const cte_en_metadata_map_iterate int = 79
const cte_en_main int = 1


//line cte.rl:248

type Parser struct {
    cs int // Current Ragel state
    ts int // Position: start of token
    te int // Position: end of token
    top int // Index of top of stack
    stack []int
    data []byte
    arrayStart int // Start of the current item of interest
    commentDepth int
    exponent uint64
    exponentSign int
    significand uint64
    significandSign int
}

func (this *Parser) Init(maxDepth int) {
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
//  if this.ts > 0 {
        // TODO: Read from undeflow buffer
//  }
    this.data = src
    p := 0 // Position: current
    pe := len(this.data) // Position: end of buffer
    // TODO: Change to -1 and check for end of file
    eof := pe // Position: end of file

    _ = eof
    
    
//line cte.go:328
	{
	 this.cs = cte_start
	 this.top = 0
	}

//line cte.go:334
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
//line cte.rl:45

        err = callbacks.OnNil()
        if err != nil {
            p++; goto _out

        }
    
		case 1:
//line cte.rl:52

        err = callbacks.OnBool(true)
        if err != nil {
            p++; goto _out

        }
    
		case 2:
//line cte.rl:59

        err = callbacks.OnBool(false)
        if err != nil {
            p++; goto _out

        }
    
		case 3:
//line cte.rl:72

        err = callbacks.OnListBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 18; goto _again

    
		case 4:
//line cte.rl:80

        err = callbacks.OnUnorderedMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 31; goto _again

    
		case 5:
//line cte.rl:88

        err = callbacks.OnOrderedMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 55; goto _again

    
		case 6:
//line cte.rl:96

        err = callbacks.OnMetadataMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 79; goto _again

    
		case 7:
//line cte.rl:104

        this.arrayStart = p + 1
        err = callbacks.OnCommentBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 12; goto _again

    
		case 8:
//line cte.rl:113

        if this.commentDepth == 0 {
            err = callbacks.OnCommentBegin()
            if err != nil {
                p++; goto _out

            }
        } else {
            err = callbacks.OnArrayData(this.data[this.arrayStart:p+1])
            if err != nil {
                p++; goto _out

            }
        }
        this.arrayStart = p + 1
        this.commentDepth++
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 13; goto _again

    
		case 9:
//line cte.rl:130

        this.arrayStart = p + 1
        err = callbacks.OnStringBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 16; goto _again

    
		case 10:
//line cte.rl:149

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

        
		case 11:
//line cte.rl:162

            err = callbacks.OnArrayData(this.data[this.arrayStart:p-1])
            if err != nil {
                p++; goto _out

            }
            this.arrayStart = p-1
            this.commentDepth--
            if this.commentDepth == 0 {
                err = callbacks.OnArrayEnd()
                if err != nil {
                    p++; goto _out

                }
            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 12:
//line cte.rl:182
this.flushAndAddEscapedCharacter(p-1, '\\', callbacks)
		case 13:
//line cte.rl:183
this.flushAndAddEscapedCharacter(p-1, '\n', callbacks)
		case 14:
//line cte.rl:184
this.flushAndAddEscapedCharacter(p-1, '\r', callbacks)
		case 15:
//line cte.rl:185
this.flushAndAddEscapedCharacter(p-1, '\t', callbacks)
		case 16:
//line cte.rl:186
this.flushAndAddEscapedCharacter(p-1, '"', callbacks)
		case 17:
//line cte.rl:187
return false, fmt.Errorf("\\%c: Illegal escape encoding", this.data[p])
		case 18:
//line cte.rl:191

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

        
		case 19:
//line cte.rl:205

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 20:
//line cte.rl:215

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 21:
//line cte.rl:225

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 22:
//line cte.rl:235

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
//line cte.go:649
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
	_out: {}
	}

//line cte.rl:308


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
