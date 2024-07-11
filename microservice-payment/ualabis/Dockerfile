FROM node:12
RUN mkdir -p /opt/app
WORKDIR /opt/app
COPY *.json ./src
RUN npm install
COPY . .
EXPOSE 5000
CMD [ "npm", "start"]