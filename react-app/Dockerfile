# Imagen oficial de Node.js como base
FROM node:20-alpine

WORKDIR /app

# Copia el package.json y el package-lock.json al contenedor
COPY package*.json ./

# Instala las dependencias de la aplicación
RUN npm install

# Copia el resto del código de la aplicación al contenedor
COPY . .

# Compila la aplicación React para prod
RUN npm run build

# Expone el puerto 3000 para la app
EXPOSE 3000

# Ejecuta la app cuando el contenedor se inicia
CMD ["npm", "start"]
