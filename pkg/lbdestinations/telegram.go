package lbdestinations

import (
	"bytes"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/ethanquix/logbalancer/gen/go/pkg/model/pb_logs"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Telegram struct {
	bot *tgbotapi.BotAPI
}

func NewTelegram(token string) (*Telegram, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("while creating telegram bot: %v", err)
	}

	return &Telegram{
		bot: bot,
	}, nil
}

func (t *Telegram) SendTo(to int64) func(incomingLog *pb_logs.RuntimeLogs) error {
	return func(incomingLog *pb_logs.RuntimeLogs) error {
		return t.Send(to, incomingLog)
	}
}

func (t *Telegram) Send(to int64, incomingLog *pb_logs.RuntimeLogs) error {
	msg := tgbotapi.NewMessage(to, FormatRuntimeLogsToHTML(incomingLog))
	msg.ParseMode = tgbotapi.ModeHTML
	_, err := t.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("while sending message: %v", err)
	}
	return nil
}

// FormatSeverity converts Severity enum to string
func FormatSeverity(s pb_logs.Severity) string {
	return strings.Title(strings.ToLower(s.String()))
}

// FormatRuntimeLogsToHTML formats RuntimeLogs to compact HTML suitable for Telegram
func FormatRuntimeLogsToHTML(log *pb_logs.RuntimeLogs) string {
	var buffer bytes.Buffer
	buffer.WriteString("<b>Runtime Log</b>\n")
	buffer.WriteString(fmt.Sprintf("<b>Date:</b> %s\n", formatTimestamp(log.LogDate)))
	buffer.WriteString(fmt.Sprintf("<b>Severity:</b> %s\n", FormatSeverity(log.Severity)))
	buffer.WriteString(fmt.Sprintf("<b>Source:</b> %s\n", html.EscapeString(log.Source)))
	buffer.WriteString(fmt.Sprintf("<b>Message:</b> %s\n", html.EscapeString(log.Message)))

	if len(log.Context) > 0 {
		buffer.WriteString("<b>Context:</b>\n<pre>")
		for k, v := range log.Context {
			buffer.WriteString(fmt.Sprintf("%s: %s\n", html.EscapeString(k), html.EscapeString(v)))
		}
		buffer.WriteString("</pre>")
	}

	if log.Path != "" {
		buffer.WriteString(fmt.Sprintf("<b>Path:</b> %s\n", html.EscapeString(log.Path)))
	}
	if log.Details != "" {
		buffer.WriteString(fmt.Sprintf("<b>Details:</b> %s\n", html.EscapeString(log.Details)))
	}
	if len(log.Tags) > 0 {
		buffer.WriteString("<b>Tags:</b>\n<pre>")
		for k, v := range log.Tags {
			buffer.WriteString(fmt.Sprintf("%s: %s\n", html.EscapeString(k), html.EscapeString(v)))
		}
		buffer.WriteString("</pre>")
	}
	return buffer.String()
}

// formatTimestamp formats protobuf Timestamp to a readable string
func formatTimestamp(ts *timestamppb.Timestamp) string {
	if ts == nil {
		return "N/A"
	}
	return ts.AsTime().Format(time.RFC822)
}
