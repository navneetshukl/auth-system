package mail

import (
	"context"
	"fmt"

	mail "github.com/wneessen/go-mail"
)

// MailService defines the interface for sending emails.
type MailService interface {
	SendMail(ctx context.Context, to, subject, body string, opts ...MsgOption) error
}

// MsgOption allows customization of email messages.
type MsgOption func(*mail.Msg)

// WithHTML sets the email body as HTML content.
func WithHTML() MsgOption {
	return func(msg *mail.Msg) {
		msg.SetBodyString(mail.TypeTextHTML, "")
	}
}

// WithCC adds a CC recipient to the email.
func WithCC(cc string) MsgOption {
	return func(msg *mail.Msg) {
		if err := msg.AddCc(cc); err != nil {
			fmt.Printf("failed to add CC: %v\n", err)
		}
	}
}

// MailhogAdapter implements MailService using Mailhog (for local testing).
type MailhogAdapter struct {
	client *mail.Client
	from   string
}

// MailConfig holds configuration for the Mailhog adapter.
type MailConfig struct {
	Host string
	Port int
	From string
}

// NewMailhogAdapter creates a new Mailhog adapter that implements MailService.
func NewMailhogAdapter(config MailConfig) (MailService, error) {
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == 0 {
		config.Port = 1025
	}
	if config.From == "" {
		config.From = "test@example.com"
	}

	client, err := mail.NewClient(
		config.Host,
		mail.WithPort(config.Port),
		mail.WithTLSPolicy(mail.TLSOpportunistic),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create mail client: %w", err)
	}

	return &MailhogAdapter{
		client: client,
		from:   config.From,
	}, nil
}

// SendMail sends an email using the Mailhog SMTP server.
func (m *MailhogAdapter) SendMail(ctx context.Context, to, subject, body string, opts ...MsgOption) error {
	msg := mail.NewMsg()

	if err := msg.From(m.from); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}

	if err := msg.To(to); err != nil {
		return fmt.Errorf("failed to set recipient: %w", err)
	}

	msg.Subject(subject)

	// Apply default plain text body unless HTML is specified
	msg.SetBodyString(mail.TypeTextPlain, body)

	// Apply additional message options (e.g., HTML, CC)
	for _, opt := range opts {
		opt(msg)
	}

	if err := m.client.DialAndSend(msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}