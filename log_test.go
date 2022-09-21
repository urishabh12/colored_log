package colored_log

import (
	"fmt"
	"strings"
	"testing"
)

const (
	seperator     = "\r\n"
	log_seperator = "\n\r\n"
)

type StringWriter struct {
	data []byte
}

func (s *StringWriter) Write(p []byte) (n int, err error) {
	s.data = append(s.data, p...)
	s.data = append(s.data, []byte(seperator)...)
	return len(p), nil
}

func Test_Deafult_STD(t *testing.T) {
	out := &StringWriter{}
	SetFlags(0)
	SetOutput(out)
	output := "Hello World"
	//Print()
	Print(output)
	s := strings.Split(string(out.data), log_seperator)
	if s[0] != fmt.Sprint(ColorBlue, output) {
		t.Errorf("output are not same %s != %s", s[0], fmt.Sprint(ColorBlue, output))
	}
	//Println()
	out.data = []byte{}
	Println(output)
	s = strings.Split(string(out.data), log_seperator)
	if s[0] != fmt.Sprint(ColorBlue, output) {
		t.Errorf("output are not same %s != %s", s[0], fmt.Sprint(ColorBlue, output))
	}
	//Printf()
	out.data = []byte{}
	Printf(output)
	s = strings.Split(string(out.data), log_seperator)
	if s[0] != fmt.Sprint(ColorBlue, output) {
		t.Errorf("output are not same %s != %s", s[0], fmt.Sprint(ColorBlue, output))
	}
}

func Test_Success_STD(t *testing.T) {
	out := &StringWriter{}
	SetFlags(0)
	SetOutput(out)
	output := "Hello World"
	//Success()
	Success(output)
	s := strings.Split(string(out.data), log_seperator)
	if s[0] != fmt.Sprint(ColorGreen, output) {
		t.Errorf("output are not same %s != %s", s[0], fmt.Sprint(ColorBlue, output))
	}
	//Successln()
	out.data = []byte{}
	Successln(output)
	s = strings.Split(string(out.data), log_seperator)
	if s[0] != fmt.Sprint(ColorGreen, output) {
		t.Errorf("output are not same %s != %s", s[0], fmt.Sprint(ColorBlue, output))
	}
	//Successf()
	out.data = []byte{}
	Successf(output)
	s = strings.Split(string(out.data), log_seperator)
	if s[0] != fmt.Sprint(ColorGreen, output) {
		t.Errorf("output are not same %s != %s", s[0], fmt.Sprint(ColorBlue, output))
	}
}
