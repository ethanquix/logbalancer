package lbdestinations

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethanquix/logbalancer/gen/go/pkg/model/pb_logs"
	"github.com/ethanquix/logbalancer/pkg/utils"
	"github.com/slack-go/slack"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Slack struct {
	client *slack.Client
}

func NewSlack(token string) *Slack {
	client := slack.New(token)
	return &Slack{client: client}
}

func (s *Slack) SendTo(channel string) func(incomingLog *pb_logs.RuntimeLogs) error {
	return func(incomingLog *pb_logs.RuntimeLogs) error {
		return s.Send(channel, incomingLog)
	}
}

func (s *Slack) Send(channel string, incomingLog *pb_logs.RuntimeLogs) error {
	_, _, _, err := s.client.SendMessage(channel, slack.MsgOptionAttachments(formatLogAsSlackMessage(incomingLog)))
	if err != nil {
		return fmt.Errorf("while sending slack message: %w", err)
	}
	return nil
}

func formatLogAsSlackMessage(incomingLog *pb_logs.RuntimeLogs) slack.Attachment {
	// Map severity levels to colors
	severityColor := map[pb_logs.Severity]string{
		pb_logs.Severity_SEVERITY_INFO:  "#36a64f", // Green
		pb_logs.Severity_SEVERITY_WARN:  "#ffb000", // Yellow
		pb_logs.Severity_SEVERITY_ERROR: "#e01e5a", // Red
		pb_logs.Severity_SEVERITY_DEBUG: "#5c9ded", // Blue
	}

	// Format the incomingLog date
	logDate := time.Unix(incomingLog.LogDate.Seconds, 0).Format(time.RFC3339)

	// Convert context and tags to fields
	fields := []slack.AttachmentField{
		{Title: "Log Date", Value: logDate, Short: true},
		{Title: "Severity", Value: utils.SeverityToString(incomingLog.Severity), Short: true},
		{Title: "Source", Value: incomingLog.Source, Short: true},
		{Title: "Path", Value: incomingLog.Path, Short: true},
		{Title: "Message", Value: incomingLog.Message, Short: false},
	}

	// Append Context fields
	for k, v := range incomingLog.Context {
		fields = append(fields, slack.AttachmentField{Title: fmt.Sprintf("Context: %s", k), Value: v, Short: true})
	}

	// Append Tag fields
	for k, v := range incomingLog.Tags {
		fields = append(fields, slack.AttachmentField{Title: fmt.Sprintf("Tag: %s", k), Value: v, Short: true})
	}

	// Create a Slack attachment
	attachment := slack.Attachment{
		Color:  severityColor[incomingLog.Severity], // Apply color based on severity
		Fields: fields,
		Footer: "Runtime Logs System",
		Ts:     jsonTime(incomingLog.LogDate),
	}

	return attachment
}

func jsonTime(t *timestamppb.Timestamp) json.Number {
	return json.Number(fmt.Sprintf("%d", t.Seconds))
}
