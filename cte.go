
//line cte.rl:1
package cte

import (
    "fmt"
//    "math"
//    "time"
)

type CteDecoderCallbacks interface {
    OnNil() error
    OnBool(value bool) error
    OnPositiveInt(value uint64) error
    OnNegativeInt(value uint64) error
    OnDecimalFloat(significand int64, exponent int) error
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


//line cte.rl:297




//line cte.go:41
var _cte_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 8, 1, 9, 1, 10, 1, 11, 
	1, 12, 1, 13, 1, 14, 1, 15, 
	1, 16, 1, 17, 1, 18, 1, 19, 
	1, 20, 1, 21, 1, 22, 1, 23, 
	1, 24, 1, 25, 1, 26, 1, 27, 
	1, 28, 1, 29, 1, 30, 2, 9, 
	27, 2, 9, 28, 2, 9, 29, 2, 
	9, 30, 2, 10, 27, 2, 10, 28, 
	2, 10, 29, 2, 10, 30, 
}

var _cte_key_offsets []int16 = []int16{
	0, 0, 17, 20, 23, 27, 29, 31, 
	32, 33, 34, 35, 36, 37, 38, 39, 
	40, 41, 43, 45, 47, 49, 54, 72, 
	75, 83, 86, 95, 99, 101, 108, 113, 
	130, 132, 133, 134, 135, 136, 137, 138, 
	139, 140, 141, 155, 158, 167, 173, 175, 
	192, 195, 203, 206, 215, 219, 221, 228, 
	233, 235, 236, 237, 238, 239, 240, 241, 
	242, 243, 244, 247, 257, 261, 263, 271, 
	284, 286, 287, 288, 289, 290, 291, 292, 
	293, 307, 310, 319, 325, 327, 344, 347, 
	355, 358, 367, 371, 373, 380, 385, 387, 
	388, 389, 390, 391, 392, 393, 394, 395, 
	396, 399, 409, 413, 415, 423, 436, 438, 
	439, 440, 441, 442, 443, 444, 445, 459, 
	462, 471, 477, 479, 496, 499, 507, 510, 
	519, 523, 525, 532, 537, 539, 540, 541, 
	542, 543, 544, 545, 546, 547, 548, 551, 
	561, 565, 567, 575, 588, 590, 591, 592, 
	593, 594, 595, 596, 597, 604, 608, 616, 
	622, 622, 624, 624, 624, 624, 624, 
}

var _cte_trans_keys []byte = []byte{
	0, 13, 32, 34, 40, 45, 47, 60, 
	91, 102, 110, 116, 123, 9, 10, 48, 
	57, 0, 48, 57, 48, 49, 57, 43, 
	45, 48, 57, 48, 57, 42, 47, 97, 
	108, 115, 101, 105, 108, 114, 117, 101, 
	10, 42, 47, 42, 47, 42, 47, 34, 
	92, 34, 92, 110, 114, 116, 0, 13, 
	32, 34, 40, 45, 47, 60, 91, 93, 
	102, 110, 116, 123, 9, 10, 48, 57, 
	0, 48, 57, 13, 32, 46, 93, 9, 
	10, 48, 57, 48, 49, 57, 13, 32, 
	48, 93, 101, 9, 10, 49, 57, 43, 
	45, 48, 57, 48, 57, 13, 32, 93, 
	9, 10, 48, 57, 13, 32, 93, 9, 
	10, 0, 13, 32, 34, 40, 45, 47, 
	60, 91, 102, 110, 116, 123, 9, 10, 
	48, 57, 42, 47, 97, 108, 115, 101, 
	105, 108, 114, 117, 101, 0, 13, 32, 
	34, 40, 45, 47, 102, 116, 125, 9, 
	10, 48, 57, 0, 48, 57, 13, 32, 
	46, 47, 61, 9, 10, 48, 57, 13, 
	32, 47, 61, 9, 10, 42, 47, 0, 
	13, 32, 34, 40, 45, 47, 60, 91, 
	102, 110, 116, 123, 9, 10, 48, 57, 
	0, 48, 57, 13, 32, 46, 125, 9, 
	10, 48, 57, 48, 49, 57, 13, 32, 
	48, 101, 125, 9, 10, 49, 57, 43, 
	45, 48, 57, 48, 57, 13, 32, 125, 
	9, 10, 48, 57, 13, 32, 125, 9, 
	10, 42, 47, 97, 108, 115, 101, 105, 
	108, 114, 117, 101, 48, 49, 57, 13, 
	32, 47, 48, 61, 101, 9, 10, 49, 
	57, 43, 45, 48, 57, 48, 57, 13, 
	32, 47, 61, 9, 10, 48, 57, 0, 
	13, 32, 34, 40, 45, 47, 102, 116, 
	9, 10, 48, 57, 42, 47, 97, 108, 
	115, 101, 114, 117, 101, 0, 13, 32, 
	34, 40, 45, 47, 62, 102, 116, 9, 
	10, 48, 57, 0, 48, 57, 13, 32, 
	46, 47, 61, 9, 10, 48, 57, 13, 
	32, 47, 61, 9, 10, 42, 47, 0, 
	13, 32, 34, 40, 45, 47, 60, 91, 
	102, 110, 116, 123, 9, 10, 48, 57, 
	0, 48, 57, 13, 32, 46, 62, 9, 
	10, 48, 57, 48, 49, 57, 13, 32, 
	48, 62, 101, 9, 10, 49, 57, 43, 
	45, 48, 57, 48, 57, 13, 32, 62, 
	9, 10, 48, 57, 13, 32, 62, 9, 
	10, 42, 47, 97, 108, 115, 101, 105, 
	108, 114, 117, 101, 48, 49, 57, 13, 
	32, 47, 48, 61, 101, 9, 10, 49, 
	57, 43, 45, 48, 57, 48, 57, 13, 
	32, 47, 61, 9, 10, 48, 57, 0, 
	13, 32, 34, 40, 45, 47, 102, 116, 
	9, 10, 48, 57, 42, 47, 97, 108, 
	115, 101, 114, 117, 101, 0, 13, 32, 
	34, 40, 41, 45, 47, 102, 116, 9, 
	10, 48, 57, 0, 48, 57, 13, 32, 
	46, 47, 61, 9, 10, 48, 57, 13, 
	32, 47, 61, 9, 10, 42, 47, 0, 
	13, 32, 34, 40, 45, 47, 60, 91, 
	102, 110, 116, 123, 9, 10, 48, 57, 
	0, 48, 57, 13, 32, 41, 46, 9, 
	10, 48, 57, 48, 49, 57, 13, 32, 
	41, 48, 101, 9, 10, 49, 57, 43, 
	45, 48, 57, 48, 57, 13, 32, 41, 
	9, 10, 48, 57, 13, 32, 41, 9, 
	10, 42, 47, 97, 108, 115, 101, 105, 
	108, 114, 117, 101, 48, 49, 57, 13, 
	32, 47, 48, 61, 101, 9, 10, 49, 
	57, 43, 45, 48, 57, 48, 57, 13, 
	32, 47, 61, 9, 10, 48, 57, 0, 
	13, 32, 34, 40, 45, 47, 102, 116, 
	9, 10, 48, 57, 42, 47, 97, 108, 
	115, 101, 114, 117, 101, 13, 32, 46, 
	9, 10, 48, 57, 13, 32, 9, 10, 
	13, 32, 48, 101, 9, 10, 49, 57, 
	13, 32, 9, 10, 48, 57, 42, 47, 
	
}

var _cte_single_lengths []byte = []byte{
	0, 13, 1, 1, 2, 0, 2, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 2, 2, 2, 2, 5, 14, 1, 
	4, 1, 5, 2, 0, 3, 3, 13, 
	2, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 10, 1, 5, 4, 2, 13, 
	1, 4, 1, 5, 2, 0, 3, 3, 
	2, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 6, 2, 0, 4, 9, 
	2, 1, 1, 1, 1, 1, 1, 1, 
	10, 1, 5, 4, 2, 13, 1, 4, 
	1, 5, 2, 0, 3, 3, 2, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 6, 2, 0, 4, 9, 2, 1, 
	1, 1, 1, 1, 1, 1, 10, 1, 
	5, 4, 2, 13, 1, 4, 1, 5, 
	2, 0, 3, 3, 2, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 6, 
	2, 0, 4, 9, 2, 1, 1, 1, 
	1, 1, 1, 1, 3, 2, 4, 2, 
	0, 2, 0, 0, 0, 0, 0, 
}

var _cte_range_lengths []byte = []byte{
	0, 2, 1, 1, 1, 1, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 2, 1, 
	2, 1, 2, 1, 1, 2, 1, 2, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 2, 1, 2, 1, 0, 2, 
	1, 2, 1, 2, 1, 1, 2, 1, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 1, 2, 1, 1, 2, 2, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	2, 1, 2, 1, 0, 2, 1, 2, 
	1, 2, 1, 1, 2, 1, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	1, 2, 1, 1, 2, 2, 0, 0, 
	0, 0, 0, 0, 0, 0, 2, 1, 
	2, 1, 0, 2, 1, 2, 1, 2, 
	1, 1, 2, 1, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 1, 2, 
	1, 1, 2, 2, 0, 0, 0, 0, 
	0, 0, 0, 0, 2, 1, 2, 2, 
	0, 0, 0, 0, 0, 0, 0, 
}

var _cte_index_offsets []int16 = []int16{
	0, 0, 16, 19, 22, 26, 28, 31, 
	33, 35, 37, 39, 41, 43, 45, 47, 
	49, 51, 54, 57, 60, 63, 69, 86, 
	89, 96, 99, 107, 111, 113, 119, 124, 
	140, 143, 145, 147, 149, 151, 153, 155, 
	157, 159, 161, 174, 177, 185, 191, 194, 
	210, 213, 220, 223, 231, 235, 237, 243, 
	248, 251, 253, 255, 257, 259, 261, 263, 
	265, 267, 269, 272, 281, 285, 287, 294, 
	306, 309, 311, 313, 315, 317, 319, 321, 
	323, 336, 339, 347, 353, 356, 372, 375, 
	382, 385, 393, 397, 399, 405, 410, 413, 
	415, 417, 419, 421, 423, 425, 427, 429, 
	431, 434, 443, 447, 449, 456, 468, 471, 
	473, 475, 477, 479, 481, 483, 485, 498, 
	501, 509, 515, 518, 534, 537, 544, 547, 
	555, 559, 561, 567, 572, 575, 577, 579, 
	581, 583, 585, 587, 589, 591, 593, 596, 
	605, 609, 611, 618, 630, 633, 635, 637, 
	639, 641, 643, 645, 647, 653, 657, 664, 
	669, 670, 673, 674, 675, 676, 677, 
}

var _cte_trans_targs []byte = []byte{
	2, 1, 1, 157, 1, 2, 6, 157, 
	157, 7, 11, 13, 157, 1, 156, 0, 
	2, 156, 0, 158, 158, 0, 5, 5, 
	159, 0, 159, 0, 1, 1, 0, 8, 
	0, 9, 0, 10, 0, 157, 0, 12, 
	0, 157, 0, 14, 0, 15, 0, 157, 
	0, 160, 16, 18, 19, 17, 18, 161, 
	17, 18, 19, 17, 162, 21, 20, 20, 
	20, 20, 20, 20, 20, 23, 22, 22, 
	30, 31, 23, 32, 30, 30, 163, 33, 
	37, 39, 30, 22, 24, 0, 23, 24, 
	0, 22, 22, 25, 163, 22, 24, 0, 
	26, 26, 0, 22, 22, 26, 163, 27, 
	22, 26, 0, 28, 28, 29, 0, 29, 
	0, 22, 22, 163, 22, 29, 0, 22, 
	22, 163, 22, 0, 23, 31, 31, 30, 
	31, 23, 32, 30, 30, 33, 37, 39, 
	30, 31, 24, 0, 31, 31, 0, 34, 
	0, 35, 0, 36, 0, 30, 0, 38, 
	0, 30, 0, 40, 0, 41, 0, 30, 
	0, 43, 42, 42, 45, 71, 43, 72, 
	73, 77, 164, 42, 44, 0, 43, 44, 
	0, 45, 45, 66, 46, 47, 45, 44, 
	0, 45, 45, 46, 47, 45, 0, 45, 
	45, 0, 48, 47, 47, 55, 47, 48, 
	56, 55, 55, 57, 61, 63, 55, 47, 
	49, 0, 48, 49, 0, 42, 42, 50, 
	164, 42, 49, 0, 51, 51, 0, 42, 
	42, 51, 52, 164, 42, 51, 0, 53, 
	53, 54, 0, 54, 0, 42, 42, 164, 
	42, 54, 0, 42, 42, 164, 42, 0, 
	47, 47, 0, 58, 0, 59, 0, 60, 
	0, 55, 0, 62, 0, 55, 0, 64, 
	0, 65, 0, 55, 0, 67, 67, 0, 
	45, 45, 46, 67, 47, 68, 45, 67, 
	0, 69, 69, 70, 0, 70, 0, 45, 
	45, 46, 47, 45, 70, 0, 43, 71, 
	71, 45, 71, 43, 72, 73, 77, 71, 
	44, 0, 71, 71, 0, 74, 0, 75, 
	0, 76, 0, 45, 0, 78, 0, 79, 
	0, 45, 0, 81, 80, 80, 83, 109, 
	81, 110, 165, 111, 115, 80, 82, 0, 
	81, 82, 0, 83, 83, 104, 84, 85, 
	83, 82, 0, 83, 83, 84, 85, 83, 
	0, 83, 83, 0, 86, 85, 85, 93, 
	85, 86, 94, 93, 93, 95, 99, 101, 
	93, 85, 87, 0, 86, 87, 0, 80, 
	80, 88, 165, 80, 87, 0, 89, 89, 
	0, 80, 80, 89, 165, 90, 80, 89, 
	0, 91, 91, 92, 0, 92, 0, 80, 
	80, 165, 80, 92, 0, 80, 80, 165, 
	80, 0, 85, 85, 0, 96, 0, 97, 
	0, 98, 0, 93, 0, 100, 0, 93, 
	0, 102, 0, 103, 0, 93, 0, 105, 
	105, 0, 83, 83, 84, 105, 85, 106, 
	83, 105, 0, 107, 107, 108, 0, 108, 
	0, 83, 83, 84, 85, 83, 108, 0, 
	81, 109, 109, 83, 109, 81, 110, 111, 
	115, 109, 82, 0, 109, 109, 0, 112, 
	0, 113, 0, 114, 0, 83, 0, 116, 
	0, 117, 0, 83, 0, 119, 118, 118, 
	121, 147, 166, 119, 148, 149, 153, 118, 
	120, 0, 119, 120, 0, 121, 121, 142, 
	122, 123, 121, 120, 0, 121, 121, 122, 
	123, 121, 0, 121, 121, 0, 124, 123, 
	123, 131, 123, 124, 132, 131, 131, 133, 
	137, 139, 131, 123, 125, 0, 124, 125, 
	0, 118, 118, 166, 126, 118, 125, 0, 
	127, 127, 0, 118, 118, 166, 127, 128, 
	118, 127, 0, 129, 129, 130, 0, 130, 
	0, 118, 118, 166, 118, 130, 0, 118, 
	118, 166, 118, 0, 123, 123, 0, 134, 
	0, 135, 0, 136, 0, 131, 0, 138, 
	0, 131, 0, 140, 0, 141, 0, 131, 
	0, 143, 143, 0, 121, 121, 122, 143, 
	123, 144, 121, 143, 0, 145, 145, 146, 
	0, 146, 0, 121, 121, 122, 123, 121, 
	146, 0, 119, 147, 147, 121, 147, 119, 
	148, 149, 153, 147, 120, 0, 147, 147, 
	0, 150, 0, 151, 0, 152, 0, 121, 
	0, 154, 0, 155, 0, 121, 0, 157, 
	157, 3, 157, 156, 0, 157, 157, 157, 
	0, 157, 157, 158, 4, 157, 158, 0, 
	157, 157, 157, 159, 0, 0, 18, 19, 
	17, 0, 0, 0, 0, 0, 
}

var _cte_trans_actions []byte = []byte{
	0, 0, 0, 35, 29, 7, 0, 27, 
	23, 0, 0, 0, 25, 0, 11, 0, 
	0, 11, 0, 13, 15, 0, 0, 9, 
	17, 0, 17, 0, 33, 31, 0, 0, 
	0, 0, 0, 0, 0, 5, 0, 0, 
	0, 1, 0, 0, 0, 0, 0, 3, 
	0, 37, 0, 0, 0, 0, 0, 39, 
	0, 33, 0, 0, 53, 0, 0, 49, 
	41, 43, 45, 47, 51, 0, 0, 0, 
	35, 29, 7, 0, 27, 23, 55, 0, 
	0, 0, 25, 0, 11, 0, 0, 11, 
	0, 19, 19, 0, 63, 19, 11, 0, 
	13, 15, 0, 21, 21, 13, 75, 0, 
	21, 15, 0, 0, 9, 17, 0, 17, 
	0, 21, 21, 75, 21, 17, 0, 0, 
	0, 55, 0, 0, 0, 0, 0, 35, 
	29, 7, 0, 27, 23, 0, 0, 0, 
	25, 0, 11, 0, 33, 31, 0, 0, 
	0, 0, 0, 0, 0, 5, 0, 0, 
	0, 1, 0, 0, 0, 0, 0, 3, 
	0, 0, 0, 0, 35, 29, 7, 0, 
	0, 0, 57, 0, 11, 0, 0, 11, 
	0, 19, 19, 0, 19, 19, 19, 11, 
	0, 0, 0, 0, 0, 0, 0, 33, 
	31, 0, 0, 0, 0, 35, 29, 7, 
	0, 27, 23, 0, 0, 0, 25, 0, 
	11, 0, 0, 11, 0, 19, 19, 0, 
	66, 19, 11, 0, 13, 15, 0, 21, 
	21, 13, 0, 78, 21, 15, 0, 0, 
	9, 17, 0, 17, 0, 21, 21, 78, 
	21, 17, 0, 0, 0, 57, 0, 0, 
	33, 31, 0, 0, 0, 0, 0, 0, 
	0, 5, 0, 0, 0, 1, 0, 0, 
	0, 0, 0, 3, 0, 13, 15, 0, 
	21, 21, 21, 13, 21, 0, 21, 15, 
	0, 0, 9, 17, 0, 17, 0, 21, 
	21, 21, 21, 21, 17, 0, 0, 0, 
	0, 35, 29, 7, 0, 0, 0, 0, 
	11, 0, 33, 31, 0, 0, 0, 0, 
	0, 0, 0, 5, 0, 0, 0, 0, 
	0, 3, 0, 0, 0, 0, 35, 29, 
	7, 0, 59, 0, 0, 0, 11, 0, 
	0, 11, 0, 19, 19, 0, 19, 19, 
	19, 11, 0, 0, 0, 0, 0, 0, 
	0, 33, 31, 0, 0, 0, 0, 35, 
	29, 7, 0, 27, 23, 0, 0, 0, 
	25, 0, 11, 0, 0, 11, 0, 19, 
	19, 0, 69, 19, 11, 0, 13, 15, 
	0, 21, 21, 13, 81, 0, 21, 15, 
	0, 0, 9, 17, 0, 17, 0, 21, 
	21, 81, 21, 17, 0, 0, 0, 59, 
	0, 0, 33, 31, 0, 0, 0, 0, 
	0, 0, 0, 5, 0, 0, 0, 1, 
	0, 0, 0, 0, 0, 3, 0, 13, 
	15, 0, 21, 21, 21, 13, 21, 0, 
	21, 15, 0, 0, 9, 17, 0, 17, 
	0, 21, 21, 21, 21, 21, 17, 0, 
	0, 0, 0, 35, 29, 7, 0, 0, 
	0, 0, 11, 0, 33, 31, 0, 0, 
	0, 0, 0, 0, 0, 5, 0, 0, 
	0, 0, 0, 3, 0, 0, 0, 0, 
	35, 29, 61, 7, 0, 0, 0, 0, 
	11, 0, 0, 11, 0, 19, 19, 0, 
	19, 19, 19, 11, 0, 0, 0, 0, 
	0, 0, 0, 33, 31, 0, 0, 0, 
	0, 35, 29, 7, 0, 27, 23, 0, 
	0, 0, 25, 0, 11, 0, 0, 11, 
	0, 19, 19, 72, 0, 19, 11, 0, 
	13, 15, 0, 21, 21, 84, 13, 0, 
	21, 15, 0, 0, 9, 17, 0, 17, 
	0, 21, 21, 84, 21, 17, 0, 0, 
	0, 61, 0, 0, 33, 31, 0, 0, 
	0, 0, 0, 0, 0, 5, 0, 0, 
	0, 1, 0, 0, 0, 0, 0, 3, 
	0, 13, 15, 0, 21, 21, 21, 13, 
	21, 0, 21, 15, 0, 0, 9, 17, 
	0, 17, 0, 21, 21, 21, 21, 21, 
	17, 0, 0, 0, 0, 35, 29, 7, 
	0, 0, 0, 0, 11, 0, 33, 31, 
	0, 0, 0, 0, 0, 0, 0, 5, 
	0, 0, 0, 0, 0, 3, 0, 19, 
	19, 0, 19, 11, 0, 0, 0, 0, 
	0, 21, 21, 13, 0, 21, 15, 0, 
	21, 21, 21, 17, 0, 0, 33, 0, 
	0, 0, 0, 0, 0, 0, 
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
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 19, 0, 21, 21, 
	0, 0, 0, 0, 0, 0, 0, 
}

const cte_start int = 1
const cte_error int = 0

const cte_en_comment_iterate int = 16
const cte_en_multiline_comment_iterate int = 17
const cte_en_string_iterate int = 20
const cte_en_list_iterate int = 22
const cte_en_unordered_map_iterate int = 42
const cte_en_ordered_map_iterate int = 80
const cte_en_metadata_map_iterate int = 118
const cte_en_main int = 1


//line cte.rl:301

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
    zeroAdjust int
    zeroMagnitude uint64
}

func (this *Parser) Init(maxDepth int) {
    this.stack = make([]int, maxDepth)
    this.significandSign = 1
    this.exponentSign = 1
    this.zeroMagnitude = 1
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
    
    
//line cte.go:512
	{
	 this.cs = cte_start
	 this.top = 0
	}

//line cte.go:518
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
//line cte.rl:47

        err = callbacks.OnNil()
        if err != nil {
            p++; goto _out

        }
    
		case 1:
//line cte.rl:54

        err = callbacks.OnBool(true)
        if err != nil {
            p++; goto _out

        }
    
		case 2:
//line cte.rl:61

        err = callbacks.OnBool(false)
        if err != nil {
            p++; goto _out

        }
    
		case 3:
//line cte.rl:70

        this.significandSign = -1
    
		case 4:
//line cte.rl:74

        this.exponentSign = -1
    
		case 5:
//line cte.rl:78

        this.significand = this.significand * 10 + uint64( this.data[p] - '0')
    
		case 6:
//line cte.rl:82

        this.zeroAdjust--
        this.zeroMagnitude *= 10
    
		case 7:
//line cte.rl:85

        this.significand *= this.zeroMagnitude
        this.exponentAdjust += this.zeroAdjust
        this.zeroMagnitude = 1
        this.zeroAdjust = 0
        this.significand = this.significand * 10 + uint64( this.data[p] - '0')
        this.exponentAdjust--
    
		case 8:
//line cte.rl:94

        this.exponent = this.exponent * 10 + int( this.data[p] - '0')
    
		case 9:
//line cte.rl:103

        if this.significandSign >= 0 {
            callbacks.OnPositiveInt(this.significand)
        } else {
            callbacks.OnNegativeInt(this.significand)
        }
        this.significand = 0
    
		case 10:
//line cte.rl:112

        callbacks.OnDecimalFloat(int64(this.significand) * int64(this.significandSign), (this.exponent+this.exponentAdjust) * this.exponentSign)
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.zeroAdjust = 0
        this.zeroMagnitude = 1
        this.exponentSign = 1
        this.exponent = 0
    
		case 11:
//line cte.rl:123

        err = callbacks.OnListBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 22; goto _again

    
		case 12:
//line cte.rl:131

        err = callbacks.OnUnorderedMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 42; goto _again

    
		case 13:
//line cte.rl:139

        err = callbacks.OnOrderedMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 80; goto _again

    
		case 14:
//line cte.rl:147

        err = callbacks.OnMetadataMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 118; goto _again

    
		case 15:
//line cte.rl:155

        this.arrayStart = p + 1
        err = callbacks.OnCommentBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 16; goto _again

    
		case 16:
//line cte.rl:164

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
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 17; goto _again

    
		case 17:
//line cte.rl:181

        this.arrayStart = p + 1
        err = callbacks.OnStringBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 20; goto _again

    
		case 18:
//line cte.rl:200

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
//line cte.rl:213

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

        
		case 20:
//line cte.rl:233
this.flushAndAddEscapedCharacter(p-1, '\\', callbacks)
		case 21:
//line cte.rl:234
this.flushAndAddEscapedCharacter(p-1, '\n', callbacks)
		case 22:
//line cte.rl:235
this.flushAndAddEscapedCharacter(p-1, '\r', callbacks)
		case 23:
//line cte.rl:236
this.flushAndAddEscapedCharacter(p-1, '\t', callbacks)
		case 24:
//line cte.rl:237
this.flushAndAddEscapedCharacter(p-1, '"', callbacks)
		case 25:
//line cte.rl:238
return false, fmt.Errorf("\\%c: Illegal escape encoding", this.data[p])
		case 26:
//line cte.rl:242

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

        
		case 27:
//line cte.rl:256

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 28:
//line cte.rl:266

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 29:
//line cte.rl:276

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 30:
//line cte.rl:286

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
//line cte.go:891
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
			case 9:
//line cte.rl:103

        if this.significandSign >= 0 {
            callbacks.OnPositiveInt(this.significand)
        } else {
            callbacks.OnNegativeInt(this.significand)
        }
        this.significand = 0
    
			case 10:
//line cte.rl:112

        callbacks.OnDecimalFloat(int64(this.significand) * int64(this.significandSign), (this.exponent+this.exponentAdjust) * this.exponentSign)
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.zeroAdjust = 0
        this.zeroMagnitude = 1
        this.exponentSign = 1
        this.exponent = 0
    
//line cte.go:932
			}
		}
	}

	_out: {}
	}

//line cte.rl:367


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
