package main

//model for course -file
type Course struct {
	CourseId    string  `json:courseid`
	CourseName  string  "json:coursename"
	CoursePrice string  `json:price`
	Author      *Author `json:author`
}

type Author struct {
	FullName string `json:"fullname`
	Website  string `json:website`
}

//fake DB

var courses []Course

//midleware ,helper

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}
