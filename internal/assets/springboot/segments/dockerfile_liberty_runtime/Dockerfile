# Run App
FROM registry.access.redhat.com/jboss-eap-6/eap64-openshift:latest
{{range $port := .port}}EXPOSE {{$port}}
{{end}}
COPY --from=build_base /{{ .app_name }}/target/* /opt/eap/standalone/deployments/