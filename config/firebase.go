package config

// Go puede importar automaticamente los paquetes necesarios, pero tambien eliminarlos

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

const serviceAccountKey = `{
  "type": "service_account",
  "project_id": "modernlanguagesad2024",
  "private_key_id": "11e1a695a012dfbc3ce21550673e2de216c623ec",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC8C//paaq2xFmZ\nSLJy642QNiCmK4xQ9oH/gEfSJTemjgzdVcdqkyqHf5iz0rZCn+jhabuLcym3sfv7\niUL3cYaAQAPUgWie+MPBed37ZE7bhxugxwtaA5TCQJ4NVqLRNAi9Qylpz+tgGiL7\ny+AznREWLek2qoYzW5LOR5c+0veJz3TVYaoE2/EKlrssUbmF9zSIr41iHdDwKyWi\n2N8rYXxVIj9ZU258K7DO6BcBoldf2Xrpg6hstCzy1Ogqm5awnxI0F2wGOcYZxRNx\n4pu98SLsd6AHy1aFvT4wRxJtbAtPiiRCvI8fEV5Z0GrtRorEGbrEkiOD/YoRL/RJ\nGgFjULztAgMBAAECggEACo8CmGNlAQHRusV2CPWIoNjRM34Ne7xOUf+CHjQVxbCD\nML2wxp9Vy37oHuy4FtHuHGi4mEdn8mxi0YhZrhV977STInr+WggYzG511DhDz0LC\nU+A+lxgitwI0KFvoJNFZo5riTT0t97BFV5hmfITCDASZltwIpgNAdnxW6zb4TDtf\nraR1RpZVoHhm0R8XLGXu8YklwLzarV9E+tKmTsAqiPpcBkpQstLhDzmZASIyyHTQ\nYR/jNVYRfHxjaIDd9Z0May8MMoB+k+lrDYAKRcpoP++7MGbj3uczeXa/ZwVigKra\nXFlZxDPQAZ4082PRHs2uF4H4KsxXPFKf7Qb1qINjYQKBgQDdGo02TtsOnmDQq7mB\nkbi6GqNM6Es1bxhjspriZRD3rEjDaUip2I8I2Zh1hMq7uWPeGXQmI1Ngb3e0wQoH\nZkD43Wt2NVt62v4hBoYiePVYHEJCMRT/CjHA+F9G+r5j7W2Pq5MlpkRhMprfhaHO\nhnS2G16b0pGRP/+Gu6ofcBZ7kQKBgQDZudLA1f9K7Y810fAk4axoQ1mIqTngQbwn\nvs4H/mNxsTGIySocBNcFPf/DD76yXXm3iAED+zXdRn0m0/o6WVAIRzTp8H88G+l9\neBL9k7vqAJ4mpnAa5yaK16FsjLM6sM2zFQfxC5uRuKN7IXcWDomanTVEyNHCfAYl\n2govF3slnQKBgQCmmC72UIw1g5gnInzv58BsWU3Z0N16MW8LD7QsCUKzmxhoRJf/\ngbCkVm/NfHTU5WceHceqRMu5ycNEan8H/D/QQ+WId8Op2f/JGoVOzBzwerph7WGI\nmUHLyFjhvmReHrOCoSCKKbE0G1EeQi4c0qp13EF43zvXWYVLG3OfpHEMQQKBgDY8\n6lmLdctunoxjvhUR+ucGUBmRPo+1EHA+QgkPwCokYBiZNFpCgmYV6c7n7zNzwcyM\nnuQlAmgbFTZE/ELo2N0XEAFvHeMVePb+oIx27wr+GGe8cpThHGLeEKJ/8m4eeT+x\nb4cGwzxr6J3V2lquSG0x2IFIHaj7SHTJelv813qpAoGBALV4bOhg0aGxKnQs2PQg\n2OrBuel+LaORXW0Go0MtTSUsfs0HplQYaPppJLmpxCzTglPyqlxrHeqEXZDvgpvt\nSx0+wNQKfg0fpkM4b/R0GSBFN0Iukk83NcV0aNjbPxqNtbv9MdNg1X6gxUOkpaRF\nCX2WDLEm5X0u9T+3dXuKI6P3\n-----END PRIVATE KEY-----\n",
  "client_email": "firebase-adminsdk-7vryn@modernlanguagesad2024.iam.gserviceaccount.com",
  "client_id": "114299507961950141172",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-7vryn%40modernlanguagesad2024.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

func InitializeFirebaseApp() {
	opt := option.WithCredentialsJSON([]byte(serviceAccountKey))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("Failed to Initialize Firebase App: %v ", err)
	}
	FirebaseApp = app
}

func GetAuthClient(app *firebase.App) *auth.Client {
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal("Failed to Initialize Firebase Auth: %v ", err)
	}
	return client
}
