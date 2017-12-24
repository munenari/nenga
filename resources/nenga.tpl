<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="UTF-8">
<link rel="stylesheet" href="nenga.css" type="text/css">
<title>年賀状宛名</title>
</head>
<body>

{{ range .Destinations }}
<section class="sheet">
	<p class="郵便番号">{{ .Postcode }}</p>
	<p class="住所">{{ noescape .Address }}</p>
	<div class="宛名">
		{{ range .Names }}
			<p>{{ . }}</p>
		{{ end }}
	</div>
	<div class="差出人住所氏名">
		<p class="差出人住所">{{ noescape $.Sender.Address }}</p>
		{{ range $.Sender.Names }}
			<p class="差出人">{{ . }}</p>
		{{ end }}
	</div>
		<p class="差出人郵便番号">{{ $.Sender.Postcode }}</p>
</section>
{{ end }}

</body>
</html>

