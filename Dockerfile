FROM scratch

ADD ./bin/simulation-linux-amd64 /simulation

CMD ["/simulation"]
