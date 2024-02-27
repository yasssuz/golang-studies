package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println(githubInfo("tebeka"))
	fmt.Println(githubInfo("yasssuz"))
}

func validateRequest(resp *http.Response, err error) {
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}
}

func githubInfo(login string) (string, int, error) {
	url, _ := url.JoinPath("https://api.github.com/users/", url.PathEscape(login))
	resp, err := http.Get(url)

	validateRequest(resp, err)

	// fmt.Printf("content type: %s\n", resp.Header.Get("Content-Type"))
	// fmt.Println("response: ", resp.Body)

	var decodedResp struct {
		Name         string
		Public_Repos int
		// NumRepos int `json:"public_repos"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&decodedResp); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}

	return decodedResp.Name, decodedResp.Public_Repos, err
}

// Normal struct
// type Reply struct {
// 	Name         string
// 	Public_Repos int
// }

/*

JSON -> Go
true/false <-> true/false
string <-> string
null <-> nil
///array <-> []any ([]interface{})
object <-> map[string]any, struct

encoding/json API
JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal
Go -> io.Writer -> JSON: json.Encoder
Go -> []byte -> JSON: json.Marshal
*/
