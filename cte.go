
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
//    OnURIBegin() error
    OnCommentBegin() error
    OnArrayData(bytes []byte) error
    OnArrayEnd() error
}


//line cte.rl:325




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
	1, 36, 2, 14, 33, 2, 14, 34, 
	2, 14, 35, 2, 14, 36, 2, 15, 
	33, 2, 15, 34, 2, 15, 35, 2, 
	15, 36, 2, 16, 33, 2, 16, 34, 
	2, 16, 35, 2, 16, 36, 
}

var _cte_key_offsets []int16 = []int16{
	0, 0, 18, 21, 23, 27, 29, 33, 
	35, 37, 41, 45, 50, 54, 56, 58, 
	59, 60, 61, 62, 63, 64, 65, 66, 
	67, 68, 70, 72, 74, 76, 81, 100, 
	103, 111, 113, 121, 125, 127, 134, 139, 
	157, 161, 172, 174, 181, 183, 190, 194, 
	204, 208, 213, 217, 219, 226, 228, 229, 
	230, 231, 232, 233, 234, 235, 236, 237, 
	252, 255, 264, 270, 272, 290, 293, 301, 
	303, 311, 315, 317, 324, 329, 333, 344, 
	346, 353, 355, 362, 366, 376, 380, 385, 
	389, 391, 398, 400, 401, 402, 403, 404, 
	405, 406, 407, 408, 409, 411, 420, 424, 
	426, 434, 448, 452, 464, 466, 474, 476, 
	484, 488, 499, 503, 508, 512, 514, 522, 
	524, 525, 526, 527, 528, 529, 530, 531, 
	546, 549, 558, 564, 566, 584, 587, 595, 
	597, 605, 609, 611, 618, 623, 627, 638, 
	640, 647, 649, 656, 660, 670, 674, 679, 
	683, 685, 692, 694, 695, 696, 697, 698, 
	699, 700, 701, 702, 703, 705, 714, 718, 
	720, 728, 742, 746, 758, 760, 768, 770, 
	778, 782, 793, 797, 802, 806, 808, 816, 
	818, 819, 820, 821, 822, 823, 824, 825, 
	840, 843, 852, 858, 860, 878, 881, 889, 
	891, 899, 903, 905, 912, 917, 921, 932, 
	934, 941, 943, 950, 954, 964, 968, 973, 
	977, 979, 986, 988, 989, 990, 991, 992, 
	993, 994, 995, 996, 997, 999, 1008, 1012, 
	1014, 1022, 1036, 1040, 1052, 1054, 1062, 1064, 
	1072, 1076, 1087, 1091, 1096, 1100, 1102, 1110, 
	1112, 1113, 1114, 1115, 1116, 1117, 1118, 1119, 
	1126, 1130, 1137, 1143, 1153, 1159, 1165, 1174, 
	1180, 1180, 1182, 1182, 1182, 1182, 1182, 
}

var _cte_trans_keys []byte = []byte{
	0, 13, 32, 34, 40, 45, 47, 48, 
	60, 91, 102, 110, 116, 123, 9, 10, 
	49, 57, 0, 48, 57, 48, 57, 43, 
	45, 48, 57, 48, 57, 0, 48, 49, 
	57, 48, 49, 48, 55, 48, 57, 97, 
	102, 48, 57, 97, 102, 112, 48, 57, 
	97, 102, 43, 45, 48, 57, 48, 57, 
	42, 47, 97, 108, 115, 101, 105, 108, 
	114, 117, 101, 10, 42, 47, 42, 47, 
	42, 47, 34, 92, 34, 92, 110, 114, 
	116, 0, 13, 32, 34, 40, 45, 47, 
	48, 60, 91, 93, 102, 110, 116, 123, 
	9, 10, 49, 57, 0, 48, 57, 13, 
	32, 46, 93, 9, 10, 48, 57, 48, 
	57, 13, 32, 93, 101, 9, 10, 48, 
	57, 43, 45, 48, 57, 48, 57, 13, 
	32, 93, 9, 10, 48, 57, 13, 32, 
	93, 9, 10, 0, 13, 32, 34, 40, 
	45, 47, 48, 60, 91, 102, 110, 116, 
	123, 9, 10, 49, 57, 0, 48, 49, 
	57, 13, 32, 46, 93, 98, 111, 120, 
	9, 10, 48, 57, 48, 49, 13, 32, 
	93, 9, 10, 48, 49, 48, 55, 13, 
	32, 93, 9, 10, 48, 55, 48, 57, 
	97, 102, 13, 32, 46, 93, 9, 10, 
	48, 57, 97, 102, 48, 57, 97, 102, 
	112, 48, 57, 97, 102, 43, 45, 48, 
	57, 48, 57, 13, 32, 93, 9, 10, 
	48, 57, 42, 47, 97, 108, 115, 101, 
	105, 108, 114, 117, 101, 0, 13, 32, 
	34, 40, 45, 47, 48, 102, 116, 125, 
	9, 10, 49, 57, 0, 48, 57, 13, 
	32, 46, 47, 61, 9, 10, 48, 57, 
	13, 32, 47, 61, 9, 10, 42, 47, 
	0, 13, 32, 34, 40, 45, 47, 48, 
	60, 91, 102, 110, 116, 123, 9, 10, 
	49, 57, 0, 48, 57, 13, 32, 46, 
	125, 9, 10, 48, 57, 48, 57, 13, 
	32, 101, 125, 9, 10, 48, 57, 43, 
	45, 48, 57, 48, 57, 13, 32, 125, 
	9, 10, 48, 57, 13, 32, 125, 9, 
	10, 0, 48, 49, 57, 13, 32, 46, 
	98, 111, 120, 125, 9, 10, 48, 57, 
	48, 49, 13, 32, 125, 9, 10, 48, 
	49, 48, 55, 13, 32, 125, 9, 10, 
	48, 55, 48, 57, 97, 102, 13, 32, 
	46, 125, 9, 10, 48, 57, 97, 102, 
	48, 57, 97, 102, 112, 48, 57, 97, 
	102, 43, 45, 48, 57, 48, 57, 13, 
	32, 125, 9, 10, 48, 57, 42, 47, 
	97, 108, 115, 101, 105, 108, 114, 117, 
	101, 48, 57, 13, 32, 47, 61, 101, 
	9, 10, 48, 57, 43, 45, 48, 57, 
	48, 57, 13, 32, 47, 61, 9, 10, 
	48, 57, 0, 13, 32, 34, 40, 45, 
	47, 48, 102, 116, 9, 10, 49, 57, 
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
	114, 117, 101, 0, 13, 32, 34, 40, 
	45, 47, 48, 62, 102, 116, 9, 10, 
	49, 57, 0, 48, 57, 13, 32, 46, 
	47, 61, 9, 10, 48, 57, 13, 32, 
	47, 61, 9, 10, 42, 47, 0, 13, 
	32, 34, 40, 45, 47, 48, 60, 91, 
	102, 110, 116, 123, 9, 10, 49, 57, 
	0, 48, 57, 13, 32, 46, 62, 9, 
	10, 48, 57, 48, 57, 13, 32, 62, 
	101, 9, 10, 48, 57, 43, 45, 48, 
	57, 48, 57, 13, 32, 62, 9, 10, 
	48, 57, 13, 32, 62, 9, 10, 0, 
	48, 49, 57, 13, 32, 46, 62, 98, 
	111, 120, 9, 10, 48, 57, 48, 49, 
	13, 32, 62, 9, 10, 48, 49, 48, 
	55, 13, 32, 62, 9, 10, 48, 55, 
	48, 57, 97, 102, 13, 32, 46, 62, 
	9, 10, 48, 57, 97, 102, 48, 57, 
	97, 102, 112, 48, 57, 97, 102, 43, 
	45, 48, 57, 48, 57, 13, 32, 62, 
	9, 10, 48, 57, 42, 47, 97, 108, 
	115, 101, 105, 108, 114, 117, 101, 48, 
	57, 13, 32, 47, 61, 101, 9, 10, 
	48, 57, 43, 45, 48, 57, 48, 57, 
	13, 32, 47, 61, 9, 10, 48, 57, 
	0, 13, 32, 34, 40, 45, 47, 48, 
	102, 116, 9, 10, 49, 57, 0, 48, 
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
	101, 0, 13, 32, 34, 40, 41, 45, 
	47, 48, 102, 116, 9, 10, 49, 57, 
	0, 48, 57, 13, 32, 46, 47, 61, 
	9, 10, 48, 57, 13, 32, 47, 61, 
	9, 10, 42, 47, 0, 13, 32, 34, 
	40, 45, 47, 48, 60, 91, 102, 110, 
	116, 123, 9, 10, 49, 57, 0, 48, 
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
	105, 108, 114, 117, 101, 48, 57, 13, 
	32, 47, 61, 101, 9, 10, 48, 57, 
	43, 45, 48, 57, 48, 57, 13, 32, 
	47, 61, 9, 10, 48, 57, 0, 13, 
	32, 34, 40, 45, 47, 48, 102, 116, 
	9, 10, 49, 57, 0, 48, 49, 57, 
	13, 32, 46, 47, 61, 98, 111, 120, 
	9, 10, 48, 57, 48, 49, 13, 32, 
	47, 61, 9, 10, 48, 49, 48, 55, 
	13, 32, 47, 61, 9, 10, 48, 55, 
	48, 57, 97, 102, 13, 32, 46, 47, 
	61, 9, 10, 48, 57, 97, 102, 48, 
	57, 97, 102, 112, 48, 57, 97, 102, 
	43, 45, 48, 57, 48, 57, 13, 32, 
	47, 61, 9, 10, 48, 57, 42, 47, 
	97, 108, 115, 101, 114, 117, 101, 13, 
	32, 46, 9, 10, 48, 57, 13, 32, 
	9, 10, 13, 32, 101, 9, 10, 48, 
	57, 13, 32, 9, 10, 48, 57, 13, 
	32, 46, 98, 111, 120, 9, 10, 48, 
	57, 13, 32, 9, 10, 48, 49, 13, 
	32, 9, 10, 48, 55, 13, 32, 46, 
	9, 10, 48, 57, 97, 102, 13, 32, 
	9, 10, 48, 57, 42, 47, 
}

var _cte_single_lengths []byte = []byte{
	0, 14, 1, 0, 2, 0, 2, 0, 
	0, 0, 0, 1, 2, 0, 2, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 2, 2, 2, 2, 5, 15, 1, 
	4, 0, 4, 2, 0, 3, 3, 14, 
	2, 7, 0, 3, 0, 3, 0, 4, 
	0, 1, 2, 0, 3, 2, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 11, 
	1, 5, 4, 2, 14, 1, 4, 0, 
	4, 2, 0, 3, 3, 2, 7, 0, 
	3, 0, 3, 0, 4, 0, 1, 2, 
	0, 3, 2, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 0, 5, 2, 0, 
	4, 10, 2, 8, 0, 4, 0, 4, 
	0, 5, 0, 1, 2, 0, 4, 2, 
	1, 1, 1, 1, 1, 1, 1, 11, 
	1, 5, 4, 2, 14, 1, 4, 0, 
	4, 2, 0, 3, 3, 2, 7, 0, 
	3, 0, 3, 0, 4, 0, 1, 2, 
	0, 3, 2, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 0, 5, 2, 0, 
	4, 10, 2, 8, 0, 4, 0, 4, 
	0, 5, 0, 1, 2, 0, 4, 2, 
	1, 1, 1, 1, 1, 1, 1, 11, 
	1, 5, 4, 2, 14, 1, 4, 0, 
	4, 2, 0, 3, 3, 2, 7, 0, 
	3, 0, 3, 0, 4, 0, 1, 2, 
	0, 3, 2, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 0, 5, 2, 0, 
	4, 10, 2, 8, 0, 4, 0, 4, 
	0, 5, 0, 1, 2, 0, 4, 2, 
	1, 1, 1, 1, 1, 1, 1, 3, 
	2, 3, 2, 6, 2, 2, 3, 2, 
	0, 2, 0, 0, 0, 0, 0, 
}

var _cte_range_lengths []byte = []byte{
	0, 2, 1, 1, 1, 1, 1, 1, 
	1, 2, 2, 2, 1, 1, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 2, 1, 
	2, 1, 2, 1, 1, 2, 1, 2, 
	1, 2, 1, 2, 1, 2, 2, 3, 
	2, 2, 1, 1, 2, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 2, 
	1, 2, 1, 0, 2, 1, 2, 1, 
	2, 1, 1, 2, 1, 1, 2, 1, 
	2, 1, 2, 2, 3, 2, 2, 1, 
	1, 2, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 1, 2, 1, 1, 
	2, 2, 1, 2, 1, 2, 1, 2, 
	2, 3, 2, 2, 1, 1, 2, 0, 
	0, 0, 0, 0, 0, 0, 0, 2, 
	1, 2, 1, 0, 2, 1, 2, 1, 
	2, 1, 1, 2, 1, 1, 2, 1, 
	2, 1, 2, 2, 3, 2, 2, 1, 
	1, 2, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 1, 2, 1, 1, 
	2, 2, 1, 2, 1, 2, 1, 2, 
	2, 3, 2, 2, 1, 1, 2, 0, 
	0, 0, 0, 0, 0, 0, 0, 2, 
	1, 2, 1, 0, 2, 1, 2, 1, 
	2, 1, 1, 2, 1, 1, 2, 1, 
	2, 1, 2, 2, 3, 2, 2, 1, 
	1, 2, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 1, 2, 1, 1, 
	2, 2, 1, 2, 1, 2, 1, 2, 
	2, 3, 2, 2, 1, 1, 2, 0, 
	0, 0, 0, 0, 0, 0, 0, 2, 
	1, 2, 2, 2, 2, 2, 3, 2, 
	0, 0, 0, 0, 0, 0, 0, 
}

var _cte_index_offsets []int16 = []int16{
	0, 0, 17, 20, 22, 26, 28, 32, 
	34, 36, 39, 42, 46, 50, 52, 55, 
	57, 59, 61, 63, 65, 67, 69, 71, 
	73, 75, 78, 81, 84, 87, 93, 111, 
	114, 121, 123, 130, 134, 136, 142, 147, 
	164, 168, 178, 180, 186, 188, 194, 197, 
	205, 208, 212, 216, 218, 224, 227, 229, 
	231, 233, 235, 237, 239, 241, 243, 245, 
	259, 262, 270, 276, 279, 296, 299, 306, 
	308, 315, 319, 321, 327, 332, 336, 346, 
	348, 354, 356, 362, 365, 373, 376, 380, 
	384, 386, 392, 395, 397, 399, 401, 403, 
	405, 407, 409, 411, 413, 415, 423, 427, 
	429, 436, 449, 453, 464, 466, 473, 475, 
	482, 485, 494, 497, 501, 505, 507, 514, 
	517, 519, 521, 523, 525, 527, 529, 531, 
	545, 548, 556, 562, 565, 582, 585, 592, 
	594, 601, 605, 607, 613, 618, 622, 632, 
	634, 640, 642, 648, 651, 659, 662, 666, 
	670, 672, 678, 681, 683, 685, 687, 689, 
	691, 693, 695, 697, 699, 701, 709, 713, 
	715, 722, 735, 739, 750, 752, 759, 761, 
	768, 771, 780, 783, 787, 791, 793, 800, 
	803, 805, 807, 809, 811, 813, 815, 817, 
	831, 834, 842, 848, 851, 868, 871, 878, 
	880, 887, 891, 893, 899, 904, 908, 918, 
	920, 926, 928, 934, 937, 945, 948, 952, 
	956, 958, 964, 967, 969, 971, 973, 975, 
	977, 979, 981, 983, 985, 987, 995, 999, 
	1001, 1008, 1021, 1025, 1036, 1038, 1045, 1047, 
	1054, 1057, 1066, 1069, 1073, 1077, 1079, 1086, 
	1089, 1091, 1093, 1095, 1097, 1099, 1101, 1103, 
	1109, 1113, 1119, 1124, 1133, 1138, 1143, 1150, 
	1155, 1156, 1159, 1160, 1161, 1162, 1163, 
}

var _cte_trans_targs []int16 = []int16{
	2, 1, 1, 256, 1, 6, 14, 259, 
	256, 256, 15, 19, 21, 256, 1, 255, 
	0, 2, 255, 0, 257, 0, 5, 5, 
	258, 0, 258, 0, 2, 259, 255, 0, 
	260, 0, 261, 0, 262, 262, 0, 11, 
	11, 0, 12, 11, 11, 0, 13, 13, 
	263, 0, 263, 0, 1, 1, 0, 16, 
	0, 17, 0, 18, 0, 256, 0, 20, 
	0, 256, 0, 22, 0, 23, 0, 256, 
	0, 264, 24, 26, 27, 25, 26, 265, 
	25, 26, 27, 25, 266, 29, 28, 28, 
	28, 28, 28, 28, 28, 31, 30, 30, 
	38, 39, 40, 53, 41, 38, 38, 267, 
	54, 58, 60, 38, 30, 32, 0, 31, 
	32, 0, 30, 30, 33, 267, 30, 32, 
	0, 34, 0, 30, 30, 267, 35, 30, 
	34, 0, 36, 36, 37, 0, 37, 0, 
	30, 30, 267, 30, 37, 0, 30, 30, 
	267, 30, 0, 31, 39, 39, 38, 39, 
	40, 53, 41, 38, 38, 54, 58, 60, 
	38, 39, 32, 0, 31, 41, 32, 0, 
	30, 30, 33, 267, 42, 44, 46, 30, 
	32, 0, 43, 0, 30, 30, 267, 30, 
	43, 0, 45, 0, 30, 30, 267, 30, 
	45, 0, 47, 47, 0, 30, 30, 48, 
	267, 30, 47, 47, 0, 49, 49, 0, 
	50, 49, 49, 0, 51, 51, 52, 0, 
	52, 0, 30, 30, 267, 30, 52, 0, 
	39, 39, 0, 55, 0, 56, 0, 57, 
	0, 38, 0, 59, 0, 38, 0, 61, 
	0, 62, 0, 38, 0, 64, 63, 63, 
	66, 105, 106, 119, 107, 120, 124, 268, 
	63, 65, 0, 64, 65, 0, 66, 66, 
	100, 67, 68, 66, 65, 0, 66, 66, 
	67, 68, 66, 0, 66, 66, 0, 69, 
	68, 68, 76, 68, 77, 90, 78, 76, 
	76, 91, 95, 97, 76, 68, 70, 0, 
	69, 70, 0, 63, 63, 71, 268, 63, 
	70, 0, 72, 0, 63, 63, 73, 268, 
	63, 72, 0, 74, 74, 75, 0, 75, 
	0, 63, 63, 268, 63, 75, 0, 63, 
	63, 268, 63, 0, 69, 78, 70, 0, 
	63, 63, 71, 79, 81, 83, 268, 63, 
	70, 0, 80, 0, 63, 63, 268, 63, 
	80, 0, 82, 0, 63, 63, 268, 63, 
	82, 0, 84, 84, 0, 63, 63, 85, 
	268, 63, 84, 84, 0, 86, 86, 0, 
	87, 86, 86, 0, 88, 88, 89, 0, 
	89, 0, 63, 63, 268, 63, 89, 0, 
	68, 68, 0, 92, 0, 93, 0, 94, 
	0, 76, 0, 96, 0, 76, 0, 98, 
	0, 99, 0, 76, 0, 101, 0, 66, 
	66, 67, 68, 102, 66, 101, 0, 103, 
	103, 104, 0, 104, 0, 66, 66, 67, 
	68, 66, 104, 0, 64, 105, 105, 66, 
	105, 106, 119, 107, 120, 124, 105, 65, 
	0, 64, 107, 65, 0, 66, 66, 100, 
	67, 68, 108, 110, 112, 66, 65, 0, 
	109, 0, 66, 66, 67, 68, 66, 109, 
	0, 111, 0, 66, 66, 67, 68, 66, 
	111, 0, 113, 113, 0, 66, 66, 114, 
	67, 68, 66, 113, 113, 0, 115, 115, 
	0, 116, 115, 115, 0, 117, 117, 118, 
	0, 118, 0, 66, 66, 67, 68, 66, 
	118, 0, 105, 105, 0, 121, 0, 122, 
	0, 123, 0, 66, 0, 125, 0, 126, 
	0, 66, 0, 128, 127, 127, 130, 169, 
	170, 183, 171, 269, 184, 188, 127, 129, 
	0, 128, 129, 0, 130, 130, 164, 131, 
	132, 130, 129, 0, 130, 130, 131, 132, 
	130, 0, 130, 130, 0, 133, 132, 132, 
	140, 132, 141, 154, 142, 140, 140, 155, 
	159, 161, 140, 132, 134, 0, 133, 134, 
	0, 127, 127, 135, 269, 127, 134, 0, 
	136, 0, 127, 127, 269, 137, 127, 136, 
	0, 138, 138, 139, 0, 139, 0, 127, 
	127, 269, 127, 139, 0, 127, 127, 269, 
	127, 0, 133, 142, 134, 0, 127, 127, 
	135, 269, 143, 145, 147, 127, 134, 0, 
	144, 0, 127, 127, 269, 127, 144, 0, 
	146, 0, 127, 127, 269, 127, 146, 0, 
	148, 148, 0, 127, 127, 149, 269, 127, 
	148, 148, 0, 150, 150, 0, 151, 150, 
	150, 0, 152, 152, 153, 0, 153, 0, 
	127, 127, 269, 127, 153, 0, 132, 132, 
	0, 156, 0, 157, 0, 158, 0, 140, 
	0, 160, 0, 140, 0, 162, 0, 163, 
	0, 140, 0, 165, 0, 130, 130, 131, 
	132, 166, 130, 165, 0, 167, 167, 168, 
	0, 168, 0, 130, 130, 131, 132, 130, 
	168, 0, 128, 169, 169, 130, 169, 170, 
	183, 171, 184, 188, 169, 129, 0, 128, 
	171, 129, 0, 130, 130, 164, 131, 132, 
	172, 174, 176, 130, 129, 0, 173, 0, 
	130, 130, 131, 132, 130, 173, 0, 175, 
	0, 130, 130, 131, 132, 130, 175, 0, 
	177, 177, 0, 130, 130, 178, 131, 132, 
	130, 177, 177, 0, 179, 179, 0, 180, 
	179, 179, 0, 181, 181, 182, 0, 182, 
	0, 130, 130, 131, 132, 130, 182, 0, 
	169, 169, 0, 185, 0, 186, 0, 187, 
	0, 130, 0, 189, 0, 190, 0, 130, 
	0, 192, 191, 191, 194, 233, 270, 234, 
	247, 235, 248, 252, 191, 193, 0, 192, 
	193, 0, 194, 194, 228, 195, 196, 194, 
	193, 0, 194, 194, 195, 196, 194, 0, 
	194, 194, 0, 197, 196, 196, 204, 196, 
	205, 218, 206, 204, 204, 219, 223, 225, 
	204, 196, 198, 0, 197, 198, 0, 191, 
	191, 270, 199, 191, 198, 0, 200, 0, 
	191, 191, 270, 201, 191, 200, 0, 202, 
	202, 203, 0, 203, 0, 191, 191, 270, 
	191, 203, 0, 191, 191, 270, 191, 0, 
	197, 206, 198, 0, 191, 191, 270, 199, 
	207, 209, 211, 191, 198, 0, 208, 0, 
	191, 191, 270, 191, 208, 0, 210, 0, 
	191, 191, 270, 191, 210, 0, 212, 212, 
	0, 191, 191, 270, 213, 191, 212, 212, 
	0, 214, 214, 0, 215, 214, 214, 0, 
	216, 216, 217, 0, 217, 0, 191, 191, 
	270, 191, 217, 0, 196, 196, 0, 220, 
	0, 221, 0, 222, 0, 204, 0, 224, 
	0, 204, 0, 226, 0, 227, 0, 204, 
	0, 229, 0, 194, 194, 195, 196, 230, 
	194, 229, 0, 231, 231, 232, 0, 232, 
	0, 194, 194, 195, 196, 194, 232, 0, 
	192, 233, 233, 194, 233, 234, 247, 235, 
	248, 252, 233, 193, 0, 192, 235, 193, 
	0, 194, 194, 228, 195, 196, 236, 238, 
	240, 194, 193, 0, 237, 0, 194, 194, 
	195, 196, 194, 237, 0, 239, 0, 194, 
	194, 195, 196, 194, 239, 0, 241, 241, 
	0, 194, 194, 242, 195, 196, 194, 241, 
	241, 0, 243, 243, 0, 244, 243, 243, 
	0, 245, 245, 246, 0, 246, 0, 194, 
	194, 195, 196, 194, 246, 0, 233, 233, 
	0, 249, 0, 250, 0, 251, 0, 194, 
	0, 253, 0, 254, 0, 194, 0, 256, 
	256, 3, 256, 255, 0, 256, 256, 256, 
	0, 256, 256, 4, 256, 257, 0, 256, 
	256, 256, 258, 0, 256, 256, 3, 7, 
	8, 9, 256, 255, 0, 256, 256, 256, 
	260, 0, 256, 256, 256, 261, 0, 256, 
	256, 10, 256, 262, 262, 0, 256, 256, 
	256, 263, 0, 0, 26, 27, 25, 0, 
	0, 0, 0, 0, 
}

var _cte_trans_actions []byte = []byte{
	0, 0, 0, 47, 41, 7, 0, 11, 
	39, 35, 0, 0, 0, 37, 0, 11, 
	0, 0, 11, 0, 17, 0, 0, 9, 
	23, 0, 23, 0, 0, 11, 11, 0, 
	25, 0, 27, 0, 13, 15, 0, 19, 
	21, 0, 0, 19, 21, 0, 0, 9, 
	23, 0, 23, 0, 45, 43, 0, 0, 
	0, 0, 0, 0, 0, 5, 0, 0, 
	0, 1, 0, 0, 0, 0, 0, 3, 
	0, 49, 0, 0, 0, 0, 0, 51, 
	0, 45, 0, 0, 65, 0, 0, 61, 
	53, 55, 57, 59, 63, 0, 0, 0, 
	47, 41, 7, 0, 11, 39, 35, 67, 
	0, 0, 0, 37, 0, 11, 0, 0, 
	11, 0, 29, 29, 0, 75, 29, 11, 
	0, 17, 0, 31, 31, 87, 0, 31, 
	17, 0, 0, 9, 23, 0, 23, 0, 
	31, 31, 87, 31, 23, 0, 0, 0, 
	67, 0, 0, 0, 0, 0, 47, 41, 
	7, 0, 11, 39, 35, 0, 0, 0, 
	37, 0, 11, 0, 0, 11, 11, 0, 
	29, 29, 0, 75, 0, 0, 0, 29, 
	11, 0, 25, 0, 29, 29, 75, 29, 
	25, 0, 27, 0, 29, 29, 75, 29, 
	27, 0, 13, 15, 0, 29, 29, 0, 
	75, 29, 13, 15, 0, 19, 21, 0, 
	0, 19, 21, 0, 0, 9, 23, 0, 
	23, 0, 33, 33, 99, 33, 23, 0, 
	45, 43, 0, 0, 0, 0, 0, 0, 
	0, 5, 0, 0, 0, 1, 0, 0, 
	0, 0, 0, 3, 0, 0, 0, 0, 
	47, 41, 7, 0, 11, 0, 0, 69, 
	0, 11, 0, 0, 11, 0, 29, 29, 
	0, 29, 29, 29, 11, 0, 0, 0, 
	0, 0, 0, 0, 45, 43, 0, 0, 
	0, 0, 47, 41, 7, 0, 11, 39, 
	35, 0, 0, 0, 37, 0, 11, 0, 
	0, 11, 0, 29, 29, 0, 78, 29, 
	11, 0, 17, 0, 31, 31, 0, 90, 
	31, 17, 0, 0, 9, 23, 0, 23, 
	0, 31, 31, 90, 31, 23, 0, 0, 
	0, 69, 0, 0, 0, 11, 11, 0, 
	29, 29, 0, 0, 0, 0, 78, 29, 
	11, 0, 25, 0, 29, 29, 78, 29, 
	25, 0, 27, 0, 29, 29, 78, 29, 
	27, 0, 13, 15, 0, 29, 29, 0, 
	78, 29, 13, 15, 0, 19, 21, 0, 
	0, 19, 21, 0, 0, 9, 23, 0, 
	23, 0, 33, 33, 102, 33, 23, 0, 
	45, 43, 0, 0, 0, 0, 0, 0, 
	0, 5, 0, 0, 0, 1, 0, 0, 
	0, 0, 0, 3, 0, 17, 0, 31, 
	31, 31, 31, 0, 31, 17, 0, 0, 
	9, 23, 0, 23, 0, 31, 31, 31, 
	31, 31, 23, 0, 0, 0, 0, 47, 
	41, 7, 0, 11, 0, 0, 0, 11, 
	0, 0, 11, 11, 0, 29, 29, 0, 
	29, 29, 0, 0, 0, 29, 11, 0, 
	25, 0, 29, 29, 29, 29, 29, 25, 
	0, 27, 0, 29, 29, 29, 29, 29, 
	27, 0, 13, 15, 0, 29, 29, 0, 
	29, 29, 29, 13, 15, 0, 19, 21, 
	0, 0, 19, 21, 0, 0, 9, 23, 
	0, 23, 0, 33, 33, 33, 33, 33, 
	23, 0, 45, 43, 0, 0, 0, 0, 
	0, 0, 0, 5, 0, 0, 0, 0, 
	0, 3, 0, 0, 0, 0, 47, 41, 
	7, 0, 11, 71, 0, 0, 0, 11, 
	0, 0, 11, 0, 29, 29, 0, 29, 
	29, 29, 11, 0, 0, 0, 0, 0, 
	0, 0, 45, 43, 0, 0, 0, 0, 
	47, 41, 7, 0, 11, 39, 35, 0, 
	0, 0, 37, 0, 11, 0, 0, 11, 
	0, 29, 29, 0, 81, 29, 11, 0, 
	17, 0, 31, 31, 93, 0, 31, 17, 
	0, 0, 9, 23, 0, 23, 0, 31, 
	31, 93, 31, 23, 0, 0, 0, 71, 
	0, 0, 0, 11, 11, 0, 29, 29, 
	0, 81, 0, 0, 0, 29, 11, 0, 
	25, 0, 29, 29, 81, 29, 25, 0, 
	27, 0, 29, 29, 81, 29, 27, 0, 
	13, 15, 0, 29, 29, 0, 81, 29, 
	13, 15, 0, 19, 21, 0, 0, 19, 
	21, 0, 0, 9, 23, 0, 23, 0, 
	33, 33, 105, 33, 23, 0, 45, 43, 
	0, 0, 0, 0, 0, 0, 0, 5, 
	0, 0, 0, 1, 0, 0, 0, 0, 
	0, 3, 0, 17, 0, 31, 31, 31, 
	31, 0, 31, 17, 0, 0, 9, 23, 
	0, 23, 0, 31, 31, 31, 31, 31, 
	23, 0, 0, 0, 0, 47, 41, 7, 
	0, 11, 0, 0, 0, 11, 0, 0, 
	11, 11, 0, 29, 29, 0, 29, 29, 
	0, 0, 0, 29, 11, 0, 25, 0, 
	29, 29, 29, 29, 29, 25, 0, 27, 
	0, 29, 29, 29, 29, 29, 27, 0, 
	13, 15, 0, 29, 29, 0, 29, 29, 
	29, 13, 15, 0, 19, 21, 0, 0, 
	19, 21, 0, 0, 9, 23, 0, 23, 
	0, 33, 33, 33, 33, 33, 23, 0, 
	45, 43, 0, 0, 0, 0, 0, 0, 
	0, 5, 0, 0, 0, 0, 0, 3, 
	0, 0, 0, 0, 47, 41, 73, 7, 
	0, 11, 0, 0, 0, 11, 0, 0, 
	11, 0, 29, 29, 0, 29, 29, 29, 
	11, 0, 0, 0, 0, 0, 0, 0, 
	45, 43, 0, 0, 0, 0, 47, 41, 
	7, 0, 11, 39, 35, 0, 0, 0, 
	37, 0, 11, 0, 0, 11, 0, 29, 
	29, 84, 0, 29, 11, 0, 17, 0, 
	31, 31, 96, 0, 31, 17, 0, 0, 
	9, 23, 0, 23, 0, 31, 31, 96, 
	31, 23, 0, 0, 0, 73, 0, 0, 
	0, 11, 11, 0, 29, 29, 84, 0, 
	0, 0, 0, 29, 11, 0, 25, 0, 
	29, 29, 84, 29, 25, 0, 27, 0, 
	29, 29, 84, 29, 27, 0, 13, 15, 
	0, 29, 29, 84, 0, 29, 13, 15, 
	0, 19, 21, 0, 0, 19, 21, 0, 
	0, 9, 23, 0, 23, 0, 33, 33, 
	108, 33, 23, 0, 45, 43, 0, 0, 
	0, 0, 0, 0, 0, 5, 0, 0, 
	0, 1, 0, 0, 0, 0, 0, 3, 
	0, 17, 0, 31, 31, 31, 31, 0, 
	31, 17, 0, 0, 9, 23, 0, 23, 
	0, 31, 31, 31, 31, 31, 23, 0, 
	0, 0, 0, 47, 41, 7, 0, 11, 
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
	0, 0, 0, 0, 0, 3, 0, 29, 
	29, 0, 29, 11, 0, 0, 0, 0, 
	0, 31, 31, 0, 31, 17, 0, 31, 
	31, 31, 23, 0, 29, 29, 0, 0, 
	0, 0, 29, 11, 0, 29, 29, 29, 
	25, 0, 29, 29, 29, 27, 0, 29, 
	29, 0, 29, 13, 15, 0, 33, 33, 
	33, 23, 0, 0, 45, 0, 0, 0, 
	0, 0, 0, 0, 
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
	0, 0, 0, 0, 0, 0, 0, 29, 
	0, 31, 31, 29, 29, 29, 29, 33, 
	0, 0, 0, 0, 0, 0, 0, 
}

const cte_start int = 1
const cte_error int = 0

const cte_en_comment_iterate int = 24
const cte_en_multiline_comment_iterate int = 25
const cte_en_string_iterate int = 28
const cte_en_list_iterate int = 30
const cte_en_unordered_map_iterate int = 63
const cte_en_ordered_map_iterate int = 127
const cte_en_metadata_map_iterate int = 191
const cte_en_main int = 1


//line cte.rl:329

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
    
    
//line cte.go:768
	{
	 this.cs = cte_start
	 this.top = 0
	}

//line cte.go:774
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

        this.significand = (this.significand << 4) | uint64( this.data[p] - '0')
    
		case 7:
//line cte.rl:84

        this.significand = (this.significand << 4) | uint64( this.data[p] - 'a' + 10)
    
		case 8:
//line cte.rl:88

        this.significand = this.significand * 10 + uint64( this.data[p] - '0')
        this.exponentAdjust--
    
		case 9:
//line cte.rl:93

        this.significand = (this.significand << 4) | uint64( this.data[p] - '0')
        this.exponentAdjust -= 4
    
		case 10:
//line cte.rl:96

        this.significand = (this.significand << 4) | uint64( this.data[p] - 'a' + 10)
        this.exponentAdjust -= 4
    
		case 11:
//line cte.rl:101

        this.exponent = this.exponent * 10 + int( this.data[p] - '0')
    
		case 12:
//line cte.rl:110

        this.significand = (this.significand << 1) | uint64( this.data[p] - '0')
    
		case 13:
//line cte.rl:114

        this.significand = (this.significand << 3) | uint64( this.data[p] - '0')
    
		case 14:
//line cte.rl:120

        if this.significandSign >= 0 {
            callbacks.OnPositiveInt(this.significand)
        } else {
            callbacks.OnNegativeInt(this.significand)
        }
        this.significandSign = 1
        this.significand = 0
    
		case 15:
//line cte.rl:130

        callbacks.OnDecimalFloat(int64(this.significand) * int64(this.significandSign), (this.exponent+this.exponentAdjust) * this.exponentSign)
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.exponentSign = 1
        this.exponent = 0
    
		case 16:
//line cte.rl:139

        exponent := float64((this.exponent * this.exponentSign + this.exponentAdjust))
        callbacks.OnFloat(float64(this.significandSign) * float64(this.significand) * math.Pow(2.0, exponent))
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.exponentSign = 1
        this.exponent = 0
    
		case 17:
//line cte.rl:151

        err = callbacks.OnListBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 30; goto _again

    
		case 18:
//line cte.rl:159

        err = callbacks.OnUnorderedMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 63; goto _again

    
		case 19:
//line cte.rl:167

        err = callbacks.OnOrderedMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 127; goto _again

    
		case 20:
//line cte.rl:175

        err = callbacks.OnMetadataMapBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 191; goto _again

    
		case 21:
//line cte.rl:183

        this.arrayStart = p + 1
        err = callbacks.OnCommentBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 24; goto _again

    
		case 22:
//line cte.rl:192

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
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 25; goto _again

    
		case 23:
//line cte.rl:209

        this.arrayStart = p + 1
        err = callbacks.OnStringBegin()
        if err != nil {
            p++; goto _out

        }
         this.stack[ this.top] =  this.cs;  this.top++;  this.cs = 28; goto _again

    
		case 24:
//line cte.rl:228

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

        
		case 25:
//line cte.rl:241

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

        
		case 26:
//line cte.rl:261
this.flushAndAddEscapedCharacter(p-1, '\\', callbacks)
		case 27:
//line cte.rl:262
this.flushAndAddEscapedCharacter(p-1, '\n', callbacks)
		case 28:
//line cte.rl:263
this.flushAndAddEscapedCharacter(p-1, '\r', callbacks)
		case 29:
//line cte.rl:264
this.flushAndAddEscapedCharacter(p-1, '\t', callbacks)
		case 30:
//line cte.rl:265
this.flushAndAddEscapedCharacter(p-1, '"', callbacks)
		case 31:
//line cte.rl:266
return false, fmt.Errorf("\\%c: Illegal escape encoding", this.data[p])
		case 32:
//line cte.rl:270

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

        
		case 33:
//line cte.rl:284

            err = callbacks.OnContainerEnd()
            if err != nil {
                p++; goto _out

            }
             this.top--;  this.cs =  this.stack[ this.top]
goto _again

        
		case 34:
//line cte.rl:294

            err = callbacks.OnContainerEnd()
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

        
//line cte.go:1179
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
//line cte.rl:120

        if this.significandSign >= 0 {
            callbacks.OnPositiveInt(this.significand)
        } else {
            callbacks.OnNegativeInt(this.significand)
        }
        this.significandSign = 1
        this.significand = 0
    
			case 15:
//line cte.rl:130

        callbacks.OnDecimalFloat(int64(this.significand) * int64(this.significandSign), (this.exponent+this.exponentAdjust) * this.exponentSign)
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.exponentSign = 1
        this.exponent = 0
    
			case 16:
//line cte.rl:139

        exponent := float64((this.exponent * this.exponentSign + this.exponentAdjust))
        callbacks.OnFloat(float64(this.significandSign) * float64(this.significand) * math.Pow(2.0, exponent))
        this.significandSign = 1
        this.significand = 0
        this.exponentAdjust = 0
        this.exponentSign = 1
        this.exponent = 0
    
//line cte.go:1230
			}
		}
	}

	_out: {}
	}

//line cte.rl:392


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
