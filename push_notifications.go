package pusher

import "errors"

const (
	WebhookLvlInfo  = "INFO"
	WebhookLvlDebug = "DEBUG"
)

// PushNotification is a type for requesting push notifications
type PushNotification struct {
	Interests    []string `json:"interests"`
	WebhookURL   string   `json:"webhook_url,omitempty"`
	WebhookLevel string   `json:"webhook_level,omitempty"`
	APNS         []byte   `json:"apns,omitempty"`
	GCM          []byte   `json:"gcm,omitempty"`
	FCM          []byte   `json:"fcm,omitempty"`
}

// validate checks the PushNotification has 0<Interests<11 and has a
// APNS, GCM or FCM payload
func (pN *PushNotification) validate() error {
	if 0 == len(pN.Interests) || len(pN.Interests) > 10 {
		return errors.New("PushNotification must contain between 1 and 10 Interests")
	}

	if pN.APNS == nil && pN.GCM == nil && pN.FCM == nil {
		return errors.New("PushNotification must contain a GCM, FCM or APNS payload")
	}

	return nil
}
