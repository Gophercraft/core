package realm

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/text"
	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/superp00t/etc"
)

var (
	wordType = reflect.TypeOf((*text.Word)(nil)).Elem()
)

func parseBool(value reflect.Value, str string) error {
	if str == "on" {
		value.SetBool(true)
		return nil
	} else if str == "off" {
		value.SetBool(false)
		return nil
	}

	on, err := strconv.ParseBool(str)
	if err != nil {
		return err
	}

	value.SetBool(on)
	return nil
}

func parseFloat(value reflect.Value, str string) error {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}
	value.SetFloat(f)
	return nil
}

func parseUint(value reflect.Value, str string) error {
	u, err := strconv.ParseUint(str, 0, 64)
	if err != nil {
		return err
	}
	value.SetUint(u)
	return nil
}

func parseInt(value reflect.Value, str string) error {
	i, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		return err
	}

	value.SetInt(i)
	return nil
}

func (s *Session) getCommandPrivileges() CommandPrivileges {
	if s.Tier == rpcnet.Tier_Admin {
		// All privileges enabled
		return 0xFF
	}

	cp := CommandPrivileges(0)

	if s.Tier == rpcnet.Tier_GameMaster {
		cp |= GameMaster
	}

	return cp
}

func (s *Session) HandleCommand(c string) {
	args, err := parseCmd(c)
	if err != nil {
		log.Warn(err)
		return
	}

cmd:
	for _, command := range s.Server.CommandHandlers {
		sig := strings.Split(command.Signature, " ")

		for idx, sigPart := range sig {
			if idx >= len(args) {
				break
			}

			if !strings.HasPrefix(sigPart, args[idx]) && !(sigPart == args[idx]) {
				continue cmd
			}
		}

		cp := s.getCommandPrivileges()

		// Check if allowed
		if cp&command.Requires == 0 && !s.IsAdmin() {
			s.Warnf("You do not have the required permissions to use this command.")
			return
		}

		c := reflect.ValueOf(command.Function)
		commandArgs := []reflect.Value{reflect.ValueOf(s)}

		// Number of strings passed into arguments after command signature
		// This MAY be less arguments than the function specifies, in which case we need to call it with zero values
		numPassedArgs := len(args) - len(sig)
		if numPassedArgs < 0 {
			continue
		}

		paramType := reflect.TypeOf([]string{})

		// Some functions just accept a slice of strings.
		// In which case we can just slice the argument strings and pass them into the function.
		if c.Type().NumIn() > 1 && c.Type().In(1) == paramType {
			params := args[len(sig):]
			commandArgs = append(commandArgs, reflect.ValueOf(params))
			c.Call(commandArgs)
			return
		}

		for idx := 1; idx < c.Type().NumIn(); idx++ {
			if idx-1 >= numPassedArgs {
				// Create zero value for omitted argument
				zero := reflect.New(c.Type().In(idx)).Elem()
				commandArgs = append(commandArgs, zero)
			} else {
				str := args[len(sig)-1+idx]
				value := reflect.New(c.Type().In(idx)).Elem()

				// Some arguments are actually text.Words (i.e. have string conversion methods built in)
				// Exploit this to make input easier
				var w text.Word

				if value.Type().Implements(wordType) {
					w = value.Interface().(text.Word)
				} else if reflect.PtrTo(value.Type()).Implements(wordType) {
					w = value.Addr().Interface().(text.Word)
				}

				if w != nil {
					err := w.DecodeWord(str)
					if err == nil {
						commandArgs = append(commandArgs, value)
						continue
					}
				}

				var err error
				switch value.Kind() {
				case reflect.Bool:
					err = parseBool(value, str)
				case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
					err = parseUint(value, str)
				case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
					err = parseInt(value, str)
				case reflect.String:
					value.SetString(str)
				case reflect.Float32, reflect.Float64:
					err = parseFloat(value, str)
				default:
					panic(value.Kind())
				}
				if err != nil {
					log.Warn(err)
					s.Warnf("%s", err)
					return
				}

				commandArgs = append(commandArgs, value)
			}
		}

		c.Call(commandArgs)
		return
	}

	s.Warnf("Command not found.")

	s.showCommandSuggestions(c)
}

type commandSuggest struct {
	Signature  string
	Similarity float64
}

func (s *Session) showCommandSuggestions(c string) {
	lev := metrics.NewLevenshtein()

	var suggestions []commandSuggest
	for _, helpHandler := range s.Server.CommandHandlers {
		similarity := strutil.Similarity(c, helpHandler.Signature, lev)

		suggestions = append(suggestions, commandSuggest{
			Signature:  helpHandler.Signature,
			Similarity: similarity,
		})
	}

	if len(suggestions) == 0 {
		return
	}

	sort.Slice(suggestions, func(i, j int) bool {
		return suggestions[i].Similarity >= suggestions[j].Similarity
	})

	log.Dump("suggestions", suggestions)

	s.Warnf("Did you mean: |c%s.%s|r ?", HelpColor, suggestions[0].Signature)
}

func parseCmd(s string) ([]string, error) {
	e := etc.FromString(s)

	if rn, _, _ := e.ReadRune(); rn != '.' {
		return nil, fmt.Errorf("not a command")
	}

	var args []string

argScan:
	for {
		argBuf := etc.NewBuffer()

		for x := 0; ; x++ {
			rn, _, _ := e.ReadRune()
			if rn == 0 {
				args = append(args, argBuf.ToString())
				goto endScan
			}

			if argBuf.Len() == 0 && rn == ' ' {
				continue
			}

			if rn == ' ' {
				args = append(args, argBuf.ToString())
				argBuf = etc.NewBuffer()
				continue argScan
			}

			// Don't split markup block
			if rn == '|' {
				markupCode, _, _ := e.ReadRune()
				if markupCode == 'c' {
					e.Jump(-2)

					var markupText string
					for {
						r, _, _ := e.ReadRune()
						if r == 0 {
							argBuf.Write([]byte(markupText))
							break
						}

						if r == '|' {
							r2, _, _ := e.ReadRune()
							if r2 == 0 {
								argBuf.WriteRune(r)
								break
							}

							if r2 == 'r' {
								argBuf.WriteRune(r2)
								break
							}

							argBuf.WriteRune(r)
							argBuf.WriteRune(r2)
						} else {
							argBuf.WriteRune(r)
						}
					}
				} else {
					argBuf.WriteRune(rn)
					argBuf.WriteRune(markupCode)
				}
			} else {
				argBuf.WriteRune(rn)
			}
		}
	}
endScan:

	return args, nil
}
