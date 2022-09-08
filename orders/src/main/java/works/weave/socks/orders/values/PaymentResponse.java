package works.weave.socks.orders.values;

public class PaymentResponse {
    private boolean authorised = false;
    private String  message;

    private String isFraud;

    // For jackson
    public PaymentResponse() {
    }

    public PaymentResponse(boolean authorised, String message, String isFraud) {
        this.authorised = authorised;
        this.message = message;
        this.isFraud = isFraud;
    }

    @Override
    public String toString() {
        return "PaymentResponse{" +
                "authorised=" + authorised +
                ", message=" + message +
                ", isFraud=" + isFraud +
                '}';
    }

    public boolean isAuthorised() {
        return authorised;
    }

    public void setAuthorised(boolean authorised) {
        this.authorised = authorised;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public String getMessage() {
        return message;
    }

    public String getIsFraud() {
        return isFraud;
    }

   public String setIsFraud(String isFraud) {
        this.isFraud = isFraud;
    }

    public String getIsFraud() {
        return isFraud;
    }

}