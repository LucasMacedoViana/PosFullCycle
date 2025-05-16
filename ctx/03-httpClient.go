package ctx

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func Client() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req, r := http.NewRequestWithContext(ctx, "GET", "http://localhost:8000", nil)
	if r != nil {
		panic(r)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
