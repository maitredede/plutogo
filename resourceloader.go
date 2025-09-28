package plutogo

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"sync"
)

type CustomResourceLoader func(url string) (*ResourceData, error)

type resourceLoaderData struct {
	err    error
	mutex  sync.Mutex
	loader CustomResourceLoader
}

type ResourceData struct {
	Bin          []byte
	Mime         string
	TextEncoding string
}

func DefaultHttpLoader(url string) (*ResourceData, error) {
	slog.Debug(fmt.Sprintf("go: loadResource url=%s", url))

	cookies, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("cookie jar create failed: %w", err)
	}
	client := &http.Client{
		Jar: cookies,
	}
	ctx := context.TODO()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("new request failed: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}
	defer res.Body.Close()

	slog.Debug(fmt.Sprintf("go: loadResource url=%s status=%s", url, res.Status))
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unhandled status : %s", res.Status)
	}

	data := ResourceData{}
	headerCT := res.Header.Values("Content-Type")
	if len(headerCT) > 0 {
		//TODO get content-type and text encoding
		data.Mime = headerCT[0]
	}
	slog.Debug(fmt.Sprintf("go: loadResource url=%s mime=%s textEncoding=%s", url, data.Mime, data.TextEncoding))
	bin, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("body read failed: %w", err)
	}
	data.Bin = bin
	return &data, nil
}
