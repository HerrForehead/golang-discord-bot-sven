package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token     string
	replySelf string
	DeepAIkey string
)

func init() {
	flag.StringVar(&replySelf, "rs", "nil", "Allow bot to reply to self.")
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	configInfo := config{}

	config, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(config, &configInfo)

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + configInfo.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	fmt.Println("Using API key: ", configInfo.Token)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if replySelf != "nil" {
		if m.Author.ID == s.State.User.ID {
			return
		}
	}

	// If the message is "ping" reply with "Pong!"
	if strings.ToLower(m.Content) == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

	// If the message is "pong" reply with "Ping!"
	if strings.ToLower(m.Content) == "pong" {
		s.ChannelMessageSend(m.ChannelID, "ping")
	}

	if strings.ToLower(m.Content) == "hello" {
		s.ChannelMessageSend(m.ChannelID, "How about no?")
	}

	if strings.ToLower(m.Content) == "activity" {
		idea := activiteit{}

		url := "https://www.boredapi.com/api/activity"

		response, err := http.Get(url)
		if err != nil {
			log.Fatal("kon de api niet aanroepen:", err)
		}
		if response.StatusCode != http.StatusOK {
			log.Fatal("De api heeft een andere statuscode teruggegeven dan 200, namelijk ", response.Status)
		}

		defer response.Body.Close()

		data, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal("Kon de data van de api niet lezen:", err)
		}

		err = json.Unmarshal(data, &idea)
		if err != nil {
			log.Fatal()
		}

		activity := "Activity: " + idea.Activity
		pa := fmt.Sprint(idea.Participants)
		participants := "No. of people: " + string(pa)
		f := fmt.Sprint(idea.Price)
		price := "Price of activity: " + string(f)

		s.ChannelMessageSend(m.ChannelID, activity)
		s.ChannelMessageSend(m.ChannelID, participants)
		s.ChannelMessageSend(m.ChannelID, price)
	}

	if strings.ToLower(m.Content) == "catfact" {
		feitje := catfact{}

		url := "https://catfact.ninja/fact"

		response, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		if response.StatusCode != http.StatusOK {
			log.Fatal("De api heeft een andere statuscode teruggegeven dan 200, namelijk ", response.Status)
		}

		defer response.Body.Close()

		data, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(data, &feitje)
		if err != nil {
			log.Fatal(err)
		}

		s.ChannelMessageSend(m.ChannelID, feitje.Fact)
	}

	if strings.ToLower(m.Content) == "doggy" {
		doggy := doggyphoto{}

		url := "https://dog.ceo/api/breeds/image/random"

		response, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		if response.StatusCode != http.StatusOK {
			log.Fatal("De api heeft een andere statuscode teruggegeven dan 200, namelijk ", response.Status)
		}

		defer response.Body.Close()

		data, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(data, &doggy)
		if err != nil {
			log.Fatal(err)
		}

		s.ChannelMessageSend(m.ChannelID, doggy.Message)
	}

	if strings.ToLower(m.Content) == "fortunecookie" {
		cookie := fortunecookie{}

		url := "http://www.yerkee.com/api/fortune/wisdom"

		response, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		if response.StatusCode != http.StatusOK {
			log.Fatal("De api heeft een andere statuscode teruggegeven dan 200, namelijk ", response.Status)
		}

		defer response.Body.Close()

		data, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(data, &cookie)
		if err != nil {
			log.Fatal("Fout bij het openen van het bestand: ", err)
		}

		s.ChannelMessageSend(m.ChannelID, cookie.Fortune)
	}

	if strings.Contains(m.Content, "text2img") {
		configInfo := config{}

		config, err := ioutil.ReadFile("./config.json")
		if err != nil {
			log.Fatal("Fout bij het openen van het bestand: ", err)
		}

		err = json.Unmarshal(config, &configInfo)

		t2i := text2img{}

		word := strings.Split(m.Content, " ")
		fmt.Println(word[1])

		client := &http.Client{}
		form := url.Values{}
		form.Set("text", word[1])
		req, _ := http.NewRequest("GET", "https://api.deepai.org/api/text2img", strings.NewReader(form.Encode()))
		req.Header.Set("api-key", configInfo.DeepAIkey)
		res, _ := client.Do(req)

		//fmt.Println(res.Body)

		data, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("API key for T2I is: ", configInfo.DeepAIkey)
		fmt.Println(string(data))

		err = json.Unmarshal(data, &t2i)

		s.ChannelMessageSend(m.ChannelID, t2i.Output_url)
	}

}
