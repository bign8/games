FROM scratch
ADD cmd/server/www/ /cmd/server/www/
ADD cmd/server/tpl/ /cmd/server/tpl/
ADD server /
ENTRYPOINT ["/server"]
EXPOSE 4000
