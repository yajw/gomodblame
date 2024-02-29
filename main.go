package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	cur string

	rev    map[string][]string
	chains [][]string
)

func buildRev(m string) {
	if len(m) == 0 {
		return
	}

	parts := strings.Split(m, " ")
	mod, dep := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

	if len(rev[dep]) == 0 {
		rev[dep] = make([]string, 0)
	}

	// break cycle
	for _, x := range rev[mod] {
		if x == dep {
			return
		}
	}

	rev[dep] = append(rev[dep], mod)
}

func blame(m string, s []string, d int) {
	s[d] = m
	ups, ok := rev[m]

	if !ok || len(ups) == 0 || m == cur {
		chain := make([]string, d+1)
		for i := 0; i <= d; i++ {
			chain[i] = s[i]
		}
		chains = append(chains, chain)
		return
	}

	for _, up := range ups {
		blame(up, s, d+1)
	}
}

func main() {
	target := os.Args[1]

	out, err := exec.Command("go", "list", "-m").Output()
	if err != nil {
		panic(err)
	}

	cur = strings.Trim(string(out), "\n")

	out, err = exec.Command("go", "mod", "graph").Output()
	if err != nil {
		panic(err)
	}

	rev = make(map[string][]string)
	chains = make([][]string, 0)

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		buildRev(line)
	}

	_, ok := rev[target]
	if !ok {
		fmt.Printf("%v is  not a dependency of current module", target)
		return
	}

	blame(target, make([]string, 100000), 0)

	for _, chain := range chains {
		println(strings.Join(chain, " <- "))
	}
}
