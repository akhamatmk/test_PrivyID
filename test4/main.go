package main

import "fmt"

func main() {
	// initiate slice array string
	s := []string{"foo", "bar", "baz"}
	fmt.Println(fungsiA(s))
}

func fungsiA(slice []string) []string {

	// create slice string of struct
	fungsiMap := make(map[string]struct{})
	// loop sluce from parameter
	for _, v := range slice {
		// set slice string value of empty struct
		fungsiMap[v] = struct{}{}
	}

	// i dont know this function because in test unqmap and this variable not found so i change
	fungsiSlice := make([]string, 0, len(fungsiMap))
	// looping this slice fungsimap
	for v := range fungsiMap {
		// insert array in fungsislice
		fungsiSlice = append(fungsiSlice, v)
	}

	// make return fungsi slice
	return fungsiSlice
}
