// You can edit this code!

// Click here and start typing.

package main

import "fmt"

func main() {

	fmt.Println("PROGRAM START")

	var NumberOfPatients [7]int = [7]int{25, 86, 45, 10, 20, 58, 78}

	TotalNumberOfPatientsİnTheHospital(NumberOfPatients)

}

func TotalNumberOfPatientsİnTheHospital(NumberOfPatients [7]int) {

	var MaximumPatient int = 200

	var TotalPatients int = 10

	for j := 0; j < 7; j++ {

		var NextPatients int = NumberOfPatients[j]

		{

			if MaximumPatient == TotalPatients {

				break

			}

		}

		{

			if (MaximumPatient - TotalPatients) >= NextPatients {

				TotalPatients += NextPatients

				fmt.Println(TotalPatients)

			} else {

				fmt.Println("Number Of Patients")

			}

			if TotalPatients < 40 {

				fmt.Println("PLEASE REGISTER PATIENT")

			}

		}

	}

}

