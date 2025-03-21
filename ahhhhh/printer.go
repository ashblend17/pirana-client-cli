package ahhhhh

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// test for semwise table
func PrintSGPATable(data StudentData) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Roll No", "Semester", "SGPA"})

	// Iterate over the SGPA data
	for roll, entries := range data.SemwiseSGPA {
		for _, entry := range entries {
			table.Append([]string{roll, entry.Semester, entry.SGPA})
		}
	}
	fmt.Println("Semester-wise SGPA:")
	table.Render()
}

func PrintPass(password string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Password"})
	table.Append([]string{password})
	table.Render()
}

// test for course wise data
func PrintCourseTable(data StudentData) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Roll No", "Semester", "Course Code", "Course Name", "Grade"})

	// Iterate over the course data
	for roll, courses := range data.CourseData {
		for _, course := range courses {
			table.Append([]string{roll, course.Semester, course.Code, course.Name, course.Grade})
		}
	}
	fmt.Println("Course Data:")
	table.Render()
}

func PrintGroupedTable(data StudentData) {
	grouped := make(map[string]*GroupData)

	// Group courses (iterate over each roll’s course data).
	for _, courses := range data.CourseData {
		for _, course := range courses {
			if _, exists := grouped[course.Semester]; !exists {
				grouped[course.Semester] = &GroupData{}
			}
			grouped[course.Semester].Courses = append(grouped[course.Semester].Courses, course)
		}
	}

	// Assign SGPA values to the corresponding semester group.
	for _, sgpaEntries := range data.SemwiseSGPA {
		for _, sgpaEntry := range sgpaEntries {
			if _, exists := grouped[sgpaEntry.Semester]; !exists {
				grouped[sgpaEntry.Semester] = &GroupData{}
			}
			grouped[sgpaEntry.Semester].SGPA = sgpaEntry.SGPA
		}
	}

	// Sort the semesters numerically.
	var semesters []int
	for semStr := range grouped {
		semInt, err := strconv.Atoi(semStr)
		if err != nil {
			continue
		}
		semesters = append(semesters, semInt)
	}
	sort.Ints(semesters)

	// Create and configure the table.
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Semester", "Course", "Grade", "SGPA"})

	// For each sorted semester, print rows for each course.
	for _, sem := range semesters {
		semStr := strconv.Itoa(sem)
		group := grouped[semStr]
		if len(group.Courses) == 0 {
			// In case there are no courses, still print the semester and SGPA.
			row := []string{semStr, "", "", group.SGPA}
			table.Append(row)
		} else {
			for i, course := range group.Courses {
				var semDisplay, sgpaDisplay string
				if i == 0 {
					semDisplay = semStr
					sgpaDisplay = group.SGPA
				}
				// Format the course column as "code - name".
				courseDisplay := fmt.Sprintf("%s - %s", course.Code, course.Name)
				row := []string{semDisplay, courseDisplay, course.Grade, sgpaDisplay}
				table.Append(row)
			}
		}
	}
	table.Render()
}

// printCombinedTable prints one table that combines course data and matching SGPA.
// For each roll number, it prints each course record and looks up the SGPA (by semester) if available.
func printCombinedTable(data StudentData, rollNumbers []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Roll No", "Semester", "Course Code", "Course Name", "Grade", "SGPA"})

	sort.Strings(rollNumbers)
	for _, roll := range rollNumbers {
		courses, okC := data.CourseData[roll]
		sgpaEntries, okS := data.SemwiseSGPA[roll]
		// Build a lookup: semester -> sgpa.
		sgpaMap := make(map[string]string)
		if okS {
			for _, entry := range sgpaEntries {
				sgpaMap[entry.Semester] = entry.SGPA
			}
		}
		if !okC || len(courses) == 0 {
			// If there are no course records, print SGPA rows only.
			for sem, sgpa := range sgpaMap {
				table.Append([]string{roll, sem, "", "", "", sgpa})
				roll = "" // only print roll number on first row
			}
		} else {
			// For each course, print its details and, if available, the corresponding SGPA.
			for i, course := range courses {
				rollDisplay := ""
				if i == 0 {
					rollDisplay = roll
				}
				sgpaDisplay := ""
				if sg, exists := sgpaMap[course.Semester]; exists {
					sgpaDisplay = sg
				}
				table.Append([]string{rollDisplay, course.Semester, course.Code, course.Name, course.Grade, sgpaDisplay})
			}
		}
	}
	table.Render()
}

// printPasswordTable prints a table of roll numbers and their password.
func printPasswordTable(data StudentData, rollNumbers []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Roll No", "Password"})

	sort.Strings(rollNumbers)
	for _, roll := range rollNumbers {
		pass, ok := data.Password[roll]
		if !ok {
			pass = ""
		}
		table.Append([]string{roll, pass})
	}
	table.Render()
}

// PrintGroupedTableByRoll groups the courses and SGPA by semester for each roll number,
// and then prints one table per roll number.
func PrintGroupedTableByRoll(data StudentData, rollNumbers []string) {
	// Iterate over each roll number provided.
	for _, roll := range rollNumbers {
		// Create a map: semester (string) -> GroupData for this roll.
		grouped := make(map[string]*GroupData)

		// Group the courses by semester.
		if courses, ok := data.CourseData[roll]; ok {
			for _, course := range courses {
				if _, exists := grouped[course.Semester]; !exists {
					grouped[course.Semester] = &GroupData{}
				}
				grouped[course.Semester].Courses = append(grouped[course.Semester].Courses, course)
			}
		}

		// Assign the SGPA value(s) for each semester.
		if sgpaEntries, ok := data.SemwiseSGPA[roll]; ok {
			for _, sgpaEntry := range sgpaEntries {
				if _, exists := grouped[sgpaEntry.Semester]; !exists {
					grouped[sgpaEntry.Semester] = &GroupData{}
				}
				// If there are multiple entries for the same semester,
				// you may choose to adjust this logic.
				grouped[sgpaEntry.Semester].SGPA = sgpaEntry.SGPA
			}
		}

		// Sort the semester keys numerically.
		var semesters []int
		for semStr := range grouped {
			semInt, err := strconv.Atoi(semStr)
			if err != nil {
				continue
			}
			semesters = append(semesters, semInt)
		}
		sort.Ints(semesters)

		// Print a header indicating which roll number is being printed.
		fmt.Printf("Roll Number: %s\n", roll)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Semester", "Course", "Grade", "SGPA"})

		// For each semester (in sorted order), append rows.
		for _, semInt := range semesters {
			semStr := strconv.Itoa(semInt)
			group := grouped[semStr]
			if group == nil {
				continue
			}
			// If there are no courses for this semester, still print a row.
			if len(group.Courses) == 0 {
				table.Append([]string{semStr, "", "", group.SGPA})
			} else {
				for i, course := range group.Courses {
					var semDisplay, sgpaDisplay string
					if i == 0 {
						semDisplay = semStr
						sgpaDisplay = group.SGPA
					}
					// Format the course column as "code - name".
					courseDisplay := fmt.Sprintf("%s - %s", course.Code, course.Name)
					table.Append([]string{semDisplay, courseDisplay, course.Grade, sgpaDisplay})
				}
			}
		}
		table.Render()
		fmt.Println() // Blank line to separate tables for different roll numbers.
	}
}

func PrintGroupedCourseTable(data StudentData, rollNumbers []string) {
	// Iterate over each roll number.
	for _, roll := range rollNumbers {
		courses, ok := data.CourseData[roll]
		if !ok || len(courses) == 0 {
			continue
		}

		// Group courses by semester.
		groupedCourses := make(map[string][]Course)
		for _, course := range courses {
			groupedCourses[course.Semester] = append(groupedCourses[course.Semester], course)
		}

		// Sort semester keys numerically.
		var semesters []int
		for semStr := range groupedCourses {
			semInt, err := strconv.Atoi(semStr)
			if err != nil {
				continue
			}
			semesters = append(semesters, semInt)
		}
		sort.Ints(semesters)

		// Print a header for this roll number.
		fmt.Printf("Roll Number: %s\n", roll)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Semester", "Course Code", "Course Name", "Grade"})

		// For each sorted semester, print the rows.
		for _, semInt := range semesters {
			semStr := strconv.Itoa(semInt)
			courseGroup := groupedCourses[semStr]
			if len(courseGroup) == 0 {
				table.Append([]string{semStr, "", "", ""})
			} else {
				for i, course := range courseGroup {
					semDisplay := ""
					if i == 0 {
						semDisplay = semStr
					}
					table.Append([]string{semDisplay, course.Code, course.Name, course.Grade})
				}
			}
		}
		table.Render()
		fmt.Println() // Blank line between roll numbers.
	}
}

func PrintGroupedSGPATable(data StudentData, rollNumbers []string) {
	// Sort the roll numbers in ascending order.
	sort.Strings(rollNumbers)

	// Iterate over each provided roll number.
	for _, roll := range rollNumbers {
		entries, ok := data.SemwiseSGPA[roll]
		if !ok || len(entries) == 0 {
			continue
		}

		// Sort the SGPA entries by semester numerically.
		sort.SliceStable(entries, func(i, j int) bool {
			// Convert semester strings to integers for proper numeric comparison.
			si, err1 := strconv.Atoi(entries[i].Semester)
			sj, err2 := strconv.Atoi(entries[j].Semester)
			if err1 != nil || err2 != nil {
				return entries[i].Semester < entries[j].Semester
			}
			return si < sj
		})

		// Print a header for the current roll number.
		fmt.Printf("Roll Number: %s\n", roll)

		// Create a new table for this roll.
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Roll No", "Semester", "SGPA"})

		// Append rows; only the first row will include the roll number.
		for i, entry := range entries {
			rollDisplay := ""
			if i == 0 {
				rollDisplay = roll
			}
			table.Append([]string{rollDisplay, entry.Semester, entry.SGPA})
		}

		// Render the table and print a blank line afterward.
		table.Render()
		fmt.Println()
	}
}

func PrintStudentTable(students []Student) {
	sort.Slice(students, func(i, j int) bool {
		return students[i].Roll < students[j].Roll
	})
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Roll"})

	for _, s := range students {
		table.Append([]string{s.Name, s.Roll})
	}
	table.Render()
}
