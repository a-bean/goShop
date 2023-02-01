package main

import "fmt"

/*
 在小明的学校，每一年开学都会发教材，
主要包括语文书、数学书、英语书，还有各种练习试卷。
这一天，小明去领了三本教材，分别是语文书、数学书和英语书，老师忙不过来，指定某个同学去发书，
同学们都去这个同学这里去领书。这个同学就是工厂。
*/

type Book interface {
	Name() string
}

type Paper interface {
	Name() string
}

type chineseBook struct {
	name string
}

type chinesePaper struct {
	name string
}

func (cb *chineseBook) Name() string {
	return cb.name
}

type mathBook struct {
	name string
}

func (mb *mathBook) Name() string {
	return mb.name
}

type englishBook struct {
	name string
}

func (eb *englishBook) Name() string {
	return eb.name
}

// person具体指定的是某个类型的人，我现在想抽象出一个角色来，这个角色就叫发书人
type Person struct{}

// 发书人
type Assigner interface {
	GetBook(name string) Book
	GetPaper(string) Paper
}

type assigner struct{}

func (a *assigner) GetBook(name string) Book {
	if name == "语文书" {
		return &chineseBook{name: "语文书"}
	} else if name == "数学书" {
		return &mathBook{name: "数学书"}
	} else if name == "英语书" {
		return &englishBook{name: "英语书"}
	}
	return nil
}

type chineseBookAssigner struct{}

func (cba *chineseBookAssigner) GetBook(name string) Book {
	if name == "语文书" {
		return &chineseBook{
			name: "语文书",
		}
	}
	return nil
}

// 责任链模式 - 行为模式
/*
例如，采购审批流程、请假流程等。公司员工请假，可批假的领导有部门负责人、副总经理、总经理等，但每个领导能批准的天数不同，
员工必须根据需要请假的天数去找不同的领导签名，也就是说员工必须记住每个领导的姓名、电话和地址等信息，这无疑增加了难度。
员工只需要发起一个请求， 我不管应该找谁来负责， 交给我的leader， 由leader自己去找应该发给谁， 职责链
*/
func main() {
	//暴露出了chineseBook， 这个实例化过程简单， 实际开发中这个创建过程可能很复杂， 我现在有个结构体，里边存储了底层的dbclient，redis的client
	var a chineseBookAssigner
	fmt.Println(a.GetBook("语文书").Name())

	//简答工厂，抽象工厂
}
