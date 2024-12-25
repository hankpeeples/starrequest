// Package parser provides functionality to parse star request files
package parser

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	charmLog "github.com/charmbracelet/log"
)

// Request represents the root structure of a starrequest file
type Request struct {
	Requests []Requests `yaml:"requests"`
}

// Requests defines an individual HTTP request configuration
type Requests struct {
	Method  string            `yaml:"method"`
	URL     string            `yaml:"url"`
	Headers map[string]string `yaml:"headers,omitempty"`
	Body    interface{}       `yaml:"body,omitempty"`
}

// Parser defines the interface for parsing star request files
type Parser interface {
	// Parse reads a file at the given path and returns a parsed Request
	Parse(string) (Request, error)
}

// yamlParser implements the Parser interface for YAML files
type yamlParser struct {
	log *charmLog.Logger
}

// NewParser creates a new YAML parser instance
func NewParser(log *charmLog.Logger) Parser {
	return &yamlParser{
		log: log,
	}
}

// Parse reads a YAML file and returns a Request
func (p *yamlParser) Parse(filename string) (Request, error) {
	p.log.Debug("parsing request file.", "path", filename)

	// Read the file
	yamlText, err := os.ReadFile(filename)
	if err != nil {
		return Request{}, fmt.Errorf("failed to read file: %w", err)
	}

	// Unmarshal the YAML
	var req Request
	if err = yaml.Unmarshal(yamlText, &req); err != nil {
		return Request{}, fmt.Errorf("failed to parse yaml: %w", err)
	}

	// Validate the request
	if err = p.validate(req); err != nil {
		return Request{}, fmt.Errorf("invalid request: %w", err)
	}

	p.log.Debug("successfully parsed request file.", "requests", len(req.Requests))
	return req, nil
}

// validate performs basic validation of the request structure
func (p *yamlParser) validate(req Request) error {
	if len(req.Requests) == 0 {
		return fmt.Errorf("no requests defined")
	}

	for i, r := range req.Requests {
		if r.Method == "" {
			return fmt.Errorf("request %d: method is required", i)
		}
		if r.URL == "" {
			return fmt.Errorf("request %d: url is required", i)
		}
	}

	return nil
}
