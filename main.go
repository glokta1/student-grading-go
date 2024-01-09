package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func (s student) CalculateMeanScore() float32 {
	return float32(s.test1Score+s.test2Score+s.test3Score+s.test4Score) / 4
}

func (s student) String() string {
	return fmt.Sprintf("Name: %s %s\nUniversity: %s\nScores: [%d, %d, %d, %d]", s.firstName, s.lastName, s.university, s.test1Score, s.test2Score, s.test3Score, s.test4Score)
}

func parseCSV(filePath string) []student {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	defer csvFile.Close()
	fileScanner := bufio.NewScanner(csvFile)
	// skip CSV header
	fileScanner.Scan()

	// TODO: dynamically input length of students slice
	students := make([]student, 0, 30)
	for fileScanner.Scan() {
		row := fileScanner.Text()

		// split row by commas
		studentInfo := strings.Split(row, ",")
		firstName, lastName, university := studentInfo[0], studentInfo[1], studentInfo[2]

		// TODO: cleanup error handling here
		test1Score, err := strconv.Atoi(studentInfo[3])
		if err != nil {
			fmt.Println(err)
		}
		test2Score, err := strconv.Atoi(studentInfo[4])
		if err != nil {
			fmt.Println(err)
		}
		test3Score, err := strconv.Atoi(studentInfo[5])
		if err != nil {
			fmt.Println(err)
		}
		test4Score, err := strconv.Atoi(studentInfo[6])
		if err != nil {
			fmt.Println(err)
		}
		student := student{firstName, lastName, university, test1Score, test2Score, test3Score, test4Score}
		students = append(students, student)
	}

	if fileScanner.Err() != nil {
		fmt.Println(fileScanner.Err().Error())
	}

	return students
}

func calculateGrade(students []student) []studentStat {
	studentStats := make([]studentStat, 0, len(students))
	for _, student := range students {
		meanScore := student.CalculateMeanScore()
		var grade string
		switch {
		case meanScore < 35:
			grade = "F"
		case meanScore >= 35 && meanScore < 50:
			grade = "C"
		case meanScore >= 50 && meanScore < 70:
			grade = "B"
		case meanScore >= 70:
			grade = "A"
		}

		stat := studentStat{student, meanScore, Grade(grade)}
		studentStats = append(studentStats, stat)
	}

	return studentStats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	var topper studentStat
	var maxFinalScore float32 = 0.0
	for _, studentStat := range gradedStudents {
		if studentStat.finalScore > float32(maxFinalScore) {
			maxFinalScore = studentStat.finalScore
			topper = studentStat
		}
	}

	return topper
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	universities := make(map[string][]studentStat)
	for _, stat := range gs {
		universities[stat.student.university] = append(universities[stat.student.university], stat)
	}

	toppers := make(map[string]studentStat)
	for uni, studentsFromUni := range universities {
		toppers[uni] = findOverallTopper(studentsFromUni)
	}

	return toppers
}

func main() {
	students := parseCSV("grades.csv")
	fmt.Println(students[0])
}
