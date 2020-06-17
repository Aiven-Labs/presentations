# Aiven Demo

## What we will do:

1.	Set up a Kafka cluster with Aiven
2.	Use a Twitter Producer to create real time content
3.	Use Kafdrop as a Web Dashboard

###	Creating the cluster

1.	Click to add a Service into your project and select **Kafka (version 2.5)**
2.	Choose any plan and/or region (you have credits so go wild if you like!)
3.	Once the cluster is ready, download the:
>  1.	Access Key
>  2.	Access Secret
>  3.	Certificate
4.	Make a note of the `Service URI`
5.	Scroll to the bottom and under User Configuration, set auto create topics to true

###	Setting up the Twitter Producer

1.	Download the code from our Github Presentations repo
2.	Create an application at http://dev.twitter.com
3.	Add the details to config.yaml
>  1.	consumerKey
>  2.	consumerSecret
> 3.	accessToken
>  4.	accessSecret
>  5.	kafkaURI (Your Service URI)
>  6.	kafkaCertPath
>  7.	kafkaKeyPath
>  8.	kafkaPEMPath
>  9.	kafkaTopic
>  10.	terms (an array of things you want to watch Twitter for)
  4.	Run `go build -o twitterer`
  5.	Run twitterer and expect to see **all of the output**
  
###	Using the Aiven Console

1.	Log on at http://console.aiven.com
2.	Click on your Kafka Service -> Topics
3.	Select your Twitter topic
4.	Click Fetch Messages
> You may need to select your format (json) and an offset
    
###	Using Kafdrop for an overview

1.	Follow the instructions at [https://help.aiven.io/en/articles/3988677-using-kafdrop-web-ui-with-aiven-kafka](https://help.aiven.io/en/articles/3988677-using-kafdrop-web-ui-with-aiven-kafka) inside the `kafdrop` folder
2. Head to localhost:9000 and explore!

