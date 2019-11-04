
//line cte.rl:1
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


//line cte.rl:345




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
	1, 32, 1, 33, 1, 34, 1, 35, 
	1, 36, 1, 37, 1, 38, 2, 14, 
	35, 2, 14, 36, 2, 14, 37, 2, 
	14, 38, 2, 15, 35, 2, 15, 36, 
	2, 15, 37, 2, 15, 38, 2, 16, 
	35, 2, 16, 36, 2, 16, 37, 2, 
	16, 38, 
}

var _cte_key_offsets []int16 = []int16{
	0, 0, 19, 22, 24, 28, 30, 34, 
	36, 38, 42, 46, 51, 55, 57, 59, 
	60, 61, 62, 63, 64, 65, 66, 67, 
	68, 69, 70, 72, 74, 76, 78, 83, 
	87, 90, 110, 113, 121, 123, 131, 135, 
	137, 144, 149, 168, 172, 183, 185, 192, 
	194, 201, 205, 215, 219, 224, 228, 230, 
	237, 239, 240, 241, 242, 243, 244, 245, 
	246, 247, 248, 249, 265, 268, 277, 283, 
	285, 304, 307, 315, 317, 325, 329, 331, 
	338, 343, 347, 358, 360, 367, 369, 376, 
	380, 390, 394, 399, 403, 405, 412, 414, 
	415, 416, 417, 418, 419, 420, 421, 422, 
	423, 424, 426, 435, 439, 441, 449, 464, 
	468, 480, 482, 490, 492, 500, 504, 515, 
	519, 524, 528, 530, 538, 540, 541, 542, 
	543, 544, 545, 546, 547, 548, 564, 567, 
	576, 582, 584, 603, 606, 614, 616, 624, 
	628, 630, 637, 642, 646, 657, 659, 666, 
	668, 675, 679, 689, 693, 698, 702, 704, 
	711, 713, 714, 715, 716, 717, 718, 719, 
	720, 721, 722, 723, 725, 734, 738, 740, 
	748, 763, 767, 779, 781, 789, 791, 799, 
	803, 814, 818, 823, 827, 829, 837, 839, 
	840, 841, 842, 843, 844, 845, 846, 847, 
	863, 866, 875, 881, 883, 902, 905, 913, 
	915, 923, 927, 929, 936, 941, 945, 956, 
	958, 965, 967, 974, 978, 988, 992, 997, 
	1001, 1003, 1010, 1012, 1013, 1014, 1015, 1016, 
	1017, 1018, 1019, 1020, 1021, 1022, 1024, 1033, 
	1037, 1039, 1047, 1062, 1066, 1078, 1080, 1088, 
	1090, 1098, 1102, 1113, 1117, 1122, 1126, 1128, 
	1136, 1138, 1139, 1140, 1141, 1142, 1143, 1144, 
	1145, 1146, 1153, 1157, 1164, 1170, 1180, 1186, 
	1192, 1201, 1207, 1207, 1209, 1209, 1209, 1209, 
	1209, 1209, 
}

var _cte_trans_keys []byte = []byte{
	0, 13, 32, 34, 40, 45, 47, 48, 
	60, 91, 102, 110, 116, 117, 123, 9, 
	10, 49, 57, 0, 48, 57, 48, 57, 
	43, 45, 48, 57, 48, 57, 0, 48, 
	49, 57, 48, 49, 48, 55, 48, 57, 
	97, 102, 48, 57, 97, 102, 112, 48, 
	57, 97, 102, 43, 45, 48, 57, 48, 
	57, 42, 47, 97, 108, 115, 101, 105, 
	108, 114, 117, 101, 34, 10, 42, 47, 
	42, 47, 42, 47, 34, 92, 34, 92, 
	110, 114, 116, 32, 33, 35, 126, 34, 
	32, 126, 0, 13, 32, 34, 40, 45, 
	47, 48, 60, 91, 93, 102, 110, 116, 
	117, 123, 9, 10, 49, 57, 0, 48, 
	57, 13, 32, 46, 93, 9, 10, 48, 
	57, 48, 57, 13, 32, 93, 101, 9, 
	10, 48, 57, 43, 45, 48, 57, 48, 
	57, 13, 32, 93, 9, 10, 48, 57, 
	13, 32, 93, 9, 10, 0, 13, 32, 
	34, 40, 45, 47, 48, 60, 91, 102, 
	110, 116, 117, 123, 9, 10, 49, 57, 
	0, 48, 49, 57, 13, 32, 46, 93, 
	98, 111, 120, 9, 10, 48, 57, 48, 
	49, 13, 32, 93, 9, 10, 48, 49, 
	48, 55, 13, 32, 93, 9, 10, 48, 
	55, 48, 57, 97, 102, 13, 32, 46, 
	93, 9, 10, 48, 57, 97, 102, 48, 
	57, 97, 102, 112, 48, 57, 97, 102, 
	43, 45, 48, 57, 48, 57, 13, 32, 
	93, 9, 10, 48, 57, 42, 47, 97, 
	108, 115, 101, 105, 108, 114, 117, 101, 
	34, 0, 13, 32, 34, 40, 45, 47, 
	48, 102, 116, 117, 125, 9, 10, 49, 
	57, 0, 48, 57, 13, 32, 46, 47, 
	61, 9, 10, 48, 57, 13, 32, 47, 
	61, 9, 10, 42, 47, 0, 13, 32, 
	34, 40, 45, 47, 48, 60, 91, 102, 
	110, 116, 117, 123, 9, 10, 49, 57, 
	0, 48, 57, 13, 32, 46, 125, 9, 
	10, 48, 57, 48, 57, 13, 32, 101, 
	125, 9, 10, 48, 57, 43, 45, 48, 
	57, 48, 57, 13, 32, 125, 9, 10, 
	48, 57, 13, 32, 125, 9, 10, 0, 
	48, 49, 57, 13, 32, 46, 98, 111, 
	120, 125, 9, 10, 48, 57, 48, 49, 
	13, 32, 125, 9, 10, 48, 49, 48, 
	55, 13, 32, 125, 9, 10, 48, 55, 
	48, 57, 97, 102, 13, 32, 46, 125, 
	9, 10, 48, 57, 97, 102, 48, 57, 
	97, 102, 112, 48, 57, 97, 102, 43, 
	45, 48, 57, 48, 57, 13, 32, 125, 
	9, 10, 48, 57, 42, 47, 97, 108, 
	115, 101, 105, 108, 114, 117, 101, 34, 
	48, 57, 13, 32, 47, 61, 101, 9, 
	10, 48, 57, 43, 45, 48, 57, 48, 
	57, 13, 32, 47, 61, 9, 10, 48, 
	57, 0, 13, 32, 34, 40, 45, 47, 
	48, 102, 116, 117, 9, 10, 49, 57, 
	0, 48, 49, 57, 13, 32, 46, 47, 
	61, 98, 111, 120, 9, 10, 48, 57, 
	48, 49, 13, 32, 47, 61, 9, 10, 
	48, 49, 48, 55, 13, 32, 47, 61, 
	9, 10, 48, 55, 48, 57, 97, 102, 
	13, 32, 46, 47, 61, 9, 10, 48, 
	57, 97, 102, 48, 57, 97, 102, 112, 
	48, 57, 97, 102, 43, 45, 48, 57, 
	48, 57, 13, 32, 47, 61, 9, 10, 
	48, 57, 42, 47, 97, 108, 115, 101, 
	114, 117, 101, 34, 0, 13, 32, 34, 
	40, 45, 47, 48, 62, 102, 116, 117, 
	9, 10, 49, 57, 0, 48, 57, 13, 
	32, 46, 47, 61, 9, 10, 48, 57, 
	13, 32, 47, 61, 9, 10, 42, 47, 
	0, 13, 32, 34, 40, 45, 47, 48, 
	60, 91, 102, 110, 116, 117, 123, 9, 
	10, 49, 57, 0, 48, 57, 13, 32, 
	46, 62, 9, 10, 48, 57, 48, 57, 
	13, 32, 62, 101, 9, 10, 48, 57, 
	43, 45, 48, 57, 48, 57, 13, 32, 
	62, 9, 10, 48, 57, 13, 32, 62, 
	9, 10, 0, 48, 49, 57, 13, 32, 
	46, 62, 98, 111, 120, 9, 10, 48, 
	57, 48, 49, 13, 32, 62, 9, 10, 
	48, 49, 48, 55, 13, 32, 62, 9, 
	10, 48, 55, 48, 57, 97, 102, 13, 
	32, 46, 62, 9, 10, 48, 57, 97, 
	102, 48, 57, 97, 102, 112, 48, 57, 
	97, 102, 43, 45, 48, 57, 48, 57, 
	13, 32, 62, 9, 10, 48, 57, 42, 
	47, 97, 108, 115, 101, 105, 108, 114, 
	117, 101, 34, 48, 57, 13, 32, 47, 
	61, 101, 9, 10, 48, 57, 43, 45, 
	48, 57, 48, 57, 13, 32, 47, 61, 
	9, 10, 48, 57, 0, 13, 32, 34, 
	40, 45, 47, 48, 102, 116, 117, 9, 
	10, 49, 57, 0, 48, 49, 57, 13, 
	32, 46, 47, 61, 98, 111, 120, 9, 
	10, 48, 57, 48, 49, 13, 32, 47, 
	61, 9, 10, 48, 49, 48, 55, 13, 
	32, 47, 61, 9, 10, 48, 55, 48, 
	57, 97, 102, 13, 32, 46, 47, 61, 
	9, 10, 48, 57, 97, 102, 48, 57, 
	97, 102, 112, 48, 57, 97, 102, 43, 
	45, 48, 57, 48, 57, 13, 32, 47, 
	61, 9, 10, 48, 57, 42, 47, 97, 
	108, 115, 101, 114, 117, 101, 34, 0, 
	13, 32, 34, 40, 41, 45, 47, 48, 
	102, 116, 117, 9, 10, 49, 57, 0, 
	48, 57, 13, 32, 46, 47, 61, 9, 
	10, 48, 57, 13, 32, 47, 61, 9, 
	10, 42, 47, 0, 13, 32, 34, 40, 
	45, 47, 48, 60, 91, 102, 110, 116, 
	117, 123, 9, 10, 49, 57, 0, 48, 
	57, 13, 32, 41, 46, 9, 10, 48, 
	57, 48, 57, 13, 32, 41, 101, 9, 
	10, 48, 57, 43, 45, 48, 57, 48, 
	57, 13, 32, 41, 9, 10, 48, 57, 
	13, 32, 41, 9, 10, 0, 48, 49, 
	57, 13, 32, 41, 46, 98, 111, 120, 
	9, 10, 48, 57, 48, 49, 13, 32, 
	41, 9, 10, 48, 49, 48, 55, 13, 
	32, 41, 9, 10, 48, 55, 48, 57, 
	97, 102, 13, 32, 41, 46, 9, 10, 
	48, 57, 97, 102, 48, 57, 97, 102, 
	112, 48, 57, 97, 102, 43, 45, 48, 
	57, 48, 57, 13, 32, 41, 9, 10, 
	48, 57, 42, 47, 97, 108, 115, 101, 
	105, 108, 114, 117, 101, 34, 48, 57, 
	13, 32, 47, 61, 101, 9, 10, 48, 
	57, 43, 45, 48, 57, 48, 57, 13, 
	32, 47, 61, 9, 10, 48, 57, 0, 
	13, 32, 34, 40, 45, 47, 48, 102, 
	116, 117, 9, 10, 49, 57, 0, 48, 
	49, 57, 13, 32, 46, 47, 61, 98, 
	111, 120, 9, 10, 48, 57, 48, 49, 
	13, 32, 47, 61, 9, 10, 48, 49, 
	48, 55, 13, 32, 47, 61, 9, 10, 
	48, 55, 48, 57, 97, 102, 13, 32, 
	46, 47, 61, 9, 10, 48, 57, 97, 
	102, 48, 57, 97, 102, 112, 48, 57, 
	97, 102, 43, 45, 48, 57, 48, 57, 
	13, 32, 47, 61, 9, 10, 48, 57, 
	42, 47, 97, 108, 115, 101, 114, 117, 
	101, 34, 13, 32, 46, 9, 10, 48, 
	57, 13, 32, 9, 10, 13, 32, 101, 
	9, 10, 48, 57, 13, 32, 9, 10, 
	48, 57, 13, 32, 46, 98, 111, 120, 
	9, 10, 48, 57, 13, 32, 9, 10, 
	48, 49, 13, 32, 9, 10, 48, 55, 
	13, 32, 46, 9, 10, 48, 57, 97, 
	102, 13, 32, 9, 10, 48, 57, 42, 
	47, 
}

var _cte_single_lengths []byte = []byte{
	0, 15, 1, 0, 2, 0, 2, 0, 
	0, 0, 0, 1, 2, 0, 2, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 2, 2, 2, 2, 5, 0, 
	1, 16, 1, 4, 0, 4, 2, 0, 
	3, 3, 15, 2, 7, 0, 3, 0, 
	3, 0, 4, 0, 1, 2, 0, 3, 
	2, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 12, 1, 5, 4, 2, 
	15, 1, 4, 0, 4, 2, 0, 3, 
	3, 2, 7, 0, 3, 0, 3, 0, 
	4, 0, 1, 2, 0, 3, 2, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 0, 5, 2, 0, 4, 11, 2, 
	8, 0, 4, 0, 4, 0, 5, 0, 
	1, 2, 0, 4, 2, 1, 1, 1, 
	1, 1, 1, 1, 1, 12, 1, 5, 
	4, 2, 15, 1, 4, 0, 4, 2, 
	0, 3, 3, 2, 7, 0, 3, 0, 
	3, 0, 4, 0, 1, 2, 0, 3, 
	2, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 0, 5, 2, 0, 4, 
	11, 2, 8, 0, 4, 0, 4, 0, 
	5, 0, 1, 2, 0, 4, 2, 1, 
	1, 1, 1, 1, 1, 1, 1, 12, 
	1, 5, 4, 2, 15, 1, 4, 0, 
	4, 2, 0, 3, 3, 2, 7, 0, 
	3, 0, 3, 0, 4, 0, 1, 2, 
	0, 3, 2, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 0, 5, 2, 
	0, 4, 11, 2, 8, 0, 4, 0, 
	4, 0, 5, 0, 1, 2, 0, 4, 
	2, 1, 1, 1, 1, 1, 1, 1, 
	1, 3, 2, 3, 2, 6, 2, 2, 
	3, 2, 0, 2, 0, 0, 0, 0, 
	0, 0, 
}

var _cte_range_lengths []byte = []byte{
	0, 2, 1, 1, 1, 1, 1, 1, 
	1, 2, 2, 2, 1, 1, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 2, 
	1, 2, 1, 2, 1, 2, 1, 1, 
	2, 1, 2, 1, 2, 1, 2, 1, 
	2, 2, 3, 2, 2, 1, 1, 2, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 2, 1, 2, 1, 0, 
	2, 1, 2, 1, 2, 1, 1, 2, 
	1, 1, 2, 1, 2, 1, 2, 2, 
	3, 2, 2, 1, 1, 2, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 1, 2, 1, 1, 2, 2, 1, 
	2, 1, 2, 1, 2, 2, 3, 2, 
	2, 1, 1, 2, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 2, 1, 2, 
	1, 0, 2, 1, 2, 1, 2, 1, 
	1, 2, 1, 1, 2, 1, 2, 1, 
	2, 2, 3, 2, 2, 1, 1, 2, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 1, 2, 1, 1, 2, 
	2, 1, 2, 1, 2, 1, 2, 2, 
	3, 2, 2, 1, 1, 2, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 2, 
	1, 2, 1, 0, 2, 1, 2, 1, 
	2, 1, 1, 2, 1, 1, 2, 1, 
	2, 1, 2, 2, 3, 2, 2, 1, 
	1, 2, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 1, 2, 1, 
	1, 2, 2, 1, 2, 1, 2, 1, 
	2, 2, 3, 2, 2, 1, 1, 2, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 2, 1, 2, 2, 2, 2, 2, 
	3, 2, 0, 0, 0, 0, 0, 0, 
	0, 0, 
}

var _cte_index_offsets []int16 = []int16{
	0, 0, 18, 21, 23, 27, 29, 33, 
	35, 37, 40, 43, 47, 51, 53, 56, 
	58, 60, 62, 64, 66, 68, 70, 72, 
	74, 76, 78, 81, 84, 87, 90, 96, 
	99, 102, 121, 124, 131, 133, 140, 144, 
	146, 152, 157, 175, 179, 189, 191, 197, 
	199, 205, 208, 216, 219, 223, 227, 229, 
	235, 238, 240, 242, 244, 246, 248, 250, 
	252, 254, 256, 258, 273, 276, 284, 290, 
	293, 311, 314, 321, 323, 330, 334, 336, 
	342, 347, 351, 361, 363, 369, 371, 377, 
	380, 388, 391, 395, 399, 401, 407, 410, 
	412, 414, 416, 418, 420, 422, 424, 426, 
	428, 430, 432, 440, 444, 446, 453, 467, 
	471, 482, 484, 491, 493, 500, 503, 512, 
	515, 519, 523, 525, 532, 535, 537, 539, 
	541, 543, 545, 547, 549, 551, 566, 569, 
	577, 583, 586, 604, 607, 614, 616, 623, 
	627, 629, 635, 640, 644, 654, 656, 662, 
	664, 670, 673, 681, 684, 688, 692, 694, 
	700, 703, 705, 707, 709, 711, 713, 715, 
	717, 719, 721, 723, 725, 733, 737, 739, 
	746, 760, 764, 775, 777, 784, 786, 793, 
	796, 805, 808, 812, 816, 818, 825, 828, 
	830, 832, 834, 836, 838, 840, 842, 844, 
	859, 862, 870, 876, 879, 897, 900, 907, 
	909, 916, 920, 922, 928, 933, 937, 947, 
	949, 955, 957, 963, 966, 974, 977, 981, 
	985, 987, 993, 996, 998, 1000, 1002, 1004, 
	1006, 1008, 1010, 1012, 1014, 1016, 1018, 1026, 
	1030, 1032, 1039, 1053, 1057, 1068, 1070, 1077, 
	1079, 1086, 1089, 1098, 1101, 1105, 1109, 1111, 
	1118, 1121, 1123, 1125, 1127, 1129, 1131, 1133, 
	1135, 1137, 1143, 1147, 1153, 1158, 1167, 1172, 
	1177, 1184, 1189, 1190, 1193, 1194, 1195, 1196, 
	1197, 1198, 
}

var _cte_trans_targs []int16 = []int16{
	2, 1, 1, 266, 1, 6, 14, 269, 
	266, 266, 15, 19, 21, 24, 266, 1, 
	265, 0, 2, 265, 0, 267, 0, 5, 
	5, 268, 0, 268, 0, 2, 269, 265, 
	0, 270, 0, 271, 0, 272, 272, 0, 
	11, 11, 0, 12, 11, 11, 0, 13, 
	13, 273, 0, 273, 0, 1, 1, 0, 
	16, 0, 17, 0, 18, 0, 266, 0, 
	20, 0, 266, 0, 22, 0, 23, 0, 
	266, 0, 266, 0, 274, 25, 27, 28, 
	26, 27, 275, 26, 27, 28, 26, 276, 
	30, 29, 29, 29, 29, 29, 29, 29, 
	32, 32, 0, 277, 32, 0, 34, 33, 
	33, 41, 42, 43, 56, 44, 41, 41, 
	278, 57, 61, 63, 66, 41, 33, 35, 
	0, 34, 35, 0, 33, 33, 36, 278, 
	33, 35, 0, 37, 0, 33, 33, 278, 
	38, 33, 37, 0, 39, 39, 40, 0, 
	40, 0, 33, 33, 278, 33, 40, 0, 
	33, 33, 278, 33, 0, 34, 42, 42, 
	41, 42, 43, 56, 44, 41, 41, 57, 
	61, 63, 66, 41, 42, 35, 0, 34, 
	44, 35, 0, 33, 33, 36, 278, 45, 
	47, 49, 33, 35, 0, 46, 0, 33, 
	33, 278, 33, 46, 0, 48, 0, 33, 
	33, 278, 33, 48, 0, 50, 50, 0, 
	33, 33, 51, 278, 33, 50, 50, 0, 
	52, 52, 0, 53, 52, 52, 0, 54, 
	54, 55, 0, 55, 0, 33, 33, 278, 
	33, 55, 0, 42, 42, 0, 58, 0, 
	59, 0, 60, 0, 41, 0, 62, 0, 
	41, 0, 64, 0, 65, 0, 41, 0, 
	41, 0, 68, 67, 67, 70, 110, 111, 
	124, 112, 125, 129, 132, 279, 67, 69, 
	0, 68, 69, 0, 70, 70, 105, 71, 
	72, 70, 69, 0, 70, 70, 71, 72, 
	70, 0, 70, 70, 0, 73, 72, 72, 
	80, 72, 81, 94, 82, 80, 80, 95, 
	99, 101, 104, 80, 72, 74, 0, 73, 
	74, 0, 67, 67, 75, 279, 67, 74, 
	0, 76, 0, 67, 67, 77, 279, 67, 
	76, 0, 78, 78, 79, 0, 79, 0, 
	67, 67, 279, 67, 79, 0, 67, 67, 
	279, 67, 0, 73, 82, 74, 0, 67, 
	67, 75, 83, 85, 87, 279, 67, 74, 
	0, 84, 0, 67, 67, 279, 67, 84, 
	0, 86, 0, 67, 67, 279, 67, 86, 
	0, 88, 88, 0, 67, 67, 89, 279, 
	67, 88, 88, 0, 90, 90, 0, 91, 
	90, 90, 0, 92, 92, 93, 0, 93, 
	0, 67, 67, 279, 67, 93, 0, 72, 
	72, 0, 96, 0, 97, 0, 98, 0, 
	80, 0, 100, 0, 80, 0, 102, 0, 
	103, 0, 80, 0, 80, 0, 106, 0, 
	70, 70, 71, 72, 107, 70, 106, 0, 
	108, 108, 109, 0, 109, 0, 70, 70, 
	71, 72, 70, 109, 0, 68, 110, 110, 
	70, 110, 111, 124, 112, 125, 129, 132, 
	110, 69, 0, 68, 112, 69, 0, 70, 
	70, 105, 71, 72, 113, 115, 117, 70, 
	69, 0, 114, 0, 70, 70, 71, 72, 
	70, 114, 0, 116, 0, 70, 70, 71, 
	72, 70, 116, 0, 118, 118, 0, 70, 
	70, 119, 71, 72, 70, 118, 118, 0, 
	120, 120, 0, 121, 120, 120, 0, 122, 
	122, 123, 0, 123, 0, 70, 70, 71, 
	72, 70, 123, 0, 110, 110, 0, 126, 
	0, 127, 0, 128, 0, 70, 0, 130, 
	0, 131, 0, 70, 0, 70, 0, 134, 
	133, 133, 136, 176, 177, 190, 178, 280, 
	191, 195, 198, 133, 135, 0, 134, 135, 
	0, 136, 136, 171, 137, 138, 136, 135, 
	0, 136, 136, 137, 138, 136, 0, 136, 
	136, 0, 139, 138, 138, 146, 138, 147, 
	160, 148, 146, 146, 161, 165, 167, 170, 
	146, 138, 140, 0, 139, 140, 0, 133, 
	133, 141, 280, 133, 140, 0, 142, 0, 
	133, 133, 280, 143, 133, 142, 0, 144, 
	144, 145, 0, 145, 0, 133, 133, 280, 
	133, 145, 0, 133, 133, 280, 133, 0, 
	139, 148, 140, 0, 133, 133, 141, 280, 
	149, 151, 153, 133, 140, 0, 150, 0, 
	133, 133, 280, 133, 150, 0, 152, 0, 
	133, 133, 280, 133, 152, 0, 154, 154, 
	0, 133, 133, 155, 280, 133, 154, 154, 
	0, 156, 156, 0, 157, 156, 156, 0, 
	158, 158, 159, 0, 159, 0, 133, 133, 
	280, 133, 159, 0, 138, 138, 0, 162, 
	0, 163, 0, 164, 0, 146, 0, 166, 
	0, 146, 0, 168, 0, 169, 0, 146, 
	0, 146, 0, 172, 0, 136, 136, 137, 
	138, 173, 136, 172, 0, 174, 174, 175, 
	0, 175, 0, 136, 136, 137, 138, 136, 
	175, 0, 134, 176, 176, 136, 176, 177, 
	190, 178, 191, 195, 198, 176, 135, 0, 
	134, 178, 135, 0, 136, 136, 171, 137, 
	138, 179, 181, 183, 136, 135, 0, 180, 
	0, 136, 136, 137, 138, 136, 180, 0, 
	182, 0, 136, 136, 137, 138, 136, 182, 
	0, 184, 184, 0, 136, 136, 185, 137, 
	138, 136, 184, 184, 0, 186, 186, 0, 
	187, 186, 186, 0, 188, 188, 189, 0, 
	189, 0, 136, 136, 137, 138, 136, 189, 
	0, 176, 176, 0, 192, 0, 193, 0, 
	194, 0, 136, 0, 196, 0, 197, 0, 
	136, 0, 136, 0, 200, 199, 199, 202, 
	242, 281, 243, 256, 244, 257, 261, 264, 
	199, 201, 0, 200, 201, 0, 202, 202, 
	237, 203, 204, 202, 201, 0, 202, 202, 
	203, 204, 202, 0, 202, 202, 0, 205, 
	204, 204, 212, 204, 213, 226, 214, 212, 
	212, 227, 231, 233, 236, 212, 204, 206, 
	0, 205, 206, 0, 199, 199, 281, 207, 
	199, 206, 0, 208, 0, 199, 199, 281, 
	209, 199, 208, 0, 210, 210, 211, 0, 
	211, 0, 199, 199, 281, 199, 211, 0, 
	199, 199, 281, 199, 0, 205, 214, 206, 
	0, 199, 199, 281, 207, 215, 217, 219, 
	199, 206, 0, 216, 0, 199, 199, 281, 
	199, 216, 0, 218, 0, 199, 199, 281, 
	199, 218, 0, 220, 220, 0, 199, 199, 
	281, 221, 199, 220, 220, 0, 222, 222, 
	0, 223, 222, 222, 0, 224, 224, 225, 
	0, 225, 0, 199, 199, 281, 199, 225, 
	0, 204, 204, 0, 228, 0, 229, 0, 
	230, 0, 212, 0, 232, 0, 212, 0, 
	234, 0, 235, 0, 212, 0, 212, 0, 
	238, 0, 202, 202, 203, 204, 239, 202, 
	238, 0, 240, 240, 241, 0, 241, 0, 
	202, 202, 203, 204, 202, 241, 0, 200, 
	242, 242, 202, 242, 243, 256, 244, 257, 
	261, 264, 242, 201, 0, 200, 244, 201, 
	0, 202, 202, 237, 203, 204, 245, 247, 
	249, 202, 201, 0, 246, 0, 202, 202, 
	203, 204, 202, 246, 0, 248, 0, 202, 
	202, 203, 204, 202, 248, 0, 250, 250, 
	0, 202, 202, 251, 203, 204, 202, 250, 
	250, 0, 252, 252, 0, 253, 252, 252, 
	0, 254, 254, 255, 0, 255, 0, 202, 
	202, 203, 204, 202, 255, 0, 242, 242, 
	0, 258, 0, 259, 0, 260, 0, 202, 
	0, 262, 0, 263, 0, 202, 0, 202, 
	0, 266, 266, 3, 266, 265, 0, 266, 
	266, 266, 0, 266, 266, 4, 266, 267, 
	0, 266, 266, 266, 268, 0, 266, 266, 
	3, 7, 8, 9, 266, 265, 0, 266, 
	266, 266, 270, 0, 266, 266, 266, 271, 
	0, 266, 266, 10, 266, 272, 272, 0, 
	266, 266, 266, 273, 0, 0, 27, 28, 
	26, 0, 0, 0, 0, 0, 0, 
}

var _cte_trans_actions []byte = []byte{
	0, 0, 0, 47, 41, 7, 0, 11, 
	39, 35, 0, 0, 0, 0, 37, 0, 
	11, 0, 0, 11, 0, 17, 0, 0, 
	9, 23, 0, 23, 0, 0, 11, 11, 
	0, 25, 0, 27, 0, 13, 15, 0, 
	19, 21, 0, 0, 19, 21, 0, 0, 
	9, 23, 0, 23, 0, 45, 43, 0, 
	0, 0, 0, 0, 0, 0, 5, 0, 
	0, 0, 1, 0, 0, 0, 0, 0, 
	3, 0, 49, 0, 51, 0, 0, 0, 
	0, 0, 53, 0, 45, 0, 0, 67, 
	0, 0, 63, 55, 57, 59, 61, 65, 
	0, 0, 0, 69, 0, 0, 0, 0, 
	0, 47, 41, 7, 0, 11, 39, 35, 
	71, 0, 0, 0, 0, 37, 0, 11, 
	0, 0, 11, 0, 29, 29, 0, 79, 
	29, 11, 0, 17, 0, 31, 31, 91, 
	0, 31, 17, 0, 0, 9, 23, 0, 
	23, 0, 31, 31, 91, 31, 23, 0, 
	0, 0, 71, 0, 0, 0, 0, 0, 
	47, 41, 7, 0, 11, 39, 35, 0, 
	0, 0, 0, 37, 0, 11, 0, 0, 
	11, 11, 0, 29, 29, 0, 79, 0, 
	0, 0, 29, 11, 0, 25, 0, 29, 
	29, 79, 29, 25, 0, 27, 0, 29, 
	29, 79, 29, 27, 0, 13, 15, 0, 
	29, 29, 0, 79, 29, 13, 15, 0, 
	19, 21, 0, 0, 19, 21, 0, 0, 
	9, 23, 0, 23, 0, 33, 33, 103, 
	33, 23, 0, 45, 43, 0, 0, 0, 
	0, 0, 0, 0, 5, 0, 0, 0, 
	1, 0, 0, 0, 0, 0, 3, 0, 
	49, 0, 0, 0, 0, 47, 41, 7, 
	0, 11, 0, 0, 0, 73, 0, 11, 
	0, 0, 11, 0, 29, 29, 0, 29, 
	29, 29, 11, 0, 0, 0, 0, 0, 
	0, 0, 45, 43, 0, 0, 0, 0, 
	47, 41, 7, 0, 11, 39, 35, 0, 
	0, 0, 0, 37, 0, 11, 0, 0, 
	11, 0, 29, 29, 0, 82, 29, 11, 
	0, 17, 0, 31, 31, 0, 94, 31, 
	17, 0, 0, 9, 23, 0, 23, 0, 
	31, 31, 94, 31, 23, 0, 0, 0, 
	73, 0, 0, 0, 11, 11, 0, 29, 
	29, 0, 0, 0, 0, 82, 29, 11, 
	0, 25, 0, 29, 29, 82, 29, 25, 
	0, 27, 0, 29, 29, 82, 29, 27, 
	0, 13, 15, 0, 29, 29, 0, 82, 
	29, 13, 15, 0, 19, 21, 0, 0, 
	19, 21, 0, 0, 9, 23, 0, 23, 
	0, 33, 33, 106, 33, 23, 0, 45, 
	43, 0, 0, 0, 0, 0, 0, 0, 
	5, 0, 0, 0, 1, 0, 0, 0, 
	0, 0, 3, 0, 49, 0, 17, 0, 
	31, 31, 31, 31, 0, 31, 17, 0, 
	0, 9, 23, 0, 23, 0, 31, 31, 
	31, 31, 31, 23, 0, 0, 0, 0, 
	47, 41, 7, 0, 11, 0, 0, 0, 
	0, 11, 0, 0, 11, 11, 0, 29, 
	29, 0, 29, 29, 0, 0, 0, 29, 
	11, 0, 25, 0, 29, 29, 29, 29, 
	29, 25, 0, 27, 0, 29, 29, 29, 
	29, 29, 27, 0, 13, 15, 0, 29, 
	29, 0, 29, 29, 29, 13, 15, 0, 
	19, 21, 0, 0, 19, 21, 0, 0, 
	9, 23, 0, 23, 0, 33, 33, 33, 
	33, 33, 23, 0, 45, 43, 0, 0, 
	0, 0, 0, 0, 0, 5, 0, 0, 
	0, 0, 0, 3, 0, 49, 0, 0, 
	0, 0, 47, 41, 7, 0, 11, 75, 
	0, 0, 0, 0, 11, 0, 0, 11, 
	0, 29, 29, 0, 29, 29, 29, 11, 
	0, 0, 0, 0, 0, 0, 0, 45, 
	43, 0, 0, 0, 0, 47, 41, 7, 
	0, 11, 39, 35, 0, 0, 0, 0, 
	37, 0, 11, 0, 0, 11, 0, 29, 
	29, 0, 85, 29, 11, 0, 17, 0, 
	31, 31, 97, 0, 31, 17, 0, 0, 
	9, 23, 0, 23, 0, 31, 31, 97, 
	31, 23, 0, 0, 0, 75, 0, 0, 
	0, 11, 11, 0, 29, 29, 0, 85, 
	0, 0, 0, 29, 11, 0, 25, 0, 
	29, 29, 85, 29, 25, 0, 27, 0, 
	29, 29, 85, 29, 27, 0, 13, 15, 
	0, 29, 29, 0, 85, 29, 13, 15, 
	0, 19, 21, 0, 0, 19, 21, 0, 
	0, 9, 23, 0, 23, 0, 33, 33, 
	109, 33, 23, 0, 45, 43, 0, 0, 
	0, 0, 0, 0, 0, 5, 0, 0, 
	0, 1, 0, 0, 0, 0, 0, 3, 
	0, 49, 0, 17, 0, 31, 31, 31, 
	31, 0, 31, 17, 0, 0, 9, 23, 
	0, 23, 0, 31, 31, 31, 31, 31, 
	23, 0, 0, 0, 0, 47, 41, 7, 
	0, 11, 0, 0, 0, 0, 11, 0, 
	0, 11, 11, 0, 29, 29, 0, 29, 
	29, 0, 0, 0, 29, 11, 0, 25, 
	0, 29, 29, 29, 29, 29, 25, 0, 
	27, 0, 29, 29, 29, 29, 29, 27, 
	0, 13, 15, 0, 29, 29, 0, 29, 
	29, 29, 13, 15, 0, 19, 21, 0, 
	0, 19, 21, 0, 0, 9, 23, 0, 
	23, 0, 33, 33, 33, 33, 33, 23, 
	0, 45, 43, 0, 0, 0, 0, 0, 
	0, 0, 5, 0, 0, 0, 0, 0, 
	3, 0, 49, 0, 0, 0, 0, 47, 
	41, 77, 7, 0, 11, 0, 0, 0, 
	0, 11, 0, 0, 11, 0, 29, 29, 
	0, 29, 29, 29, 11, 0, 0, 0, 
	0, 0, 0, 0, 45, 43, 0, 0, 
	0, 0, 47, 41, 7, 0, 11, 39, 
	35, 0, 0, 0, 0, 37, 0, 11, 
	0, 0, 11, 0, 29, 29, 88, 0, 
	29, 11, 0, 17, 0, 31, 31, 100, 
	0, 31, 17, 0, 0, 9, 23, 0, 
	23, 0, 31, 31, 100, 31, 23, 0, 
	0, 0, 77, 0, 0, 0, 11, 11, 
	0, 29, 29, 88, 0, 0, 0, 0, 
	29, 11, 0, 25, 0, 29, 29, 88, 
	29, 25, 0, 27, 0, 29, 29, 88, 
	29, 27, 0, 13, 15, 0, 29, 29, 
	88, 0, 29, 13, 15, 0, 19, 21, 
	0, 0, 19, 21, 0, 0, 9, 23, 
	0, 23, 0, 33, 33, 112, 33, 23, 
	0, 45, 43, 0, 0, 0, 0, 0, 
	0, 0, 5, 0, 0, 0, 1, 0, 
	0, 0, 0, 0, 3, 0, 49, 0, 
	17, 0, 31, 31, 31, 31, 0, 31, 
	17, 0, 0, 9, 23, 0, 23, 0, 
	31, 31, 31, 31, 31, 23, 0, 0, 
	0, 0, 47, 41, 7, 0, 11, 0, 
	0, 0, 0, 11, 0, 0, 11, 11, 
	0, 29, 29, 0, 29, 29, 0, 0, 
	0, 29, 11, 0, 25, 0, 29, 29, 
	29, 29, 29, 25, 0, 27, 0, 29, 
	29, 29, 29, 29, 27, 0, 13, 15, 
	0, 29, 29, 0, 29, 29, 29, 13, 
	15, 0, 19, 21, 0, 0, 19, 21, 
	0, 0, 9, 23, 0, 23, 0, 33, 
	33, 33, 33, 33, 23, 0, 45, 43, 
	0, 0, 0, 0, 0, 0, 0, 5, 
	0, 0, 0, 0, 0, 3, 0, 49, 
	0, 29, 29, 0, 29, 11, 0, 0, 
	0, 0, 0, 31, 31, 0, 31, 17, 
	0, 31, 31, 31, 23, 0, 29, 29, 
	0, 0, 0, 0, 29, 11, 0, 29, 
	29, 29, 25, 0, 29, 29, 29, 27, 
	0, 29, 29, 0, 29, 13, 15, 0, 
	33, 33, 33, 23, 0, 0, 45, 0, 
	0, 0, 0, 0, 0, 0, 0, 
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
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 29, 0, 31, 31, 29, 29, 29, 
	29, 33, 0, 0, 0, 0, 0, 0, 
	0, 0, 
}

const cte_start int = 1
const cte_error int = 0

const cte_en_comment_iterate int = 25
const cte_en_multiline_comment_iterate int = 26
const cte_en_string_iterate int = 29
const cte_en_uri_iterate int = 31
const cte_en_list_iterate int = 33
const cte_en_unordered_map_iterate int = 67
const cte_en_ordered_map_iterate int = 133
const cte_en_metadata_map_iterate int = 199
const cte_en_main int = 1


//line cte.rl:349

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
    
    
//line cte.go:792
	{
	 this.cs = cte_start
	 this.top = 0
	}

//line cte.go:798
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
//line cte.rl:46

        err = callbacks.OnNil()
        if err != nil {
            p++; goto _out

        }
    
		case 1:
//line cte.rl:53

        err = callbacks.OnBool(true)
        if err != nil {
            p++; goto _out

        }
    
		case 2:
//line cte.rl:60

        err = callbacks.OnBool(false)
        if err != nil {
            p++; goto _out

        }
    
		case 3:
//line cte.rl:69

        this.significandSign = -1
    
		case 4:
//line cte.rl:73

        this.exponentSign = -1
    
		case 5:
//line cte.rl:77

        this.significand = this.significand * 10 + uint64( this.data[p] - '0')
    
		case 6:
//line cte.rl:81

        this.significand = (this.significand << 4) | uint64( this.data[p] - '0')
    
		case 7:
//line cte.rl:83

        this.significand = (this.significand << 4) | uint64( this.data[p] - 'a' + 10)
    
		case 8:
//line cte.rl:87

        this.significand = this.significand * 10 + uint64( this.data[p] - '0')
        this.exponentAdjust--
    
		case 9:
//line cte.rl:92

        this.significand = (this.significand << 4) | uint64( this.data[p] - '0')
        this.exponentAdjust -= 4
    
		case 10:
//line cte.rl:95

        this.significand = (this.significand << 4) | uint64( this.data[p] - 'a' + 10)
        this.exponentAdjust -= 4
    
		case 11:
//line cte.rl:100

        this.exponent = this.exponent * 10 + int( this.data[p] - '0')
    
		case 12:
//line cte.rl:109

        this.significand = (this.significand << 1) | uint64( this.data[p] - '0')
    
		case 13:
//line cte.rl:113

        this.significand = (this.significand << 3) | uint64( this.data[p] - '0')
    
		case 14:
//line cte.rl:119

        if this.significandSign >= 0 {
            callbacks.OnPositiveInt(this.significand)
        } else {
            callbacks.OnNegativeInt(this.significand)
        }
        this.significandSign = 1
        this.significand = 0
    
		case 15:
//line cte.rl:129

        callbacks.OnDecimalFloat(int64(this.significand) * int64(this.significandSign), (this.exponent+this.exponentAdjust) * this.exponentSign)
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.exponentSign = 1
        this.exponent = 0
    
		case 16:
//line cte.rl:138

        exponent := float64((this.exponent * this.exponentSign + this.exponentAdjust))
        callbacks.OnFloat(float64(this.significandSign) * float64(this.significand) * math.Pow(2.0, exponent))
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.exponentSign = 1
        this.exponent = 0
    
		case 17:
//line cte.rl:150

        err = callbacks.OnListBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 33; goto _again

    
		case 18:
//line cte.rl:158

        err = callbacks.OnUnorderedMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 67; goto _again

    
		case 19:
//line cte.rl:166

        err = callbacks.OnOrderedMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 133; goto _again

    
		case 20:
//line cte.rl:174

        err = callbacks.OnMetadataMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 199; goto _again

    
		case 21:
//line cte.rl:182

        this.arrayStart = p + 1
        err = callbacks.OnCommentBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 25; goto _again

    
		case 22:
//line cte.rl:191

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
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 26; goto _again

    
		case 23:
//line cte.rl:208

        this.arrayStart = p + 1
        err = callbacks.OnStringBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 29; goto _again

    
		case 24:
//line cte.rl:217

        this.arrayStart = p + 1
        err = callbacks.OnURIBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 31; goto _again

    
		case 25:
//line cte.rl:236

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

        
		case 26:
//line cte.rl:249

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

        
		case 27:
//line cte.rl:269
this.flushAndAddEscapedCharacter(p-1, '\\', callbacks)
		case 28:
//line cte.rl:270
this.flushAndAddEscapedCharacter(p-1, '\n', callbacks)
		case 29:
//line cte.rl:271
this.flushAndAddEscapedCharacter(p-1, '\r', callbacks)
		case 30:
//line cte.rl:272
this.flushAndAddEscapedCharacter(p-1, '\t', callbacks)
		case 31:
//line cte.rl:273
this.flushAndAddEscapedCharacter(p-1, '"', callbacks)
		case 32:
//line cte.rl:274
return false, fmt.Errorf("\\%c: Illegal escape encoding", this.data[p])
		case 33:
//line cte.rl:278

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

        
		case 34:
//line cte.rl:290

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

        
		case 35:
//line cte.rl:304

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 36:
//line cte.rl:314

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 37:
//line cte.rl:324

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 38:
//line cte.rl:334

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
//line cte.go:1232
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
			case 14:
//line cte.rl:119

        if this.significandSign >= 0 {
            callbacks.OnPositiveInt(this.significand)
        } else {
            callbacks.OnNegativeInt(this.significand)
        }
        this.significandSign = 1
        this.significand = 0
    
			case 15:
//line cte.rl:129

        callbacks.OnDecimalFloat(int64(this.significand) * int64(this.significandSign), (this.exponent+this.exponentAdjust) * this.exponentSign)
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.exponentSign = 1
        this.exponent = 0
    
			case 16:
//line cte.rl:138

        exponent := float64((this.exponent * this.exponentSign + this.exponentAdjust))
        callbacks.OnFloat(float64(this.significandSign) * float64(this.significand) * math.Pow(2.0, exponent))
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.exponentSign = 1
        this.exponent = 0
    
//line cte.go:1283
			}
		}
	}

	_out: {}
	}

//line cte.rl:412


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
