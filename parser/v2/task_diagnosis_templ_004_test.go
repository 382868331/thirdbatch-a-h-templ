package parser

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisTempl004SourceContract(t *testing.T) {
    source, err := os.ReadFile("raw.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if e.Contents, ok, err = parse.StringUntil(end).Parse(pi); err != nil || !ok {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if e.Contents, ok, err = parse.StringUntil(end).Parse(pi); err == nil || !ok {") {
        t.Fatalf("mutated source contract is still present")
    }
}
