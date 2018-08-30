package evaluate

import (
	"testing"

	"github.com/mohuishou/scu/jwc"
	"github.com/mohuishou/scu/test"
)

func TestGetEvaList(t *testing.T) {
	c, _ := jwc.Login(test.JwcStudentID, test.JwcPassword)
	res, err := GetEvaList(c)
	if err != nil {
		panic(err)
	}

	t.Log(len(res), err, res[1])
	// c, _ = jwc.Login(test.JwcStudentID, test.JwcPassword)
	// r := res[len(res)-2]
	// r.Comment = "超级棒的老师"
	// r.Star = 5
	// t.Log(AddEvaluate(c, &r))
}
