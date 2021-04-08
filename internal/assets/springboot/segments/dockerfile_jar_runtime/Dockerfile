
FROM openjdk:8
COPY --from=build_base /app/target/{{ .app_file }} app.jar
{{range $port := .port}}EXPOSE {{$port}}{{end}}
ENTRYPOINT ["java","-jar","/app.jar"]