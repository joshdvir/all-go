FROM scratch

ADD main /all-go

EXPOSE 9292

CMD ["/all-go"]
