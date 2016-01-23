package main

const BaseTmplStr = `
{{ define "base" }}
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    {{ template "title" . }}
    <link href="/assets/css/bootstrap.min.css" rel="stylesheet">
    <link href="/assets/css/skeddy.css" rel="stylesheet">
  </head>

  <body>
    <nav class="navbar navbar-default navbar-fixed-top">
      <div class="container">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="#">Skeddy</a>
        </div>
        <div id="navbar" class="collapse navbar-collapse pull-right">
          <ul class="nav navbar-nav"></ul>
        </div>
      </div>
    </nav>
    <div class="container">
    {{ template "content" . }}
    </div>

    <footer class="footer">
      <div class="container">
        <p class="text-muted">Copyright © skeddy All Right Reserved 2016</p>
      </div>
    </footer>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
    <script src="/assets/js/bootstrap.min.js"></script>
</body>
</html>
{{ end }}
`

const ViewTmplStr = `
{{ define "title"}}<title>Skeddy - Entries</title>{{ end }}
{{ define "content" }}
  <div class="page-header">
    <div class="row">
    <h4 class="col-lg-11">Scheduler entries</h4>
    <div class="pull-right col-lg-1"><a href="/new" class="btn btn-primary">Add</a></div>
    </div>
  </div>
  <p class="lead">
    <table class="table table-condensed table-hover">
    {{ range $key, $value := . }}
        <tr>
          <td>{{$value.Expression}}</td>
          <td>{{$value.Endpoint}}</td>
          <td>{{$value.Payload}}</td>
          <td><a href="/edit/{{$value.ID}}" class="btn btn-default">Edit</a></td>
          <td><a href="/delete/{{$value.ID}}" class="btn btn-default">Delete</a></td>
        </tr>
    {{ end }}
    </table>
  </p>
{{ end }}
`

const EditTmplStr = `
{{ define "title"}}<title>Skeddy - Edit Entries</title>{{ end }}
{{ define "content" }}
<div class="page-header"> <h4>Edit entries</h4> </div>
<p class="lead">
<form class="form-horizontal" action="/save/" method="post">
  <input type="hidden" class="form-control" name="id" value="{{.ID}}">
  <div class="form-group">
    <label class="col-sm-2 control-label">Cron Expression</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" placeholder="* * * * *" name="expression" value="{{.Expression}}">
    </div>
  </div>

  <div class="form-group">
    <label class="col-sm-2 control-label">Endpoint</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" placeholder="Endpoint URL" name="endpoint" value="{{.Endpoint}}">
    </div>
  </div>

  <div class="form-group">
    <label class="col-sm-2 control-label">Payload</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" placeholder="Endpoint URL" name="payload" value="{{.Payload}}">
    </div>
  </div>

  <div class="form-group">
    <div class="col-sm-offset-2 col-sm-10">
      <input type="submit" class="btn btn-default"/>
      <a href="/" class="btn btn-default">Cancel</a>
    </div>
  </div>
</form>
</p>
{{ end }}
`
