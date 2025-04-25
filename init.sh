#!/bin/bash

#Estructura de ficheros y directorios para comenzar la programación Web en Go
#Crear los directorios
mkdir -p \
    cmd/main \
    internal/config \
    internal/handlers \
    internal/models \
    internal/routes \
    web/templates \
    web/statics/css \
    web/statics/js \

#Para crear los archivos
touch \
    cmd/main/main.go \
    internal/config/config.go \
    internal/handlers/handlers.go \
    internal/models/models.go \
    internal/routes/routes.go \
    web/main.go \
    web/templates/home.html \
    web/templates/error.html \
    web/statics/css/custom.css \
    web/statics/js/custom.js \



