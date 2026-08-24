package proxy

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixTempl014SourceContract(t *testing.T) {
    source, err := os.ReadFile("server.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "p.Log.Info(\"client -> server: FoldingRanges\")") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "p.Log.Info(\"client ->= server: FoldingRanges\")") {
        t.Fatalf("mutated source contract is still present")
    }
}
