package version

import (
	"fmt"
	"reflect"
	"sort"
)

var typeBuildRange = reflect.TypeOf(BuildRange{})

func BuildRangeOverlap(r1, r2 BuildRange) bool {
	return (r1[0] >= r2[0] && r1[0] <= r2[1]) ||
		(r1[1] >= r2[0] && r1[1] <= r2[1]) ||
		(r2[0] >= r1[0] && r2[0] <= r1[1]) ||
		(r2[1] >= r1[0] && r2[1] <= r1[1])
}

// QueryDescriptors expects
//   - A map[version.BuildRange]<your type>
//     In this case, it will return one match after finding a build range.
//     if any overlaps are detected, it will return an error
func QueryDescriptors[T any](build Build, descriptors map[BuildRange]T, to *T) error {
	// t := time.Now()

	var ranges []BuildRange
	for key := range descriptors {
		ranges = append(ranges, key)
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	for i := 0; i < len(ranges)-1; i++ {
		if BuildRangeOverlap(ranges[i], ranges[i+1]) {
			return fmt.Errorf("version: QueryDescriptors: overlap detected (%s, %s)", ranges[i], ranges[i+1])
		}
	}

	// Look for unary ranges (preferable to ranges sometimes)
	for _, r := range ranges {
		if r.Unary() && build == r[0] {
			return nil
		}
	}

	// look for binary ranges
	for _, r := range ranges {
		if r.Contains(build) {
			// found!
			*to = descriptors[r]
			return nil
		}
	}

	return nil

	// descType := reflect.TypeOf(desc)
	// outType := reflect.TypeOf(to)

	// descriptor := reflect.ValueOf(desc)

	// if outType.Kind() == reflect.Ptr {
	// 	outType = outType.Elem()
	// }
	// if descType.Kind() == reflect.Ptr {
	// 	descType = descType.Elem()
	// 	descriptor = descriptor.Elem()
	// }
	// if descType.Kind() != reflect.Map {
	// 	return fmt.Errorf("vsn: bad descriptor layout")
	// }
	// if descType.Key() != typeBuildRange {
	// 	return fmt.Errorf("vsn: bad descriptor keys, replace %s with version.BuildRange", descType.Key())
	// }

	// var ranges []BuildRange
	// for _, k := range descriptor.MapKeys() {
	// 	ranges = append(ranges, k.Interface().(BuildRange))
	// }
	// sort.Slice(ranges, func(i, j int) bool {
	// 	return ranges[i][0] < ranges[j][0]
	// })
	// for i := 0; i < len(ranges)-1; i++ {
	// 	if BuildRangeOverlap(ranges[i], ranges[i+1]) {
	// 		return fmt.Errorf("overlap detected (%s, %s)", ranges[i], ranges[i+1])
	// 	}
	// }

	// // fmt.Println("Sanity checks take", time.Since(t))

	// out := reflect.ValueOf(to)

	// // setValue := out

	// // // setRaw := bool

	// // // Just a pointer
	// // if out.Kind() == reflect.Ptr {
	// // 	setValue = out.Elem()

	// // 	// if descriptor.Type().Elem() == setValue.Type() {
	// // 	// 	setRaw = true
	// // 	// }
	// // }

	// for _, r := range ranges {
	// 	if r.Contains(v) {
	// 		mapvalue := descriptor.MapIndex(reflect.ValueOf(r))
	// 		// This should follow the normal Go rules of assignment, I.E slices and maps will be passed by reference
	// 		// and structs will be copied
	// 		copyvalue := reflect.New(mapvalue.Type()).Elem()
	// 		copyvalue.Set(mapvalue)

	// 		if out.Kind() == reflect.Ptr {
	// 			// out type is a pointer to a pointer
	// 			if out.Elem().Kind() == reflect.Ptr {
	// 				if out.Elem().Type() == copyvalue.Type() {
	// 					out.Elem().Set(copyvalue)
	// 				} else if out.Elem().Type() == reflect.PtrTo(copyvalue.Type()) {
	// 					out.Elem().Set(copyvalue.Addr())
	// 				}
	// 			} else if out.Elem().Kind() != reflect.Ptr {
	// 				out.Elem().Set(copyvalue)
	// 			}
	// 		}

	// 		return nil
	// 	}
	// }

	// return fmt.Errorf("vsn: QueryDescriptors: no matching descriptor for build %s", v)
}
