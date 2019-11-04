
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


//line cte.rl:312




//line cte.go:41
var _cte_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 8, 1, 9, 1, 10, 1, 11, 
	1, 12, 1, 13, 1, 14, 1, 15, 
	1, 16, 1, 17, 1, 18, 1, 19, 
	1, 20, 1, 21, 1, 22, 1, 23, 
	1, 24, 1, 25, 1, 26, 1, 27, 
	1, 28, 1, 29, 1, 30, 1, 31, 
	1, 32, 1, 33, 1, 34, 2, 13, 
	31, 2, 13, 32, 2, 13, 33, 2, 
	13, 34, 2, 14, 31, 2, 14, 32, 
	2, 14, 33, 2, 14, 34, 
}

var _cte_key_offsets []int16 = []int16{
	0, 0, 18, 21, 24, 28, 30, 34, 
	36, 38, 42, 44, 45, 46, 47, 48, 
	49, 50, 51, 52, 53, 54, 56, 58, 
	60, 62, 67, 86, 89, 97, 100, 109, 
	113, 115, 122, 127, 145, 149, 160, 162, 
	169, 171, 178, 182, 191, 193, 194, 195, 
	196, 197, 198, 199, 200, 201, 202, 217, 
	220, 229, 235, 237, 255, 258, 266, 269, 
	278, 282, 284, 291, 296, 300, 311, 313, 
	320, 322, 329, 333, 342, 344, 345, 346, 
	347, 348, 349, 350, 351, 352, 353, 356, 
	366, 370, 372, 380, 394, 398, 410, 412, 
	420, 422, 430, 434, 444, 446, 447, 448, 
	449, 450, 451, 452, 453, 468, 471, 480, 
	486, 488, 506, 509, 517, 520, 529, 533, 
	535, 542, 547, 551, 562, 564, 571, 573, 
	580, 584, 593, 595, 596, 597, 598, 599, 
	600, 601, 602, 603, 604, 607, 617, 621, 
	623, 631, 645, 649, 661, 663, 671, 673, 
	681, 685, 695, 697, 698, 699, 700, 701, 
	702, 703, 704, 719, 722, 731, 737, 739, 
	757, 760, 768, 771, 780, 784, 786, 793, 
	798, 802, 813, 815, 822, 824, 831, 835, 
	844, 846, 847, 848, 849, 850, 851, 852, 
	853, 854, 855, 858, 868, 872, 874, 882, 
	896, 900, 912, 914, 922, 924, 932, 936, 
	946, 948, 949, 950, 951, 952, 953, 954, 
	955, 962, 966, 974, 980, 990, 996, 1002, 
	1010, 1010, 1012, 1012, 1012, 1012, 1012, 
}

var _cte_trans_keys []byte = []byte{
	0, 13, 32, 34, 40, 45, 47, 48, 
	60, 91, 102, 110, 116, 123, 9, 10, 
	49, 57, 0, 48, 57, 48, 49, 57, 
	43, 45, 48, 57, 48, 57, 0, 48, 
	49, 57, 48, 49, 48, 55, 48, 57, 
	97, 102, 42, 47, 97, 108, 115, 101, 
	105, 108, 114, 117, 101, 10, 42, 47, 
	42, 47, 42, 47, 34, 92, 34, 92, 
	110, 114, 116, 0, 13, 32, 34, 40, 
	45, 47, 48, 60, 91, 93, 102, 110, 
	116, 123, 9, 10, 49, 57, 0, 48, 
	57, 13, 32, 46, 93, 9, 10, 48, 
	57, 48, 49, 57, 13, 32, 48, 93, 
	101, 9, 10, 49, 57, 43, 45, 48, 
	57, 48, 57, 13, 32, 93, 9, 10, 
	48, 57, 13, 32, 93, 9, 10, 0, 
	13, 32, 34, 40, 45, 47, 48, 60, 
	91, 102, 110, 116, 123, 9, 10, 49, 
	57, 0, 48, 49, 57, 13, 32, 46, 
	93, 98, 111, 120, 9, 10, 48, 57, 
	48, 49, 13, 32, 93, 9, 10, 48, 
	49, 48, 55, 13, 32, 93, 9, 10, 
	48, 55, 48, 57, 97, 102, 13, 32, 
	93, 9, 10, 48, 57, 97, 102, 42, 
	47, 97, 108, 115, 101, 105, 108, 114, 
	117, 101, 0, 13, 32, 34, 40, 45, 
	47, 48, 102, 116, 125, 9, 10, 49, 
	57, 0, 48, 57, 13, 32, 46, 47, 
	61, 9, 10, 48, 57, 13, 32, 47, 
	61, 9, 10, 42, 47, 0, 13, 32, 
	34, 40, 45, 47, 48, 60, 91, 102, 
	110, 116, 123, 9, 10, 49, 57, 0, 
	48, 57, 13, 32, 46, 125, 9, 10, 
	48, 57, 48, 49, 57, 13, 32, 48, 
	101, 125, 9, 10, 49, 57, 43, 45, 
	48, 57, 48, 57, 13, 32, 125, 9, 
	10, 48, 57, 13, 32, 125, 9, 10, 
	0, 48, 49, 57, 13, 32, 46, 98, 
	111, 120, 125, 9, 10, 48, 57, 48, 
	49, 13, 32, 125, 9, 10, 48, 49, 
	48, 55, 13, 32, 125, 9, 10, 48, 
	55, 48, 57, 97, 102, 13, 32, 125, 
	9, 10, 48, 57, 97, 102, 42, 47, 
	97, 108, 115, 101, 105, 108, 114, 117, 
	101, 48, 49, 57, 13, 32, 47, 48, 
	61, 101, 9, 10, 49, 57, 43, 45, 
	48, 57, 48, 57, 13, 32, 47, 61, 
	9, 10, 48, 57, 0, 13, 32, 34, 
	40, 45, 47, 48, 102, 116, 9, 10, 
	49, 57, 0, 48, 49, 57, 13, 32, 
	46, 47, 61, 98, 111, 120, 9, 10, 
	48, 57, 48, 49, 13, 32, 47, 61, 
	9, 10, 48, 49, 48, 55, 13, 32, 
	47, 61, 9, 10, 48, 55, 48, 57, 
	97, 102, 13, 32, 47, 61, 9, 10, 
	48, 57, 97, 102, 42, 47, 97, 108, 
	115, 101, 114, 117, 101, 0, 13, 32, 
	34, 40, 45, 47, 48, 62, 102, 116, 
	9, 10, 49, 57, 0, 48, 57, 13, 
	32, 46, 47, 61, 9, 10, 48, 57, 
	13, 32, 47, 61, 9, 10, 42, 47, 
	0, 13, 32, 34, 40, 45, 47, 48, 
	60, 91, 102, 110, 116, 123, 9, 10, 
	49, 57, 0, 48, 57, 13, 32, 46, 
	62, 9, 10, 48, 57, 48, 49, 57, 
	13, 32, 48, 62, 101, 9, 10, 49, 
	57, 43, 45, 48, 57, 48, 57, 13, 
	32, 62, 9, 10, 48, 57, 13, 32, 
	62, 9, 10, 0, 48, 49, 57, 13, 
	32, 46, 62, 98, 111, 120, 9, 10, 
	48, 57, 48, 49, 13, 32, 62, 9, 
	10, 48, 49, 48, 55, 13, 32, 62, 
	9, 10, 48, 55, 48, 57, 97, 102, 
	13, 32, 62, 9, 10, 48, 57, 97, 
	102, 42, 47, 97, 108, 115, 101, 105, 
	108, 114, 117, 101, 48, 49, 57, 13, 
	32, 47, 48, 61, 101, 9, 10, 49, 
	57, 43, 45, 48, 57, 48, 57, 13, 
	32, 47, 61, 9, 10, 48, 57, 0, 
	13, 32, 34, 40, 45, 47, 48, 102, 
	116, 9, 10, 49, 57, 0, 48, 49, 
	57, 13, 32, 46, 47, 61, 98, 111, 
	120, 9, 10, 48, 57, 48, 49, 13, 
	32, 47, 61, 9, 10, 48, 49, 48, 
	55, 13, 32, 47, 61, 9, 10, 48, 
	55, 48, 57, 97, 102, 13, 32, 47, 
	61, 9, 10, 48, 57, 97, 102, 42, 
	47, 97, 108, 115, 101, 114, 117, 101, 
	0, 13, 32, 34, 40, 41, 45, 47, 
	48, 102, 116, 9, 10, 49, 57, 0, 
	48, 57, 13, 32, 46, 47, 61, 9, 
	10, 48, 57, 13, 32, 47, 61, 9, 
	10, 42, 47, 0, 13, 32, 34, 40, 
	45, 47, 48, 60, 91, 102, 110, 116, 
	123, 9, 10, 49, 57, 0, 48, 57, 
	13, 32, 41, 46, 9, 10, 48, 57, 
	48, 49, 57, 13, 32, 41, 48, 101, 
	9, 10, 49, 57, 43, 45, 48, 57, 
	48, 57, 13, 32, 41, 9, 10, 48, 
	57, 13, 32, 41, 9, 10, 0, 48, 
	49, 57, 13, 32, 41, 46, 98, 111, 
	120, 9, 10, 48, 57, 48, 49, 13, 
	32, 41, 9, 10, 48, 49, 48, 55, 
	13, 32, 41, 9, 10, 48, 55, 48, 
	57, 97, 102, 13, 32, 41, 9, 10, 
	48, 57, 97, 102, 42, 47, 97, 108, 
	115, 101, 105, 108, 114, 117, 101, 48, 
	49, 57, 13, 32, 47, 48, 61, 101, 
	9, 10, 49, 57, 43, 45, 48, 57, 
	48, 57, 13, 32, 47, 61, 9, 10, 
	48, 57, 0, 13, 32, 34, 40, 45, 
	47, 48, 102, 116, 9, 10, 49, 57, 
	0, 48, 49, 57, 13, 32, 46, 47, 
	61, 98, 111, 120, 9, 10, 48, 57, 
	48, 49, 13, 32, 47, 61, 9, 10, 
	48, 49, 48, 55, 13, 32, 47, 61, 
	9, 10, 48, 55, 48, 57, 97, 102, 
	13, 32, 47, 61, 9, 10, 48, 57, 
	97, 102, 42, 47, 97, 108, 115, 101, 
	114, 117, 101, 13, 32, 46, 9, 10, 
	48, 57, 13, 32, 9, 10, 13, 32, 
	48, 101, 9, 10, 49, 57, 13, 32, 
	9, 10, 48, 57, 13, 32, 46, 98, 
	111, 120, 9, 10, 48, 57, 13, 32, 
	9, 10, 48, 49, 13, 32, 9, 10, 
	48, 55, 13, 32, 9, 10, 48, 57, 
	97, 102, 42, 47, 
}

var _cte_single_lengths []byte = []byte{
	0, 14, 1, 1, 2, 0, 2, 0, 
	0, 0, 2, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 2, 2, 2, 
	2, 5, 15, 1, 4, 1, 5, 2, 
	0, 3, 3, 14, 2, 7, 0, 3, 
	0, 3, 0, 3, 2, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 11, 1, 
	5, 4, 2, 14, 1, 4, 1, 5, 
	2, 0, 3, 3, 2, 7, 0, 3, 
	0, 3, 0, 3, 2, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 6, 
	2, 0, 4, 10, 2, 8, 0, 4, 
	0, 4, 0, 4, 2, 1, 1, 1, 
	1, 1, 1, 1, 11, 1, 5, 4, 
	2, 14, 1, 4, 1, 5, 2, 0, 
	3, 3, 2, 7, 0, 3, 0, 3, 
	0, 3, 2, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 6, 2, 0, 
	4, 10, 2, 8, 0, 4, 0, 4, 
	0, 4, 2, 1, 1, 1, 1, 1, 
	1, 1, 11, 1, 5, 4, 2, 14, 
	1, 4, 1, 5, 2, 0, 3, 3, 
	2, 7, 0, 3, 0, 3, 0, 3, 
	2, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 6, 2, 0, 4, 10, 
	2, 8, 0, 4, 0, 4, 0, 4, 
	2, 1, 1, 1, 1, 1, 1, 1, 
	3, 2, 4, 2, 6, 2, 2, 2, 
	0, 2, 0, 0, 0, 0, 0, 
}

var _cte_range_lengths []byte = []byte{
	0, 2, 1, 1, 1, 1, 1, 1, 
	1, 2, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 2, 1, 2, 1, 2, 1, 
	1, 2, 1, 2, 1, 2, 1, 2, 
	1, 2, 2, 3, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 2, 1, 
	2, 1, 0, 2, 1, 2, 1, 2, 
	1, 1, 2, 1, 1, 2, 1, 2, 
	1, 2, 2, 3, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 1, 2, 
	1, 1, 2, 2, 1, 2, 1, 2, 
	1, 2, 2, 3, 0, 0, 0, 0, 
	0, 0, 0, 0, 2, 1, 2, 1, 
	0, 2, 1, 2, 1, 2, 1, 1, 
	2, 1, 1, 2, 1, 2, 1, 2, 
	2, 3, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 1, 2, 1, 1, 
	2, 2, 1, 2, 1, 2, 1, 2, 
	2, 3, 0, 0, 0, 0, 0, 0, 
	0, 0, 2, 1, 2, 1, 0, 2, 
	1, 2, 1, 2, 1, 1, 2, 1, 
	1, 2, 1, 2, 1, 2, 2, 3, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 1, 2, 1, 1, 2, 2, 
	1, 2, 1, 2, 1, 2, 2, 3, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	2, 1, 2, 2, 2, 2, 2, 3, 
	0, 0, 0, 0, 0, 0, 0, 
}

var _cte_index_offsets []int16 = []int16{
	0, 0, 17, 20, 23, 27, 29, 33, 
	35, 37, 40, 43, 45, 47, 49, 51, 
	53, 55, 57, 59, 61, 63, 66, 69, 
	72, 75, 81, 99, 102, 109, 112, 120, 
	124, 126, 132, 137, 154, 158, 168, 170, 
	176, 178, 184, 187, 194, 197, 199, 201, 
	203, 205, 207, 209, 211, 213, 215, 229, 
	232, 240, 246, 249, 266, 269, 276, 279, 
	287, 291, 293, 299, 304, 308, 318, 320, 
	326, 328, 334, 337, 344, 347, 349, 351, 
	353, 355, 357, 359, 361, 363, 365, 368, 
	377, 381, 383, 390, 403, 407, 418, 420, 
	427, 429, 436, 439, 447, 450, 452, 454, 
	456, 458, 460, 462, 464, 478, 481, 489, 
	495, 498, 515, 518, 525, 528, 536, 540, 
	542, 548, 553, 557, 567, 569, 575, 577, 
	583, 586, 593, 596, 598, 600, 602, 604, 
	606, 608, 610, 612, 614, 617, 626, 630, 
	632, 639, 652, 656, 667, 669, 676, 678, 
	685, 688, 696, 699, 701, 703, 705, 707, 
	709, 711, 713, 727, 730, 738, 744, 747, 
	764, 767, 774, 777, 785, 789, 791, 797, 
	802, 806, 816, 818, 824, 826, 832, 835, 
	842, 845, 847, 849, 851, 853, 855, 857, 
	859, 861, 863, 866, 875, 879, 881, 888, 
	901, 905, 916, 918, 925, 927, 934, 937, 
	945, 948, 950, 952, 954, 956, 958, 960, 
	962, 968, 972, 979, 984, 993, 998, 1003, 
	1009, 1010, 1013, 1014, 1015, 1016, 1017, 
}

var _cte_trans_targs []byte = []byte{
	2, 1, 1, 217, 1, 6, 10, 220, 
	217, 217, 11, 15, 17, 217, 1, 216, 
	0, 2, 216, 0, 218, 218, 0, 5, 
	5, 219, 0, 219, 0, 2, 220, 216, 
	0, 221, 0, 222, 0, 223, 223, 0, 
	1, 1, 0, 12, 0, 13, 0, 14, 
	0, 217, 0, 16, 0, 217, 0, 18, 
	0, 19, 0, 217, 0, 224, 20, 22, 
	23, 21, 22, 225, 21, 22, 23, 21, 
	226, 25, 24, 24, 24, 24, 24, 24, 
	24, 27, 26, 26, 34, 35, 36, 44, 
	37, 34, 34, 227, 45, 49, 51, 34, 
	26, 28, 0, 27, 28, 0, 26, 26, 
	29, 227, 26, 28, 0, 30, 30, 0, 
	26, 26, 30, 227, 31, 26, 30, 0, 
	32, 32, 33, 0, 33, 0, 26, 26, 
	227, 26, 33, 0, 26, 26, 227, 26, 
	0, 27, 35, 35, 34, 35, 36, 44, 
	37, 34, 34, 45, 49, 51, 34, 35, 
	28, 0, 27, 37, 28, 0, 26, 26, 
	29, 227, 38, 40, 42, 26, 28, 0, 
	39, 0, 26, 26, 227, 26, 39, 0, 
	41, 0, 26, 26, 227, 26, 41, 0, 
	43, 43, 0, 26, 26, 227, 26, 43, 
	43, 0, 35, 35, 0, 46, 0, 47, 
	0, 48, 0, 34, 0, 50, 0, 34, 
	0, 52, 0, 53, 0, 34, 0, 55, 
	54, 54, 57, 91, 92, 100, 93, 101, 
	105, 228, 54, 56, 0, 55, 56, 0, 
	57, 57, 86, 58, 59, 57, 56, 0, 
	57, 57, 58, 59, 57, 0, 57, 57, 
	0, 60, 59, 59, 67, 59, 68, 76, 
	69, 67, 67, 77, 81, 83, 67, 59, 
	61, 0, 60, 61, 0, 54, 54, 62, 
	228, 54, 61, 0, 63, 63, 0, 54, 
	54, 63, 64, 228, 54, 63, 0, 65, 
	65, 66, 0, 66, 0, 54, 54, 228, 
	54, 66, 0, 54, 54, 228, 54, 0, 
	60, 69, 61, 0, 54, 54, 62, 70, 
	72, 74, 228, 54, 61, 0, 71, 0, 
	54, 54, 228, 54, 71, 0, 73, 0, 
	54, 54, 228, 54, 73, 0, 75, 75, 
	0, 54, 54, 228, 54, 75, 75, 0, 
	59, 59, 0, 78, 0, 79, 0, 80, 
	0, 67, 0, 82, 0, 67, 0, 84, 
	0, 85, 0, 67, 0, 87, 87, 0, 
	57, 57, 58, 87, 59, 88, 57, 87, 
	0, 89, 89, 90, 0, 90, 0, 57, 
	57, 58, 59, 57, 90, 0, 55, 91, 
	91, 57, 91, 92, 100, 93, 101, 105, 
	91, 56, 0, 55, 93, 56, 0, 57, 
	57, 86, 58, 59, 94, 96, 98, 57, 
	56, 0, 95, 0, 57, 57, 58, 59, 
	57, 95, 0, 97, 0, 57, 57, 58, 
	59, 57, 97, 0, 99, 99, 0, 57, 
	57, 58, 59, 57, 99, 99, 0, 91, 
	91, 0, 102, 0, 103, 0, 104, 0, 
	57, 0, 106, 0, 107, 0, 57, 0, 
	109, 108, 108, 111, 145, 146, 154, 147, 
	229, 155, 159, 108, 110, 0, 109, 110, 
	0, 111, 111, 140, 112, 113, 111, 110, 
	0, 111, 111, 112, 113, 111, 0, 111, 
	111, 0, 114, 113, 113, 121, 113, 122, 
	130, 123, 121, 121, 131, 135, 137, 121, 
	113, 115, 0, 114, 115, 0, 108, 108, 
	116, 229, 108, 115, 0, 117, 117, 0, 
	108, 108, 117, 229, 118, 108, 117, 0, 
	119, 119, 120, 0, 120, 0, 108, 108, 
	229, 108, 120, 0, 108, 108, 229, 108, 
	0, 114, 123, 115, 0, 108, 108, 116, 
	229, 124, 126, 128, 108, 115, 0, 125, 
	0, 108, 108, 229, 108, 125, 0, 127, 
	0, 108, 108, 229, 108, 127, 0, 129, 
	129, 0, 108, 108, 229, 108, 129, 129, 
	0, 113, 113, 0, 132, 0, 133, 0, 
	134, 0, 121, 0, 136, 0, 121, 0, 
	138, 0, 139, 0, 121, 0, 141, 141, 
	0, 111, 111, 112, 141, 113, 142, 111, 
	141, 0, 143, 143, 144, 0, 144, 0, 
	111, 111, 112, 113, 111, 144, 0, 109, 
	145, 145, 111, 145, 146, 154, 147, 155, 
	159, 145, 110, 0, 109, 147, 110, 0, 
	111, 111, 140, 112, 113, 148, 150, 152, 
	111, 110, 0, 149, 0, 111, 111, 112, 
	113, 111, 149, 0, 151, 0, 111, 111, 
	112, 113, 111, 151, 0, 153, 153, 0, 
	111, 111, 112, 113, 111, 153, 153, 0, 
	145, 145, 0, 156, 0, 157, 0, 158, 
	0, 111, 0, 160, 0, 161, 0, 111, 
	0, 163, 162, 162, 165, 199, 230, 200, 
	208, 201, 209, 213, 162, 164, 0, 163, 
	164, 0, 165, 165, 194, 166, 167, 165, 
	164, 0, 165, 165, 166, 167, 165, 0, 
	165, 165, 0, 168, 167, 167, 175, 167, 
	176, 184, 177, 175, 175, 185, 189, 191, 
	175, 167, 169, 0, 168, 169, 0, 162, 
	162, 230, 170, 162, 169, 0, 171, 171, 
	0, 162, 162, 230, 171, 172, 162, 171, 
	0, 173, 173, 174, 0, 174, 0, 162, 
	162, 230, 162, 174, 0, 162, 162, 230, 
	162, 0, 168, 177, 169, 0, 162, 162, 
	230, 170, 178, 180, 182, 162, 169, 0, 
	179, 0, 162, 162, 230, 162, 179, 0, 
	181, 0, 162, 162, 230, 162, 181, 0, 
	183, 183, 0, 162, 162, 230, 162, 183, 
	183, 0, 167, 167, 0, 186, 0, 187, 
	0, 188, 0, 175, 0, 190, 0, 175, 
	0, 192, 0, 193, 0, 175, 0, 195, 
	195, 0, 165, 165, 166, 195, 167, 196, 
	165, 195, 0, 197, 197, 198, 0, 198, 
	0, 165, 165, 166, 167, 165, 198, 0, 
	163, 199, 199, 165, 199, 200, 208, 201, 
	209, 213, 199, 164, 0, 163, 201, 164, 
	0, 165, 165, 194, 166, 167, 202, 204, 
	206, 165, 164, 0, 203, 0, 165, 165, 
	166, 167, 165, 203, 0, 205, 0, 165, 
	165, 166, 167, 165, 205, 0, 207, 207, 
	0, 165, 165, 166, 167, 165, 207, 207, 
	0, 199, 199, 0, 210, 0, 211, 0, 
	212, 0, 165, 0, 214, 0, 215, 0, 
	165, 0, 217, 217, 3, 217, 216, 0, 
	217, 217, 217, 0, 217, 217, 218, 4, 
	217, 218, 0, 217, 217, 217, 219, 0, 
	217, 217, 3, 7, 8, 9, 217, 216, 
	0, 217, 217, 217, 221, 0, 217, 217, 
	217, 222, 0, 217, 217, 217, 223, 223, 
	0, 0, 22, 23, 21, 0, 0, 0, 
	0, 0, 
}

var _cte_trans_actions []byte = []byte{
	0, 0, 0, 43, 37, 7, 0, 11, 
	35, 31, 0, 0, 0, 33, 0, 11, 
	0, 0, 11, 0, 13, 15, 0, 0, 
	9, 17, 0, 17, 0, 0, 11, 11, 
	0, 19, 0, 21, 0, 23, 25, 0, 
	41, 39, 0, 0, 0, 0, 0, 0, 
	0, 5, 0, 0, 0, 1, 0, 0, 
	0, 0, 0, 3, 0, 45, 0, 0, 
	0, 0, 0, 47, 0, 41, 0, 0, 
	61, 0, 0, 57, 49, 51, 53, 55, 
	59, 0, 0, 0, 43, 37, 7, 0, 
	11, 35, 31, 63, 0, 0, 0, 33, 
	0, 11, 0, 0, 11, 0, 27, 27, 
	0, 71, 27, 11, 0, 13, 15, 0, 
	29, 29, 13, 83, 0, 29, 15, 0, 
	0, 9, 17, 0, 17, 0, 29, 29, 
	83, 29, 17, 0, 0, 0, 63, 0, 
	0, 0, 0, 0, 43, 37, 7, 0, 
	11, 35, 31, 0, 0, 0, 33, 0, 
	11, 0, 0, 11, 11, 0, 27, 27, 
	0, 71, 0, 0, 0, 27, 11, 0, 
	19, 0, 27, 27, 71, 27, 19, 0, 
	21, 0, 27, 27, 71, 27, 21, 0, 
	23, 25, 0, 27, 27, 71, 27, 23, 
	25, 0, 41, 39, 0, 0, 0, 0, 
	0, 0, 0, 5, 0, 0, 0, 1, 
	0, 0, 0, 0, 0, 3, 0, 0, 
	0, 0, 43, 37, 7, 0, 11, 0, 
	0, 65, 0, 11, 0, 0, 11, 0, 
	27, 27, 0, 27, 27, 27, 11, 0, 
	0, 0, 0, 0, 0, 0, 41, 39, 
	0, 0, 0, 0, 43, 37, 7, 0, 
	11, 35, 31, 0, 0, 0, 33, 0, 
	11, 0, 0, 11, 0, 27, 27, 0, 
	74, 27, 11, 0, 13, 15, 0, 29, 
	29, 13, 0, 86, 29, 15, 0, 0, 
	9, 17, 0, 17, 0, 29, 29, 86, 
	29, 17, 0, 0, 0, 65, 0, 0, 
	0, 11, 11, 0, 27, 27, 0, 0, 
	0, 0, 74, 27, 11, 0, 19, 0, 
	27, 27, 74, 27, 19, 0, 21, 0, 
	27, 27, 74, 27, 21, 0, 23, 25, 
	0, 27, 27, 74, 27, 23, 25, 0, 
	41, 39, 0, 0, 0, 0, 0, 0, 
	0, 5, 0, 0, 0, 1, 0, 0, 
	0, 0, 0, 3, 0, 13, 15, 0, 
	29, 29, 29, 13, 29, 0, 29, 15, 
	0, 0, 9, 17, 0, 17, 0, 29, 
	29, 29, 29, 29, 17, 0, 0, 0, 
	0, 43, 37, 7, 0, 11, 0, 0, 
	0, 11, 0, 0, 11, 11, 0, 27, 
	27, 0, 27, 27, 0, 0, 0, 27, 
	11, 0, 19, 0, 27, 27, 27, 27, 
	27, 19, 0, 21, 0, 27, 27, 27, 
	27, 27, 21, 0, 23, 25, 0, 27, 
	27, 27, 27, 27, 23, 25, 0, 41, 
	39, 0, 0, 0, 0, 0, 0, 0, 
	5, 0, 0, 0, 0, 0, 3, 0, 
	0, 0, 0, 43, 37, 7, 0, 11, 
	67, 0, 0, 0, 11, 0, 0, 11, 
	0, 27, 27, 0, 27, 27, 27, 11, 
	0, 0, 0, 0, 0, 0, 0, 41, 
	39, 0, 0, 0, 0, 43, 37, 7, 
	0, 11, 35, 31, 0, 0, 0, 33, 
	0, 11, 0, 0, 11, 0, 27, 27, 
	0, 77, 27, 11, 0, 13, 15, 0, 
	29, 29, 13, 89, 0, 29, 15, 0, 
	0, 9, 17, 0, 17, 0, 29, 29, 
	89, 29, 17, 0, 0, 0, 67, 0, 
	0, 0, 11, 11, 0, 27, 27, 0, 
	77, 0, 0, 0, 27, 11, 0, 19, 
	0, 27, 27, 77, 27, 19, 0, 21, 
	0, 27, 27, 77, 27, 21, 0, 23, 
	25, 0, 27, 27, 77, 27, 23, 25, 
	0, 41, 39, 0, 0, 0, 0, 0, 
	0, 0, 5, 0, 0, 0, 1, 0, 
	0, 0, 0, 0, 3, 0, 13, 15, 
	0, 29, 29, 29, 13, 29, 0, 29, 
	15, 0, 0, 9, 17, 0, 17, 0, 
	29, 29, 29, 29, 29, 17, 0, 0, 
	0, 0, 43, 37, 7, 0, 11, 0, 
	0, 0, 11, 0, 0, 11, 11, 0, 
	27, 27, 0, 27, 27, 0, 0, 0, 
	27, 11, 0, 19, 0, 27, 27, 27, 
	27, 27, 19, 0, 21, 0, 27, 27, 
	27, 27, 27, 21, 0, 23, 25, 0, 
	27, 27, 27, 27, 27, 23, 25, 0, 
	41, 39, 0, 0, 0, 0, 0, 0, 
	0, 5, 0, 0, 0, 0, 0, 3, 
	0, 0, 0, 0, 43, 37, 69, 7, 
	0, 11, 0, 0, 0, 11, 0, 0, 
	11, 0, 27, 27, 0, 27, 27, 27, 
	11, 0, 0, 0, 0, 0, 0, 0, 
	41, 39, 0, 0, 0, 0, 43, 37, 
	7, 0, 11, 35, 31, 0, 0, 0, 
	33, 0, 11, 0, 0, 11, 0, 27, 
	27, 80, 0, 27, 11, 0, 13, 15, 
	0, 29, 29, 92, 13, 0, 29, 15, 
	0, 0, 9, 17, 0, 17, 0, 29, 
	29, 92, 29, 17, 0, 0, 0, 69, 
	0, 0, 0, 11, 11, 0, 27, 27, 
	80, 0, 0, 0, 0, 27, 11, 0, 
	19, 0, 27, 27, 80, 27, 19, 0, 
	21, 0, 27, 27, 80, 27, 21, 0, 
	23, 25, 0, 27, 27, 80, 27, 23, 
	25, 0, 41, 39, 0, 0, 0, 0, 
	0, 0, 0, 5, 0, 0, 0, 1, 
	0, 0, 0, 0, 0, 3, 0, 13, 
	15, 0, 29, 29, 29, 13, 29, 0, 
	29, 15, 0, 0, 9, 17, 0, 17, 
	0, 29, 29, 29, 29, 29, 17, 0, 
	0, 0, 0, 43, 37, 7, 0, 11, 
	0, 0, 0, 11, 0, 0, 11, 11, 
	0, 27, 27, 0, 27, 27, 0, 0, 
	0, 27, 11, 0, 19, 0, 27, 27, 
	27, 27, 27, 19, 0, 21, 0, 27, 
	27, 27, 27, 27, 21, 0, 23, 25, 
	0, 27, 27, 27, 27, 27, 23, 25, 
	0, 41, 39, 0, 0, 0, 0, 0, 
	0, 0, 5, 0, 0, 0, 0, 0, 
	3, 0, 27, 27, 0, 27, 11, 0, 
	0, 0, 0, 0, 29, 29, 13, 0, 
	29, 15, 0, 29, 29, 29, 17, 0, 
	27, 27, 0, 0, 0, 0, 27, 11, 
	0, 27, 27, 27, 19, 0, 27, 27, 
	27, 21, 0, 27, 27, 27, 23, 25, 
	0, 0, 41, 0, 0, 0, 0, 0, 
	0, 0, 
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
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	27, 0, 29, 29, 27, 27, 27, 27, 
	0, 0, 0, 0, 0, 0, 0, 
}

const cte_start int = 1
const cte_error int = 0

const cte_en_comment_iterate int = 20
const cte_en_multiline_comment_iterate int = 21
const cte_en_string_iterate int = 24
const cte_en_list_iterate int = 26
const cte_en_unordered_map_iterate int = 54
const cte_en_ordered_map_iterate int = 108
const cte_en_metadata_map_iterate int = 162
const cte_en_main int = 1


//line cte.rl:316

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
    
    
//line cte.go:687
	{
	 this.cs = cte_start
	 this.top = 0
	}

//line cte.go:693
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

        this.significand = (this.significand << 1) | uint64( this.data[p] - '0')
    
		case 10:
//line cte.rl:107

        this.significand = (this.significand << 3) | uint64( this.data[p] - '0')
    
		case 11:
//line cte.rl:111

        this.significand = (this.significand << 4) | uint64( this.data[p] - '0')
    
		case 12:
//line cte.rl:113

        this.significand = (this.significand << 4) | uint64( this.data[p] - 'a' + 10)
    
		case 13:
//line cte.rl:117

        if this.significandSign >= 0 {
            callbacks.OnPositiveInt(this.significand)
        } else {
            callbacks.OnNegativeInt(this.significand)
        }
        this.significandSign = 1
        this.significand = 0
    
		case 14:
//line cte.rl:127

        callbacks.OnDecimalFloat(int64(this.significand) * int64(this.significandSign), (this.exponent+this.exponentAdjust) * this.exponentSign)
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.zeroAdjust = 0
        this.zeroMagnitude = 1
        this.exponentSign = 1
        this.exponent = 0
    
		case 15:
//line cte.rl:138

        err = callbacks.OnListBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 26; goto _again

    
		case 16:
//line cte.rl:146

        err = callbacks.OnUnorderedMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 54; goto _again

    
		case 17:
//line cte.rl:154

        err = callbacks.OnOrderedMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 108; goto _again

    
		case 18:
//line cte.rl:162

        err = callbacks.OnMetadataMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 162; goto _again

    
		case 19:
//line cte.rl:170

        this.arrayStart = p + 1
        err = callbacks.OnCommentBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 20; goto _again

    
		case 20:
//line cte.rl:179

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
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 21; goto _again

    
		case 21:
//line cte.rl:196

        this.arrayStart = p + 1
        err = callbacks.OnStringBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 24; goto _again

    
		case 22:
//line cte.rl:215

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

        
		case 23:
//line cte.rl:228

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

        
		case 24:
//line cte.rl:248
this.flushAndAddEscapedCharacter(p-1, '\\', callbacks)
		case 25:
//line cte.rl:249
this.flushAndAddEscapedCharacter(p-1, '\n', callbacks)
		case 26:
//line cte.rl:250
this.flushAndAddEscapedCharacter(p-1, '\r', callbacks)
		case 27:
//line cte.rl:251
this.flushAndAddEscapedCharacter(p-1, '\t', callbacks)
		case 28:
//line cte.rl:252
this.flushAndAddEscapedCharacter(p-1, '"', callbacks)
		case 29:
//line cte.rl:253
return false, fmt.Errorf("\\%c: Illegal escape encoding", this.data[p])
		case 30:
//line cte.rl:257

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

        
		case 31:
//line cte.rl:271

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 32:
//line cte.rl:281

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 33:
//line cte.rl:291

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 34:
//line cte.rl:301

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
//line cte.go:1087
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
			case 13:
//line cte.rl:117

        if this.significandSign >= 0 {
            callbacks.OnPositiveInt(this.significand)
        } else {
            callbacks.OnNegativeInt(this.significand)
        }
        this.significandSign = 1
        this.significand = 0
    
			case 14:
//line cte.rl:127

        callbacks.OnDecimalFloat(int64(this.significand) * int64(this.significandSign), (this.exponent+this.exponentAdjust) * this.exponentSign)
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.zeroAdjust = 0
        this.zeroMagnitude = 1
        this.exponentSign = 1
        this.exponent = 0
    
//line cte.go:1129
			}
		}
	}

	_out: {}
	}

//line cte.rl:382


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
