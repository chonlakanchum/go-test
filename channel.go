package main

import (
	"fmt"
	"time"
)

func buyGlassesAtSevenEleven() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อแว่น : ที่เซเว่น : เสร็จแล้ว")
}
func buyWatchAtCentral() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อนาฬิกา : ที่เซ็นทรัล : เสร็จแล้ว")
}
func buyFruitAtSiamParagon() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อผลไม้ : ที่สยามพารากอน : เสร็จแล้ว")
}
func buyCarAtToyota() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อรถ : ที่ศุนย์โตโยต้า : เสร็จแล้ว")
}
func sendToMisterA(message chan<- string) { //สร้าง Function ส่งหา A
	time.Sleep(1 * time.Second)
	message <- "กำลังส่งของให้ นาย A" // นำค่าเข้าท่อ Channel
}
func main() {
	start := time.Now()     // เริ่มจับเวลาตอนโปรแกรมทำงาน
	ch := make(chan string) // สร้างท่อ Channel เอาไว้ส่งข้อมูล
	go buyGlassesAtSevenEleven()
	go buyWatchAtCentral()
	go sendToMisterA(ch) // ส่งท่อ ch เข้าไปใน Function ที่ใช้ Go Routine
	buyFruitAtSiamParagon()
	buyCarAtToyota()

	messageFromMisterB := <-ch // ค่าจากท่อ Channel จะออกตรงนี้
	if messageFromMisterB == "กำลังส่งของให้ นาย A" {
		fmt.Println("นาย A ได้รับของแล้ว")
		fmt.Println("ใช้เวลาในการ Run ทั้งสิ้น : ", time.Since(start), "วินาที")
	}

}
