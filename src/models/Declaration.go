package models

type Declaration struct {
	Name       string
	Type       string
	Path       string
	Extends    string
	Implements []string
	Properties []Property
	Methods    []Method
}

type Extend struct {
	Name string
}

type Property struct {
	Modifier string
	Type     string
	Name     string
}

type Method struct {
	Modifier   string
	ReturnType string
	Name       string
	Parameters []string
}
