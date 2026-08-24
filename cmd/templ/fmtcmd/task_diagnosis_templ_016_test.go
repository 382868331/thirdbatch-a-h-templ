package fmtcmd

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisTempl016SourceContract(t *testing.T) {
    source, err := os.ReadFile("main.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
