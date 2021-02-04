# Event

|Event Title|Date|Location|
|---|---|---|
|[Python Web Conference](ttps://pythonwebconference.com/) | 22nd-26th March | Online |

# Talk

## Elevator Pitch
What’s better than a pizza example to show how Python and Apache Kafka, a streaming platform, work together to enable reliable real-time data integration for your event-driven application? Lets dig into problems Kafka is solving, its Python libraries and prebuilt connectors together!

## Description
Code and data go together like tomato and basil; not many applications work without moving data in some way. As our applications modernise and evolve to become more event-driven, the requirements for data are changing. In this session we will explore Apache Kafka, a data streaming platform, to enable reliable real-time data integration for your applications.

We will look at the types of problems that Kafka is best at solving, and show how to use it in your own applications. Whether you have a new application or are looking to upgrade an existing one, this session includes advice on adding Kafka using the Python libraries and includes code examples (with bonus discussion of pizza toppings) to use.

With Kafka in place, many things are possible so this session also introduces Kafka Connect, a selection of pre-built connectors that you can use to route events between systems and integrate with other tools. This session is recommended for engineers and architects whose applications are ready for next-level data abilities.

## Talk Structure

* What is Kafka - Working with the pizzeria analogy
  * What is an event?
    * Receiving an Order on the phone
    * New material arrives
    * Working hours start!
    * Pizza is ready!
  * Basic setup:
    * one person who replies to calls and make pizzas - not enough - monolith
    * add person who replies to calls, pizza maker, distributor
    * sync calls, show waiting time, service unavailable
    * add person that does receipts, same message goes to two people
    * add person who does inventory, same message goes to 3 people
  * A distributed log - Topic idea
    * Multiple producers/consumers   
      * Avoid Spaghetti
  * Replay events
  * Add more people to reply to the phone and to distribute pizzas
    * Partitioning - increase producer/consumer pool
  * Microservices
    * They first talk directly to each other
      * What if something fails?
      * What if it's busy
      * Burden is in the producer/consumer     
  * Kafka enables Decoupling Producer/Consumers
    * Producer can send and do other work
    * consumer can wait and execute
    * more consumers can read the same data
    * a consumer pool can divide the work

* How do I work with Kafka and Python
  * Administration
    * Create a topic
    * Create partitions
  * Sample writer
    * parameters
    * topic registration (if not done above)
    * partition subscription
  * Sample reader
    * parameters
    * timeout
    * partition assignment
  * Brief Fake data
  * Faust - Kafka streams in Python example --> Too Much?
  * Kafkacat to read? Optional only mention

* Kafka Connect
  * What is?
    * pre-built connectors, show list
  * Sources and Sinks
  * Use cases
  * Demo with pg

## references

* [Blog Post](https://github.com/aiven/blog-posts/tree/ft_python_fake_producer/2021/01/python_fake_producer)
* [Code](https://github.com/aiven/kafka-python-fake-data-producer)


## Demo

```
FOLDER_NAME=~/kafkacerts
PROJECT_NAME=dev-advocates
CLOUD=aws-eu-south-1
KAFKA_NAME=fafka-ft
POSTGRES_NAME=pg-ft
AIVEN_PLAN_NAME=business-4

avn service create -p $AIVEN_PLAN_NAME -t kafka $KAFKA_NAME --cloud $CLOUD --project $PROJECT_NAME -c kafka_rest=true -c kafka.auto_create_topics_enable=true -c schema_registry=true -c kafka_connect=true

# Download all certificates
mkdir $FOLDER_NAME
avn service user-creds-download $KAFKA_NAME --project $PROJECT_NAME -d $FOLDER_NAME --username avnadmin

# get KAFKA URL
avn service get $KAFKA_NAME --project=$PROJECT_NAME --format '{service_uri}'

# Do Python Magic
...

python FakerKafka.py --cert-folder ~/kafkacerts --host fafka-ft-dev-advocates.aivencloud.com --port 13041 --topic-name pizza-orders --nr-messages 0 --max-waiting-time 0


# create postgreSQL
avn service create -p startup-4 -t pg $POSTGRES_NAME --cloud $CLOUD --project $PROJECT_NAME

# create connector


# terminate
avn service terminate $KAFKA_NAME --project $PROJECT_NAME --force
avn service terminate $POSTGRES_NAME --project $PROJECT_NAME --force
```
