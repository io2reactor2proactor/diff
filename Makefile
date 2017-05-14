all:
  go build -x -o main *.go
clean:
  rm -fr main
