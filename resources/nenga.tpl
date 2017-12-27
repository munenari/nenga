<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="UTF-8">
<title>年賀状宛名</title>
</head>
<body>
<style>
	@font-face {
		font-family: 'Noto Serif Japanese';
		font-style: normal;
		font-weight: 400;
		src: url(https://yourfontserver/NotoSerifCJKjp-hinted/NotoSerifCJKjp-Light.otf) format('opentype');
	}
	@page {
		size: 100mm 148mm;
		margin: 0;
		padding: 0;
	}

	body {
		margin: 0;
		padding: 0;
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
		font-family: "Noto Serif Japanese";
		margin: 0;
		padding: 0;
		position: absolute;
		top: 14mm;
		left: 46mm;
		letter-spacing: 5.0mm;
	}

	.住所 {
		font-family: "Noto Serif Japanese";
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
		right: 10mm;
	}

	.宛名 {
		font-family: "Noto Serif Japanese";
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
		left: 5mm;
		writing-mode: vertical-rl;
		-webkit-writing-mode: vertical-rl;
		text-orientation: upright;
		-webkit-text-orientation: upright;
	}

	.差出人住所 {
		font-family: "Noto Serif Japanese";
		font-size: 10pt;
		margin: 0;
		padding-left: 0.3em;
	}

	.差出人 {
		font-family: "Noto Serif Japanese";
		font-size: 12pt;
		letter-spacing: 0.1em;
		margin: 0;
		padding-top: 0.5em;
	}

	.差出人郵便番号 {
		font-size: 10pt;
		font-family: "Noto Serif Japanese";
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
		{{ range $.Sender.Names }}
			<p class="差出人">{{ . }}</p>
		{{ end }}
	</div>
	<p class="差出人郵便番号">{{ $.Sender.Postcode }}</p>
</section>
{{ end }}

</body>
</html>

