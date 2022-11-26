package sentry

import (
	"github.com/getsentry/sentry-go"
	"github.com/nasermirzaei89/env"
	"github.com/pkg/errors"
)

func initSentry() { //nolint:deadcode
	//nolint:exhaustivestruct
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         env.GetString("SENTRY_DSN", "https://41bd81d832a8486f995b8d694c143cd5@o1211294.ingest.sentry.io/6347462"),
		SampleRate:  env.GetFloat64("SENTRY_SAMPLE_RATE", 1),
		Environment: string(env.Environment()),
	})
	if err != nil {
		panic(errors.Wrap(err, "error on init sentry"))
	}
}
