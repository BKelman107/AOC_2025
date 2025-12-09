// src/utils.go
package src

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
