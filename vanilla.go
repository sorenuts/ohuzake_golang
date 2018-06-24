package main

import(
	"fmt"
	"math/rand"
	"time"
	"reflect"
)

//掛け声の方の「コール」
var StrFirstCall = []string{"V!","a!","n!","i!","l!","la!"}
var StrSecondCall = []string{"バーニラ","バニラ","バニラで"}
//コールの順番の定義
var SequenceStage3 = []int{1, 2, 1}
var SequenceStage4 = []int{1, 2}
var SequenceStage6 = []int{1, 3}

func main(){
	rand.Seed(time.Now().UnixNano())

	fmt.Println("[Stage:1]")
	stage1()
	fmt.Println("[Stage:2]")
	stage2()
	fmt.Println("[Stage:3]")
	stage3()
	fmt.Println("[Stage:4]")
	stage4()
	fmt.Println("[Stage:5]")
	stage5()
	fmt.Println("[Stage:6]")
	stage6()
}

/*
* Stage1:"V!a!n!i!l!la!"文字列完成で「バニラ！」
*/
func stage1(){
	x, stage := 0, 0

	//Vanilla文字列完成シーケンス
	for stage != len(StrFirstCall){
		x = rand.Intn(len(StrFirstCall))
		fmt.Print(StrFirstCall[x]," ")
		if x == stage{
			stage++
		}else{
			stage = 0
		}
	}

	//文字列完成後の処理
	fmt.Println("\nバニラ！")
}

/*
* Stage2:Stage1と同じ
*/
func stage2(){
	stage1()
}

/*
* 「バーニラ　バニラ　バーニラ」文字列完成で「求人」
*/
func stage3(){
	//x:乱数の格納, sequence:直近3回の乱数の履歴
	x, sequence := 0, make([]int,3)

	for reflect.DeepEqual(sequence, SequenceStage3) == false{
		x = rand.Intn(len(StrSecondCall))
		fmt.Print(StrSecondCall[x]," ")
		if len(sequence) == len(SequenceStage3){
			sequence = sequence[1:]
		}
		sequence = append(sequence, x+1)
	}
	
	fmt.Println("\n求人！")
}

/*
* 「バーニラ　バニラ」文字列完成で「高収入」
*/
func stage4(){
	//x:乱数の格納, sequence:直近2回の乱数の履歴
	x, sequence := 0, make([]int,2)
	fmt.Println(sequence)

	for reflect.DeepEqual(sequence, SequenceStage4) == false{
		x = rand.Intn(len(StrSecondCall))
		fmt.Print(StrSecondCall[x]," ")
		if len(sequence) == len(SequenceStage4){
			sequence = sequence[1:]
		}
		sequence = append(sequence, x+1)
	}
	
	fmt.Println("\n高〜収〜入〜！")
}

func stage5(){
	stage3()
}

func stage6(){
	//x:乱数の格納, sequence:直近2回の乱数の履歴
	x, sequence := 0, make([]int,2)

	for reflect.DeepEqual(sequence, SequenceStage6) == false{
		x = rand.Intn(len(StrSecondCall))
		fmt.Print(StrSecondCall[x]," ")
		if len(sequence) == len(SequenceStage6){
			sequence = sequence[1:]
		}
		sequence = append(sequence, x+1)
	}
	
	fmt.Println("\nアルバイト〜！")
}
