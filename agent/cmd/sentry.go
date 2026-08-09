package cmd

import (
	"github.com/asiffer/situation/agent/config"
	"github.com/asiffer/situation/pkg/modules"
	"github.com/getsentry/sentry-go"
)

func initSentry(dsn string) error {
	return sentry.Init(
		sentry.ClientOptions{
			Dsn:              dsn,
			EnableTracing:    true,
			TracesSampleRate: 1.0,
			DisableLogs:      false,
			ServerName:       config.AgentString(),
			Release:          config.Version,
			Dist:             config.Commit,
		},
	)
}

type sentrySupervisor struct {
	span *sentry.Span
}

func (s *sentrySupervisor) StartChild(name string) modules.SchedulerSupervisor {
	return &sentrySupervisor{span: s.span.StartChild(name)}
}

func (s *sentrySupervisor) Finish() {
	s.span.Finish()
}

func (s *sentrySupervisor) SetStatus(err error) {
	if err != nil {
		s.span.Status = sentry.SpanStatusInternalError
	} else {
		s.span.Status = sentry.SpanStatusOK
	}
}

func newSentrySupervisor(span *sentry.Span) modules.SchedulerSupervisor {
	return &sentrySupervisor{span: span}
}
