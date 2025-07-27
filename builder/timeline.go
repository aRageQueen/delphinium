package builder

import (
    "sort"
    "delphinium/internal"
)

func BuildTimeline(events []internal.Event) []internal.Event {
    sort.Slice(events, func(i, j int) bool {
        return events[i].Timestamp.Before(events[j].Timestamp)
    })
    return events
}
