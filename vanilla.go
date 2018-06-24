package main

import(
	"fmt"
	"math/rand"
	"time"
	"reflect"
)

StrCall := []string{"V!","a!","n!","i!","l!","la!"}

func main(){
	rand.Seed(time.Now().UnixNano())
	stage1()
	stage2()
	stage3()
	stage4()
	stage5()
	stage6()
}

/*
* Stage1:"V!a!n!i!l!la!"文字列完成で「バニラ！」
*/
func stage1(){
	fmt.Println("[Stage:1]")

	x, stage := 0, 0

	//Vanilla文字列完成シーケンス
	for stage != len(StrCall){
		x = rand.Intn(len(StrCall))
		fmt.Print(StrCall[x]," ")
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
	fmt.Println("[Stage:2]")
	stage1()
}

/*
* 「バーニラ　バニラ　バーニラ」文字列完成で「求人」
*/
func stage3(){
	str := []string{"バーニラ","バニラ","バニラで"}
	//x:乱数の格納, sequence:直近3回の乱数の履歴
	x, sequence := 0, make([]int,3)
	SequenceStage3 := []int{0, 1, 0}

	for reflect.DeepEqual(sequence, SequenceStage3) == false{
		x = rand.Intn(len(str))
		fmt.Print(str[x]," ")
		if len(sequence) == len(SequenceStage3){
			sequence = sequence[1:]
		}
		sequence = append(sequence, x)
	}
	
	fmt.Println("\n求人！")
}

/*
* 「バーニラ　バニラ」文字列完成で「高収入」
*/
func stage4(){
	str := []string{"バーニラ","バニラ","バニラで"}
	//x:乱数の格納, sequence:直近2回の乱数の履歴
	x, sequence := 0, make([]int,2)
	SequenceStage4 := []int{0, 1}

	for reflect.DeepEqual(sequence, SequenceStage4) == false{
		x = rand.Intn(len(str))
		fmt.Print(str[x]," ")
		if len(sequence) == len(SequenceStage4){
			sequence = sequence[1:]
		}
		sequence = append(sequence, x)
	}
	
	fmt.Println("\n高〜収〜入〜！")
}

func stage5(){
	stage3()
}

func stage6(){
	str := []string{"バーニラ","バニラ","バニラで"}
	//x:乱数の格納, sequence:直近2回の乱数の履歴
	x, sequence := 0, make([]int,2)
	SequenceStage6 := []int{0, 2}

	for reflect.DeepEqual(sequence, SequenceStage6) == false{
		x = rand.Intn(len(str))
		fmt.Print(str[x]," ")
		if len(sequence) == len(SequenceStage6){
			sequence = sequence[1:]
		}
		sequence = append(sequence, x)
	}
	
	fmt.Println("\nアルバイト〜！")
}
