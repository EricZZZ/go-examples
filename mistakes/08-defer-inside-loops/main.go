package main

import "os"

func wrong(paths []string) error {
	for _, p := range paths {
		f, err := os.Open(p)
		if err != nil {
			return err
		}
		defer f.Close() // all N handles stay open until the function returns
	}
	return nil
}

func processOne(p string) error {
	f, err := os.Open(p)
	if err != nil {
		return err
	}
	defer f.Close() // the handle is closed when the function returns
	return nil
}

func right(paths []string) error {
	for _, p := range paths {
		if err := processOne(p); err != nil {
			return err
		}
	}
	return nil
}

func main() {}
