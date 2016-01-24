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
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.5.0/css/font-awesome.min.css">
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
        <p class="text-muted">Contribute to skeddy <a href="http://github.com" id="ghlogo"><span class="fa fa-github" style="font-size:20px"></i></a></p>
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
          <td><a href="/edit/{{$value.ID}}" class="btn btn-info">Edit</a></td>
          <td><a href="/delete/{{$value.ID}}" class="btn btn-danger" onclick="return confirm('Are you sure you want to delete {{$value.Expression}}')">Delete</a></td>
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
<div class="col-sm-5 pull-right" style="border:solid">
  <h4>Help document for Cron expression</h4>
  <div>
    A cron expression represents a set of times:
    <li> using 6 space-separated fields
    <br/>
      &nbsp;[seconds] [minutes] [hour] [day of month] [month] [day of week]
    <li> using predefined schedules
    <br/>
      &nbsp;@yearly (or @annually), @monthly, @weekly, @daily (or @midnight), @hourly
    <br/>
      For further information, check the documentation <a href="https://godoc.org/github.com/robfig/cron">here</a>
  </div>
</div>
<p class="lead">
<form enctype="multipart/form-data" name="edit-form" class="form-horizontal col-sm-7 pull-left" action="/save/" method="post" >
  <input type="hidden" class="form-control" name="id" value="{{.ID}}">
  <div class="form-group">
    <label class="col-sm-5 control-label">Cron Expression</label>
    <div class="col-sm-5">
      <input type="text" id="expression" class="form-control" placeholder="* * * * *" name="expression" value="{{.Expression}}" onfocusout="validateExpression()" required>
      <div id="validate"></div>
    </div>
  </div>

  <div class="form-group">
    <label class="col-sm-5 control-label">Endpoint</label>
    <div class="col-sm-5">
      <input type="url" id="url" class="form-control" placeholder="Endpoint URL" name="endpoint" value="{{.Endpoint}}" onfocusout="validateURL()" required>
      <div id="validate_url"></div>
    </div>
  </div>

  <div class="form-group">
    <label class="col-sm-5 control-label">Payload</label>
    <div class="col-sm-5">
      <input type="text" id="payload" class="form-control" placeholder="Payload" name="payload" value="{{.Payload}}" onfocusout="enableSubmission()">
      <br/>
      <input id="fileupload" type="file" title="add files" name="files" onchange="getFile()"/>
    </div>
  </div>

  <div class="form-group">
    <div class="col-sm-offset-5 col-sm-5">
      <input type="submit" class="btn btn-success" id="submitBtn" disabled/>
      <a href="/" class="btn btn-danger">Cancel</a>
    </div>
  </div>
</form>
</p>
<script src="/assets/js/skeddy.js"></script>
<script>
  enableSubmission();
</script>
{{ end }}
`

const AddTmplStr = `
{{ define "title"}}<title>Skeddy - Add Entries</title>{{ end }}
{{ define "content" }}
<div class="page-header"> <h4>Add entries</h4> </div>
<div class="col-sm-5 pull-right" style="border:solid">
  <h4>Help document for Cron expression</h4>
  <div>
    A cron expression represents a set of times:
    <li> using 6 space-separated fields
    <br/>
      &nbsp;[seconds] [minutes] [hour] [day of month] [month] [day of week]
    <li> using predefined schedules
    <br/>
      &nbsp;@yearly (or @annually), @monthly, @weekly, @daily (or @midnight), @hourly
    <br/>
      For further information, check the documentation <a href="https://godoc.org/github.com/robfig/cron">here</a>
  </div>
</div>
<p class="lead">
<form enctype="multipart/form-data" class="form-horizontal col-sm-7 pull-left" action="/add/" method="post">
  <input type="hidden" class="form-control" name="id" value="{{.ID}}">
  <div class="form-group">
    <label class="col-sm-5 control-label">Cron Expression</label>
    <div class="col-sm-5">
      <input type="text" id="expression" class="form-control" placeholder="* * * * *" name="expression" value="{{.Expression}}" onfocusout="validateExpression()" required>
      <div id="validate"></div>
    </div>
  </div>

  <div class="form-group">
    <label class="col-sm-5 control-label">Endpoint</label>
    <div class="col-sm-5">
      <input type="url" id="url" class="form-control" placeholder="Endpoint URL" name="endpoint" value="{{.Endpoint}}" onfocusout="validateURL()" required>
      <div id="validate_url"></div>
    </div>
  </div>

  <div class="form-group">
    <label class="col-sm-5 control-label">Payload</label>
    <div class="col-sm-5">
      <input type="text" id="payload" class="form-control" placeholder="Payload" name="payload" value="{{.Payload}}" onfocusout="enableSubmission()">
      <br/>
      <input id="fileupload" type="file" title="add files" name="files" onchange="getFile()"/>
    </div>
  </div>

  <div class="form-group">
    <div class="col-sm-offset-5 col-sm-5">
      <input type="submit" class="btn btn-success" id="submitBtn" disabled/>
      <a href="/" class="btn btn-danger">Cancel</a>
    </div>
  </div>
</form>
</p>

<script src="/assets/js/skeddy.js"></script>

{{ end }}
`
