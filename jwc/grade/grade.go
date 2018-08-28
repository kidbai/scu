package grade

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	"github.com/mohuishou/scu/jwc"

	"github.com/gocolly/colly"

	"github.com/PuerkitoBio/goquery"
)

//Grade 成绩
type Grade struct {
	// 新版新增，获取到之后会赋值给之前的字段
	ID struct {
		CourseID string `json:"courseNumber"`
		LessonID string `json:"coureSequenceNumber"`
	} `json:"id"`
	CourseID          string  `json:"-"`
	LessonID          string  `json:"-"`
	CourseName        string  `json:"courseName"`
	CourseEnglishName string  `json:"englishCourseName"`
	Credit            string  `json:"credit"`
	CourseType        string  `json:"courseAttributeName"`
	GradeShow         string  `json:"cj"`
	Grade             float64 `json:"courseScore"`
	GPA               float64 `json:"gradePointScore"`
	TermCode          string  `json:"termCode"` // 1: 秋季学期, 2: 春季学期
	Term              int     `json:"-"`        //0: 秋季学期, 1: 春季学期
	YearCode          string  `json:"academicYearCode"`
	Year              int     `json:"-"` //-1: 尚不及格，-2: 曾不及格
	TermName          string  `json:"-"`
}

// Grades 成绩列表
type Grades []Grade

// Term 一个学期的成绩
type Term struct {
	Grades    Grades  `json:"cjList"`
	TermName  string  `json:"cjbh"`
	AllCredit float64 `json:"yxxf"`
}

// Terms terms
type Terms []Term

func (terms Terms) getGrades() Grades {
	var grades Grades

	for _, term := range terms {
		grades = append(grades, term.getGrades()...)
	}

	return grades
}

func (term Term) getGrades() Grades {
	for _, grade := range term.Grades {
		g := &grade
		g.update()
		g.TermName = term.TermName
	}
	return term.Grades
}

func (grade *Grade) update() {
	grade.CourseID = grade.ID.CourseID
	grade.LessonID = grade.ID.LessonID

	// 学期代码转换
	if grade.TermCode == "2" {
		grade.Term = 1
	}

	grade.Year, _ = strconv.Atoi(strings.Split(grade.YearCode, "-")[0])
}

func get(doc *goquery.Selection, year, term int, termName string) Grades {
	grades := make(Grades, 0)
	//抓取数据
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return
		}
		grade := Grade{Term: term, Year: year, TermName: termName}
		v := reflect.ValueOf(&grade)
		elem := v.Elem()
		for k := 0; k < elem.NumField(); k++ {
			if k > 6 {
				break
			}
			elem.Field(k).SetString(strings.TrimSpace(s.Find("td").Eq(k).Text()))
		}
		switch termName {
		case "尚不及格":
			grade.Year = -1
		case "曾不及格":
			grade.Year = -2
		}
		grades = append(grades, grade)
	})
	return grades
}

// GetNow 获取本学期成绩
func GetNow(c *colly.Collector) Grades {
	var grades Grades
	c.OnHTML("#user", func(e *colly.HTMLElement) {
		grades = get(e.DOM, 0, 0, "本学期成绩")
	})
	c.Visit(jwc.DOMAIN + "/bxqcjcxAction.do?pageSize=200")
	return grades
}

// GetALL 获取所有及格成绩
func GetALL(c *colly.Collector) (grades Grades, err error) {
	var terms Terms
	c.OnResponse(func(r *colly.Response) {
		type tmp struct {
			Terms Terms `json:"lnList"`
		}
		data := &tmp{}
		err = json.Unmarshal(r.Body, data)
		terms = data.Terms
	})

	if err != nil {
		return nil, err
	}

	c.Visit(jwc.DOMAIN + "/student/integratedQuery/scoreQuery/allPassingScores/callback")

	return terms.getGrades(), nil
}

// GetNotPass 获取所有不及格成绩
func GetNotPass(c *colly.Collector) Grades {
	termNames := []string{"尚不及格", "曾不及格"}
	i := 0
	var grades Grades
	c.OnHTML("#user", func(e *colly.HTMLElement) {
		grades = append(grades, get(e.DOM, 0, 0, termNames[i])...)
		i++
	})
	c.Visit(jwc.DOMAIN + "/gradeLnAllAction.do?type=ln&oper=bjg")
	return grades
}
