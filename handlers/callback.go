package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"twigram-go/data"
)

func HandleCallBack(rw http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")
	rw.Header().Set("Content-Type", "text/html")

	if code == "" {
		http.ServeFile(rw, r, data.FailurePath)
	} else {
		if err := GetAuthorizationToken(code); err != nil {
			fmt.Printf("Error: %v", err)
			http.ServeFile(rw, r, data.FailurePath)
		}
		http.ServeFile(rw, r, data.SuccessPath)
		// TODO: Add the user to the db, with his ChatID, AccessToken, RefreshToken, and all of his twitter info as well.
		// 		 Show the message of success, and send him a message with this success.
		// 	     Spin up a goroutine that fetches all his followers, waits on their tweets.
	}
}

func GetAuthorizationToken(code string) error {

	client := http.Client{}
	urlData := url.Values{
		"code":          {code},
		"code_verifier": {data.CodeVerifier},
		"grant_type":    {"authorization_code"},
		"client_id":     {data.ClientID},
		"redirect_uri":  {data.RedirectURI},
	}

	r, err := http.NewRequest("POST", "https://api.twitter.com/2/oauth2/token", strings.NewReader(urlData.Encode()))
	if err != nil {
		return err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Authorization", data.ClientBasicAuth)

	res, err := client.Do(r)
	if err != nil {
		return err
	}

	if !strings.Contains(res.Status, strconv.Itoa(http.StatusOK)) {
		return fmt.Errorf("something wrong happened. Status of response: %s", res.Status)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}
