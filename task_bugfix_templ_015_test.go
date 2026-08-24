package templ

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixTempl015SourceContract(t *testing.T) {
    source, err := os.ReadFile("scripttemplate.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(c.Call) > 0 {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if len(c.Call) >= 0 {") {
        t.Fatalf("mutated source contract is still present")
    }
}
