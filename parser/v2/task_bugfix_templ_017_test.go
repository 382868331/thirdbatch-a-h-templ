package parser

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixTempl017SourceContract(t *testing.T) {
    source, err := os.ReadFile("templatefile.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "hasTemplatePrefix := strings.HasPrefix(l, \"templ \") || strings.HasPrefix(l, \"css \") || strings.HasPrefix(l, \"script \")") {
        t.Fatalf("expected source contract is missing")
    }
}
