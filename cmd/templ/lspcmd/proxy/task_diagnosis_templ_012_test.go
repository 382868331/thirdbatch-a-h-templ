package proxy

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisTempl012SourceContract(t *testing.T) {
    source, err := os.ReadFile("server.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil || result == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if err != nil && result == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
