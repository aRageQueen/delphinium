package output

import (
    "fmt"
    "os"
    "delphinium/internal"
)

func RenderTimeline(events []internal.Event, format string, path string) error {
    f, err := os.Create(path)
    if err != nil {
        return err
    }
    defer f.Close()

    for _, e := range events {
        line := fmt.Sprintf("### %s - %s\n*%s*\n\n%s\n\n", e.Timestamp.Format("2006-01-02"), e.Source, e.Title, e.Detail)
        f.WriteString(line)
    }

    return nil
}
