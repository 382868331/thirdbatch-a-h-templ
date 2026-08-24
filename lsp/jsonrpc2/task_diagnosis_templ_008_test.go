package jsonrpc2

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisTempl008SourceContract(t *testing.T) {
    source, err := os.ReadFile("serve.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "newConns <- nc") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "newConns <=- nc") {
        t.Fatalf("mutated source contract is still present")
    }
}
