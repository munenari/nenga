<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="UTF-8">
<link rel="stylesheet" href="nenga.css" type="text/css">
<title>年賀状宛名</title>
</head>
<body>

{{ range  }}
<section class="sheet">
	<p class="郵便番号">1234567</p>
	<p class="住所">鹿児島県鹿児島市○○町一―一<br />桜島マンション一〇一号</p>
	<div class="宛名">
		<p>山田　太郎　様</p>
		<p>　　　花子　様</p>
	</div>
	<div class="差出人住所氏名">
		<p class="差出人住所">{{ .Sender.Address }}</p>
		{{ range .Sender.Names }}
			<p class="差出人">{{ . }}</p>
		{{ end }}
	</div>
		<p class="差出人郵便番号">{{ .Sender.Postcode }}</p>
</section>

</body>
</html>

