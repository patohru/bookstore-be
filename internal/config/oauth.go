package config

type OauthConfig struct {
	CLIENT_ID				string `env:"CLIENT_ID"`
	CLIENT_SECRET			string `env:"CLIENT_SECRET"`
	CLIENT_CALLBACK_URL		string  `env:"CLIENT_CALLBACK_URL"`
}
