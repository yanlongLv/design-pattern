package bridge

type IMsgSender interface {
	Send(msg string) error
}

type EmailSender struct {
	email []string
}

func NewEmailSender(email []string) *EmailSender {
	return &EmailSender{email: email}
}

func (s *EmailSender) Send(msg string) error {
	return nil
}

type Inotification interface {
	Notify(msg string) error
}

type ErrorInotifyCation struct {
	sender IMsgSender
}

func NewErrorInotifyCation(sender IMsgSender) *ErrorInotifyCation {
	return &ErrorInotifyCation{sender: sender}
}

func (e *ErrorInotifyCation) Send(msg string) error {
	return e.sender.Send(msg)
}
