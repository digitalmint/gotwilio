package gotwilio

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestTWiMLSmsRenderBasic
// example from: https://www.twilio.com/docs/sms/twiml#a-basic-twiml-sms-response-example
func TestTWiMLSmsRenderBasic(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, testLookupResponse)
	}))
	defer srv.Close()

	twilio := NewTwilioClient("", "")
	twilio.LookupURL = srv.URL

	// init
	var mr MessagingResponse

	// add message
	body := "hello world!"
	redirect := "https://demo.twilio.com/welcome/sms/"
	mr.Message(&TWiMLSmsMessage{
		Body:     &body,
		Redirect: &redirect,
	})

	xml, err := mr.TWiMLSmsRender()
	if err != nil {
		t.Fatalf("failed to render xml: %+v", err)
	}
	t.Logf("xml: %+v", xml)
}

func TestTWiMLSmsRenderSend2Messages(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, testLookupResponse)
	}))
	defer srv.Close()

	twilio := NewTwilioClient("", "")
	twilio.LookupURL = srv.URL

	// init
	var mr MessagingResponse

	// message 1
	mr.Message(&TWiMLSmsMessage{
		Message: "This is message 1 of 2.",
	})

	// message 2
	mr.Message(&TWiMLSmsMessage{
		Message: "This is message 2 of 2.",
	})

	xml, err := mr.TWiMLSmsRender()
	if err != nil {
		t.Fatalf("failed to render xml: %+v", err)
	}

	t.Logf("xml: %+v", xml)
}

func TestTWiMLSmsRenderSendingMMS(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, testLookupResponse)
	}))
	defer srv.Close()

	twilio := NewTwilioClient("", "")
	twilio.LookupURL = srv.URL

	// init
	var mr MessagingResponse

	// add message
	body := "Store Location: 123 Easy St."
	redirect := "https://demo.twilio.com/owl.png"
	mr.Message(&TWiMLSmsMessage{
		Body:     &body,
		Redirect: &redirect,
	})

	xml, err := mr.TWiMLSmsRender()
	if err != nil {
		t.Fatalf("failed to render message with MMS: %+v", err)
	}
	t.Logf("xml: %+v", xml)
}

func TestTWiMLSmsRenderMessageStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, testLookupResponse)
	}))
	defer srv.Close()

	twilio := NewTwilioClient("", "")
	twilio.LookupURL = srv.URL

	// init
	var mr MessagingResponse

	// add message
	action := "/SmsHandler.php"
	method := "POST"
	mr.Message(&TWiMLSmsMessage{
		Message: "Store Location: 123 Easy St.",
		Action:  &action,
		Method:  &method,
	})

	xml, err := mr.TWiMLSmsRender()
	if err != nil {
		t.Fatalf("failed to render message status: %+v", err)
	}
	t.Logf("xml: %+v", xml)
}
