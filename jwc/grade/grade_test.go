package grade

import (
	"log"
	"testing"

	"github.com/gocolly/colly"
	"github.com/mohuishou/scu/jwc"
	"github.com/mohuishou/scu/test"
)

func TestGetNow(t *testing.T) {
	c, _ := jwc.Login(test.LibStudentID, test.LibPassword)
	type args struct {
		c *colly.Collector
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				c: c,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grades := GetNow(tt.args.c)
			log.Println(grades)
		})
	}
}

func TestGetALL(t *testing.T) {
	c, err := jwc.Login(test.LibStudentID, test.LibPassword)
	if err != nil {
		panic(err)
	}
	type args struct {
		c *colly.Collector
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				c: c,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(GetALL(tt.args.c))
		})
	}
}

func TestGetNotPass(t *testing.T) {
	c, err := jwc.Login(test.JwcStudentID, test.JwcPassword)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		c *colly.Collector
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				c: c,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(GetNotPass(tt.args.c))
		})
	}
}
