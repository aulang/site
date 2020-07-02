FROM scratch

ADD site site
ADD config.yml config.yml

EXPOSE 8081

ENTRYPOINT ["/site"]
