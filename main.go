package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token     string
	replySelf string
)

func init() {
	flag.StringVar(&replySelf, "rs", "nil", "Allow bot to reply to self.")
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	Token = "OTUxMTYzNDQ1MjAwNDMzMjEy.Yijd_Q.LjrBL8_tBU_eBU4t7XCCfeI1Bug"
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
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
}
