//go:build integration_test || chrome_test

package webdriver_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/dsmontoya/webdriver"
	"github.com/dsmontoya/webdriver/api"
	"github.com/dsmontoya/webdriver/test/assert"
	"github.com/dsmontoya/webdriver/test/testserver"
)

var platforms = []string{
	"mac64_m1",
}

func TestChrome(t *testing.T) {
	ctx := context.Background()
	chromeDriverPath := os.Getenv("CHROME_DRIVER_PATH")
	chrome := webdriver.NewChrome(chromeDriverPath)
	err := chrome.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer chrome.Stop()

	client := webdriver.NewAPIClient()

	sessionID, deleteSession := newSession(t, ctx, client, api.NewNewSessionRequestCapabilities())
	defer deleteSession()

	testServer := testserver.NewServer("test/testserver/html")
	testServer.Start()
	defer testServer.Stop()

	navigation := client.NavigationApi

	_, httpResponse, err := navigation.NavigateTo(ctx, sessionID).UrlRequest(*api.NewUrlRequest(testServer.WhereIs("click/submit_form"))).Execute()
	if err != nil {
		t.Fatal(err)
	}
	assert.HTTPResponseIsValid(t, httpResponse)

	nameInputElementID := findElement(t, ctx, client, sessionID, api.LOCATORSTRATEGY_CSS_SELECTOR, "#name")

	_, httpResponse, err = client.ElementsApi.ElementSendKeys(ctx, sessionID, nameInputElementID).ElementSendKeysRequest(*api.NewElementSendKeysRequest("hello")).Execute()
	if err != nil {
		t.Fatalf("unable to send keys: %v", err)
	}
	assert.HTTPResponseIsValid(t, httpResponse)

	submitButtonElementID := findElement(t, ctx, client, sessionID, api.LOCATORSTRATEGY_CSS_SELECTOR, "#submit")

	_, httpResponse, err = client.ElementsApi.ElementClick(ctx, sessionID, submitButtonElementID).RequestBody(map[string]map[string]interface{}{}).Execute()
	if err != nil {
		t.Fatalf("failed to click submit button: %v", err)
	}
	assert.HTTPResponseIsValid(t, httpResponse)

	waitForCondition(t, 1*time.Millisecond, 10*time.Millisecond, checkTitle(ctx, "Submitted!", sessionID, navigation))

}

type waitCondition func(next bool) (isValid bool, err error)

func waitForCondition(t *testing.T, pollInterval time.Duration, timeout time.Duration, condition waitCondition) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	next := true
	for {
		isValid, err := condition(next)
		if err != nil {
			t.Fatalf("error waiting for condition: %v", err)
		}
		if isValid {
			break
		}

		select {
		case <-time.After(pollInterval):
		case <-ctx.Done():
			next = false
		}
	}
}

func checkTitle(ctx context.Context, title string, sessionID string, navigation api.NavigationApi) func(bool) (bool, error) {
	return func(next bool) (bool, error) {
		titleResp, httpResponse, err := navigation.GetPageTitle(ctx, sessionID).Execute()
		if err != nil {
			return false, fmt.Errorf("error trying to get page title: %w", err)
		}

		if statusCode := httpResponse.StatusCode; statusCode > 399 {
			return false, fmt.Errorf("http status code %d", statusCode)
		}

		gotTitle := titleResp.GetValue()
		isValid := title == gotTitle
		if !isValid && !next {
			return false, fmt.Errorf("title should be %s, but got %s", title, gotTitle)
		}

		return isValid, nil
	}
}

func newSession(t *testing.T, ctx context.Context, client *api.APIClient, capabilities *api.NewSessionRequestCapabilities) (string, func()) {
	sessions := client.SessionsApi

	newSessionRequest := *api.NewNewSessionRequest(*capabilities)
	apiResponse, httpResponse, err := sessions.CreateSession(ctx).NewSessionRequest(newSessionRequest).Execute()
	if err != nil {
		t.Fatalf("error when creating a new session %v", err)
	}
	if statusCode := httpResponse.StatusCode; statusCode > 399 {
		t.Fatalf("got http status code %d", statusCode)
	}
	sessionID := apiResponse.Value.SessionId
	if sessionID == "" {
		t.Fatal("empty session id")
	}
	cleanup := func() {
		sessions.DeleteSession(ctx, sessionID).Execute()
	}
	return sessionID, cleanup
}

func findElement(t *testing.T, ctx context.Context, client *api.APIClient, sessionID string, using api.LocatorStrategy, value string) string {
	findElement, httpResponse, err := client.ElementsApi.FindElement(ctx, sessionID).FindElementRequest(*api.NewFindElementRequest(using, value)).Execute()
	if err != nil {
		t.Fatal(err)
	}
	assert.HTTPResponseIsValid(t, httpResponse)

	return findElement.Value.GetElement606611e4A52e4f735466cecf()
}
