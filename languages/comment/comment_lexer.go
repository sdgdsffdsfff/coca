// Code generated from CommentLexer.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 65, 532,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 148, 10, 2, 12, 2, 14,
	2, 151, 11, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7,
	3, 162, 10, 3, 12, 3, 14, 3, 165, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 7, 4,
	171, 10, 4, 12, 4, 14, 4, 174, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 5,
	5, 181, 10, 5, 3, 5, 6, 5, 184, 10, 5, 13, 5, 14, 5, 185, 3, 5, 5, 5, 189,
	10, 5, 5, 5, 191, 10, 5, 3, 5, 5, 5, 194, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	7, 6, 200, 10, 6, 12, 6, 14, 6, 203, 11, 6, 3, 6, 5, 6, 206, 10, 6, 3,
	6, 5, 6, 209, 10, 6, 3, 7, 3, 7, 7, 7, 213, 10, 7, 12, 7, 14, 7, 216, 11,
	7, 3, 7, 3, 7, 7, 7, 220, 10, 7, 12, 7, 14, 7, 223, 11, 7, 3, 7, 5, 7,
	226, 10, 7, 3, 7, 5, 7, 229, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 235,
	10, 8, 12, 8, 14, 8, 238, 11, 8, 3, 8, 5, 8, 241, 10, 8, 3, 8, 5, 8, 244,
	10, 8, 3, 9, 3, 9, 3, 9, 5, 9, 249, 10, 9, 3, 9, 3, 9, 5, 9, 253, 10, 9,
	3, 9, 5, 9, 256, 10, 9, 3, 9, 5, 9, 259, 10, 9, 3, 9, 3, 9, 3, 9, 5, 9,
	264, 10, 9, 3, 9, 5, 9, 267, 10, 9, 5, 9, 269, 10, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 5, 10, 275, 10, 10, 3, 10, 5, 10, 278, 10, 10, 3, 10, 3, 10,
	5, 10, 282, 10, 10, 3, 10, 3, 10, 5, 10, 286, 10, 10, 3, 10, 3, 10, 5,
	10, 290, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 5, 11, 301, 10, 11, 3, 12, 3, 12, 3, 12, 5, 12, 306, 10, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 13, 7, 13, 313, 10, 13, 12, 13, 14, 13, 316,
	11, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3,
	31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34,
	3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 38, 3,
	38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43,
	3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3,
	48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51,
	3, 51, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3,
	55, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57,
	3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3, 58, 7, 58, 438, 10, 58, 12, 58, 14,
	58, 441, 11, 58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60,
	3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 3, 62, 3, 63, 6, 63, 458, 10, 63, 13,
	63, 14, 63, 459, 3, 63, 3, 63, 3, 64, 3, 64, 7, 64, 466, 10, 64, 12, 64,
	14, 64, 469, 11, 64, 3, 65, 3, 65, 5, 65, 473, 10, 65, 3, 65, 3, 65, 3,
	66, 3, 66, 3, 66, 3, 66, 5, 66, 481, 10, 66, 3, 66, 5, 66, 484, 10, 66,
	3, 66, 3, 66, 3, 66, 6, 66, 489, 10, 66, 13, 66, 14, 66, 490, 3, 66, 3,
	66, 3, 66, 3, 66, 3, 66, 5, 66, 498, 10, 66, 3, 67, 3, 67, 3, 67, 7, 67,
	503, 10, 67, 12, 67, 14, 67, 506, 11, 67, 3, 67, 5, 67, 509, 10, 67, 3,
	68, 3, 68, 3, 69, 3, 69, 7, 69, 515, 10, 69, 12, 69, 14, 69, 518, 11, 69,
	3, 69, 5, 69, 521, 10, 69, 3, 70, 3, 70, 5, 70, 525, 10, 70, 3, 71, 3,
	71, 3, 71, 3, 71, 5, 71, 531, 10, 71, 3, 149, 2, 72, 3, 3, 5, 4, 7, 5,
	9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63,
	33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81,
	42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 50, 99,
	51, 101, 52, 103, 53, 105, 54, 107, 55, 109, 56, 111, 57, 113, 58, 115,
	59, 117, 60, 119, 61, 121, 62, 123, 63, 125, 64, 127, 65, 129, 2, 131,
	2, 133, 2, 135, 2, 137, 2, 139, 2, 141, 2, 3, 2, 30, 5, 2, 12, 12, 15,
	15, 8234, 8235, 4, 2, 12, 12, 14, 15, 3, 2, 51, 59, 4, 2, 78, 78, 110,
	110, 4, 2, 90, 90, 122, 122, 5, 2, 50, 59, 67, 72, 99, 104, 6, 2, 50, 59,
	67, 72, 97, 97, 99, 104, 3, 2, 50, 57, 4, 2, 50, 57, 97, 97, 4, 2, 68,
	68, 100, 100, 3, 2, 50, 51, 4, 2, 50, 51, 97, 97, 6, 2, 70, 70, 72, 72,
	102, 102, 104, 104, 4, 2, 82, 82, 114, 114, 4, 2, 45, 45, 47, 47, 6, 2,
	12, 12, 15, 15, 41, 41, 94, 94, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 3,
	2, 98, 98, 5, 2, 11, 12, 14, 15, 34, 34, 4, 2, 71, 71, 103, 103, 10, 2,
	36, 36, 41, 41, 94, 94, 100, 100, 104, 104, 112, 112, 116, 116, 118, 118,
	3, 2, 50, 53, 3, 2, 50, 59, 4, 2, 50, 59, 97, 97, 6, 2, 38, 38, 67, 92,
	97, 97, 99, 124, 4, 2, 2, 129, 55298, 56321, 3, 2, 55298, 56321, 3, 2,
	56322, 57345, 2, 576, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2,
	2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2,
	2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3,
	2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31,
	3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2,
	39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2,
	2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2,
	2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2,
	2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3,
	2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77,
	3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2,
	85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2,
	2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2,
	2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107,
	3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2,
	2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3,
	2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2, 3,
	143, 3, 2, 2, 2, 5, 157, 3, 2, 2, 2, 7, 168, 3, 2, 2, 2, 9, 190, 3, 2,
	2, 2, 11, 195, 3, 2, 2, 2, 13, 210, 3, 2, 2, 2, 15, 230, 3, 2, 2, 2, 17,
	268, 3, 2, 2, 2, 19, 270, 3, 2, 2, 2, 21, 300, 3, 2, 2, 2, 23, 302, 3,
	2, 2, 2, 25, 309, 3, 2, 2, 2, 27, 319, 3, 2, 2, 2, 29, 324, 3, 2, 2, 2,
	31, 326, 3, 2, 2, 2, 33, 328, 3, 2, 2, 2, 35, 330, 3, 2, 2, 2, 37, 332,
	3, 2, 2, 2, 39, 334, 3, 2, 2, 2, 41, 336, 3, 2, 2, 2, 43, 338, 3, 2, 2,
	2, 45, 340, 3, 2, 2, 2, 47, 342, 3, 2, 2, 2, 49, 344, 3, 2, 2, 2, 51, 346,
	3, 2, 2, 2, 53, 348, 3, 2, 2, 2, 55, 350, 3, 2, 2, 2, 57, 352, 3, 2, 2,
	2, 59, 354, 3, 2, 2, 2, 61, 356, 3, 2, 2, 2, 63, 359, 3, 2, 2, 2, 65, 362,
	3, 2, 2, 2, 67, 365, 3, 2, 2, 2, 69, 368, 3, 2, 2, 2, 71, 371, 3, 2, 2,
	2, 73, 374, 3, 2, 2, 2, 75, 377, 3, 2, 2, 2, 77, 380, 3, 2, 2, 2, 79, 382,
	3, 2, 2, 2, 81, 384, 3, 2, 2, 2, 83, 386, 3, 2, 2, 2, 85, 388, 3, 2, 2,
	2, 87, 390, 3, 2, 2, 2, 89, 392, 3, 2, 2, 2, 91, 394, 3, 2, 2, 2, 93, 396,
	3, 2, 2, 2, 95, 399, 3, 2, 2, 2, 97, 402, 3, 2, 2, 2, 99, 405, 3, 2, 2,
	2, 101, 408, 3, 2, 2, 2, 103, 411, 3, 2, 2, 2, 105, 414, 3, 2, 2, 2, 107,
	417, 3, 2, 2, 2, 109, 420, 3, 2, 2, 2, 111, 424, 3, 2, 2, 2, 113, 428,
	3, 2, 2, 2, 115, 433, 3, 2, 2, 2, 117, 444, 3, 2, 2, 2, 119, 447, 3, 2,
	2, 2, 121, 450, 3, 2, 2, 2, 123, 452, 3, 2, 2, 2, 125, 457, 3, 2, 2, 2,
	127, 463, 3, 2, 2, 2, 129, 470, 3, 2, 2, 2, 131, 497, 3, 2, 2, 2, 133,
	499, 3, 2, 2, 2, 135, 510, 3, 2, 2, 2, 137, 512, 3, 2, 2, 2, 139, 524,
	3, 2, 2, 2, 141, 530, 3, 2, 2, 2, 143, 144, 7, 49, 2, 2, 144, 145, 7, 44,
	2, 2, 145, 149, 3, 2, 2, 2, 146, 148, 11, 2, 2, 2, 147, 146, 3, 2, 2, 2,
	148, 151, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 150,
	152, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 152, 153, 7, 44, 2, 2, 153, 154,
	7, 49, 2, 2, 154, 155, 3, 2, 2, 2, 155, 156, 8, 2, 2, 2, 156, 4, 3, 2,
	2, 2, 157, 158, 7, 49, 2, 2, 158, 159, 7, 49, 2, 2, 159, 163, 3, 2, 2,
	2, 160, 162, 10, 2, 2, 2, 161, 160, 3, 2, 2, 2, 162, 165, 3, 2, 2, 2, 163,
	161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 166, 3, 2, 2, 2, 165, 163,
	3, 2, 2, 2, 166, 167, 8, 3, 2, 2, 167, 6, 3, 2, 2, 2, 168, 172, 7, 37,
	2, 2, 169, 171, 10, 3, 2, 2, 170, 169, 3, 2, 2, 2, 171, 174, 3, 2, 2, 2,
	172, 170, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 175, 3, 2, 2, 2, 174,
	172, 3, 2, 2, 2, 175, 176, 8, 4, 2, 2, 176, 8, 3, 2, 2, 2, 177, 191, 7,
	50, 2, 2, 178, 188, 9, 4, 2, 2, 179, 181, 5, 137, 69, 2, 180, 179, 3, 2,
	2, 2, 180, 181, 3, 2, 2, 2, 181, 189, 3, 2, 2, 2, 182, 184, 7, 97, 2, 2,
	183, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 185,
	186, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 189, 5, 137, 69, 2, 188, 180,
	3, 2, 2, 2, 188, 183, 3, 2, 2, 2, 189, 191, 3, 2, 2, 2, 190, 177, 3, 2,
	2, 2, 190, 178, 3, 2, 2, 2, 191, 193, 3, 2, 2, 2, 192, 194, 9, 5, 2, 2,
	193, 192, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 10, 3, 2, 2, 2, 195, 196,
	7, 50, 2, 2, 196, 197, 9, 6, 2, 2, 197, 205, 9, 7, 2, 2, 198, 200, 9, 8,
	2, 2, 199, 198, 3, 2, 2, 2, 200, 203, 3, 2, 2, 2, 201, 199, 3, 2, 2, 2,
	201, 202, 3, 2, 2, 2, 202, 204, 3, 2, 2, 2, 203, 201, 3, 2, 2, 2, 204,
	206, 9, 7, 2, 2, 205, 201, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 208,
	3, 2, 2, 2, 207, 209, 9, 5, 2, 2, 208, 207, 3, 2, 2, 2, 208, 209, 3, 2,
	2, 2, 209, 12, 3, 2, 2, 2, 210, 214, 7, 50, 2, 2, 211, 213, 7, 97, 2, 2,
	212, 211, 3, 2, 2, 2, 213, 216, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2, 214,
	215, 3, 2, 2, 2, 215, 217, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 217, 225,
	9, 9, 2, 2, 218, 220, 9, 10, 2, 2, 219, 218, 3, 2, 2, 2, 220, 223, 3, 2,
	2, 2, 221, 219, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 224, 3, 2, 2, 2,
	223, 221, 3, 2, 2, 2, 224, 226, 9, 9, 2, 2, 225, 221, 3, 2, 2, 2, 225,
	226, 3, 2, 2, 2, 226, 228, 3, 2, 2, 2, 227, 229, 9, 5, 2, 2, 228, 227,
	3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 14, 3, 2, 2, 2, 230, 231, 7, 50,
	2, 2, 231, 232, 9, 11, 2, 2, 232, 240, 9, 12, 2, 2, 233, 235, 9, 13, 2,
	2, 234, 233, 3, 2, 2, 2, 235, 238, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 236,
	237, 3, 2, 2, 2, 237, 239, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 239, 241,
	9, 12, 2, 2, 240, 236, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 243, 3, 2,
	2, 2, 242, 244, 9, 5, 2, 2, 243, 242, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2,
	244, 16, 3, 2, 2, 2, 245, 246, 5, 137, 69, 2, 246, 248, 7, 48, 2, 2, 247,
	249, 5, 137, 69, 2, 248, 247, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 253,
	3, 2, 2, 2, 250, 251, 7, 48, 2, 2, 251, 253, 5, 137, 69, 2, 252, 245, 3,
	2, 2, 2, 252, 250, 3, 2, 2, 2, 253, 255, 3, 2, 2, 2, 254, 256, 5, 129,
	65, 2, 255, 254, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 258, 3, 2, 2, 2,
	257, 259, 9, 14, 2, 2, 258, 257, 3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259,
	269, 3, 2, 2, 2, 260, 266, 5, 137, 69, 2, 261, 263, 5, 129, 65, 2, 262,
	264, 9, 14, 2, 2, 263, 262, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 267,
	3, 2, 2, 2, 265, 267, 9, 14, 2, 2, 266, 261, 3, 2, 2, 2, 266, 265, 3, 2,
	2, 2, 267, 269, 3, 2, 2, 2, 268, 252, 3, 2, 2, 2, 268, 260, 3, 2, 2, 2,
	269, 18, 3, 2, 2, 2, 270, 271, 7, 50, 2, 2, 271, 281, 9, 6, 2, 2, 272,
	274, 5, 133, 67, 2, 273, 275, 7, 48, 2, 2, 274, 273, 3, 2, 2, 2, 274, 275,
	3, 2, 2, 2, 275, 282, 3, 2, 2, 2, 276, 278, 5, 133, 67, 2, 277, 276, 3,
	2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 280, 7, 48, 2,
	2, 280, 282, 5, 133, 67, 2, 281, 272, 3, 2, 2, 2, 281, 277, 3, 2, 2, 2,
	282, 283, 3, 2, 2, 2, 283, 285, 9, 15, 2, 2, 284, 286, 9, 16, 2, 2, 285,
	284, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 289,
	5, 137, 69, 2, 288, 290, 9, 14, 2, 2, 289, 288, 3, 2, 2, 2, 289, 290, 3,
	2, 2, 2, 290, 20, 3, 2, 2, 2, 291, 292, 7, 118, 2, 2, 292, 293, 7, 116,
	2, 2, 293, 294, 7, 119, 2, 2, 294, 301, 7, 103, 2, 2, 295, 296, 7, 104,
	2, 2, 296, 297, 7, 99, 2, 2, 297, 298, 7, 110, 2, 2, 298, 299, 7, 117,
	2, 2, 299, 301, 7, 103, 2, 2, 300, 291, 3, 2, 2, 2, 300, 295, 3, 2, 2,
	2, 301, 22, 3, 2, 2, 2, 302, 305, 7, 41, 2, 2, 303, 306, 10, 17, 2, 2,
	304, 306, 5, 131, 66, 2, 305, 303, 3, 2, 2, 2, 305, 304, 3, 2, 2, 2, 306,
	307, 3, 2, 2, 2, 307, 308, 7, 41, 2, 2, 308, 24, 3, 2, 2, 2, 309, 314,
	7, 36, 2, 2, 310, 313, 10, 18, 2, 2, 311, 313, 5, 131, 66, 2, 312, 310,
	3, 2, 2, 2, 312, 311, 3, 2, 2, 2, 313, 316, 3, 2, 2, 2, 314, 312, 3, 2,
	2, 2, 314, 315, 3, 2, 2, 2, 315, 317, 3, 2, 2, 2, 316, 314, 3, 2, 2, 2,
	317, 318, 7, 36, 2, 2, 318, 26, 3, 2, 2, 2, 319, 320, 7, 112, 2, 2, 320,
	321, 7, 119, 2, 2, 321, 322, 7, 110, 2, 2, 322, 323, 7, 110, 2, 2, 323,
	28, 3, 2, 2, 2, 324, 325, 7, 42, 2, 2, 325, 30, 3, 2, 2, 2, 326, 327, 7,
	43, 2, 2, 327, 32, 3, 2, 2, 2, 328, 329, 7, 125, 2, 2, 329, 34, 3, 2, 2,
	2, 330, 331, 7, 127, 2, 2, 331, 36, 3, 2, 2, 2, 332, 333, 7, 93, 2, 2,
	333, 38, 3, 2, 2, 2, 334, 335, 7, 95, 2, 2, 335, 40, 3, 2, 2, 2, 336, 337,
	7, 61, 2, 2, 337, 42, 3, 2, 2, 2, 338, 339, 7, 46, 2, 2, 339, 44, 3, 2,
	2, 2, 340, 341, 7, 48, 2, 2, 341, 46, 3, 2, 2, 2, 342, 343, 7, 63, 2, 2,
	343, 48, 3, 2, 2, 2, 344, 345, 7, 64, 2, 2, 345, 50, 3, 2, 2, 2, 346, 347,
	7, 62, 2, 2, 347, 52, 3, 2, 2, 2, 348, 349, 7, 35, 2, 2, 349, 54, 3, 2,
	2, 2, 350, 351, 7, 128, 2, 2, 351, 56, 3, 2, 2, 2, 352, 353, 7, 65, 2,
	2, 353, 58, 3, 2, 2, 2, 354, 355, 7, 60, 2, 2, 355, 60, 3, 2, 2, 2, 356,
	357, 7, 63, 2, 2, 357, 358, 7, 63, 2, 2, 358, 62, 3, 2, 2, 2, 359, 360,
	7, 62, 2, 2, 360, 361, 7, 63, 2, 2, 361, 64, 3, 2, 2, 2, 362, 363, 7, 64,
	2, 2, 363, 364, 7, 63, 2, 2, 364, 66, 3, 2, 2, 2, 365, 366, 7, 35, 2, 2,
	366, 367, 7, 63, 2, 2, 367, 68, 3, 2, 2, 2, 368, 369, 7, 40, 2, 2, 369,
	370, 7, 40, 2, 2, 370, 70, 3, 2, 2, 2, 371, 372, 7, 126, 2, 2, 372, 373,
	7, 126, 2, 2, 373, 72, 3, 2, 2, 2, 374, 375, 7, 45, 2, 2, 375, 376, 7,
	45, 2, 2, 376, 74, 3, 2, 2, 2, 377, 378, 7, 47, 2, 2, 378, 379, 7, 47,
	2, 2, 379, 76, 3, 2, 2, 2, 380, 381, 7, 45, 2, 2, 381, 78, 3, 2, 2, 2,
	382, 383, 7, 47, 2, 2, 383, 80, 3, 2, 2, 2, 384, 385, 7, 44, 2, 2, 385,
	82, 3, 2, 2, 2, 386, 387, 7, 49, 2, 2, 387, 84, 3, 2, 2, 2, 388, 389, 7,
	40, 2, 2, 389, 86, 3, 2, 2, 2, 390, 391, 7, 126, 2, 2, 391, 88, 3, 2, 2,
	2, 392, 393, 7, 96, 2, 2, 393, 90, 3, 2, 2, 2, 394, 395, 7, 39, 2, 2, 395,
	92, 3, 2, 2, 2, 396, 397, 7, 45, 2, 2, 397, 398, 7, 63, 2, 2, 398, 94,
	3, 2, 2, 2, 399, 400, 7, 47, 2, 2, 400, 401, 7, 63, 2, 2, 401, 96, 3, 2,
	2, 2, 402, 403, 7, 44, 2, 2, 403, 404, 7, 63, 2, 2, 404, 98, 3, 2, 2, 2,
	405, 406, 7, 49, 2, 2, 406, 407, 7, 63, 2, 2, 407, 100, 3, 2, 2, 2, 408,
	409, 7, 40, 2, 2, 409, 410, 7, 63, 2, 2, 410, 102, 3, 2, 2, 2, 411, 412,
	7, 126, 2, 2, 412, 413, 7, 63, 2, 2, 413, 104, 3, 2, 2, 2, 414, 415, 7,
	96, 2, 2, 415, 416, 7, 63, 2, 2, 416, 106, 3, 2, 2, 2, 417, 418, 7, 39,
	2, 2, 418, 419, 7, 63, 2, 2, 419, 108, 3, 2, 2, 2, 420, 421, 7, 62, 2,
	2, 421, 422, 7, 62, 2, 2, 422, 423, 7, 63, 2, 2, 423, 110, 3, 2, 2, 2,
	424, 425, 7, 64, 2, 2, 425, 426, 7, 64, 2, 2, 426, 427, 7, 63, 2, 2, 427,
	112, 3, 2, 2, 2, 428, 429, 7, 64, 2, 2, 429, 430, 7, 64, 2, 2, 430, 431,
	7, 64, 2, 2, 431, 432, 7, 63, 2, 2, 432, 114, 3, 2, 2, 2, 433, 439, 7,
	98, 2, 2, 434, 435, 7, 94, 2, 2, 435, 438, 7, 98, 2, 2, 436, 438, 10, 19,
	2, 2, 437, 434, 3, 2, 2, 2, 437, 436, 3, 2, 2, 2, 438, 441, 3, 2, 2, 2,
	439, 437, 3, 2, 2, 2, 439, 440, 3, 2, 2, 2, 440, 442, 3, 2, 2, 2, 441,
	439, 3, 2, 2, 2, 442, 443, 7, 98, 2, 2, 443, 116, 3, 2, 2, 2, 444, 445,
	7, 47, 2, 2, 445, 446, 7, 64, 2, 2, 446, 118, 3, 2, 2, 2, 447, 448, 7,
	60, 2, 2, 448, 449, 7, 60, 2, 2, 449, 120, 3, 2, 2, 2, 450, 451, 7, 66,
	2, 2, 451, 122, 3, 2, 2, 2, 452, 453, 7, 48, 2, 2, 453, 454, 7, 48, 2,
	2, 454, 455, 7, 48, 2, 2, 455, 124, 3, 2, 2, 2, 456, 458, 9, 20, 2, 2,
	457, 456, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 457, 3, 2, 2, 2, 459,
	460, 3, 2, 2, 2, 460, 461, 3, 2, 2, 2, 461, 462, 8, 63, 2, 2, 462, 126,
	3, 2, 2, 2, 463, 467, 5, 141, 71, 2, 464, 466, 5, 139, 70, 2, 465, 464,
	3, 2, 2, 2, 466, 469, 3, 2, 2, 2, 467, 465, 3, 2, 2, 2, 467, 468, 3, 2,
	2, 2, 468, 128, 3, 2, 2, 2, 469, 467, 3, 2, 2, 2, 470, 472, 9, 21, 2, 2,
	471, 473, 9, 16, 2, 2, 472, 471, 3, 2, 2, 2, 472, 473, 3, 2, 2, 2, 473,
	474, 3, 2, 2, 2, 474, 475, 5, 137, 69, 2, 475, 130, 3, 2, 2, 2, 476, 477,
	7, 94, 2, 2, 477, 498, 9, 22, 2, 2, 478, 483, 7, 94, 2, 2, 479, 481, 9,
	23, 2, 2, 480, 479, 3, 2, 2, 2, 480, 481, 3, 2, 2, 2, 481, 482, 3, 2, 2,
	2, 482, 484, 9, 9, 2, 2, 483, 480, 3, 2, 2, 2, 483, 484, 3, 2, 2, 2, 484,
	485, 3, 2, 2, 2, 485, 498, 9, 9, 2, 2, 486, 488, 7, 94, 2, 2, 487, 489,
	7, 119, 2, 2, 488, 487, 3, 2, 2, 2, 489, 490, 3, 2, 2, 2, 490, 488, 3,
	2, 2, 2, 490, 491, 3, 2, 2, 2, 491, 492, 3, 2, 2, 2, 492, 493, 5, 135,
	68, 2, 493, 494, 5, 135, 68, 2, 494, 495, 5, 135, 68, 2, 495, 496, 5, 135,
	68, 2, 496, 498, 3, 2, 2, 2, 497, 476, 3, 2, 2, 2, 497, 478, 3, 2, 2, 2,
	497, 486, 3, 2, 2, 2, 498, 132, 3, 2, 2, 2, 499, 508, 5, 135, 68, 2, 500,
	503, 5, 135, 68, 2, 501, 503, 7, 97, 2, 2, 502, 500, 3, 2, 2, 2, 502, 501,
	3, 2, 2, 2, 503, 506, 3, 2, 2, 2, 504, 502, 3, 2, 2, 2, 504, 505, 3, 2,
	2, 2, 505, 507, 3, 2, 2, 2, 506, 504, 3, 2, 2, 2, 507, 509, 5, 135, 68,
	2, 508, 504, 3, 2, 2, 2, 508, 509, 3, 2, 2, 2, 509, 134, 3, 2, 2, 2, 510,
	511, 9, 7, 2, 2, 511, 136, 3, 2, 2, 2, 512, 520, 9, 24, 2, 2, 513, 515,
	9, 25, 2, 2, 514, 513, 3, 2, 2, 2, 515, 518, 3, 2, 2, 2, 516, 514, 3, 2,
	2, 2, 516, 517, 3, 2, 2, 2, 517, 519, 3, 2, 2, 2, 518, 516, 3, 2, 2, 2,
	519, 521, 9, 24, 2, 2, 520, 516, 3, 2, 2, 2, 520, 521, 3, 2, 2, 2, 521,
	138, 3, 2, 2, 2, 522, 525, 5, 141, 71, 2, 523, 525, 9, 24, 2, 2, 524, 522,
	3, 2, 2, 2, 524, 523, 3, 2, 2, 2, 525, 140, 3, 2, 2, 2, 526, 531, 9, 26,
	2, 2, 527, 531, 10, 27, 2, 2, 528, 529, 9, 28, 2, 2, 529, 531, 9, 29, 2,
	2, 530, 526, 3, 2, 2, 2, 530, 527, 3, 2, 2, 2, 530, 528, 3, 2, 2, 2, 531,
	142, 3, 2, 2, 2, 53, 2, 149, 163, 172, 180, 185, 188, 190, 193, 201, 205,
	208, 214, 221, 225, 228, 236, 240, 243, 248, 252, 255, 258, 263, 266, 268,
	274, 277, 281, 285, 289, 300, 305, 312, 314, 437, 439, 459, 467, 472, 480,
	483, 490, 497, 502, 504, 508, 516, 520, 524, 530, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "'null'", "'('", "')'",
	"'{'", "'}'", "'['", "']'", "';'", "','", "'.'", "'='", "'>'", "'<'", "'!'",
	"'~'", "'?'", "':'", "'=='", "'<='", "'>='", "'!='", "'&&'", "'||'", "'++'",
	"'--'", "'+'", "'-'", "'*'", "'/'", "'&'", "'|'", "'^'", "'%'", "'+='",
	"'-='", "'*='", "'/='", "'&='", "'|='", "'^='", "'%='", "'<<='", "'>>='",
	"'>>>='", "", "'->'", "'::'", "'@'", "'...'",
}

var lexerSymbolicNames = []string{
	"", "COMMENT", "LINE_COMMENT", "PYTHON_COMMENT", "DECIMAL_LITERAL", "HEX_LITERAL",
	"OCT_LITERAL", "BINARY_LITERAL", "FLOAT_LITERAL", "HEX_FLOAT_LITERAL",
	"BOOL_LITERAL", "CHAR_LITERAL", "STRING_LITERAL", "NULL_LITERAL", "LPAREN",
	"RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA", "DOT",
	"ASSIGN", "GT", "LT", "BANG", "TILDE", "QUESTION", "COLON", "EQUAL", "LE",
	"GE", "NOTEQUAL", "AND", "OR", "INC", "DEC", "ADD", "SUB", "MUL", "DIV",
	"BITAND", "BITOR", "CARET", "MOD", "ADD_ASSIGN", "SUB_ASSIGN", "MUL_ASSIGN",
	"DIV_ASSIGN", "AND_ASSIGN", "OR_ASSIGN", "XOR_ASSIGN", "MOD_ASSIGN", "LSHIFT_ASSIGN",
	"RSHIFT_ASSIGN", "URSHIFT_ASSIGN", "TemplateStringLiteral", "ARROW", "COLONCOLON",
	"AT", "ELLIPSIS", "WS", "IDENTIFIER",
}

var lexerRuleNames = []string{
	"COMMENT", "LINE_COMMENT", "PYTHON_COMMENT", "DECIMAL_LITERAL", "HEX_LITERAL",
	"OCT_LITERAL", "BINARY_LITERAL", "FLOAT_LITERAL", "HEX_FLOAT_LITERAL",
	"BOOL_LITERAL", "CHAR_LITERAL", "STRING_LITERAL", "NULL_LITERAL", "LPAREN",
	"RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA", "DOT",
	"ASSIGN", "GT", "LT", "BANG", "TILDE", "QUESTION", "COLON", "EQUAL", "LE",
	"GE", "NOTEQUAL", "AND", "OR", "INC", "DEC", "ADD", "SUB", "MUL", "DIV",
	"BITAND", "BITOR", "CARET", "MOD", "ADD_ASSIGN", "SUB_ASSIGN", "MUL_ASSIGN",
	"DIV_ASSIGN", "AND_ASSIGN", "OR_ASSIGN", "XOR_ASSIGN", "MOD_ASSIGN", "LSHIFT_ASSIGN",
	"RSHIFT_ASSIGN", "URSHIFT_ASSIGN", "TemplateStringLiteral", "ARROW", "COLONCOLON",
	"AT", "ELLIPSIS", "WS", "IDENTIFIER", "ExponentPart", "EscapeSequence",
	"HexDigits", "HexDigit", "Digits", "LetterOrDigit", "Letter",
}

type CommentLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewCommentLexer(input antlr.CharStream) *CommentLexer {

	l := new(CommentLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "CommentLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CommentLexer tokens.
const (
	CommentLexerCOMMENT               = 1
	CommentLexerLINE_COMMENT          = 2
	CommentLexerPYTHON_COMMENT        = 3
	CommentLexerDECIMAL_LITERAL       = 4
	CommentLexerHEX_LITERAL           = 5
	CommentLexerOCT_LITERAL           = 6
	CommentLexerBINARY_LITERAL        = 7
	CommentLexerFLOAT_LITERAL         = 8
	CommentLexerHEX_FLOAT_LITERAL     = 9
	CommentLexerBOOL_LITERAL          = 10
	CommentLexerCHAR_LITERAL          = 11
	CommentLexerSTRING_LITERAL        = 12
	CommentLexerNULL_LITERAL          = 13
	CommentLexerLPAREN                = 14
	CommentLexerRPAREN                = 15
	CommentLexerLBRACE                = 16
	CommentLexerRBRACE                = 17
	CommentLexerLBRACK                = 18
	CommentLexerRBRACK                = 19
	CommentLexerSEMI                  = 20
	CommentLexerCOMMA                 = 21
	CommentLexerDOT                   = 22
	CommentLexerASSIGN                = 23
	CommentLexerGT                    = 24
	CommentLexerLT                    = 25
	CommentLexerBANG                  = 26
	CommentLexerTILDE                 = 27
	CommentLexerQUESTION              = 28
	CommentLexerCOLON                 = 29
	CommentLexerEQUAL                 = 30
	CommentLexerLE                    = 31
	CommentLexerGE                    = 32
	CommentLexerNOTEQUAL              = 33
	CommentLexerAND                   = 34
	CommentLexerOR                    = 35
	CommentLexerINC                   = 36
	CommentLexerDEC                   = 37
	CommentLexerADD                   = 38
	CommentLexerSUB                   = 39
	CommentLexerMUL                   = 40
	CommentLexerDIV                   = 41
	CommentLexerBITAND                = 42
	CommentLexerBITOR                 = 43
	CommentLexerCARET                 = 44
	CommentLexerMOD                   = 45
	CommentLexerADD_ASSIGN            = 46
	CommentLexerSUB_ASSIGN            = 47
	CommentLexerMUL_ASSIGN            = 48
	CommentLexerDIV_ASSIGN            = 49
	CommentLexerAND_ASSIGN            = 50
	CommentLexerOR_ASSIGN             = 51
	CommentLexerXOR_ASSIGN            = 52
	CommentLexerMOD_ASSIGN            = 53
	CommentLexerLSHIFT_ASSIGN         = 54
	CommentLexerRSHIFT_ASSIGN         = 55
	CommentLexerURSHIFT_ASSIGN        = 56
	CommentLexerTemplateStringLiteral = 57
	CommentLexerARROW                 = 58
	CommentLexerCOLONCOLON            = 59
	CommentLexerAT                    = 60
	CommentLexerELLIPSIS              = 61
	CommentLexerWS                    = 62
	CommentLexerIDENTIFIER            = 63
)
