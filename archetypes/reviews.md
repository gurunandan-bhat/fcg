{{ $title := replace (replaceRE `-\d+$` "" .File.ContentBaseName)  "-" " " | title }}
+++
title = '{{  $title }}'
date = {{ .Date }}
draft = false
mreviews = ['{{ $title }}']
critics = ['']
subtitle = ""
opening = ""
img = '{{ .File.ContentBaseName }}.'
type = 'print'
source = ""
+++
