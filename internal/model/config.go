package model

type Config struct {
	DbConnString       string
	SentryDNS          string
	Environment        string
	RestPort           int
	IsHeroku           string
	AwsAccessKeyID     string
	AwsSecretAccessKey string
	Region             string
	Bucket             string
}
