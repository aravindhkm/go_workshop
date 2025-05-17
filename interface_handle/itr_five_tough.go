package interfacehandle

import (
	"errors"
	"fmt"
)

// Reader defines the ability to read data
type Reader interface {
	Read() (string, error)
}

// Processor defines the ability to process data
type Processor interface {
	Process(input string) (string, error)
}

// Reporter defines the ability to report results
type Reporter interface {
	Report(result string) error
}

// Plugin is a compound interface
type Plugin interface {
	Reader
	Processor
	Reporter
	Name() string
}

// ErrorPlugin is a plugin that intentionally returns errors
type ErrorPlugin struct{}

func (e ErrorPlugin) Read() (string, error) {
	return "", errors.New("error reading data")
}

func (e ErrorPlugin) Process(input string) (string, error) {
	return "", errors.New("cannot process empty input")
}

func (e ErrorPlugin) Report(result string) error {
	return errors.New("reporting failed")
}

func (e ErrorPlugin) Name() string {
	return "ErrorPlugin"
}

// TextPlugin is a concrete implementation of Plugin
type TextPlugin struct{}

func (t TextPlugin) Read() (string, error) {
	return "hello world", nil
}

func (t TextPlugin) Process(input string) (string, error) {
	return fmt.Sprintf("processed: %s", input), nil
}

func (t TextPlugin) Report(result string) error {
	fmt.Println("Reporting:", result)
	return nil
}

func (t TextPlugin) Name() string {
	return "TextPlugin"
}

// runPlugin runs the full plugin pipeline
func runPlugin(p Plugin) {
	fmt.Printf("Running plugin: %s\n", p.Name())

	data, err := p.Read()
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	result, err := p.Process(data)
	if err != nil {
		fmt.Println("Process error:", err)
		return
	}

	if err := p.Report(result); err != nil {
		fmt.Println("Report error:", err)
	}
}

// inspectInterface uses empty interface and type assertions
func inspectInterface(i interface{}) {
	switch v := i.(type) {
	case Plugin:
		fmt.Println("This is a Plugin:", v.Name())
	case Reporter:
		fmt.Println("This is only a Reporter")
	default:
		fmt.Println("Unknown type")
	}
}

func ItrFive() {
	plugins := []Plugin{
		TextPlugin{},
		ErrorPlugin{},
	}

	for _, plugin := range plugins {
		runPlugin(plugin)
		inspectInterface(plugin)
		fmt.Println("----")
	}
}
