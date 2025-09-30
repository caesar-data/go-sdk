# Research

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchNewResponse">ResearchNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchGetResponse">ResearchGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchListResponse">ResearchListResponse</a>

Methods:

- <code title="post /research">client.Research.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchNewParams">ResearchNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchNewResponse">ResearchNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /research/{id}">client.Research.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchGetResponse">ResearchGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /research">client.Research.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchListParams">ResearchListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go/packages/pagination#Pagination">Pagination</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchListResponse">ResearchListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Files

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchFileNewResponse">ResearchFileNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchFileListResponse">ResearchFileListResponse</a>

Methods:

- <code title="post /research/files">client.Research.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchFileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchFileNewParams">ResearchFileNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchFileNewResponse">ResearchFileNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /research/files">client.Research.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchFileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchFileListParams">ResearchFileListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go/packages/pagination#Pagination">Pagination</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchFileListResponse">ResearchFileListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Results

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchResultGetContentResponse">ResearchResultGetContentResponse</a>

Methods:

- <code title="get /research/{id}/results/{resultId}/content">client.Research.Results.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchResultService.GetContent">GetContent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, resultID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchResultGetContentParams">ResearchResultGetContentParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go">caesar</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/caesar-go#ResearchResultGetContentResponse">ResearchResultGetContentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
