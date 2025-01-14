## Stage 1 build
FROM golang:1.22.3-alpine AS build

# Set working directory to /src
WORKDIR /src

# Copy nessecary files for build process
COPY app ./


# Build the golang application into /src
RUN go build -o /ascii-art-web .



## Stage 2
FROM alpine


# Set working directory to /app
WORKDIR /app

# Copy artifacts from stage 1
COPY --from=build /ascii-art-web .

RUN apk add bash

# Copy assets from host
COPY assets assets/
COPY Banners Banners/
COPY templates templates/

# Expose port 6500
EXPOSE 6500


## Start Command
CMD [ "./ascii-art-web" ]




## Metadata
LABEL \
  version="1.0.0" \
  description="A Docker image for the ASCII Art Web application"
