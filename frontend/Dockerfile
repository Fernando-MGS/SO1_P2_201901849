# syntax=docker/dockerfile:1
FROM node:16

# Creando el espacio de trabajo
WORKDIR /app

# Copiando dependencias
COPY package.json ./
COPY package-lock.json ./

# Ejecucion de npm i
RUN npm install

#
COPY . .

EXPOSE  3000

#Levantar front
CMD ["npm", "start"]