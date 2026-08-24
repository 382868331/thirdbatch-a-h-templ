package parser

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixTempl006SourceContract(t *testing.T) {
    source, err := os.ReadFile("types.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if c.GoCode == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
