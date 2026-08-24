package parser

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisTempl020SourceContract(t *testing.T) {
    source, err := os.ReadFile("goparser.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "parenDepth--") {
        t.Fatalf("expected source contract is missing")
    }
}
