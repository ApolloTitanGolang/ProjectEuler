/*
	Some examples from https://projecteuler.net/ in go
	June 2016

 */

package main

import (
	"fmt"
	"strconv"
	"time"
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







