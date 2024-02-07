{{/* Generate basic labels */}}
{{- define "simplebank.v1.labels" }}
    labels:
        generator: helm
        deployedby: bruce
        date: {{ now | htmlDate }}
{{- end }}