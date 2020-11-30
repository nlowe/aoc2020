package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"text/template"

	"github.com/nlowe/aoc2020/util"
	"github.com/zellyn/kooky"
	"github.com/zellyn/kooky/chrome"
)

const (
	glueTemplate = `package day{{ .N }}

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "{{ .N }}",
		Short: "Problems for Day {{ .N }}",
	}

	//day.AddCommand(aCommand())
    //day.AddCommand(bCommand())

	root.AddCommand(day)
}
`

	problemTemplate = `package day{{ .N }}

import (
    "fmt"

    "github.com/nlowe/aoc2020/challenge"
	"github.com/nlowe/aoc2020/util"
	"github.com/spf13/cobra"
)

func {{ .AB | toLower }}Command() *cobra.Command {
    return &cobra.Command{
        Use:   "{{ .AB | toLower }}",
        Short: "Day {{ .N }}, Problem {{ .AB }}",
        Run: func(_ *cobra.Command, _ []string) {
            fmt.Printf("Answer: %d\n", {{ .AB | toLower }}(challenge.FromFile()))
        },
    }
}

func {{ .AB | toLower }}(challenge *challenge.Input) int {
    return 0
}
`

	testTemplate = `package day{{ .N }}

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

func Test{{ .AB }}(t *testing.T) {
	input := challenge.FromLiteral("foobar")

	result := {{ .AB | toLower }}(input)

	require.Equal(t, 42, result)
}
`
)

type metadata struct {
	N  int
	AB string
}

func main() {
	if len(os.Args) != 3 {
		abort(fmt.Errorf("expected 3 args but got %d", len(os.Args)))
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		abort(err)
	}

	ab := strings.ToUpper(os.Args[2])
	if !strings.ContainsAny(ab, "AB") {
		abort(fmt.Errorf("unknown problem segment: %s", ab))
	}

	p, err := util.PkgPath(n)
	if err != nil {
		abort(err)
	}

	if err := os.MkdirAll(p, 0744); err != nil {
		abort(err)
	}

	m := metadata{N: n, AB: ab}
	funcs := template.FuncMap{
		"toLower": strings.ToLower,
	}

	gluePath := filepath.Join(p, "import.go")
	if _, stat := os.Stat(gluePath); stat != nil && os.IsNotExist(stat) {
		genFile(gluePath, glueTemplate, funcs, m)
	}

	genFile(filepath.Join(p, fmt.Sprintf("%s.go", strings.ToLower(ab))), problemTemplate, funcs, m)
	genFile(filepath.Join(p, fmt.Sprintf("%s_test.go", strings.ToLower(ab))), testTemplate, funcs, m)

	goimports := exec.Command("goimports", "-w", p)
	if err := goimports.Run(); err != nil {
		abort(err)
	}

	if _, stat := os.Stat(filepath.Join(p, "input.txt")); os.IsNotExist(stat) {
		fmt.Println("fetching input for day", n)
		problemInput, err := getInput(n)
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile(filepath.Join(p, "input.txt"), problemInput, 0644); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("input already downloaded, skipping...")
	}

	fmt.Printf("Generated problem %s for day %d. Be sure to add it to main.go\n", ab, n)

	// TODO: Can we modify main.go easily?
}

func genFile(path, t string, funcs template.FuncMap, m metadata) {
	if _, stat := os.Stat(path); os.IsNotExist(stat) {
		fmt.Println("creating", path)
		t := template.Must(template.New(path).Funcs(funcs).Parse(t))
		cf, err := os.Create(path)
		if err != nil {
			abort(err)
		}

		defer mustClose(cf)
		if err := t.Execute(cf, m); err != nil {
			abort(err)
		}
	} else {
		fmt.Println(path, "already exists, skipping...")
	}
}

func chromeCookiePath() (string, error) {
	if p, set := os.LookupEnv("CHROME_PROFILE_PATH"); set {
		return filepath.Join(p, "Cookies"), nil
	}

	if runtime.GOOS == "windows" {
		localAppData, err := os.UserCacheDir()
		return filepath.Join(localAppData, "Google", "Chrome", "User Data", "Default", "Cookies"), err
	}

	return "", fmt.Errorf("chrome cookie path for GOOS %s not implemented, set CHROME_PROFILE_PATH instead", runtime.GOOS)
}

func getInput(day int) ([]byte, error) {
	_, _ = os.UserConfigDir()
	_, _ = os.UserCacheDir()

	cookiePath, err := chromeCookiePath()
	if err != nil {
		return nil, err
	}

	cookies, err := chrome.ReadCookies(cookiePath, kooky.Valid, kooky.Name("session"), kooky.Domain(".adventofcode.com"))
	if err != nil {
		return nil, err
	}

	if len(cookies) != 1 {
		return nil, fmt.Errorf("session cookie not found or too many results. Got %d, want 1, ensure that you are logged in", len(cookies))
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2020/day/%d/input", day), nil)
	if err != nil {
		return nil, err
	}

	sessionToken := cookies[0].HTTPCookie()
	req.AddCookie(&sessionToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer mustClose(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code %s: %s", resp.Status, body)
	}

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func mustClose(c io.Closer) {
	if c == nil {
		return
	}

	if err := c.Close(); err != nil {
		panic(fmt.Errorf("error closing io.Closer: %w", err))
	}
}

func abort(err error) {
	fmt.Printf("%s\n\nsyntax: go run gen/problem.go <day> <a|b>\n", err.Error())
	os.Exit(1)
}
