package day6

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"io"
	"strings"
)

type OrbitItem struct {
	id string
	parent *OrbitItem
}

type OrbitTree map[string]*OrbitItem

type Day6 struct {
	event.Solvable
	orbits OrbitTree
}

func New(path string, output io.Writer) event.Day {
	return Day6{
		Solvable:event.Solvable{Path:path, Output:output},
		orbits:make(OrbitTree),
	}
}

func (d Day6) Part1() {
	entries := parseEntries(input.Lines(d.Path))
	_, _ = fmt.Fprintf(d.Output, "Found %d entries\n", len(entries))

	d.orbits.buildOrbitTree(entries)
	_, _ = fmt.Fprintf(d.Output, "Found %d unique objects\n", len(d.orbits))

	totalLen := 0
	for _, item := range d.orbits {
		len := item.PathLen()
		totalLen += len
	}
	_, _ = fmt.Fprintf(d.Output, "Checksum: %d", totalLen)
}

func (d Day6) Part2() {
	entries := parseEntries(input.Lines(d.Path))
	d.orbits.buildOrbitTree(entries)

	youNode := d.orbits["YOU"]
	sanNode := d.orbits["SAN"]
	uncommon := UnCommonPaths(youNode, sanNode)
	count := TransferCount(uncommon[0], uncommon[1])

	_, _ = fmt.Fprintf(d.Output, "Minimum Transfer: %d\n", count)
}

type Entry struct {
	parent, child string
}

func (t OrbitTree) buildOrbitTree(entries []Entry) {
	for _, e := range entries {
		parent := t.get(e.parent)
		child := t.get(e.child)
		child.parent = parent
	}
}

func (t OrbitTree) get(id string) *OrbitItem {
	if o, ok := t[id]; ok {
		return o
	}
	p := &OrbitItem{id: id}
	t[id] = p
	return p
}

func (o OrbitItem) Path() []string {
	if o.parent == nil {
		return []string{o.id}
	}
	return append(o.parent.Path(), o.id)
}

func (o OrbitItem) PathLen() int {
	if o.parent == nil {
		return 0
	}
	return o.parent.PathLen() + 1
}

func UnCommonPaths(a, b *OrbitItem) [2][]string {
	aPath := a.Path()
	bPath := b.Path()

	stop := len(aPath)
	if len(bPath) < stop {
		stop = len(bPath)
	}

	for i := 0; i < stop; i++ {
		if aPath[i] != bPath[i] {
			return [2][]string{aPath[i-1:], bPath[i-1:]}
		}
	}
	return [2][]string{}
}

func TransferCount(a, b []string) int {
	ta := a[1:len(a)-1]
	tb := b[1:len(b)-1]
	return len(ta) + len(tb)
}

func parseEntries(lines []string) []Entry {
	entries := make([]Entry, len(lines))
	for i, l := range lines {
		entries[i] = parseEntry(l)
	}
	return entries
}

func parseEntry(line string) Entry {
	items := strings.Split(line, ")")
	return Entry{items[0], items[1]}
}