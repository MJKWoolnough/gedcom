package gedcom

type Line struct {
	line
	Sub []Line
}

func parseLines(lines []line) Line {
	l := Line{
		line: lines[0],
	}
	level := l.level + 1
	last := 0
	for n := range lines {
		if lines[n].level == level {
			if last != 0 {
				l.Sub = append(l.Sub, parseLines(lines[last:n]))
			}
			last = n
		}
	}
	if last != 0 {
		l.Sub = append(l.Sub, parseLines(lines[last:]))
	}
	return l
}
