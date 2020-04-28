include .env

export GO111MODULE=on

init: .env

.env:
		cp $@.example $@

build:
	$ env GOOS=linux GOARCH=amd64 go build main.go

gomod:
	go mod init github.com/mizuho1998/nature_remo_light

get_devices: .env
	curl -X GET "https://api.nature.global/1/devices" -H "accept: application/json" -H "Authorization: Bearer ${TOKEN}" | jq .

get_appliances: .env
	curl -X GET "https://api.nature.global/1/appliances" -H "accept: application/json" -H "Authorization: Bearer ${TOKEN}" | jq .

post_signal:
	curl -X POST "https://api.nature.global/1/signals/${SIGNAL_ID}/send" -H "accept: application/json" -H "Authorization: Bearer ${TOKEN}" | jq .
