package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)


type elem struct{
	value uint8
	times int64
	position int64
	width int64
}

var dict []elem
var enScripts [3]int64


func calTimes(scripts string) map[uint8]int64{
	var elem int64

	Freq := make(map[uint8]int64)
	elem = 1

	for i:=0;i<len(scripts);i++{
		_,ok := Freq[scripts[i]]

		if ok {
			Freq[scripts[i]] = Freq[scripts[i]] + elem
		} else{
			Freq[scripts[i]] = elem
		}
	}
	return Freq
}

func buildDict(times map[uint8]int64)int64{
	var i int64 =0
	var position int64 = 0
	for one := range times {
		dict = append(dict,elem{one,times[one],0,times[one]})
		i++
	}
	lens := i
	sort.Slice(dict, func(i, j int) bool {
		return dict[i].times > dict[j].times  // 降序
		// return ss[i].Value > ss[j].Value  // 升序
	})

	for i=0;i<lens;i++{
		dict[i].position = position
		position = dict[i].times + position
	}

	return lens
}

func convertToBin(num int64) string {
	s := ""

	if num == 0 {
		return "0"
	}

	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ;num > 0 ; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(int(lsb)) + s
	}
	return s
}

func powerf(x int64, n int64) int64 {
	if n == 0 {
		return 1
	} else {
		return x * powerf(x, n-1)
	}
}

func encode(scripts string){

	var frac int64 = 0
	var width int64 = 1
	var j int



	for i:=0;i<len(scripts);i++{
		for j=0;scripts[i] != dict[j].value;j++{}
		frac = frac * int64(len(scripts)) + width*dict[j].position
		width = width*dict[j].width
	}
	//fmt.Println(frac,"||",width)
	fmt.Println("fraction is from: ",frac,"/",powerf(int64(len(scripts)),int64(len(scripts))),"to",frac+width,"/",powerf(int64(len(scripts)),int64(len(scripts))))
	fmt.Println(frac,":",convertToBin(frac))
	fmt.Println(powerf(int64(len(scripts)),int64(len(scripts))),":",convertToBin(powerf(int64(len(scripts)),int64(len(scripts)))))
	N := math.Ceil(math.Log2(float64(powerf(int64(len(scripts)),int64(len(scripts)))/width)))

	binFrac := math.Ceil(float64(frac)/math.Pow(float64(len(scripts)),float64(len(scripts)))*math.Pow(2.0,N))

	fmt.Println(binFrac,":",convertToBin(int64(binFrac)))
	fmt.Println(math.Pow(2,N),":",convertToBin(int64(math.Pow(2,N))))
	enScripts[0] = int64(binFrac)
	enScripts[1] = int64(math.Pow(2,N))
	enScripts[2] = int64(len(scripts))

	enScript := math.Float64bits(float64(enScripts[0])/float64(enScripts[1]))
	fmt.Println("编码后的文本：")
	fmt.Printf("%b\n",enScript)

}

func decode(){
	var i int64
	var j int64
	var deScripts =  ""
	var slen int64
	var start float64 = 0
	var width float64 = 1
	var enScript = float64(enScripts[0])/float64(enScripts[1])

	slen = enScripts[2]
	for i = 0;i<slen;i++{
		for j = 0;(enScript<start + width*float64(dict[j].position)/float64(slen))||(enScript>=start + width*float64(dict[j].position+dict[j].width)/float64(slen));j++{}
		deScripts = deScripts + string(dict[j].value)
		start = start + width*float64(dict[j].position) / float64(slen)
		width = width * float64(dict[j].width) / float64(slen)
	}
	fmt.Println("解码后的文本：")
	fmt.Println(deScripts)
}




func main(){
	var scripts string

	fmt.Println("请输入要编码压缩的文本：")

	inputReader := bufio.NewReader(os.Stdin)
	scripts,_ = inputReader.ReadString('\n')

	times := calTimes(scripts)
	fmt.Println(times)

	slen := buildDict(times)
	fmt.Println(slen)
	fmt.Println(dict)
	encode(scripts)
	decode()

}
