package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	vkOauthConfig *oauth2.Config
	randomState   string
)

func convertUserId(token *oauth2.Token) (int, error) {
	userId, err := strconv.ParseFloat(fmt.Sprintf("%v", token.Extra("user_id")), 64)

	if err != nil {
		return 0, err
	}

	return int(userId), nil
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `<html><body><a href="/login"> vk log in </a></body></html`
	w.Write([]byte(html))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	codeURL := vkOauthConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, codeURL, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != randomState {
		log.Println("state is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}

	token, err := vkOauthConfig.Exchange(context.TODO(), r.FormValue("code"))

	if err != nil {
		log.Println("couldnot get token")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	userId, err := convertUserId(token)

	log.Println("User authenticated: " + strconv.Itoa(userId))

	if err != nil {
		log.Println("cant convert user id")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	params := "v=5.52&access_token=" + token.AccessToken + "&user_ids=" + strconv.Itoa(userId)
	resp, err := http.Get("https://api.vk.com/method/users.get?" + params)

	if err != nil {
		panic("cant get token")
	}

	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, "Response: %s", content)

}

func init() {
	err := godotenv.Load()

	vkOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("VK_CLIENT_ID"),
		ClientSecret: os.Getenv("VK_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/auth",
		Scopes:       []string{},
		Endpoint:     vk.Endpoint,
	}

	randomState = uuid.New().String()

	if err != nil {
		panic("Error loand .env file")
	}

	log.Println("starting....")
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/auth", handleCallback)
	http.ListenAndServe(":8080", nil)

}
