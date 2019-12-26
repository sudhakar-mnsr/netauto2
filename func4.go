// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.

func CountWordsAndImages(url string) (words, images int, err error) {
resp, err := http.Get(url)
if err != nil {
return 
}

doc, err := html.Parse(resp.Body)
resp.Body.Close()
if err != nil {
err = fmt.Errorf("parsing HTML: %s", err)
return
}
words, images = countWordsAndImages(doc)
return
}

