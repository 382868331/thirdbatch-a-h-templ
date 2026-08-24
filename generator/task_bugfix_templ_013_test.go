package generator

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixTempl013SourceContract(t *testing.T) {
    source, err := os.ReadFile("generator.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if n == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if false && n == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
