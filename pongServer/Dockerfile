FROM eclipse-temurin:21-jdk-alpine AS build

WORKDIR /app

COPY PongServer.java .

RUN mkdir out && javac -d out PongServer.java && jar cfe PongServer.jar PongServer -C out .


FROM gcr.io/distroless/java21-debian12

WORKDIR /app

COPY --from=build /app/PongServer.jar .

CMD ["PongServer.jar"]


