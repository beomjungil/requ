package model

type HttpRequestConfig struct {
	Method  string
	Url     string
	Headers map[string]string
	Body    string
}
