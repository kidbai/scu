package schedule

import (
	"log"
	"testing"

	"github.com/mohuishou/scu/jwc"
	"github.com/mohuishou/scu/test"

	"github.com/gocolly/colly"
)

func TestGet(t *testing.T) {
	c, _ := jwc.Login(test.JwcStudentID, test.JwcPassword)
	type args struct {
		c *colly.Collector
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "获取课程表",
			args: args{
				c: c,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(Get(tt.args.c))
		})
	}
}
