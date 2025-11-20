package main

import "fmt"

type Rect struct{
	height, width int
}

// value reciever
func (r Rect) Area()int{
	return r.height*r.width;
}

// pointer reciever
func (r *Rect) Scale(factor int){
		r.height*=factor;
		r.width*=factor;
}

func main() {

	var myHouse Rect;
	myHouse.height=10;
	myHouse.width=50;

	fmt.Println("Area of my house: ",myHouse.Area());
	
	myBuilding:=Rect{10,50};

	fmt.Println("my building area : ",myBuilding.Area())
	myBuilding.Scale(10);
	fmt.Println("my building area scale : ",myBuilding.Area());



}
