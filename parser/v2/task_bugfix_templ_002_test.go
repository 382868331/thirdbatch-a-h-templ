package parser

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixTempl002SourceContract(t *testing.T) {
    source, err := os.ReadFile("scriptparser.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return n, false, parse.Error(\"<script>: unclosed element - missing '>'\", pi.Position())") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "return n, false, parse.Error(\"<script>=: unclosed element - missing '>'\", pi.Position())") {
        t.Fatalf("mutated source contract is still present")
    }
}
