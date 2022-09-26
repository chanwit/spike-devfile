FROM node:lts-slim

# Needed for node 15+ see https://stackoverflow.com/a/65443098
WORKDIR /usr/app

RUN npm install express
RUN npm install ioredis
RUN npm install cors

COPY files/ ./files/
COPY webui.js .
CMD ["node", "webui.js"]
EXPOSE 8080
