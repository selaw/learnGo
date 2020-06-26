package main
import (
	"fmt"
)

func BuatDistOri(o_x, o_y float64)
		func (float64, float64) float64 {
			fn:= func (x, y float64) float64{
				return math.Sqrt(math.Pow(x - o_x, 2) + math.Pow(y - o_y, 2))
			}
		return fn
		}

func main() {
	Dist1 := BuatDistOri(0,0)
	Dist2 := BuatDistOri(2,2)

	fmt.Println(Dist1(2,2))
	fmt.Println(Dist2(2,2))
}