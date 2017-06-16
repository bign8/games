FROM scratch
ADD server /
ADD cmd/server/tpl/ /cmd/server/tpl/
ADD cmd/server/www/ /cmd/server/www/
CMD ["/server"]
EXPOSE 4000
