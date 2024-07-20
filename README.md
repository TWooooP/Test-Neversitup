# Test-Neversitup
## Table of Content



## Table of Content
- [Introduction](#introduction)
- [Structure Project](#structure-project)
- [Minipulate](#minipulate)
- [Odd Number](#odd-number)
- [Count Smily](#count-smily)
- 


## Introduction
จะเป็นการเขียนโปรแกรมภาษา Go บน Goland (Jetbrains)
โดยใช้ `go mod init example.com/m`
และ GOROOT ( go1.20 )

## Structure Project
```go
Test-Neversitup/
│
├── delivery/
│   └── main.go
├── environment/
│   └── main.go
├── model/
│   └── odd_number.go
├── repository/
│   └── main.go         
├── usecase/
│       │── business/
│       │   │── manipulation/
│       │   │     ├── manipulate.go
│       │   │     ├── manipulate_test.go
│       │   │── odd_number/
│       │   │     ├── odd_number.go
│       │   │     ├── odd_number_test.go
│       │   │── manipulation/
│       │   │     ├── smily.go
│       │   │     ├── smily_test.go
│       │
│       ├── port/     
│       │
│── go.mod
│── README.md
```


## Minupulate
จะเป็นการสลับข้อมูลของชุดสตริงที่อยู่ในอาร์เรย์เพื่อหาความเป็นไปได้ที่แสดงออกมาโดยไม่ซ้ำกัน
โดยจะสร้างเป็นฟังก์ชันไว้ที่ไฟล์ `manupulate.go`
ในไฟล์นี้จะเป็นฟังก์ชัน
- ทำการสร้างตัวแปรเพื่อเก็บจำนวนที่ไม่สามารถสลับกันได้โดยไม่ซ้ำกัน
- นำข้อมูลทอาร์เรย์่ที่ได้มาแยกออกเอาข้อมูลสตริงข้างในมาทำเป็นจำนวนอาร์เรย์ int32 โดยใช้ rune ในการช่วยระบุตำแหน่งของอักษรที่ต้องการ
- แล้วสร้างตัวแปรขึ้นมาเพื่อสร้างfuncในการทำงานซ้ำ โดยจะมีการทำ for loop ในการสลับตำแหน่งของข้อมูล
```go
duplicate := make(map[string]struct{})
for _, data := range text {

    strData := []rune(data)

    var Shuffle func(int, int)

    Shuffle = func(l, r int) {
        if l == r {
            duplicate[string(strData)] = struct{}{}
        } else {
            for i := l; i <= r; i++ {
            strData[l], strData[i] = strData[i], strData[l]
            Shuffle(l+1, r)
            strData[l], strData[i] = strData[i], strData[l]
            }
        }
}
Shuffle(0, len(strData)-1)
}
```
- สุดท้ายจะทำการนำข้อมูลทั้งหมดที่ได้มาเพื่อใส่ในอาร์เรย์สตริง
```go
var result []string
for dataRes := range duplicate {
    result = append(result, dataRes)
}
```
โดยทั้งหมดนี้จะถูกใช้ในไฟล์ `manupulate_test.go`
```go
func TestCheckOddNumber(t *testing.T) {
    //data := []int{7}
    //data := []int{1, 1, 2}
    //data := []int{0, 1, 0, 1, 0}
    //data := []int{0, 1, 0, 1, 0}
    data := []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}

    result := FindOddNumber(data)

    fmt.Println(result)
}
```











## Odd Number
จะเป็นการนับและหาเลขจำนวนที่แสดงในอาร์เรย์ จำนวนที่นับได้ในอาร์เรย์เป็นเลขคี่

โดยจะสร้างเป็นฟังก์ชันไว้ที่ไฟล์ `odd_number.go`
ในไฟล์นี้จะเป็นฟังก์ชัน
- นำข้อมูลที่เป็นอาร์เรย์ Int มา for loop เพื่อนำ Data ออกมาเช็ค
- มีการสร้าง model structure เข้ามาช่วยในการเก็บค่าของข้อมูล
```go
type CollectData struct {
	Value int `json:"value"`
	Count int `json:"count"`
}
```
- มี loop ในการเช็คข้อมูลที่อยู่ในตัวแปร model structure ที่สร้างไว้ และเข้าไปอัพเดตข้อมูลหากมีข้อมูลในอาร์เรย์แล้วและสร้างตัวแปรมาเพื่อเก็บค่าว่ามีการพบค่าในอาร์เรย์ที่เก็บหรือยังถ้าไม่มีให้สร้างข้อมูลชุดใหม่
```go
var ResultData []CollectData

for _, data := range text {
    var foundData bool
        ///check data from array
        for i, collect := range ResultData {
        ///if have data in array update count data
            if data == collect.Value {
                ResultData[i].Count += 1
                foundData = true
            }
        }
        //if not have data add new data
        if foundData == false {
            result := CollectData{
                Value: data,
                Count: 1,
            }
            ResultData = append(ResultData, result)
        }
}
```

- สุดท้ายจะเป็นการเช็คจำนวนของข้อมูลที่มีอยู่ที่มีจำนวนเป็นเลขคี่โดยการนำมาหาร2 และเอาจำนวนที่เป็นจำนวนเต็ม
```go
var FinishData int
/// check data from array ResultData have count is odd number
for _, jjj := range ResultData {
    if jjj.Count%2 != 0 {
    FinishData = jjj.Value
    }
}
```
โดยทั้งหมดนี้จะถูกใช้ในไฟล์ `odd_number_test.go`
```go
func TestCheckOddNumber(t *testing.T) {
    //data := []int{7}
    //data := []int{1, 1, 2}
    //data := []int{0, 1, 0, 1, 0}
    //data := []int{0, 1, 0, 1, 0}
    data := []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}

    result := FindOddNumber(data)

    fmt.Println(result)
}
```

## Count Smily
จะเป็นการนับจำนวนหน้ายิ้มตามลักษณะที่ต้องการในอาร์เรย์สตริง
เงื่อนไข :
- ตา  = :  ;
- จมูก = -  ~
- ปาก = )  D

โดยจะสร้างเป็นฟังก์ชันไว้ที่ไฟล์ `smily.go`
ในไฟล์นี้จะเป็นฟังก์ชัน
- จะมีการสร้างตัวแปรเพื่อเก็บจำนวนที่นับหน้ายิ้มได้ได้
- มีการใช้ for loop เพื่อแยกข้อมูล string ออกมาเป็นชุดๆ แล้วใช้ strings.split ทำการแยกตัวอักษรออกมาเป็นอาร์เรย์ เพื่อนำมาเช็คตา จมูก และหน้ายิ้ม
- ใช้ if ในการเช็คข้อมูลว่ามีลักษณะตรงตามที่ต้องการไหมหากให้ให้นับจำนวนเข้าไปในตัวแปรที่เราสร้างไว้

```go
var CountResultData int
for _, data := range text {
	//separate string
	sss := strings.Split(data, "")
	if ((sss[0] == ":" || sss[0] == ";") && (sss[1] == ")" || sss[1] == "D")) ||
		(sss[1] == "-" || sss[1] == "~" && (sss[1] == ")" || sss[1] == "D")) {
		CountResultData += 1
		}
}
```

โดยทั้งหมดนี้จะถูกใช้ในไฟล์ `smily_test.go`
```go
func TestCheckSmile(t *testing.T) {
    //data := []string{":)", ";(", ";}", ":-D"}
    //data := []string{";D", ":-(", ":-)", ";~)"}
    data := []string{";]", ":[", ";*", ":$", ";-D"}
    result := CountSmilyFace(data)

    fmt.Println(result)
}
```
