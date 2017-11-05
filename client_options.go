package docker

import (
	"context"
	"crypto/tls"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/Unknwon/com"
	"github.com/docker/go-connections/tlsconfig"
)

type ClientOptions struct {
	host       string
	tlsConfig  *tls.Config
	apiVersion string
	stderr     *OutStream
	stdout     *OutStream
	stdin      *InStream
	context    context.Context
}

type ClientOption func(*ClientOptions)

func NewClientOptions() *ClientOptions {
	opts := &ClientOptions{
		host:       Config.Host,
		apiVersion: Config.APIVersion,
		stderr:     NewOutStream(ioutil.Discard),
		stdout:     NewOutStream(ioutil.Discard),
		stdin:      nil,
		context:    context.Background(),
	}
	if com.IsDir(Config.CertPath) {
		TLSConfig(Config.CertPath, Config.TLSVerify)(opts)
	}
	return opts
}

func Host(s string) ClientOption {
	return func(o *ClientOptions) {
		o.host = s
	}
}

func TLSConfig(certPath string, verify bool) ClientOption {
	return func(o *ClientOptions) {
		options := tlsconfig.Options{
			CAFile:             filepath.Join(certPath, "ca.pem"),
			CertFile:           filepath.Join(certPath, "cert.pem"),
			KeyFile:            filepath.Join(certPath, "key.pem"),
			InsecureSkipVerify: verify,
		}
		tlsc, err := tlsconfig.Client(options)
		if err != nil {
			log.WithError(err).
				WithField("cert_path", certPath).
				WithField("verify", verify).
				Error("Failed to create tls configuration")
			return
		}
		o.tlsConfig = tlsc
	}
}

func APIVersion(s string) ClientOption {
	return func(o *ClientOptions) {
		o.apiVersion = s
	}
}

func Stderr(stderr io.Writer) ClientOption {
	return func(o *ClientOptions) {
		if stderr == nil {
			o.stderr = nil
			return
		}
		if s, ok := stderr.(*OutStream); ok {
			o.stderr = s
			return
		}
		o.stderr = NewOutStream(stderr)
	}
}

func Stdout(stdout io.Writer) ClientOption {
	return func(o *ClientOptions) {
		if stdout == nil {
			o.stdout = nil
			return
		}
		if s, ok := stdout.(*OutStream); ok {
			o.stdout = s
			return
		}
		o.stdout = NewOutStream(stdout)
	}
}

func Stdin(stdin io.Reader) ClientOption {
	return func(o *ClientOptions) {
		if stdin == nil {
			o.stdin = nil
			return
		}
		if s, ok := stdin.(*InStream); ok {
			o.stdin = s
			return
		}
		o.stdin = NewInStream(ioutil.NopCloser(stdin))
	}
}

func ClientContext(ctx context.Context) ClientOption {
	return func(o *ClientOptions) {
		o.context = ctx
	}
}
