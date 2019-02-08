{{define "base"}}
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
        <section class="hero is-primary">
          <div class="hero-body">
            <div class="container">
              <h1 class="title">
                Escola de barco a vela
              </h1>
              <h2 class="subtitle">
                Iate Clube Lagoa dos Ingleses
              </h2>
            </div>
          </div>
        </section>
        {{template "body" .}}
    </body>
</html>
{{end}}