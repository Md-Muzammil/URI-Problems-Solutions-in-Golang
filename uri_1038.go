package main
import (
	"fmt"
)


func main(){
    var X, Y int
    var price = 0 float64

    fmt.Scanf("%d %d", &X, &Y)
    if (X == 1)
    {
        price  = (float) (4.00 * Y)
    }
    else if (X == 2)
    {
        price  = (float) (4.50 * Y)
    }
    else if (X == 3)
    {
        price  = (float) (5.00 * Y)
    }
    else if (X == 4)
    {
        price  = (float) (2.00 * Y)
    }
    else if (X == 5)
    {
        price  = (float) (1.50 * Y)
    }

}
