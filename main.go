package main

import (
	"log"
	"os"

	alidns "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/joho/godotenv"
)

func createDnsClient() *alidns.Client {
	accessKeyId := os.Getenv("BAO_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("BAO_ACCESS_KEY_SECRET")
	endpoint := os.Getenv("BAO_ENDPOINT")

	config := openapi.Config{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
		Endpoint:        &endpoint,
	}

	client, err := alidns.NewClient(&config)

	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	return client
}

func updateDnsRecord(client *alidns.Client, value string) {
	recordId := os.Getenv("BAO_RECORD_ID")
	recordType := "TXT"
	recordRR := "_acme-challenge"
	recordTTL := int64(600)

	request := alidns.UpdateDomainRecordRequest{
		RecordId: &recordId,
		Type:     &recordType,
		Value:    &value,
		RR:       &recordRR,
		TTL:      &recordTTL,
	}
	_, err := client.UpdateDomainRecord(&request)

	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
}

func main() {
	log.SetPrefix("[BAO] ")

	certbotValidation := os.Getenv("CERTBOT_VALIDATION")

	if certbotValidation == "" {
		log.Fatalln("No validation value found")
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dnsClient := createDnsClient()

	updateDnsRecord(dnsClient, certbotValidation)
}
