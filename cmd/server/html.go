package main

import (
	"html/template"
	"net/http"

	"github.com/bign8/games/impl"
	"github.com/gorilla/mux"
)

var frame = template.Must(template.New("frame").Parse(frameHTML))

var rootTemplate = template.Must(template.Must(frame.Clone()).Parse(rootHTML))
var gameTemplate = template.Must(template.Must(frame.Clone()).Parse(gameHTML))
var aboutTemplate = template.Must(template.Must(frame.Clone()).Parse(aboutHTML))

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	rootTemplate.Execute(w, struct {
		Games interface{}
	}{
		Games: impl.Map(),
	})
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	game, ok := impl.Get(mux.Vars(r)["slug"])
	if !ok {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	gameTemplate.Execute(w, struct {
		Game  interface{}
		Board template.HTML
	}{
		Game:  game,
		Board: template.HTML(game.Board),
	})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	aboutTemplate.Execute(w, nil)
}

var gameHTML = `
{{define "body"}}
<div class="row">
  <div class="col-md-8 col-md-push-4">
    <div class="board-wrapper">
      {{ .Board }}
      <div class="board" id="game"></div>
    </div>
  </div>
  <div class="col-md-4 col-md-pull-8">
    <div class="panel panel-default chat">
      <div class="panel-heading">
        <h3 class="panel-title">Chat</h3>
      </div>
      <div class="list-group" id="output"></div>
      <input id="input" class="form-control panel-footer" type="text" placeholder="Type Message Here...">
    </div>
    <div class="panel panel-default moves">
      <div class="panel-heading">
        <h3 class="panel-title">Moves</h3>
      </div>
      <ul class="list-group" id="moves"></ul>
    </div>
  </div>
</div>
{{end}}

{{define "code"}}
<script src="/static/protocol.js"></script>
<script src="/static/app.js"></script>
{{end}}
`

var aboutHTML = `
{{define "body"}}
<div class="row about">
  <div class="col-xs-12">
    <h1 id="about">About Game Roulette</h1>
    <p class="lead">The graphic designer’s first fucking consideration is always the size and shape of the format, whether for the printed page or for digital display. Design as if your fucking life depended on it.  If you fucking give up, you will achieve nothing. Respect your fucking craft. If you fucking give up, you will achieve nothing.</p>
    <p>The details are not the details. They make the fucking design. Sometimes it is appropriate to place various typographic elements on the outside of the fucking left margin of text to maintain a strong vertical axis. This practice is referred to as exdenting and is most often used with bullets and quotations. Respect your fucking craft. If you’re not being fucking honest with yourself how could you ever hope to communicate something meaningful to someone else? Widows and orphans are terrible fucking tragedies, both in real life and definitely in typography.  You won’t get good at anything by doing it a lot fucking aimlessly.</p>
    <p>Never let your guard down by thinking you’re fucking good enough. To go partway is easy, but mastering anything requires hard fucking work. While having drinks with Tibor Kalman one night, he told me, &ldquo;When you make something no one hates, no one fucking loves it.&rdquo; Form follows fucking function.</p>
    <p>Practice won’t get you anywhere if you mindlessly fucking practice the same thing. Change only occurs when you work deliberately with purpose toward a goal. The graphic designer’s first fucking consideration is always the size and shape of the format, whether for the printed page or for digital display. You need to sit down and sketch more fucking ideas because stalking your ex on facebook isn’t going to get you anywhere. Practice won’t get you anywhere if you mindlessly fucking practice the same thing. Change only occurs when you work deliberately with purpose toward a goal. If you’re not being fucking honest with yourself how could you ever hope to communicate something meaningful to someone else? If you’re not being fucking honest with yourself how could you ever hope to communicate something meaningful to someone else? Intuition is fucking important.</p>
    <p>What’s important is the fucking drive to see a project through no matter what. Sometimes it is appropriate to place various typographic elements on the outside of the fucking left margin of text to maintain a strong vertical axis. This practice is referred to as exdenting and is most often used with bullets and quotations. If you fucking give up, you will achieve nothing. To go partway is easy, but mastering anything requires hard fucking work.</p>
  </div>
</div>

<div class="row marketing">
  <div class="col-lg-6">
    <h4>Subheading</h4>
    <p>Donec id elit non mi porta gravida at eget metus. Maecenas faucibus mollis interdum.</p>

    <h4>Subheading</h4>
    <p>Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Cras mattis consectetur purus sit amet fermentum.</p>

    <h4>Subheading</h4>
    <p>Maecenas sed diam eget risus varius blandit sit amet non magna.</p>
  </div>
  <div class="col-lg-6">
    <h4>Subheading</h4>
    <p>Donec id elit non mi porta gravida at eget metus. Maecenas faucibus mollis interdum.</p>

    <h4>Subheading</h4>
    <p>Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Cras mattis consectetur purus sit amet fermentum.</p>

    <h4>Subheading</h4>
    <p>Maecenas sed diam eget risus varius blandit sit amet non magna.</p>
  </div>
</div>
{{end}}
`

var rootHTML = `
{{define "body"}}
<div class="jumbotron">
  <h1>Game Roulette</h1>
  <p class="lead">Pick a game and be matched with an opponent!</p>
  <p><a class="btn btn-lg btn-success" href="/play/random" role="button">Pick @ Random</a></p>
</div>

<div class="row">
  {{range .Games}}
  <div class="col-xs-6 col-md-4">
    <a href="/play/{{ .Slug }}" class="thumbnail" title="{{ .Name }}">
      <img src="/static/img/{{ .Slug }}.png" alt="{{ .Name }}">
    </a>
  </div>
  {{end}}
</div>
{{end}}
`

var frameHTML = `
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Game Roulette</title>
  <meta name="description" content="xxx">
  <meta name="author" content="Nate Woods">
  <link rel="stylesheet" href="/static/bootstrap.min.css">
  <link rel="stylesheet" href="/static/css.css">
  <link href="/static/img/favicon.ico" rel="icon" type="image/x-icon" />
</head>
<body>
  <div class="container container-narrow">
    <div class="header clearfix">
      <!--<nav>
        <ul class="nav nav-pills pull-right">
          <li role="presentation">
            <a href="/about">About</a>
          </li>
        </ul>
      </nav>-->
      <h3 id="top"><a href="/" class="text-muted">Game Roulette</a></h3>
    </div>

    {{block "body" .}}{{end}}

    <footer class="footer">
      <p>&copy; <script>document.write(new Date().getFullYear())</script> <a href="https://bign8.info/contact">Nate Woods</a></p>
    </footer>
  </div>
  {{block "code" .}}{{end}}
</body>
</html>
`
