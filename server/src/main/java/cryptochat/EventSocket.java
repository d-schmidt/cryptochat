package cryptochat;

import java.io.IOException;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

import org.eclipse.jetty.websocket.api.Session;
import org.eclipse.jetty.websocket.api.WebSocketAdapter;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

public class EventSocket extends WebSocketAdapter {

    private static String CMD_MESSAGE = "msg";
    private static String CMD_REGISTER = "reg";
    private static String CMD_CONNECT = "con";
    private static String CMD_SUCCESS = "success";
    private static String CMD_ERROR = "error";
    private static String CMD_PUBLICKEY = "pk";
    private static String CMD_LIST = "list";
    private static String CMD_PING = "ping";
    private static String CMD_PONG = "pong";
    private static String CMD_NAME_IN_USE = "name in use";
    private static String CMD_USER_NOT_FOUND = "user not found";
    private static String CMD_USER_DC = "user disconnected";
    private static Gson gson = new GsonBuilder().create();
    private static Map<String, EventSocket.User> sessions = new ConcurrentHashMap<>();

    private String name;

    private static class User {
        private String pk;
        private String name;
        private Session session;
        private Long crc;

        public User(String pk, String name, Session session, Long crc) {
            this.pk = pk;
            this.name = name;
            this.session = session;
            this.crc = crc;
        }
    }

    private static class Message {
        private String Target;
        private String Name;
        private String CMD;
        private String PK;
        private Long CRC;
        private Integer ID;

        public Message(String cmd, Integer id) {
            CMD = cmd;
            ID = id;
        }

        public Message(String cmd, String pK, Integer id) {
            CMD = cmd;
            PK = pK;
            ID = id;
        }
    }

    @Override
    public void onWebSocketConnect(Session sess) {
        super.onWebSocketConnect(sess);
        System.out.println("Socket Connected: " + sess);
    }

    @Override
    public void onWebSocketText(String message) {
        super.onWebSocketText(message);
        System.out.println("Received TEXT message: " + message);
        Message msg = gson.fromJson(message, Message.class);

        try {
            if (msg != null) {
                if (CMD_REGISTER.equals(msg.CMD)) {
                    if (name != null) {
                        sessions.remove(name);
                    }

                    if (sessions.containsKey(msg.Name)) {
                        getSession().getRemote().sendString(gson.toJson(new Message(CMD_NAME_IN_USE, msg.ID)));
                    }
                    else {
                        sessions.put(msg.Name, new User(msg.PK, msg.Name, getSession(), msg.CRC));
                        name = msg.Name;
                        Message respmsg = new Message(CMD_REGISTER, msg.ID);
                        respmsg.CRC = msg.CRC;
                        getSession().getRemote().sendString(gson.toJson(respmsg));
                    }
                } else if (CMD_CONNECT.equals(msg.CMD)) {
                    User target = sessions.get(msg.Target);
                    if (target != null) {
                        getSession().getRemote().sendString(gson.toJson(new Message(CMD_PUBLICKEY, target.pk, msg.ID)));
                    }
                    else {
                        getSession().getRemote().sendString(gson.toJson(new Message(CMD_USER_NOT_FOUND, msg.ID)));
                    }
                } else if (CMD_MESSAGE.equals(msg.CMD)) {
                    User target = sessions.get(msg.Target);
                    if (target != null && msg.CRC.equals(target.crc)) {
                        target.session.getRemote().sendString(message);
                    }
                    else {
                        getSession().getRemote().sendString(gson.toJson(new Message(CMD_USER_DC, msg.ID)));
                    }
                } else if (CMD_PING.equals(msg.CMD)) {
                    getSession().getRemote().sendString(gson.toJson(new Message(CMD_PONG, msg.ID)));
                }
            }
            else {
                getSession().getRemote().sendString(gson.toJson(new Message(CMD_ERROR, null)));
            }
        } catch (IOException e) {
            System.out.println("io error for " + name + ": " + e.getMessage());
            if (getSession().isOpen()) {
                getSession().close(1, "io error " + e.getMessage());
            }
            sessions.remove(name);
        }
    }

    @Override
    public void onWebSocketClose(int statusCode, String reason) {
        super.onWebSocketClose(statusCode, reason);
        System.out.println("Socket Closed: [" + statusCode + "](" + name + ") " + reason);
        sessions.remove(name);
    }

    @Override
    public void onWebSocketError(Throwable cause) {
        super.onWebSocketError(cause);
        System.out.println("Socket Error: [](" + name + ") " + cause.getMessage());
        cause.printStackTrace(System.err);
        sessions.remove(name);
    }
}
