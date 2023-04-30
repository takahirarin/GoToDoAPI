# ベースとなるDockerイメージ指定
FROM golang:latest
# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/work

COPY ./app/go.mod /go/src/work/go.mod
COPY ./app/go.sum /go/src/work/go.sum

# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/work

RUN go mod download
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/work