# Stage 1: Build
FROM node:20-bullseye AS build

WORKDIR /app

# Copiar package.json y package-lock.json para cache
COPY package*.json ./

# Instalar dependencias
RUN npm install --legacy-peer-deps

# Copiar todo el proyecto
COPY . .

# Construir la app
RUN npm run build

# Stage 2: Nginx
FROM nginx:alpine

# Copiar build al nginx html folder
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
