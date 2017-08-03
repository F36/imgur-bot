package imgur

import (
	"net/http"
	"fmt"
)

type authConfig struct {
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	ExpiresIn       int    `json:"expires_in"`
	TokenType       string `json:"token_type"`
	AccountUsername string `json:"account_username"`
}

func (i *Imgur) AccessTokenString(state string) string {

	if state == "" {
		return "https://api.imgur.com/oauth2/authorize?client_id=" + i.Config.ClientID + "&response_type=token"
	}

	return "https://api.imgur.com/oauth2/authorize?client_id=" + i.Config.ClientID + "&response_type=token&state=" + state
}

func (i *Imgur) SetOAuthEndpoint(endpoint string) {

	http.HandleFunc(endpoint, initialOAuthEndpoint)

	http.HandleFunc("/catch_token", catchImgurOAuthResponse)
}

func initialOAuthEndpoint(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<!DOCTYPE html><html lang="en"><head> <meta charset="UTF-8"> <title>Imgur Bot</title></head><body><script type="text/javascript">var params={}, queryString=location.hash.substring(1), regex=/([^&=]+)=([^&]*)/g, m; while (m=regex.exec(queryString)){params[decodeURIComponent(m[1])]=decodeURIComponent(m[2]);}var state=window.location.search.slice(1); queryString +="&"+state; var req=new XMLHttpRequest(); req.open('POST', 'https://' + window.location.host + '/catch_token', true); req.setRequestHeader("Content-type", "application/x-www-form-urlencoded"); req.onreadystatechange=function (e){if (req.readyState==4){if (req.status==200){window.close();}else if (req.status==400){alert('There was an error processing the token.');}else{alert('An Error Occurred, Please retry');}}}; req.send(queryString);</script></body></html>`))

}

func catchImgurOAuthResponse(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	fmt.Println(r.Form)
}
