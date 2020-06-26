package main
import ("fmt"
	"math")
func main(){
	var a,vo,so,t float64
	fmt.Println("Enter value for acceleration:")
	fmt.Scan(&a)
	fmt.Println("Enter value for initial velocity:")
	fmt.Scan(&vo)
	fmt.Println("Enter value for initial displacement:")
	fmt.Scan(&so)
	fmt.Println("Enter value for time:")
	fmt.Scan(&t)
	fn:=GenDisplaceFn(a,vo,so)
	fmt.Println("Linear Displacement = ",fn(t))
}

func GenDisplaceFn(a,vo,so float64) func (t float64) float64{
	f:=func (t float64) float64{
		T:=math.Pow(t,2)
		s:=(0.5*a*T)+(vo*t)+so
		return(s)
		}
	return(f)
}