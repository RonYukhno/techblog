<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="X-UA-Compatible" content="ie=edge">	
	<title>{{ .Title }}</title>

	<link rel="stylesheet" href="/static/css/style.css">
	<!-- <script src="/static/js/vue-router.js"></script> -->
</head>
<body>
		<div class="header">
			<div class="nav">
				<ul>
					{{ range $_, $val := .topHeader }}
					<li><a href="#">{{ $val }}</a></li>
					{{ end }}
				</ul>
			</div>
		</div>
		<div class="content">
	