package runtime

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixTempl001SourceContract(t *testing.T) {
    source, err := os.ReadFile("watchmode.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if sl.watchModeRootErr != nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if sl.watchModeRootErr == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
