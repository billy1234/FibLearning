package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /ping request\n")
	io.WriteString(w, "pong\n")
}

func getFib(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /fib request\n")
	seqNo := r.URL.Query().Get("SequenceNumber")
	if seqNo == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {

		intSeqNo, err := strconv.Atoi(seqNo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			io.WriteString(w, strconv.Itoa(fibRecBetter(intSeqNo)))
		}
	}

}

func fibRecBad(seqNo int) int {
	if seqNo <= 0 {
		return 0
	} else if seqNo == 1 {
		return 1
	} else {
		return fibRecBad(seqNo-1) + fibRecBad(seqNo-2)
	}
}

func fibRecBetter(seqNo int) int {
	if seqNo <= 0 {
		return 0
	} else if seqNo == 1 {
		return 1
	}

	a := 0
	b := 1
	for i := 3; i < seqNo; i++ {
		c := a + b
		a = b
		b = c
	}
	return b
}

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/fib", getFib)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
