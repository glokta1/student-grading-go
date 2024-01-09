package main

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

func parseCSV(filePath string) []student {
	return nil
}

func calculateGrade(students []student) []studentStat {
	studentStats := make([]studentStat, 0, len(students))
	for _, student := range students {
		score := float32((student.test1Score + student.test2Score + student.test3Score + student.test4Score) / 4)
		var grade string
		switch {
		case score < 35:
			grade = "F"
		case score >= 35 && score < 50:
			grade = "D"
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
	return studentStat{}
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	return nil
}
