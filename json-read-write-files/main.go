package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/** Struct With Unexported Fields **/
type candidate struct {
	Name      string
	Expertise []string
}

/** Struct With Exported Fields **/
type candidate2 struct {
	Name      string   `json:"name"`
	Expertise []string `json:"expertise"`
}

func panicOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	/** Slice To JSON **/
	// {
	// 	expertise := []string{"Go", "PHP", "C#", "JS"}
	// 	expertiseJSON, err := json.Marshal(expertise)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Slice To JSON:", string(expertiseJSON))
	// }

	/** Map To JSON **/
	// {
	// 	expertise := map[string]string{"Go": "excellent", "PHP": "good", "JS": "great"}
	// 	expertiseJSON, err := json.Marshal(expertise)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Map To JSON:", string(expertiseJSON))
	// }

	/** Struct With Unexported Fields To JSON **/
	// {
	// 	suthinan := candidate{
	// 		Name:      "Suthinan Musitmani",
	// 		Expertise: []string{"Go", "PHP", "JS"},
	// 	}
	// 	suthinanJSON, _ := json.Marshal(suthinan)
	// 	fmt.Println("Struct With Unexported Fields To JSON:", string(suthinanJSON))
	// }

	/** Struct With Exported Fields To JSON **/
	// {
	// 	suthinan := candidate2{
	// 		Name:      "Suthinan Musitmani",
	// 		Expertise: []string{"Go", "PHP", "JS", "Ruby", "C"},
	// 	}
	// 	suthinanJSON, _ := json.Marshal(suthinan)
	// 	fmt.Println("Struct With Exported Fields To JSON:", string(suthinanJSON))
	// }

	/** Pointer To JSON **/
	// {
	// 	suthinan := &candidate2{
	// 		Name:      "Suthinan Musitmani",
	// 		Expertise: []string{"Go", "Java", "C"},
	// 	}
	// 	suthinanJSON, _ := json.Marshal(suthinan)
	// 	fmt.Println("Pointer To JSON:", string(suthinanJSON))
	// }

	/** JSON string to Struct **/
	// {
	// 	s := `{"name":"Suthinan Musitmani","expertise":["Go","Java","C","Ruby"]}`
	// 	b := []byte(s)
	// 	suthinan := candidate2{}
	// 	err := json.Unmarshal(b, &suthinan)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("JSON string to Struct: %+v\n", suthinan)
	// }

	/** Read Full File **/
	// {
	// 	bytes, err := ioutil.ReadFile("./test.json")
	// 	panicOnError(err)
	// 	fmt.Println("Read Full File:", string(bytes))
	// }

	/** Read File In Chunks **/
	// {
	// 	f, err := os.Open("./test.json")
	// 	panicOnError(err)
	// 	defer f.Close()
	// 	reader := bufio.NewReader(f)
	// 	chunk := make([]byte, 32)
	// 	fmt.Print("Read File In Chunk: ")
	// 	for {
	// 		readSize, err := reader.Read(chunk)
	// 		if err != nil {
	// 			if !errors.Is(err, io.EOF) {
	// 				fmt.Println("Error reading file", err)
	// 			}
	// 			break
	// 		}
	// 		fmt.Println(string(chunk[0:readSize]))
	// 	}
	// }

	/** Read File Line By Line **/
	// {
	// 	f, err := os.Open("./test.json")
	// 	panicOnError(err)
	// 	defer f.Close()
	// 	s := bufio.NewScanner(f)
	// 	fmt.Print("Read File Line By Line: ")
	// 	for s.Scan() {
	// 		fmt.Println(s.Text())
	// 	}
	// 	err = s.Err()
	// 	panicOnError(err)
	// }

	/** Struct To JSON File **/
	h := candidate2{
		Name:      "Suthinan Musitmani",
		Expertise: []string{"Go", "Python", "JS"},
	}
	hBytes, err := json.Marshal(h)
	panicOnError(err)
	/** Struct To JSON File 1 **/
	{
		err = ioutil.WriteFile("./struct-to-json-1.json", hBytes, os.ModePerm)
		panicOnError(err)
		fmt.Println("Struct To JSON File 1 Saved")
	}
	/** Struct To JSON File 2 **/
	{
		f, err := os.Create("./struct-to-json-2.json")
		panicOnError(err)
		defer f.Close()
		w := bufio.NewWriter(f)
		writeSize, err := w.WriteString(string(hBytes))
		panicOnError(err)
		fmt.Printf("writeSize: %v\n", writeSize)
		w.Flush()
		fmt.Println("Struct To JSON File 2 Saved")
	}

	/** Lines To Text File 1 **/
	{
		f, err := os.Create("./lines-to-text-1.txt")
		panicOnError(err)
		defer f.Close()
		writeSize, err := f.Write([]byte("First line"))
		panicOnError(err)
		fmt.Printf("writeSize: %v\n", writeSize)
		writeSize, err = f.WriteString("\nSecond line")
		panicOnError(err)
		fmt.Printf("writeSize: %v\n", writeSize)
		fmt.Println("Lines To Text File 1 Saved")
	}
	/** Lines To Text File 2 **/
	{
		f, err := os.Create("./lines-to-text-2.txt")
		panicOnError(err)
		defer f.Close()
		linesToWrite := []string{
			"First line\n",
			"Second line\n",
			"Third line\n",
		}
		// w := bufio.NewWriter(f)
		w := bufio.NewWriterSize(f, 10) // With buffer
		for _, line := range linesToWrite {
			writeSize, err := w.WriteString(line)
			panicOnError(err)
			fmt.Printf("Bytes written: %d\n", writeSize)
			fmt.Printf("Available: %d\n", w.Available())
			fmt.Printf("Buffered: %d\n", w.Buffered())
		}
		w.Flush()
		fmt.Println("Lines To Text File 2 Saved")
	}

	/** Lines To Text File 3 **/
	{
		err := ioutil.WriteFile("./lines-to-text-3.txt", []byte("Create line"), os.ModePerm)
		panicOnError(err)
		fmt.Println("Lines To Text File 3 Saved")
		f, err := os.OpenFile("./lines-to-text-3.txt", os.O_APPEND|os.O_WRONLY, 0777)
		panicOnError(err)
		defer f.Close()
		_, err = f.WriteString("\nUpdate line")
		panicOnError(err)
		fmt.Println("Lines To Text File 3 Updated")
	}
}
