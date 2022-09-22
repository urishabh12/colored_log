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

func Test_SetOutput(t *testing.T) {
	out := &StringWriter{}
	SetFlags(0)
	SetOutput(out)
	output := "Test"
	Print(output)
	s := strings.Split(string(out.data), log_seperator)
	if s[0] != fmt.Sprint(Blue, output) {
		t.Errorf("output are not same %s != %s", s[0], fmt.Sprint(Blue, output))
	}
}

func Test_Deafult_STD(t *testing.T) {
	out := &StringWriter{}
	SetFlags(0)
	SetOutput(out)
	output := "Hello World"
	//Print()
	Print(output)
	s := strings.Split(string(out.data), log_seperator)
	if s[0] != fmt.Sprint(Blue, output) {
		t.Errorf("output are not same %s != %s", s[0], fmt.Sprint(Blue, output))
	}
	//Println()
	out.data = []byte{}
	Println(output)
	s = strings.Split(string(out.data), log_seperator)
	if s[0] != fmt.Sprint(Blue, output) {
		t.Errorf("output are not same %s != %s", s[0], fmt.Sprint(Blue, output))
	}
	//Printf()
	out.data = []byte{}
	Printf(output)
	s = strings.Split(string(out.data), log_seperator)
	if s[0] != fmt.Sprint(Blue, output) {
		t.Errorf("output are not same %s != %s", s[0], fmt.Sprint(Blue, output))
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
	if s[0] != fmt.Sprint(Green, output) {
		t.Errorf("output are not same %s != %s", s[0], fmt.Sprint(Green, output))
	}
	//Successln()
	out.data = []byte{}
	Successln(output)
	s = strings.Split(string(out.data), log_seperator)
	if s[0] != fmt.Sprint(Green, output) {
		t.Errorf("output are not same %s != %s", s[0], fmt.Sprint(Green, output))
	}
	//Successf()
	out.data = []byte{}
	Successf(output)
	s = strings.Split(string(out.data), log_seperator)
	if s[0] != fmt.Sprint(Green, output) {
		t.Errorf("output are not same %s != %s", s[0], fmt.Sprint(Green, output))
	}
}

func Test_Prefix(t *testing.T) {
	prefix := "prefix"
	SetPrefix(prefix)

	if prefix != Prefix() {
		t.Errorf("output are not same %s != %s", prefix, Prefix())
	}
}

func Test_SetPrefix(t *testing.T) {
	prefix := "prefix"
	output := "Hello"
	out := &StringWriter{}
	SetFlags(0)
	SetOutput(out)
	SetPrefix(prefix)
	Print(output)
	s := strings.Split(string(out.data), log_seperator)
	if s[0] != fmt.Sprint(Blue, prefix, output) {
		t.Errorf("output are not same %s != %s", s[0], fmt.Sprint(Blue, prefix, output))
	}
}

func Test_Flags(t *testing.T) {
	flag := 1
	SetFlags(flag)

	if flag != Flags() {
		t.Errorf("output are not same %d != %d", flag, Flags())
	}
}

func Test_SetFlags(t *testing.T) {
	output := "Hello"
	out := &StringWriter{}
	SetFlags(0)
	SetOutput(out)
	Print(output)
	s := strings.Split(string(out.data), log_seperator)
	if s[0] != fmt.Sprint(Blue, output) {
		t.Errorf("output are not same %s != %s", s[0], fmt.Sprint(Blue, output))
	}
}

func Benchmark_Print(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Print("Hello World")
	}
}
