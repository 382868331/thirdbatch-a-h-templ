package generator

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixTempl003SourceContract(t *testing.T) {
    source, err := os.ReadFile("generator.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return attrName == \"href\"") {
        t.Fatalf("expected source contract is missing")
    }
}
