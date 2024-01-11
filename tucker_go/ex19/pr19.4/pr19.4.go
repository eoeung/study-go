package main

import "time"

// 택배 회사 구조체 선언
type Courier struct {
	Name string
}

// 물품 구조체 선언
type Product struct {
	Name  string
	Price int
	ID    int
}

// 소포 구조체 선언
type Parcel struct {
	Pdt           *Product
	ShippedTime   time.Time
	DeliveredTime time.Time
}

// *Courier 타입으로 SendProduct() 메서드 정의
func (c *Courier) SendProduct(p *Product) *Parcel {
	parcel := &Parcel{}
	parcel.ShippedTime = time.Now()
	parcel.Pdt = p

	return parcel
}

// *Parcel 타입으로 Delivered() 메서드 정의
func (p *Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()
	return p.Pdt
}

func main() {

}
