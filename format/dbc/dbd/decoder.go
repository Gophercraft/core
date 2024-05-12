package dbd

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/Gophercraft/core/version"
)

const eof rune = 0

var (
	columnTypeRegex      = regexp.MustCompile("^(.+)<(.+)>$")
	extractTemplateRegex = regexp.MustCompile("^(.+)<(.+)>(.*)$")
	extractArrayRegex    = regexp.MustCompile("^(.+)\\[(.+)\\]$")
)

type decoder struct {
	def *Definition
	*bufio.Reader
	Line int
}

// readline returns a series of line components with whitespace and comments removed
func (d *decoder) readLine() ([]string, error) {
	d.Line++
	line, err := d.ReadString('\n')
	if err != nil {
		return nil, err
	}

	line = strings.TrimRight(line, "\r\n")

	// remove comments
	i := strings.Index(line, "//")

	if i > -1 {
		line = line[:i]
	}

	// remove whitespace and build list of words

	elements := strings.Split(line, " ")

	newEls := make([]string, 0)

	for _, el := range elements {
		if el != "" {
			newEls = append(newEls, el)
		}
	}

	return newEls, nil
}

func columnTypeFromString(typename string) (colType ColumnType, err error) {
	switch typename {
	case "uint":
		colType = Uint
	case "int":
		colType = Int
	case "float":
		colType = Float
	case "bool":
		colType = Bool
	case "string":
		colType = String
	case "locstring":
		colType = LocString
	default:
		err = fmt.Errorf("unknown column type %s", typename)
	}
	return
}

func parseColumnType(str string) (ColumnType, string, error) {
	if match(columnTypeRegex, str) {
		w := search(columnTypeRegex, str)
		typename := w[1]
		foreign := w[2]
		colType, err := columnTypeFromString(typename)
		return colType, foreign, err
	}

	colType, err := columnTypeFromString(str)
	return colType, "", err
}

func (d *decoder) decodeColumns() error {
	for {
		w, err := d.readLine()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		if len(w) == 0 {
			break
		}

		if len(w) != 2 {
			return fmt.Errorf("%+v: bad declaration, invalid number of arguments", w)
		}

		typeID, template, err := parseColumnType(w[0])
		if err != nil {
			return err
		}

		col := ColumnDefinition{}
		fieldName := w[1]
		col.Verified = !strings.HasSuffix(fieldName, "?")
		if !col.Verified {
			fieldName = fieldName[:len(fieldName)-1]
		}

		col.Name = fieldName
		col.Type = typeID
		if template != "" {
			strs := strings.Split(template, "::")
			if len(strs) != 2 {
				return fmt.Errorf("bad foreign key declaration: %+v", w)
			}

			col.ForeignRecord = strs[0]
			col.ForeignKey = strs[1]
		}

		d.def.Columns = append(d.def.Columns, col)
	}
	return nil
}

func extractTemplate(s string) (string, string) {
	if match(extractTemplateRegex, s) {
		w := search(extractTemplateRegex, s)
		return w[2], w[1] + w[3]
	}

	return "", s
}

func extractArray(s string) (string, string) {
	if match(extractArrayRegex, s) {
		w := search(extractArrayRegex, s)
		return w[2], w[1]
	}

	return "", s
}

func (d *decoder) decodeLayout(w []string) error {
	layout := Layout{}
	var err error
	for {
		if len(w) == 0 {
			break
		}

		// Build ranges
		if w[0] == "BUILD" {
			if strings.Contains(w[1], "-") {
				brange, err := ParseBuildRange(w[1])
				if err != nil {
					return err
				}

				layout.BuildRanges = append(layout.BuildRanges, brange)
			} else if len(w) > 1 {
				// Explicit list of verified builds
				nocomma(w)
				for _, str := range w[1:] {
					v, err := version.ParseDBD(str)
					if err != nil {
						return err
					}
					layout.VerifiedBuilds = append(layout.VerifiedBuilds, v)
				}
			}
		} else if w[0] == "LAYOUT" {
			if len(layout.Hashes) > 0 {
				return fmt.Errorf("more than one LAYOUT block not allowed")
			}

			nocomma(w)
			layout.Hashes = w[1:]
		} else if w[0] != "COMMENT" {
			// Columns
			if len(w) != 1 {
				return fmt.Errorf("invalid layout column")
			}

			lc := LayoutColumn{}

			c := w[0]
			if strings.HasPrefix(c, "$") {
				list := strings.SplitN(c, "$", 3)
				lc.Options = strings.Split(list[1], ",")
				c = list[2]
			}

			var bits string
			bits, c = extractTemplate(c)

			var array string
			array, c = extractArray(c)

			col := d.def.Column(c)
			if col == nil {
				return fmt.Errorf("no such column as %s", c)
			}

			lc.Name = c

			lc.Signed = col.Type == Int
			lc.ArraySize = -1

			if bits != "" {
				if bits[0] == 'u' {
					lc.Signed = false
					bits = bits[1:]
				}

				sz, err := strconv.ParseUint(bits, 0, 64)
				if err != nil {
					return err
				}

				// DBCs necessarily are 8-bit aligned.
				if sz%8 != 0 {
					return fmt.Errorf("Invalid bit size %s, must be aligned to 8 bits", w[0])
				}

				if col.Type == Int || col.Type == Uint {
					if sz != 8 && sz != 16 && sz != 32 && sz != 64 {
						return fmt.Errorf("No int types for %d", sz)
					}
				}

				if col.Type == Float {
					if sz != 32 && sz != 64 {
						return fmt.Errorf("No float types for %d", sz)
					}
				}

				lc.Bits = int(sz)
			} else {
				if col.Type == Int || col.Type == Uint {
					return fmt.Errorf("cannot infer bit size of integer %s", col.Name)
				}

				// Most floats are 32-bit
				if col.Type == Float {
					lc.Bits = 32
				}
				// Other types are not treated in terms of bits, Bits can remain zero forever
			}

			if array != "" {
				asz, err := strconv.ParseUint(array, 0, 64)
				if err != nil {
					return err
				}
				lc.ArraySize = int(asz)
				col.HintArray = true
			}

			layout.Columns = append(layout.Columns, lc)

			if lc.Bits >= col.HintBits {
				col.HintBits = lc.Bits
			}
		}

		w, err = d.readLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}

	d.def.Layouts = append(d.def.Layouts, layout)

	return nil
}

func (d *decoder) decodeBlock() (end bool, err error) {
	var w []string
	w, err = d.readLine()
	if err != nil {
		if err == io.EOF {
			end = true
			err = nil
		}
		return
	}
	if len(w) == 0 {
		return
	}
	switch w[0] {
	case "COLUMNS":
		err = d.decodeColumns()
	case "BUILD", "LAYOUT":
		err = d.decodeLayout(w)
	case "COMMENT":
		err = nil
	default:
		err = fmt.Errorf("unknown block type %s", w[0])
	}
	return
}

func (d *decoder) Decode(name string) (*Definition, error) {
	d.def = &Definition{}
	d.def.Name = name

	for {
		end, err := d.decodeBlock()
		if err != nil {
			return nil, fmt.Errorf("Error in %s:%d: %s", name, d.Line, err)
		}

		if end {
			break
		}
	}

	return d.def, nil
}
