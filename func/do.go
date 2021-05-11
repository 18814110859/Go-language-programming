package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func do() error {
	ret, err := http.Get("http://www.google.com")

	if err != nil {
		return err
	}

	if ret != nil {
		defer ret.Body.Close()
	}

	return nil
}

func doFile() (err error) {
	f, err := os.Open("/Users/yu/code/Go-language-programming/func/book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func() {
			if ferr := f.Close(); ferr != nil {
				err = ferr
			}
		}()
	}

	return nil
}

func doFile2() error {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close book.txt err %v\n", err)
			}
		}(f)
	}

	f, err = os.Open("book1.txt")
	if err != nil {
		return err
	}
	defer func() {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close book1.txt err %v\n", err)
			}
		}(f)
	}()
	return nil
}

func main() {
	fmt.Println(do())
	fmt.Println(doFile())
	fmt.Println(doFile2())
}
