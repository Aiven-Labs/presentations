
build:
	go build -o twitterer

clean:
	rm -rf twitterer

run: build
	./twitterer

kafdrop:
	docker run -p 9000:9000 -e KAFKA_BROKERCONNECT=$SERVICEURI -e KAFKA_PROPERTIES="$(cat kafka.properties | base64)" -e KAFKA_TRUSTSTORE="$(cat kafka.truststore.jks | base64)" -e KAFKA_KEYSTORE="$(cat kafka.keystore.p12 | base64)" obsidiandynamics/kafdrop
