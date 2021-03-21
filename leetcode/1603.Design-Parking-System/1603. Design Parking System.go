package leetcode

type ParkingSystem struct {
	Big    int
	Medium int
	Small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Big:    big,
		Medium: medium,
		Small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		{
			if this.Big > 0 {
				this.Big--
				return true
			}
			return false
		}
	case 2:
		{
			if this.Medium > 0 {
				this.Medium--
				return true
			}
			return false
		}
	case 3:
		{
			if this.Small > 0 {
				this.Small--
				return true
			}
			return false
		}
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
