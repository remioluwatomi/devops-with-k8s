import com.sun.net.httpserver.HttpServer;
import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.Headers;

import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;

import java.util.concurrent.atomic.AtomicInteger;

public class PongServer {
    private HttpServer server;

    public static void main(String[] args) {
        String cliPort = (args.length > 0) ? args[0] : null;
        int port = resolvePort(cliPort);

        PongServer pongServer = new PongServer(port);
        System.out.println("server is about to start on port: " + port);
        pongServer.server.start();
    }

    public PongServer(int port) {
        try {
            server = HttpServer.create(new InetSocketAddress(port), 0);
            server.createContext("/pingpong", new MyHandler());
            server.setExecutor(null);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static int resolvePort(String portNum) {
        int port = 8881;

        if (portNum != null && !portNum.isBlank()) {
            port = Integer.parseInt(portNum);
        }
        Integer envPort = getEnvPort();
        if (envPort != null) {
            port = envPort;
        }
        return port;
    }

    // this is envPort and it will overwrite positional port if set
    static Integer getEnvPort() {
        String envPort = System.getenv("PORT");
        if (envPort != null && !envPort.isBlank()) {
            return Integer.parseInt(envPort);
        }
        return null;
    }

    static class MyHandler implements HttpHandler {
        private AtomicInteger requestCounter = new AtomicInteger(0);

        public void handle(HttpExchange t) throws IOException {
            String json = "{\"ping\": \"pong0\", \"counter\": " + requestCounter.incrementAndGet() + "}\n";

            byte[] body = json.getBytes();

            Headers jsonHeaders = t.getResponseHeaders();
            jsonHeaders.add("Content-Type", "application/json; charset=UTF-8");

            t.sendResponseHeaders(200, body.length);
            try(OutputStream os = t.getResponseBody()) {
                os.write(body);
            }

         }
    }
}
