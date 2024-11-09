# Stage 1: Build the React application
FROM node:18-alpine as build

# Set working directory
WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy all frontend source code
COPY . .

# Build the application
RUN npm run build

# Stage 2: Serve the application
FROM node:18-alpine

# Install serve to run the application
RUN npm install -g serve

WORKDIR /app

# Copy the build output from previous stage
COPY --from=build /app/build ./build

# Expose port 3000
EXPOSE 3000

# Start the application
CMD ["serve", "-s", "build", "-l", "3000"]