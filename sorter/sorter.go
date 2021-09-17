package main

import "flag"
import "fmt"
import "bufio" 
import "io" 
import "os" 
import "strconv"
import "sorter/al/bubblesort"

var inFile *string = flag.String("i", "inFile", "file1")
var outFile *string = flag.String("o", "outFile", "file2")
var al *string = flag.String("a", "al", "al1111111")

func readValues(inFile string)(values []int, err error)  {
	file, err := os.Open(inFile)
	if err != nil {
		fmt.Println("fail open file", inFile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1!= nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			 return
		}
	
		str := string(line)
		value, err1 := strconv.Atoi(str)
	
		if err1 != nil {
			err = err1
			return
		}
	
		values = append(values, value)

	}
	return
}

func writeValues(values []int, outFile string) error  {
	file, err := os.Create(outFile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outFile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()

	if inFile != nil {
		fmt.Println("infile =", *inFile, "outfile =", *outFile, "algorithm =", *al)
	}

	values, err := readValues(*inFile)
	if err == nil {
		fmt.Println("Read values:", values)
		bubblesort.Bubblesort(values)
		writeValues(values, *outFile)
	} else {
		fmt.Println(err)
	}

}
