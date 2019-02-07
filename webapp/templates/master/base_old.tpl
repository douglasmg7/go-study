{{define "base_old"}}
<!doctype html>
<html lang='en'>
    <head>
        <meta charset='utf-8'>
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>{{template "title" .}}</title>
        <link rel='stylesheet' href='/static/css/bulma.min.css'>
        <!-- <link rel='shortcut icon' href='/static/img/favicon.ico' type='image/x-icon'> -->
        <!-- <link rel='stylesheet' href='https://fonts.googleapis.com/css?family=Ubuntu+Mono:400,700'> -->
    </head>
    <body>
        <header>
            <h1><a href='/'>Snippetbox</a></h1>
        </header>
        <nav>
            <a href='/'>Home</a>
        </nav>
        <section>
            {{template "body" .}}
        </section>
        {{template "footer" .}}
        <script src="/static/js/main.js" type="text/javascript"></script>
    </body>
</html>
{{end}}