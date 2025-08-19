// Copyright (c) 2024 Neomantra Corp

package dbn_hist

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// A **half**-closed date interval with an inclusive start date and an exclusive end date.
type DateRange struct {
	// The start date (inclusive).
	Start time.Time `json:"start"`
	// The end date (exclusive).
	End time.Time `json:"end"`
}

type RequestError struct {
	Case       string `json:"case"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Docs       string `json:"docs,omitempty"`
	Payload    string `json:"payload,omitempty"`
}

type RequestErrorResp struct {
	Detail RequestError `json:"detail"`
}

// ////////////////////////////////////////////////////////////////////////////

func (c *client) databentoGetRequest(ctx context.Context, urlStr string) ([]byte, error) {
	apiUrl, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", apiUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(c.apiKey + ":"))
	req.Header.Add("Authorization", "Basic "+auth)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	badStatusCode := resp.StatusCode != http.StatusOK

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		if badStatusCode {
			return nil, fmt.Errorf("HTTP %d %s %s %w", resp.StatusCode, resp.Status, string(body), err)
		}
		return nil, err
	}

	if badStatusCode {
		return nil, fmt.Errorf("HTTP %d %s %s", resp.StatusCode, resp.Status, string(body))
	}

	return body, nil
}

// ////////////////////////////////////////////////////////////////////////////

func (c *client) databentoPostFormRequest(ctx context.Context, urlStr string, form url.Values, accept string) ([]byte, error) {
	apiUrl, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	formBody := strings.NewReader(form.Encode())
	req, err := http.NewRequestWithContext(ctx, "POST", apiUrl.String(), formBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if accept != "" {
		req.Header.Set("Accept-Encoding", accept)
	}

	auth := base64.StdEncoding.EncodeToString([]byte(c.apiKey + ":"))
	req.Header.Add("Authorization", "Basic "+auth)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	badStatusCode := resp.StatusCode != http.StatusOK

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		if badStatusCode {
			return nil, fmt.Errorf("HTTP %d %s %s %w", resp.StatusCode, resp.Status, string(body), err)
		}
		return nil, err
	}

	if badStatusCode {
		return nil, fmt.Errorf("HTTP %d %s %s", resp.StatusCode, resp.Status, string(body))
	}

	return body, nil
}
