package main

import (
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/russross/blackfriday/v2"
	"github.com/webview/webview_go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: md <markdown_file>")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read file: %s\n", err)
		return
	}

	// Convert Markdown to HTML
	htmlContent := blackfriday.Run(content)

	// Wrap the content with fonts and highlight.js
	htmlWithHighlight := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Markdown Viewer</title>
		<!-- Import Google Fonts -->
		<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;700&display=swap" rel="stylesheet">
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.5.1/styles/default.min.css">
		<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.5.1/highlight.min.js"></script>
		<script>hljs.highlightAll();</script>
		<style>
			body {
				font-family: 'Inter', sans-serif;
				line-height: 1.6;
				margin: 20px;
				background-color: #f5f5f5;
				color: #333;
			}
			code, pre {
				font-family: 'Menlo', 'Monaco', 'Courier New', monospace;
				font-size: 14px;
				background-color: #f0f0f0;
				padding: 5px;
				border-radius: 5px;
				color: #333;
			}
			pre {
				padding: 10px;
				overflow: auto;
			}
			h1, h2, h3 {
				font-weight: 600;
				margin-top: 1.5em;
				color: #222;
			}
			p {
				margin-bottom: 1.2em;
			}
			.markdown-content {
				max-width: 800px;
				margin: auto;
			}
		</style>
	</head>
	<body>
		<div class="markdown-content">` + string(htmlContent) + `</div>
	</body>
	</html>`

	// URL-encode the HTML content
	encodedHTML := url.PathEscape(htmlWithHighlight)

	// Create a WebView to display the encoded HTML content
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Markdown Viewer")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("data:text/html," + encodedHTML)
	w.Run()
}

