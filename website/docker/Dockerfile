FROM 192.168.88.1:8888/library/nginx:1.21.1

COPY --chown=nginx:nginx dist /usr/share/nginx/html/release
COPY default.conf /etc/nginx/conf.d/
WORKDIR /usr/share/nginx/html/release
