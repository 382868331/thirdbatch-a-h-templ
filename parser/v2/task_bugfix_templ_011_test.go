package parser

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixTempl011SourceContract(t *testing.T) {
    source, err := os.ReadFile("elementparser.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if out, ok, err = expressionAttributeParser.Parse(in); err != nil || ok {") {
        t.Fatalf("expected source contract is missing")
    }
}
