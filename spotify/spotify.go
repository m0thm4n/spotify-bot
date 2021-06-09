package spotify

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/m0thm4n/spotify-bot/utils"
	"github.com/sirupsen/logrus"
	"github.com/zmb3/spotify"
	"log"
	"net/http"
	"os"
)


const redirectURI = "http://localhost:8888/callback"

var (
	auth     = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate, spotify.ScopeUserLibraryModify, spotify.ScopePlaylistModifyPrivate)
	ch       = make(chan *spotify.Client)
	state    = "ringdingthing"
)

func completeAuth(w http.ResponseWriter, r *http.Request) {
	token, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatalln(err)
	}

	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// Use the token to get an authenticated client
	client := auth.NewClient(token)
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}

// Auth authenticates with Spotify and refreshes the token
func spotifyAuth(s *discordgo.Session, m *discordgo.Message) *spotify.Client {
	clientID := utils.EnvVar("SPOTIFY_CLIENTID", "")
	secret := utils.EnvVar("SPOTIFY_SECRET", "")

	fmt.Println(clientID, secret)

	if clientID == "" || secret == "" {
		fmt.Println("Please configure your Spotify client ID and secret in the config file at C:\\workspace\\go\\src\\spotify-bot\\")
		os.Exit(1)
	}

	// shouldRefresh, err := cmd.Flags().GetBool("refresh")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	fmt.Println("Getting token...")
	auth.SetAuthInfo(clientID, secret)
	http.HandleFunc("/callback", completeAuth)
	go http.ListenAndServe(":8888", nil)
	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by clicking the following link:", url)
	_, _ = s.ChannelMessageSend(m.ChannelID, "Please log in to Spotify by clicking the following link: "+url)
	//wait for auth to finish
	client := <-ch
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatalln(err)
	}

	// conf.Token = *token
	// marshalToken, err := json.Marshal(conf.Token)
	if err != nil {
		log.Fatalln(err)
	}
	// viper.Set("auth", string(marshalToken))
	// if err := viper.WriteConfigAs(cfgFile); err != nil {
	// 	glog.Fatal("Error writing config:", err)
	// }
	fmt.Println("Login successful as", user.ID)
	_, _ = s.ChannelMessageSend(m.ChannelID, "Login successful as "+user.ID)

	return client
}

func Play(s *discordgo.Session, m *discordgo.Message) {
	client := spotifyAuth(s, m)

	err := client.Play()
	if err != nil {
		logrus.Println(err)
	}
}