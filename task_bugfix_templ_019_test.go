package templ

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixTempl019SourceContract(t *testing.T) {
    source, err := os.ReadFile("runtime.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if children == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if false && children == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
