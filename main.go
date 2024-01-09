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

func (s student) String() string {
	return fmt.Sprintf("Name: %s %s\nUniversity: %s\nScores: [%d, %d, %d, %d]", s.firstName, s.lastName, s.university, s.test1Score, s.test2Score, s.test3Score, s.test4Score)
}

func parseCSV(filePath string) []student {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	fileScanner := bufio.NewScanner(f)
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
		score := float32(student.test1Score+student.test2Score+student.test3Score+student.test4Score) / 4
		var grade string
		switch {
		case score < 35:
			grade = "F"
		case score >= 35 && score < 50:
			grade = "C"
		case score >= 50 && score < 70:
			grade = "B"
		case score >= 70:
			grade = "A"
		}

		stat := studentStat{student, score, Grade(grade)}
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
	toppers := make(map[string]studentStat)
	for _, studentStat := range gs {
		if studentStat.finalScore > toppers[studentStat.student.university].finalScore {
			toppers[studentStat.student.university] = studentStat
		}
	}
	return toppers
}

func main() {
	students := parseCSV("grades.csv")
	fmt.Println(students[0])
}
