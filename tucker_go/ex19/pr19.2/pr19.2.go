package main

type ParkingLot struct {
	LotSize int // 총 주차 가능한 자동차 대수
}

func ParkCar(lot *ParkingLot, carSize int) {
	lot.LotSize -= carSize
}

// ParkCar 함수와 같은 동작을 하는 ParkCar() 메서드
func (p *ParkingLot) ParkCar(carSize int) {
	p.LotSize -= carSize
}

func main() {
	lot := &ParkingLot{100}
	ParkCar(lot, 10)

	lot.ParkCar(10)
}
