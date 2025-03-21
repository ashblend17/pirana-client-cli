package ahhhhh

type TableData struct {
	Name   string
	Roll   string
	Result string
	Flags  string
}

type SGPAEntry struct {
	Semester string `json:"semester"`
	SGPA     string `json:"sgpa"`
}

type SemwiseSGPA struct {
	Data map[string][]SGPAEntry `json:"semwiseSGPA"`
}

type Student struct {
	Name string `json:"name"`
	Roll string `json:"roll"`
}

type Course struct {
	Semester string `json:"semester"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Grade    string `json:"grade"`
}

type StudentData struct {
	CourseData  map[string][]Course    `json:"courseData"`
	Password    map[string]string      `json:"password"`
	SemwiseSGPA map[string][]SGPAEntry `json:"semwiseSGPA"`
}

type GroupData struct {
	SGPA     string
	Courses  []Course
	Password string
}
