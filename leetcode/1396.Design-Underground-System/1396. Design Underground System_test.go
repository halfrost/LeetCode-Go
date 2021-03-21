package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem(t *testing.T) {
	undergroundSystem := Constructor()
	undergroundSystem.CheckIn(45, "Leyton", 3)
	undergroundSystem.CheckIn(32, "Paradise", 8)
	undergroundSystem.CheckIn(27, "Leyton", 10)
	undergroundSystem.CheckOut(45, "Waterloo", 15)
	undergroundSystem.CheckOut(27, "Waterloo", 20)
	undergroundSystem.CheckOut(32, "Cambridge", 22)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime("Paradise", "Cambridge")) // return 14.00000. There was only one travel from "Paradise" (at time 8) to "Cambridge" (at time 22)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime("Leyton", "Waterloo"))    // return 11.00000. There were two travels from "Leyton" to "Waterloo", a customer with id=45 from time=3 to time=15 and a customer with id=27 from time=10 to time=20. So the average time is ( (15-3) + (20-10) ) / 2 = 11.00000
	undergroundSystem.CheckIn(10, "Leyton", 24)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime("Leyton", "Waterloo")) // return 11.00000
	undergroundSystem.CheckOut(10, "Waterloo", 38)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime("Leyton", "Waterloo")) // return 12.00000

	undergroundSystem.CheckIn(10, "Leyton", 3)
	undergroundSystem.CheckOut(10, "Paradise", 8)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime("Leyton", "Paradise")) // return 5.00000
	undergroundSystem.CheckIn(5, "Leyton", 10)
	undergroundSystem.CheckOut(5, "Paradise", 16)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime("Leyton", "Paradise")) // return 5.50000
	undergroundSystem.CheckIn(2, "Leyton", 21)
	undergroundSystem.CheckOut(2, "Paradise", 30)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime("Leyton", "Paradise")) // return 6.66667
}
