package notifier

import (
    "github.com/twilio/twilio-go"
    "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS(to, body string) error {
    client := twilio.NewRestClient()

    params := &v2010.CreateMessageParams{}
    params.SetTo(to)
    params.SetFrom("+1234567890") // Twilio number
    params.SetBody(body)

    _, err := client.ApiV2010.CreateMessage(params)
    if err != nil {
        return err
    }

    return nil
}