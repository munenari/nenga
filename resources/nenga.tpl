<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="UTF-8">
<title>年賀状宛名</title>
</head>
<body>
<style>
	@page {
		size: 100mm 148mm;
		margin: 0;
		padding: 0;
	}

	body {
		margin: 0;
		padding: 0;
		font-family: 'HiraMinProN-W3', serif;
		font-style: normal;
		font-weight: 400;
	}

	.sheet {
		width: 100mm;
		height: 148mm;
		margin: 0;
		padding: 0;
		position: relative;
		page-break-after: always;
	}

	p {
		margin: 0;
		padding: 0;
	}

	.郵便番号 {
		font-size: 12pt;
		margin: 0;
		padding: 0;
		position: absolute;
		top: 12mm;
		left: 46mm;
		letter-spacing: 4.8mm;
	}

	.住所 {
		font-size: 14pt;
		margin-left: 0px;
		padding-top: 0px;
		padding-bottom: 20mm;
		line-height: 1.5em;
		writing-mode: vertical-rl;
		-webkit-writing-mode: vertical-rl;
		text-orientation: upright;
		-webkit-text-orientation: upright;
		position: absolute;
		top: 30mm;
		right: 8mm;
	}

	.宛名 {
		font-size: 22pt;
		letter-spacing: 0.1em;
		position: absolute;
		top: 30%;
		left: 57%;
		-webkit-transform: translateY( -10% ) translateX( -50% );
		transform: translateY( -10% ) translateX( -50% );
		writing-mode: vertical-rl;
		-webkit-writing-mode: vertical-rl;
		text-orientation: upright;
		-webkit-text-orientation: upright;
	}

	.差出人住所氏名 {
		position: absolute;
		top: 70mm;
		left: 3mm;
		writing-mode: vertical-rl;
		-webkit-writing-mode: vertical-rl;
		text-orientation: upright;
		-webkit-text-orientation: upright;
		line-height: 1.2em;
	}

	.差出人住所 {
		font-size: 10pt;
		margin: 0;
	}

	.差出人 {
		font-size: 12pt;
		letter-spacing: 0.1em;
		margin: 0;
		padding-top: 10mm;
	}

	.差出人郵便番号 {
		font-size: 10pt;
		margin: 0;
		padding: 0;
		position: absolute;
		top: 124mm;
		left: 6mm;
		letter-spacing: 2.1mm;
	}

	@media screen {
		body {
			background: #e0e0e0;
		}
		.sheet {
			background: white;
			box-shadow: 0 .5mm 2mm rgba( 0, 0, 0, .3 );
			margin: 5mm;
		}
	}

	@media print {
		body {
			width:100mm;
			height:148mm;
		}
	}
</style>

{{ range .Destinations }}
<section class="sheet">
	<p class="郵便番号">{{ .Postcode }}</p>
	<p class="住所">{{ noescape .Address }}</p>
	<div class="宛名">
		{{ range .Names }}
			<p>{{ . }}　様</p>
		{{ end }}
	</div>
	<div class="差出人住所氏名">
		<p class="差出人住所">{{ noescape $.Sender.Address }}</p>
		<p class="差出人">
			{{ range $.Sender.Names }}
				{{ . }}<br>
			{{ end }}
		</p>
	</div>
	<p class="差出人郵便番号">{{ $.Sender.Postcode }}</p>
</section>
{{ end }}

</body>
</html>

