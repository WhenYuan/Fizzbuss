package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
	func doEvery(d time.Duration, f func(time.Time)) {
		for x := range time.Tick(d) {
			f(x)
		}
	}
*/
func main() {

	t0 := time.Now()

	a := make([]string, 0, 100)
	for i := 0; i < 100; i++ {
		var s string
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if s == "" {
			s = strconv.Itoa(i)
		}
		a = append(a, s)
	}
	fmt.Println("Hello and welcome to a game of fizzbuzz! In this game, you will be asked to type 1-100 and if the number is divisible by 3, then type Fizz, if it is divisible by 5, then type Buzz, if it is dvisible by both numbers they type FizzBuzz. You have 3 lives, if you did not input the correct answer, then you will lose one live, if all 3 lives are gone, then you die. Your goal is too finish as fast as you can. Enjoy!(Enter G for new game)")

	var b string

	/*func scam(t time.Time)  {
		fmt.Scanln(&b)
	}
	*/

	var lives int = 2
Outerloop:

	for v := 1; v < 102; v++ {
		//time.Sleep(5 * time.Second)
		//fmt.Println(a[v])
		fmt.Scanln(&b)
		if b == "G" {
			main()
			break //Outerloop

		} else if v == 100 {
			t2 := time.Now()
			fmt.Println("Congratulations!!!")
			fmt.Println("*     *                                                               ***  ***  ***")
			fmt.Println(" *   *                                                                ***  ***  ***")
			fmt.Println("  * *  —————    *      *       *              *  -----    -------     ***  ***  ***")
			fmt.Println("   *  *     *   *      *        *     **     *  *     *  *       *     *    *    *")
			fmt.Println("   *  *     *   *      *         *   *  *   *   *     *  *       *     ")
			fmt.Println("   *   -----     ------           * *    * *     -----   *       *     *    *    *")
			fmt.Println("Score:100")
			fmt.Printf("Time taken:%v \n", t2.Sub(t0))
			fmt.Println("lives remaining:", lives)
			var c string

			for {
				fmt.Print("Enter G to continue, N to not continue:")
				fmt.Scanln(&c)
				switch c {
				case "G":
					main()
					break Outerloop
				case "N":
					break Outerloop
				default:
					fmt.Print("invalid operation")
				}
			}
		} else if b == a[v] {
			continue
		} else if lives == 0 {
			fmt.Println("YOU DIED!!!")

			t1 := time.Now()

			fmt.Println("Score:", v)
			fmt.Printf("Time taken:%v \n", t1.Sub(t0))
			var C string

			for {
				fmt.Print("Enter G for new game, N for no game:")
				fmt.Scanln(&C)
				switch C {
				case "G":
					main()
					break Outerloop
				case "N":
					break Outerloop
				default:
					fmt.Print("invalid operation")
				}

			}
			//else if b == C
		} else {
			fmt.Println("Oh no, you lost a live! Be carful next time!")
			lives = lives - 1
			continue
		}
	}

}
