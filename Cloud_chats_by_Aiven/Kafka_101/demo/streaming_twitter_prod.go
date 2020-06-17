package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/viper"
)

func streamRunner(stream *twitter.Stream) {
	kProd := ConnectKafka()
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(msg *twitter.Tweet) {
		key := []byte(msg.IDStr)
		tweetBytes := new(bytes.Buffer)
		json.NewEncoder(tweetBytes).Encode(msg)
    if viper.GetBool("verbose") == true {
      fmt.Println(msg.Text)
    }
		value := tweetBytes.Bytes()
		tweetBytes.Bytes()

		go WriteMessage(kProd, key, value)
	}
	fmt.Println("Stream started")

	for message := range stream.Messages {
	 	// fmt.Println(message)
	 	demux.Handle(message)
	}
}

func main() {
	flags := struct {
		consumerKey    string
		consumerSecret string
		accessToken    string
		accessSecret   string
	}{}

	loadConfig()
	fmt.Println(viper.GetString("kafkaSecretPath"))
	flag.StringVar(&flags.consumerKey, "consumer-key", viper.GetString("consumerKey"), "Twitter Consumer Key")
	flag.StringVar(&flags.consumerSecret, "consumer-secret", viper.GetString("consumerSecret"), "Twitter Consumer Secret")
	flag.StringVar(&flags.accessToken, "access-token", viper.GetString("accessToken"), "Twitter Access Token")
	flag.StringVar(&flags.accessSecret, "access-secret", viper.GetString("accessSecret"), "Twitter Access Secret")
  viper.SetDefault("terms", []string{"brexit", "trump", "covid"})
	flag.Parse()
	flagutil.SetFlagsFromEnv(flag.CommandLine, "TWITTER")

	if flags.consumerKey == "" || flags.consumerSecret == "" {
		log.Fatal("Application Access Token required")
	}
	config := oauth1.NewConfig(flags.consumerKey, flags.consumerSecret)
	token := oauth1.NewToken(flags.accessToken, flags.accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter Client
	client := twitter.NewClient(httpClient)

	filterParams := &twitter.StreamFilterParams{
		Track: viper.GetStringSlice("terms"),
	}
	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal("Failed to open stream")
	}
	go streamRunner(stream)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
}
