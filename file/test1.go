package main

import (
	"os"
	"bufio"
	"fmt"
)

func main()  {
	filePath := "/tmp/file.log"
	//cmd := exec.Command("bash file.sh")
	//cmd.Start()
	//cmd.Run()
	for i:=0; i<500; i++{
		text :=fmt.Sprintf("xxxxx-%d", i)
		for j:=0;j<i;j++{
			text = fmt.Sprintf(text+"-%d", j)
		}
		println("write loop "+text)
		//time.Sleep(12*time.Millisecond)
		go writeLog(filePath, text)
	}

}

func writeLog(path, text string) {

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		f, err = os.Create(path)
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	w.WriteString(text+ "\n")
	w.Flush()
}