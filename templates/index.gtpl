
{{ define "index" }}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .PageTitle }}</title>
    <nav>
        <ul>
            {{ range .Navigation }}
                <li><a href="{{ .Url }}">{{ .Name }}</a></li>
            {{ end }}
        </ul>
    </nav>
</head>
<body>
    <h1>content</h1>
</body>
</html>
{{ end }}