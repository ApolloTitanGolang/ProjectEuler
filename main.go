/*
	Some examples from https://projecteuler.net/ in go
	June 2016

 */

package main

import (
	"fmt"
	"strconv"
	"time"
	"math"
)

func main() {


	fmt.Println("Project Euler - my solutions")

	fmt.Println("\nProblem 1 (Multiples of 3 and 5)")
	fmt.Println("-----------------------------------------------------------------")
	starttime := time.Now()
	problem1()
	endtime := time.Now()
	fmt.Printf("\nDuration: %v", endtime.Sub(starttime))

	fmt.Println("\n\n\nProblem 2 (Even Fibonacci numbers)")
	fmt.Println("-----------------------------------------------------------------")
	starttime = time.Now()
	problem2()
	endtime = time.Now()
	fmt.Printf("\nDuration: %v", endtime.Sub(starttime))

	fmt.Println("\n\n\nProblem 3 (larges prime factor of 600851475143)")
	fmt.Println("-----------------------------------------------------------------")
	starttime = time.Now()
	problem3()
	endtime = time.Now()
	fmt.Printf("\nDuration: %v", endtime.Sub(starttime))

	fmt.Println("\n\n\nProblem 4 (Largest palindrome product from 3-digit numbers)")
	fmt.Println("-----------------------------------------------------------------")
	starttime = time.Now()
	problem4()
	endtime = time.Now()
	fmt.Printf("\nDuration: %v", endtime.Sub(starttime))

	fmt.Println("\n\n\nProblem 5 (Smallest multiple)")
	fmt.Println("-----------------------------------------------------------------")
	starttime = time.Now()
	// problem5()	// commented because it takes to long
	endtime = time.Now()
	fmt.Printf("\nDuration: %v", endtime.Sub(starttime))

	fmt.Println("\n\n\nProblem 6 (Sum square difference)")
	fmt.Println("-----------------------------------------------------------------")
	starttime = time.Now()
	problem6()
	endtime = time.Now()
	fmt.Printf("\nDuration: %v", endtime.Sub(starttime))

	fmt.Println("\n\n\nProblem 7 (10001st prime)")
	fmt.Println("-----------------------------------------------------------------")
	starttime = time.Now()
	problem7()
	endtime = time.Now()
	fmt.Printf("\nDuration: %v", endtime.Sub(starttime))

	fmt.Println("\n\n\nProblem 8 (Largest product in a series)")
	fmt.Println("-----------------------------------------------------------------")
	starttime = time.Now()
	problem8()
	endtime = time.Now()
	fmt.Printf("\nDuration: %v", endtime.Sub(starttime))

	fmt.Println("\n\n\nProblem 9 (Special Pythagorean triplet)")
	fmt.Println("-----------------------------------------------------------------")
	starttime = time.Now()
	problem9()
	endtime = time.Now()
	fmt.Printf("\nDuration: %v", endtime.Sub(starttime))

	fmt.Println("\n\n\nProblem 10 (Summation of primes)")
	fmt.Println("-----------------------------------------------------------------")
	starttime = time.Now()
	// problem10()		// commented because it takes to long
	endtime = time.Now()
	fmt.Printf("\nDuration: %v", endtime.Sub(starttime))

	fmt.Printf("\n\n")
}


func problem1() {
	sum := 0
	for i := 0; i < 1000; i++ {
		if (i % 3 == 0) || (i % 5 == 0) {
			sum += i
		}
	}

	fmt.Printf("The sum of all multiples of 3 and 5 are: %v", sum)
}

func problem2() {
	sum := 0
	lastVib := 0
	thisVib := 1
	tempVib := 0

	for thisVib < 4000000 {			// calculate fibonacci not exceed four million
		tempVib = lastVib
		lastVib = thisVib
		thisVib = thisVib + tempVib

		if (thisVib % 2 == 0) {		// check if even
			sum += thisVib
		}
	}

	fmt.Printf("The sum is: %v", sum)
}

func problem3() {
	primeList := make([]int, 0, 500)
	isPrime := true
	num := 1
	checkNumber := 600851475143
	sum := 1

	for num < checkNumber {
		// check if it is a prime
		for _, plv := range primeList {
			if (plv != 1) && (num % plv == 0) {
				isPrime = false
			}
		}

		if isPrime {
			primeList = append(primeList, num)
			// fmt.Println(primeList)
		}

		if (isPrime) && (checkNumber % (num*sum) == 0) {
			sum = sum * num
			fmt.Printf("\nSum: %v and value: %v", sum, num)
		}

		if (sum == checkNumber) {
			break
		}


		isPrime = true
		num++
	}
}

func problem4() {
	largestPalindrome := 0
	numberToCheck := 0
	var palindromeByteSlice []byte
	isPalindrome := true
	// x := 0
	y := 0

	for i := 0; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			if i*j > largestPalindrome {
				numberToCheck = i*j

				palindromeByteSlice = []byte(strconv.Itoa(numberToCheck))
				y = len(palindromeByteSlice)

				if y % 2 == 0 {		// in this example we check even-length values only (11, 1111, 111111)

					for pos := 0; pos < y/2; pos++ {
						if palindromeByteSlice[pos] == palindromeByteSlice[y-pos-1] {
							isPalindrome = true

						} else {
							isPalindrome = false
							break
						}
					}
					if isPalindrome {
						largestPalindrome = numberToCheck
					}

				}

			}

		}
	}

	fmt.Printf("\nThe largest palindrome: %v", largestPalindrome)
}


func problem5() {

	evenlyDivisible := false

	for number := 1; number < 1000000000; number++ {
		for div:= 1; div <=20; div++ {
			if number % div == 0 {
				evenlyDivisible = true
			} else {
				evenlyDivisible = false
				break
			}
		}

		if evenlyDivisible {
			fmt.Printf("\nThe smallest multiple: %v", number)
			break
		}
	}
}


func problem6() {
	squaresSum := 0.0
	sumNatural := 0.0
	sumSquare := 0.0

	for i := 1; i <= 100; i++ {
		squaresSum += math.Pow(float64(i), 2)
		sumNatural += float64(i)
	}

	sumSquare = math.Pow(sumNatural, 2)

	fmt.Printf("\nThe square of the sum (%v) minus the sum of the squares (%v)", sumSquare, squaresSum)
	fmt.Printf("\n= %v", int(sumSquare - squaresSum))
}


func problem7() {
	primeList := make([]int, 0, 10001)
	primeCount := 0
	number := 2
	isPrime := true

	// calculate primes
	for primeCount < 10001 {
		isPrime = true
		number++

		for _, prime := range primeList {
			if number % prime == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primeList = append(primeList, number)
			primeCount++
			// fmt.Printf("\nPrimelist: %v - PrimeCount: %v", primeList, primeCount)
		}
	}

	fmt.Printf("\nThe 10001st prime number: %v", number)
}

func problem8() {

	// get byte out of the string and convert (cast) it to an int
	// i, _ := strconv.Atoi(string(stringNumber[0]))

	greatestProduct := 1

	stringNumber :=
	"73167176531330624919225119674426574742355349194934"+
	"96983520312774506326239578318016984801869478851843"+
	"85861560789112949495459501737958331952853208805511"+
	"12540698747158523863050715693290963295227443043557"+
	"66896648950445244523161731856403098711121722383113"+
	"62229893423380308135336276614282806444486645238749"+
	"30358907296290491560440772390713810515859307960866"+
	"70172427121883998797908792274921901699720888093776"+
	"65727333001053367881220235421809751254540594752243"+
	"52584907711670556013604839586446706324415722155397"+
	"53697817977846174064955149290862569321978468622482"+
	"83972241375657056057490261407972968652414535100474"+
	"82166370484403199890008895243450658541227588666881"+
	"16427171479924442928230863465674813919123162824586"+
	"17866458359124566529476545682848912883142607690042"+
	"24219022671055626321111109370544217506941658960408"+
	"07198403850962455444362981230987879927244284909188"+
	"84580156166097919133875499200524063689912560717606"+
	"05886116467109405077541002256983155200055935729725"+
	"71636269561882670428252483600823257530420752963450"

	for startValue := 0; startValue <= 987; startValue++ {

		greatestTempProduct := 1

		for valueIndex := startValue; valueIndex <= startValue + 12; valueIndex++ {
			number, _ := strconv.Atoi(string(stringNumber[valueIndex]))
			greatestTempProduct = greatestTempProduct * number
		}

		if greatestTempProduct > greatestProduct {
			greatestProduct = greatestTempProduct
		}
	}

	fmt.Printf("\nThe largest product in a series: %v", greatestProduct)
}

func problem9() {

	// a^2 + b^2 = c^2 <> a + b + c = 1000
	// a^2 + b^2 = (1000 - a - b)^2
	// SQRT(a^2 + b^2) + a + b = 1000
	a := 1.0
	b := 1.0
	c := 1.0
	found := false;

	for a = 1; a < 1000; a++ {
		for b = 1; b < 1000; b++ {

			c = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))

			if (c + a + b) == 1000 {
				found = true
				break
			}
		}
		if found {
			fmt.Printf("1")
			break
		}
	}

	p := a * b * c

	fmt.Printf("\nTriplet for 1000 is: %v,  %v, %v", a, b, c)
	fmt.Printf("\nTriplet product: %v", int(p))
}


func problem10() {

	primes := make([]int, 0, 100)
	number := 2
	isPrime := true
	primeSum := 0

	for number < 2000000 {
		primeSum = 0
		for _, prime := range primes {
			if number % prime == 0 {
				isPrime = false
				break
			} else {
				isPrime = true
			}
		}
		// fmt.Printf("\nslice: %v", primes)

		if isPrime {
			primes = append(primes, number)
		}
		number++
	}

	for _, prime := range primes {
		primeSum += prime
	}

	fmt.Printf("\nSum of all primes below two millions: %v", primeSum)
}