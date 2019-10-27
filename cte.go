
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
//  OnBytesBegin(byteCount uint64) error
//  OnStringBegin(byteCount uint64) error
//  OnURIBegin(byteCount uint64) error
//  OnCommentBegin(byteCount uint64) error
//  OnArrayData(bytes []byte) error
}


//line cte.rl:145




//line cte.go:38
var _cte_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 8, 1, 9, 1, 10, 
}

var _cte_key_offsets []byte = []byte{
	0, 0, 10, 11, 12, 13, 14, 15, 
	16, 17, 18, 19, 30, 35, 36, 37, 
	38, 39, 40, 41, 42, 43, 44, 51, 
	52, 53, 54, 55, 60, 70, 75, 76, 
	77, 78, 79, 80, 81, 82, 83, 84, 
	85, 86, 87, 94, 95, 96, 97, 98, 
	103, 113, 118, 119, 120, 121, 122, 123, 
	124, 125, 126, 127, 128, 129, 130, 137, 
	138, 139, 140, 141, 146, 156, 161, 162, 
	163, 164, 165, 166, 167, 168, 169, 170, 
	171, 172, 173, 177, 177, 177, 177, 
}

var _cte_trans_keys []byte = []byte{
	13, 32, 60, 91, 102, 110, 116, 123, 
	9, 10, 97, 108, 115, 101, 105, 108, 
	114, 117, 101, 13, 32, 60, 91, 93, 
	102, 110, 116, 123, 9, 10, 13, 32, 
	93, 9, 10, 97, 108, 115, 101, 105, 
	108, 114, 117, 101, 13, 32, 102, 116, 
	125, 9, 10, 97, 108, 115, 101, 13, 
	32, 61, 9, 10, 13, 32, 60, 91, 
	102, 110, 116, 123, 9, 10, 13, 32, 
	125, 9, 10, 97, 108, 115, 101, 105, 
	108, 114, 117, 101, 114, 117, 101, 13, 
	32, 62, 102, 116, 9, 10, 97, 108, 
	115, 101, 13, 32, 61, 9, 10, 13, 
	32, 60, 91, 102, 110, 116, 123, 9, 
	10, 13, 32, 62, 9, 10, 97, 108, 
	115, 101, 105, 108, 114, 117, 101, 114, 
	117, 101, 13, 32, 41, 102, 116, 9, 
	10, 97, 108, 115, 101, 13, 32, 61, 
	9, 10, 13, 32, 60, 91, 102, 110, 
	116, 123, 9, 10, 13, 32, 41, 9, 
	10, 97, 108, 115, 101, 105, 108, 114, 
	117, 101, 114, 117, 101, 13, 32, 9, 
	10, 
}

var _cte_single_lengths []byte = []byte{
	0, 8, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 9, 3, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 5, 1, 
	1, 1, 1, 3, 8, 3, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 5, 1, 1, 1, 1, 3, 
	8, 3, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 5, 1, 
	1, 1, 1, 3, 8, 3, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 2, 0, 0, 0, 0, 
}

var _cte_range_lengths []byte = []byte{
	0, 1, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 1, 1, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 1, 0, 
	0, 0, 0, 1, 1, 1, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 1, 0, 0, 0, 0, 1, 
	1, 1, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 1, 0, 
	0, 0, 0, 1, 1, 1, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 1, 0, 0, 0, 0, 
}

var _cte_index_offsets []int16 = []int16{
	0, 0, 10, 12, 14, 16, 18, 20, 
	22, 24, 26, 28, 39, 44, 46, 48, 
	50, 52, 54, 56, 58, 60, 62, 69, 
	71, 73, 75, 77, 82, 92, 97, 99, 
	101, 103, 105, 107, 109, 111, 113, 115, 
	117, 119, 121, 128, 130, 132, 134, 136, 
	141, 151, 156, 158, 160, 162, 164, 166, 
	168, 170, 172, 174, 176, 178, 180, 187, 
	189, 191, 193, 195, 200, 210, 215, 217, 
	219, 221, 223, 225, 227, 229, 231, 233, 
	235, 237, 239, 243, 244, 245, 246, 
}

var _cte_indicies []byte = []byte{
	0, 0, 2, 3, 4, 5, 6, 7, 
	0, 1, 8, 1, 9, 1, 10, 1, 
	11, 1, 12, 1, 13, 1, 14, 1, 
	15, 1, 16, 1, 18, 18, 19, 20, 
	21, 22, 23, 24, 25, 18, 17, 18, 
	18, 21, 18, 17, 26, 1, 27, 1, 
	28, 1, 29, 1, 30, 1, 31, 1, 
	32, 1, 33, 1, 34, 1, 35, 35, 
	36, 37, 38, 35, 17, 39, 1, 40, 
	1, 41, 1, 42, 1, 43, 43, 44, 
	43, 1, 44, 44, 45, 46, 47, 48, 
	49, 50, 44, 1, 35, 35, 38, 35, 
	17, 51, 1, 52, 1, 53, 1, 54, 
	1, 55, 1, 56, 1, 57, 1, 58, 
	1, 59, 1, 60, 1, 61, 1, 62, 
	1, 63, 63, 64, 65, 66, 63, 17, 
	67, 1, 68, 1, 69, 1, 70, 1, 
	71, 71, 72, 71, 1, 72, 72, 73, 
	74, 75, 76, 77, 78, 72, 1, 63, 
	63, 64, 63, 17, 79, 1, 80, 1, 
	81, 1, 82, 1, 83, 1, 84, 1, 
	85, 1, 86, 1, 87, 1, 88, 1, 
	89, 1, 90, 1, 91, 91, 92, 93, 
	94, 91, 17, 95, 1, 96, 1, 97, 
	1, 98, 1, 99, 99, 100, 99, 1, 
	100, 100, 101, 102, 103, 104, 105, 106, 
	100, 1, 91, 91, 92, 91, 17, 107, 
	1, 108, 1, 109, 1, 110, 1, 111, 
	1, 112, 1, 113, 1, 114, 1, 115, 
	1, 116, 1, 117, 1, 118, 1, 119, 
	119, 119, 17, 17, 17, 17, 17, 
}

var _cte_trans_targs []byte = []byte{
	1, 0, 82, 82, 2, 6, 8, 82, 
	3, 4, 5, 82, 7, 82, 9, 10, 
	82, 0, 11, 12, 12, 83, 13, 17, 
	19, 12, 14, 15, 16, 12, 18, 12, 
	20, 21, 12, 22, 23, 39, 84, 24, 
	25, 26, 27, 27, 28, 29, 29, 30, 
	34, 36, 29, 31, 32, 33, 29, 35, 
	29, 37, 38, 29, 40, 41, 27, 42, 
	85, 43, 59, 44, 45, 46, 47, 47, 
	48, 49, 49, 50, 54, 56, 49, 51, 
	52, 53, 49, 55, 49, 57, 58, 49, 
	60, 61, 47, 62, 86, 63, 79, 64, 
	65, 66, 67, 67, 68, 69, 69, 70, 
	74, 76, 69, 71, 72, 73, 69, 75, 
	69, 77, 78, 69, 80, 81, 67, 82, 
}

var _cte_trans_actions []byte = []byte{
	0, 0, 13, 9, 0, 0, 0, 11, 
	0, 0, 0, 7, 0, 3, 0, 0, 
	5, 1, 0, 13, 9, 15, 0, 0, 
	0, 11, 0, 0, 0, 7, 0, 3, 
	0, 0, 5, 0, 0, 0, 17, 0, 
	0, 0, 7, 0, 0, 13, 9, 0, 
	0, 0, 11, 0, 0, 0, 7, 0, 
	3, 0, 0, 5, 0, 0, 5, 0, 
	19, 0, 0, 0, 0, 0, 7, 0, 
	0, 13, 9, 0, 0, 0, 11, 0, 
	0, 0, 7, 0, 3, 0, 0, 5, 
	0, 0, 5, 0, 21, 0, 0, 0, 
	0, 0, 7, 0, 0, 13, 9, 0, 
	0, 0, 11, 0, 0, 0, 7, 0, 
	3, 0, 0, 5, 0, 0, 5, 0, 
}

var _cte_eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 1, 1, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 1, 0, 
	0, 0, 0, 0, 0, 1, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 1, 0, 0, 0, 0, 0, 
	0, 1, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 1, 0, 
	0, 0, 0, 0, 0, 1, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 1, 1, 1, 1, 
}

const cte_start int = 1
const cte_error int = 0

const cte_en_list_iterate int = 11
const cte_en_unordered_map_iterate int = 22
const cte_en_ordered_map_iterate int = 42
const cte_en_metadata_map_iterate int = 62
const cte_en_main int = 1


//line cte.rl:149

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
	
    
//line cte.go:261
	{
	 this.cs = cte_start
	 this.top = 0
	}

//line cte.go:267
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
	_trans = int(_cte_indicies[_trans])
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
//line cte.rl:34

    		err = fmt.Errorf("Parse Error")
        p++; goto _out

    
		case 1:
//line cte.rl:41

        err = callbacks.OnNil()
        if err != nil {
        		p++; goto _out

        }
    
		case 2:
//line cte.rl:48

        err = callbacks.OnBool(true)
        if err != nil {
        		p++; goto _out

        }
    
		case 3:
//line cte.rl:55

        err = callbacks.OnBool(false)
        if err != nil {
        		p++; goto _out

        }
    
		case 4:
//line cte.rl:62

        err = callbacks.OnListBegin()
        if err != nil {
        		p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 11; goto _again

    
		case 5:
//line cte.rl:70

        err = callbacks.OnUnorderedMapBegin()
        if err != nil {
        		p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 22; goto _again

    
		case 6:
//line cte.rl:78

        err = callbacks.OnOrderedMapBegin()
        if err != nil {
        		p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 42; goto _again

    
		case 7:
//line cte.rl:106

            err = callbacks.OnContainerEnd()
            if err != nil {
	        		p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 8:
//line cte.rl:116

            err = callbacks.OnContainerEnd()
            if err != nil {
	        		p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 9:
//line cte.rl:126

            err = callbacks.OnContainerEnd()
            if err != nil {
	        		p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 10:
//line cte.rl:136

            err = callbacks.OnContainerEnd()
            if err != nil {
	        		p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
//line cte.go:461
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
//line cte.rl:34

    		err = fmt.Errorf("Parse Error")
        p++; goto _out

    
//line cte.go:487
			}
		}
	}

	_out: {}
	}

//line cte.rl:190


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
