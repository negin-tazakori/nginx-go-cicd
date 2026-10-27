FROM nginx:alpine

COPY index.html /usr/share/ngnix/html/index.html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

ENTRYPOINT ["nginx"]
CMD ["-g", "daemon off;"]
