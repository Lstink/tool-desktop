FROM golang:windowsservercore-ltsc2022

WORKDIR /app

COPY ../../../ .

RUN go mod tidy

RUN go install fyne.io/fyne/v2/cmd/fyne

RUN fyne package -os windows -icon icon.png