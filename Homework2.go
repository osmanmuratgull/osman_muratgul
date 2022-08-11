package main

import "fmt"

func main() {

	fmt.Println("PROGRAM START")

	var numberPatients [7]int = [7]int{25, 86, 45, 10, 20, 58, 78}

	totalNumber(numberPatients)

}

func totalNumber(numberPatients [7]int) {

	var maximumPatient int = 200

	var totalPatients int = 10

	for j := 0; j < len(numberPatients); j++ {

		var nextPatients int = numberPatients[j]

		if maximumPatient == totalPatients {

			break

			if (maximumPatient - totalPatients) >= nextPatients {

				totalPatients += nextPatients

				fmt.Println(totalPatients)

			} else {

				fmt.Println("Number Of Patients")

			}

			if totalPatients < 40 {

				fmt.Println("PLEASE REGISTER PATIENT")

			}

		}

	}

}
