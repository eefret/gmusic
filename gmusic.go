/*
Copyright 2014 Kaissersoft Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gmusic

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//=======================================================================
//                          Const
//=======================================================================

const (
	servicesBase         string = "https://play.google.com/music/services/"
	searchUrl            string = "search"
	loadAllTracksUrl     string = "streamingloadalltracks"
	loadPlaylistUrl      string = "streamingloadplaylist"
	deletePlaylistUrl    string = "deleteplaylist"
	addPlaylistUrl       string = "addplaylist"
	loginUrl             string = "https://www.google.com/accounts/ClientLogin"
	playSongUrl          string = "https://play.google.com/music/play?u=0&songid=%1$s&pt=e"
	cookieFormat         string = "?u=0&xt=%1$s"
	googleLoginAuthKey   string = "Authorization"
	googleLoginAuthValue string = "GoogleLogin auth=%1$s"
)

type Gmusic struct {
	isStartup          bool
	authorizationToken string
	cookie             string
	rawCookie          string
	email              string
	passwd             string
}

func New(email string, passwd string) (g *Gmusic, err error) {
	g = &Gmusic{
		isStartup:          true,
		authorizationToken: "",
		cookie:             "",
		rawCookie:          "",
		email:              email,
		passwd:             passwd,
	}
	logged, err := g.login()
	if !logged {
		return nil, err
	}
	return g, nil
}

func (g *Gmusic) login() (logged bool, err error) {
	form := url.Values{}
	form.Add("service", "sj")
	form.Add("Email", g.email)
	form.Add("Passwd", g.passwd)

	req, err := http.NewRequest("POST", g.adjustURL(loginUrl), strings.NewReader(form.Encode()))
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	g.setCookie(resp)
	body, _ := ioutil.ReadAll(resp.Body)
	g.authorizationToken = extractAuthToken(string(body))
	return true, nil
}

func (g *Gmusic) adjustURL(baseUrl string) string {
	if strings.HasPrefix(baseUrl, servicesBase) {
		return baseUrl + fmt.Sprintf(cookieFormat, g.cookie) + "&format=jsarray"
	}
	return baseUrl
}

func extractAuthToken(response string) string {
	return response[strings.Index(response, "Auth=")+len("Auth=") : len(response)-1]
}

func (g *Gmusic) setCookie(resp *http.Response) {
	if resp.Header.Get("Set-Cookie") != "" && g.cookie == "" {
		g.rawCookie = resp.Header.Get("Set-Cookie")
		start := strings.Index(g.rawCookie, "xt=") + len("xt=")
		g.cookie = g.rawCookie[start:strings.Index(g.rawCookie[start:], ";")]
	}
}
