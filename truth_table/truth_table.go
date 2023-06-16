package truth_table

import (
	"fmt"
	"github.com/koanhealth/gotools/fixgo"
	ks "github.com/koanhealth/gotools/strings"
	"strings"
)

type TruthTable struct {
	fields []string
	lines  []*truthTableLine
}

type keyMatch string

const (
	Yes keyMatch = "Y"
	No  keyMatch = "N"
	DC  keyMatch = "--"
)

type truthTableLine struct {
	result interface{}
	values []keyMatch
}

func Create(keys ...string) *TruthTable {
	return &TruthTable{fields: keys}
}

func (tt *TruthTable) Line(result interface{}, fieldValues ...interface{}) *TruthTable {
	if len(fieldValues) != len(tt.fields) {
		panic(fmt.Errorf("count of field values (%d) does not match decared field count (%d)", len(fieldValues), len(tt.fields)))
	} else {
		tt.lines = append(tt.lines, &truthTableLine{result: result, values: tt.convertFieldValues(fieldValues)})
	}
	return tt
}

func (tt *TruthTable) Result(fieldValues ...interface{}) (found bool, result interface{}) {
	found, result, _ = tt.resultEx(false, fieldValues...)
	return
}

func (tt *TruthTable) ResultVerbose(fieldValues ...interface{}) (found bool, result interface{}, details string) {
	found, result, details = tt.resultEx(true, fieldValues...)
	return
}

func (tt *TruthTable) resultEx(verbose bool, fieldValues ...interface{}) (found bool, result interface{}, details string) {
	values := tt.convertFieldValues(fieldValues)
	if len(values) != len(tt.fields) {
		panic(fmt.Errorf("count of field values (%d) does not match decared field count (%d)", len(values), len(tt.fields)))
	}

	for _, line := range tt.lines {
		found, result = line.match(values)
		if found {
			if verbose {
				details = tt.toString(values, line)
			}
			break
		}
	}
	return
}

func (tt *TruthTable) String() string {
	return tt.toString([]keyMatch{}, nil)
}

func (tt *TruthTable) toString(input []keyMatch, matchedLine *truthTableLine) string {
	sb := ks.NewStringBuilder()

	sb.Print("|")
	widths := make([]int, 0, len(tt.fields)+1)
	for _, f := range tt.fields {
		width := fixgo.Ternary(len(f) > len(Yes), len(f), len(Yes)) + 2
		widths = append(widths, width)
		sb.Print(ks.CenterString(f, " ", width))
		sb.Print("|")
	}

	resultWidth := len("Result")
	for _, l := range tt.lines {
		lineResultWidth := len(fmt.Sprint(l.result))
		resultWidth = fixgo.Ternary(resultWidth < lineResultWidth, lineResultWidth, resultWidth)
	}

	resultWidth += 2
	widths = append(widths, resultWidth)
	sb.Print(ks.CenterString("Result", " ", resultWidth))
	sb.Print("|")
	lineWidth := len(sb.String())

	sb.Println("")
	sb.Println(strings.Repeat("-", lineWidth))
	for _, l := range tt.lines {
		l.addString(sb, widths)
	}

	sb.Println(strings.Repeat("-", lineWidth))

	if matchedLine != nil {
		sb.Println(ks.CenterString(" Input ", "=", lineWidth))
		matchedLine.addStringEx(sb, input, "", widths)
		sb.Println(ks.CenterString(" Matches ", "=", lineWidth))
		matchedLine.addString(sb, widths)
		sb.Println(strings.Repeat("=", lineWidth))
	}

	return sb.String()
}

func (tt *TruthTable) convertFieldValues(fieldValues []interface{}) (result []keyMatch) {
	result = make([]keyMatch, 0, len(fieldValues))
	for _, v := range fieldValues {
		if v == nil {
			result = append(result, DC)
		} else {
			switch tv := v.(type) {
			case keyMatch:
				result = append(result, tv)
			case bool:
				result = append(result, fixgo.Ternary(tv, Yes, No))
			default:
				panic(fmt.Errorf("unrocgonized type for field value: %v", v))
			}
		}
	}
	return
}

func (l *truthTableLine) match(fieldValues []keyMatch) (found bool, result interface{}) {
	for idx, lhs := range l.values {
		rhs := fieldValues[idx]
		if lhs == DC || rhs == DC {
			continue
		}
		if lhs != rhs {
			return
		}
	}
	found = true
	result = l.result

	return
}

func (l *truthTableLine) addStringEx(sb *ks.StringBuilder, fieldValues []keyMatch, resultString string, widths []int) {
	sb.Print("|")

	for idx, f := range fieldValues {
		width := widths[idx]
		sb.Print(ks.CenterString(string(f), " ", width))
		sb.Print("|")
	}
	sb.Print(ks.CenterString(fmt.Sprint(resultString), " ", widths[len(widths)-1]))
	sb.Print("|")
	sb.Println("")
}

func (l *truthTableLine) addString(sb *ks.StringBuilder, widths []int) {
	l.addStringEx(sb, l.values, fmt.Sprint(l.result), widths)
}
