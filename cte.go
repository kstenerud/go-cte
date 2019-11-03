
//line cte.rl:1
package cte

import (
    "fmt"
//    "time"
)

type CteDecoderCallbacks interface {
    OnNil() error
    OnBool(value bool) error
    OnPositiveInt(value uint64) error
    OnNegativeInt(value uint64) error
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


//line cte.rl:261




//line cte.go:39
var _cte_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 8, 1, 9, 1, 10, 1, 11, 
	1, 12, 1, 13, 1, 14, 1, 15, 
	1, 16, 1, 17, 1, 18, 1, 19, 
	1, 20, 1, 21, 1, 22, 1, 23, 
	1, 24, 1, 25, 2, 5, 22, 2, 
	5, 23, 2, 5, 24, 2, 5, 25, 
}

var _cte_key_offsets []int16 = []int16{
	0, 0, 16, 18, 20, 21, 22, 23, 
	24, 25, 26, 27, 28, 29, 30, 32, 
	34, 36, 38, 43, 60, 65, 81, 83, 
	90, 92, 93, 94, 95, 96, 97, 98, 
	99, 100, 101, 114, 120, 122, 138, 143, 
	145, 152, 154, 155, 156, 157, 158, 159, 
	160, 161, 162, 163, 175, 177, 185, 187, 
	188, 189, 190, 191, 192, 193, 194, 207, 
	213, 215, 231, 236, 238, 245, 247, 248, 
	249, 250, 251, 252, 253, 254, 255, 256, 
	268, 270, 278, 280, 281, 282, 283, 284, 
	285, 286, 287, 300, 306, 308, 324, 329, 
	331, 338, 340, 341, 342, 343, 344, 345, 
	346, 347, 348, 349, 361, 363, 371, 373, 
	374, 375, 376, 377, 378, 379, 380, 384, 
	390, 390, 392, 392, 392, 392, 392, 
}

var _cte_trans_keys []byte = []byte{
	13, 32, 34, 40, 45, 47, 60, 91, 
	102, 110, 116, 123, 9, 10, 48, 57, 
	48, 57, 42, 47, 97, 108, 115, 101, 
	105, 108, 114, 117, 101, 10, 42, 47, 
	42, 47, 42, 47, 34, 92, 34, 92, 
	110, 114, 116, 13, 32, 34, 40, 45, 
	47, 60, 91, 93, 102, 110, 116, 123, 
	9, 10, 48, 57, 13, 32, 93, 9, 
	10, 13, 32, 34, 40, 45, 47, 60, 
	91, 102, 110, 116, 123, 9, 10, 48, 
	57, 48, 57, 13, 32, 93, 9, 10, 
	48, 57, 42, 47, 97, 108, 115, 101, 
	105, 108, 114, 117, 101, 13, 32, 34, 
	40, 45, 47, 102, 116, 125, 9, 10, 
	48, 57, 13, 32, 47, 61, 9, 10, 
	42, 47, 13, 32, 34, 40, 45, 47, 
	60, 91, 102, 110, 116, 123, 9, 10, 
	48, 57, 13, 32, 125, 9, 10, 48, 
	57, 13, 32, 125, 9, 10, 48, 57, 
	42, 47, 97, 108, 115, 101, 105, 108, 
	114, 117, 101, 13, 32, 34, 40, 45, 
	47, 102, 116, 9, 10, 48, 57, 48, 
	57, 13, 32, 47, 61, 9, 10, 48, 
	57, 42, 47, 97, 108, 115, 101, 114, 
	117, 101, 13, 32, 34, 40, 45, 47, 
	62, 102, 116, 9, 10, 48, 57, 13, 
	32, 47, 61, 9, 10, 42, 47, 13, 
	32, 34, 40, 45, 47, 60, 91, 102, 
	110, 116, 123, 9, 10, 48, 57, 13, 
	32, 62, 9, 10, 48, 57, 13, 32, 
	62, 9, 10, 48, 57, 42, 47, 97, 
	108, 115, 101, 105, 108, 114, 117, 101, 
	13, 32, 34, 40, 45, 47, 102, 116, 
	9, 10, 48, 57, 48, 57, 13, 32, 
	47, 61, 9, 10, 48, 57, 42, 47, 
	97, 108, 115, 101, 114, 117, 101, 13, 
	32, 34, 40, 41, 45, 47, 102, 116, 
	9, 10, 48, 57, 13, 32, 47, 61, 
	9, 10, 42, 47, 13, 32, 34, 40, 
	45, 47, 60, 91, 102, 110, 116, 123, 
	9, 10, 48, 57, 13, 32, 41, 9, 
	10, 48, 57, 13, 32, 41, 9, 10, 
	48, 57, 42, 47, 97, 108, 115, 101, 
	105, 108, 114, 117, 101, 13, 32, 34, 
	40, 45, 47, 102, 116, 9, 10, 48, 
	57, 48, 57, 13, 32, 47, 61, 9, 
	10, 48, 57, 42, 47, 97, 108, 115, 
	101, 114, 117, 101, 13, 32, 9, 10, 
	13, 32, 9, 10, 48, 57, 42, 47, 
	
}

var _cte_single_lengths []byte = []byte{
	0, 12, 0, 2, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 2, 2, 
	2, 2, 5, 13, 3, 12, 0, 3, 
	2, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 9, 4, 2, 12, 3, 0, 
	3, 2, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 8, 0, 4, 2, 1, 
	1, 1, 1, 1, 1, 1, 9, 4, 
	2, 12, 3, 0, 3, 2, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 8, 
	0, 4, 2, 1, 1, 1, 1, 1, 
	1, 1, 9, 4, 2, 12, 3, 0, 
	3, 2, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 8, 0, 4, 2, 1, 
	1, 1, 1, 1, 1, 1, 2, 2, 
	0, 2, 0, 0, 0, 0, 0, 
}

var _cte_range_lengths []byte = []byte{
	0, 2, 1, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 2, 1, 2, 1, 2, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 2, 1, 0, 2, 1, 1, 
	2, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 2, 1, 2, 0, 0, 
	0, 0, 0, 0, 0, 0, 2, 1, 
	0, 2, 1, 1, 2, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 2, 
	1, 2, 0, 0, 0, 0, 0, 0, 
	0, 0, 2, 1, 0, 2, 1, 1, 
	2, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 2, 1, 2, 0, 0, 
	0, 0, 0, 0, 0, 0, 1, 2, 
	0, 0, 0, 0, 0, 0, 0, 
}

var _cte_index_offsets []int16 = []int16{
	0, 0, 15, 17, 20, 22, 24, 26, 
	28, 30, 32, 34, 36, 38, 40, 43, 
	46, 49, 52, 58, 74, 79, 94, 96, 
	102, 105, 107, 109, 111, 113, 115, 117, 
	119, 121, 123, 135, 141, 144, 159, 164, 
	166, 172, 175, 177, 179, 181, 183, 185, 
	187, 189, 191, 193, 204, 206, 213, 216, 
	218, 220, 222, 224, 226, 228, 230, 242, 
	248, 251, 266, 271, 273, 279, 282, 284, 
	286, 288, 290, 292, 294, 296, 298, 300, 
	311, 313, 320, 323, 325, 327, 329, 331, 
	333, 335, 337, 349, 355, 358, 373, 378, 
	380, 386, 389, 391, 393, 395, 397, 399, 
	401, 403, 405, 407, 418, 420, 427, 430, 
	432, 434, 436, 438, 440, 442, 444, 448, 
	453, 454, 457, 458, 459, 460, 461, 
}

var _cte_indicies []byte = []byte{
	0, 0, 2, 3, 4, 5, 7, 8, 
	9, 10, 11, 12, 0, 6, 1, 6, 
	1, 13, 14, 1, 15, 1, 16, 1, 
	17, 1, 18, 1, 19, 1, 20, 1, 
	21, 1, 22, 1, 23, 1, 25, 24, 
	27, 28, 26, 27, 29, 26, 30, 28, 
	26, 32, 33, 31, 35, 36, 37, 38, 
	39, 34, 40, 40, 41, 42, 43, 44, 
	46, 47, 48, 49, 50, 51, 52, 40, 
	45, 1, 40, 40, 48, 40, 1, 53, 
	53, 41, 42, 43, 44, 46, 47, 49, 
	50, 51, 52, 53, 45, 1, 45, 1, 
	54, 54, 55, 54, 45, 1, 56, 57, 
	1, 58, 1, 59, 1, 60, 1, 61, 
	1, 62, 1, 63, 1, 64, 1, 65, 
	1, 66, 1, 67, 67, 68, 69, 70, 
	71, 73, 74, 75, 67, 72, 1, 76, 
	76, 77, 78, 76, 1, 79, 80, 1, 
	78, 78, 81, 82, 83, 84, 86, 87, 
	88, 89, 90, 91, 78, 85, 1, 67, 
	67, 75, 67, 1, 85, 1, 92, 92, 
	93, 92, 85, 1, 94, 95, 1, 96, 
	1, 97, 1, 98, 1, 99, 1, 100, 
	1, 101, 1, 102, 1, 103, 1, 104, 
	1, 105, 105, 68, 69, 70, 71, 73, 
	74, 105, 72, 1, 72, 1, 106, 106, 
	107, 108, 106, 72, 1, 109, 110, 1, 
	111, 1, 112, 1, 113, 1, 114, 1, 
	115, 1, 116, 1, 117, 1, 118, 118, 
	119, 120, 121, 122, 124, 125, 126, 118, 
	123, 1, 127, 127, 128, 129, 127, 1, 
	130, 131, 1, 129, 129, 132, 133, 134, 
	135, 137, 138, 139, 140, 141, 142, 129, 
	136, 1, 118, 118, 124, 118, 1, 136, 
	1, 143, 143, 144, 143, 136, 1, 145, 
	146, 1, 147, 1, 148, 1, 149, 1, 
	150, 1, 151, 1, 152, 1, 153, 1, 
	154, 1, 155, 1, 156, 156, 119, 120, 
	121, 122, 125, 126, 156, 123, 1, 123, 
	1, 157, 157, 158, 159, 157, 123, 1, 
	160, 161, 1, 162, 1, 163, 1, 164, 
	1, 165, 1, 166, 1, 167, 1, 168, 
	1, 169, 169, 170, 171, 172, 173, 174, 
	176, 177, 169, 175, 1, 178, 178, 179, 
	180, 178, 1, 181, 182, 1, 180, 180, 
	183, 184, 185, 186, 188, 189, 190, 191, 
	192, 193, 180, 187, 1, 169, 169, 172, 
	169, 1, 187, 1, 194, 194, 195, 194, 
	187, 1, 196, 197, 1, 198, 1, 199, 
	1, 200, 1, 201, 1, 202, 1, 203, 
	1, 204, 1, 205, 1, 206, 1, 207, 
	207, 170, 171, 173, 174, 176, 177, 207, 
	175, 1, 175, 1, 208, 208, 209, 210, 
	208, 175, 1, 211, 212, 1, 213, 1, 
	214, 1, 215, 1, 216, 1, 217, 1, 
	218, 1, 219, 1, 220, 220, 220, 1, 
	221, 221, 221, 6, 1, 1, 30, 28, 
	26, 1, 1, 1, 1, 1, 
}

var _cte_trans_targs []byte = []byte{
	1, 0, 118, 1, 2, 3, 119, 118, 
	118, 4, 8, 10, 118, 1, 1, 5, 
	6, 7, 118, 9, 118, 11, 12, 118, 
	13, 120, 14, 15, 16, 121, 15, 17, 
	122, 18, 17, 17, 17, 17, 17, 17, 
	19, 20, 21, 22, 24, 23, 20, 20, 
	123, 25, 29, 31, 20, 21, 19, 123, 
	21, 21, 26, 27, 28, 20, 30, 20, 
	32, 33, 20, 34, 35, 51, 52, 54, 
	53, 55, 59, 124, 35, 36, 37, 35, 
	35, 38, 37, 39, 41, 40, 38, 38, 
	42, 46, 48, 38, 34, 124, 37, 37, 
	43, 44, 45, 38, 47, 38, 49, 50, 
	38, 51, 35, 36, 37, 51, 51, 56, 
	57, 58, 35, 60, 61, 35, 62, 63, 
	79, 80, 82, 81, 125, 83, 87, 63, 
	64, 65, 63, 63, 66, 65, 67, 69, 
	68, 66, 66, 70, 74, 76, 66, 62, 
	125, 65, 65, 71, 72, 73, 66, 75, 
	66, 77, 78, 66, 79, 63, 64, 65, 
	79, 79, 84, 85, 86, 63, 88, 89, 
	63, 90, 91, 107, 126, 108, 110, 109, 
	111, 115, 91, 92, 93, 91, 91, 94, 
	93, 95, 97, 96, 94, 94, 98, 102, 
	104, 94, 90, 126, 93, 93, 99, 100, 
	101, 94, 103, 94, 105, 106, 94, 107, 
	91, 92, 93, 107, 107, 112, 113, 114, 
	91, 116, 117, 91, 118, 118, 
}

var _cte_trans_actions []byte = []byte{
	0, 0, 25, 19, 7, 0, 9, 17, 
	13, 0, 0, 0, 15, 23, 21, 0, 
	0, 0, 5, 0, 1, 0, 0, 3, 
	0, 27, 0, 0, 0, 29, 23, 0, 
	43, 0, 41, 39, 31, 33, 35, 37, 
	0, 25, 19, 7, 0, 9, 17, 13, 
	45, 0, 0, 0, 15, 0, 11, 53, 
	23, 21, 0, 0, 0, 5, 0, 1, 
	0, 0, 3, 0, 25, 19, 7, 0, 
	9, 0, 0, 47, 0, 0, 0, 23, 
	21, 25, 19, 7, 0, 9, 17, 13, 
	0, 0, 0, 15, 11, 56, 23, 21, 
	0, 0, 0, 5, 0, 1, 0, 0, 
	3, 0, 11, 11, 11, 23, 21, 0, 
	0, 0, 5, 0, 0, 3, 0, 25, 
	19, 7, 0, 9, 49, 0, 0, 0, 
	0, 0, 23, 21, 25, 19, 7, 0, 
	9, 17, 13, 0, 0, 0, 15, 11, 
	59, 23, 21, 0, 0, 0, 5, 0, 
	1, 0, 0, 3, 0, 11, 11, 11, 
	23, 21, 0, 0, 0, 5, 0, 0, 
	3, 0, 25, 19, 51, 7, 0, 9, 
	0, 0, 0, 0, 0, 23, 21, 25, 
	19, 7, 0, 9, 17, 13, 0, 0, 
	0, 15, 11, 62, 23, 21, 0, 0, 
	0, 5, 0, 1, 0, 0, 3, 0, 
	11, 11, 11, 23, 21, 0, 0, 0, 
	5, 0, 0, 3, 0, 11, 
}

var _cte_eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 11, 
	0, 0, 0, 0, 0, 0, 0, 
}

const cte_start int = 1
const cte_error int = 0

const cte_en_comment_iterate int = 13
const cte_en_multiline_comment_iterate int = 14
const cte_en_string_iterate int = 17
const cte_en_list_iterate int = 19
const cte_en_unordered_map_iterate int = 34
const cte_en_ordered_map_iterate int = 62
const cte_en_metadata_map_iterate int = 90
const cte_en_main int = 1


//line cte.rl:265

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
    
    
//line cte.go:394
	{
	 this.cs = cte_start
	 this.top = 0
	}

//line cte.go:400
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
//line cte.rl:66

        this.significandSign = -1
    
		case 4:
//line cte.rl:70

        this.significand = this.significand * 10 + uint64( this.data[p]) - uint64('0')
    
		case 5:
//line cte.rl:74

        value := this.significand
        sign := this.significandSign
        this.significand = 0
        this.significandSign = 1
        if sign < 0 {
            err = callbacks.OnNegativeInt(value)
        } else {
            err = callbacks.OnPositiveInt(value)
        }
        if err != nil {
            p++; goto _out

        }
    
		case 6:
//line cte.rl:89

        err = callbacks.OnListBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 19; goto _again

    
		case 7:
//line cte.rl:97

        err = callbacks.OnUnorderedMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 34; goto _again

    
		case 8:
//line cte.rl:105

        err = callbacks.OnOrderedMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 62; goto _again

    
		case 9:
//line cte.rl:113

        err = callbacks.OnMetadataMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 90; goto _again

    
		case 10:
//line cte.rl:121

        this.arrayStart = p + 1
        err = callbacks.OnCommentBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 13; goto _again

    
		case 11:
//line cte.rl:130

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
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 14; goto _again

    
		case 12:
//line cte.rl:147

        this.arrayStart = p + 1
        err = callbacks.OnStringBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 17; goto _again

    
		case 13:
//line cte.rl:166

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

        
		case 14:
//line cte.rl:179

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

        
		case 15:
//line cte.rl:199
this.flushAndAddEscapedCharacter(p-1, '\\', callbacks)
		case 16:
//line cte.rl:200
this.flushAndAddEscapedCharacter(p-1, '\n', callbacks)
		case 17:
//line cte.rl:201
this.flushAndAddEscapedCharacter(p-1, '\r', callbacks)
		case 18:
//line cte.rl:202
this.flushAndAddEscapedCharacter(p-1, '\t', callbacks)
		case 19:
//line cte.rl:203
this.flushAndAddEscapedCharacter(p-1, '"', callbacks)
		case 20:
//line cte.rl:204
return false, fmt.Errorf("\\%c: Illegal escape encoding", this.data[p])
		case 21:
//line cte.rl:208

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

        
		case 22:
//line cte.rl:222

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 23:
//line cte.rl:232

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 24:
//line cte.rl:242

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 25:
//line cte.rl:252

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
//line cte.go:743
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
			case 5:
//line cte.rl:74

        value := this.significand
        sign := this.significandSign
        this.significand = 0
        this.significandSign = 1
        if sign < 0 {
            err = callbacks.OnNegativeInt(value)
        } else {
            err = callbacks.OnPositiveInt(value)
        }
        if err != nil {
            p++; goto _out

        }
    
//line cte.go:779
			}
		}
	}

	_out: {}
	}

//line cte.rl:325


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
