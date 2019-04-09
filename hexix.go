package main

import "fmt"

type Hex struct{
	value, load int
	hexw, hexnw, hexne, hexe, hexse, hexsw Hex
	is_inverse bool
}

var board = []Hex{
		{0,0,NULL,board[4],board[5],board[1],NULL,NULL,false}, //hex 0
		{0,0,board[0],board[5],board[6],board[2],NULL,NULL,false}, //hex 1
		{0,0,board[1],board[6],board[7],board[3],NULL,NULL,false}, //hex 2
		{6,0,board[2],board[7],board[8],NULL,NULL,NULL,false}, //hex 3
		{0,0,NULL,board[9],board[10],board[5],board[0],NULL,false}, //hex 4
		{0,0,board[4],board[10],board[11],board[6],board[1],board[0],false}, //hex 5
		{0,0,board[5],board[11],board[12],board[7],board[2],board[1],false}, //hex 6
		{0,0,board[6],board[12],board[13],board[8],board[3],board[2],false}, //hex 7
		{0,0,board[7],board[13],board[14],NULL,NULL,board[3],false}, //hex 8
		{0,0,NULL,NULL,board[15],board[10],board[4],NULL,false}, //hex 9
		{0,0,board[9],board[15],board[16],board[11],board[5],board[4],false}, //hex 10
		{0,0,board[10],board[16],board[17],board[12],board[6],board[5],false}, //hex 11
		{0,0,board[11],board[17],board[18],board[13],board[7],board[6],false}, //hex 12
		{0,0,board[12],board[18],board[19],board[14],board[8],board[7],false}, //hex 13
		{0,0,board[13],board[19],board[20],NULL,NULL,board[8],false}, //hex 14
		{0,0,NULL,NULL,board[21],board[16],board[10],board[9],false}, //hex 15
		{0,0,board[15],board[21],board[22],board[17],board[11],board[10],false}, //hex 16
		{0,0,board[16],board[22],board[23],board[18],board[12],board[11],false}, //hex 17
		{0,0,board[17],board[23],board[24],board[19],board[13],board[12],false}, //hex 18
		{0,0,board[18],board[24],board[25],board[20],board[14],board[13],false}, //hex 19
		{0,0,board[19],board[25],NULL,NULL,NULL,board[14],false}, //hex 20
		{0,0,NULL,NULL,board[26],board[22],board[16],board[15],false}, //hex 21
		{0,0,board[21],board[26],board[27],board[23],board[17],board[16],false}, //hex 22
		{0,0,board[22],board[27],board[28],board[24],board[18],board[17],false}, //hex 23
		{0,0,board[23],board[28],board[29],board[25],board[19],board[18],false}, //hex 24
		{0,0,board[24],board[29],NULL,NULL,board[20],board[19],false}, //hex 25
		{6,0,NULL,NULL,NULL,board[27],board[22],board[21],false}, //hex 26
		{0,0,board[26],NULL,NULL,board[28],board[23],board[22],false}, //hex 27
		{0,0,board[27],NULL,NULL,board[29],board[24],board[23],false}, //hex 28
		{0,0,board[28],NULL,NULL,NULL,board[25],board[24],false}, //hex 29
		}

func main() {
	    fmt.Println(len(board))
    }
