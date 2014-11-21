package martin.larka;

import java.io.*;
import java.net.*;
import javax.servlet.http.*;
 
public class SendSMS extends HttpServlet {
	private static final long serialVersionUID = 1L;
	private static String mosms_username = "";
	private static String mosms_password = "";
	private static String mosms_url = "http://www.mosms.com/se/sms-send.php";
	private static String mosms_type = "text";
 
	public void doSendSMS(Santa giver, Santa reciver) throws IOException {
		try {
			URL url = new URL(mosms_url + "?username=" + URLEncoder.encode(mosms_username, "ISO-8859-1")
					+ "&password=" + URLEncoder.encode(mosms_password, "ISO-8859-1")
					+ "&nr=" + giver.getPhoneNumber() + "&type=" + mosms_type
					+ "&allowlong=1"
					+ "&data="	+ URLEncoder.encode(generateMessage(giver.getName(), reciver.getName()), "ISO-8859-1"));
 
			URLConnection connection = url.openConnection();
			connection.setDoOutput(true);
			connection.setReadTimeout(5000);
			connection.connect();
 
			BufferedReader in = new BufferedReader(new InputStreamReader(connection.getInputStream()));
			StringWriter result = new StringWriter();
 
			while (in.ready()) {
				result.write(in.read());
			}
 
			in.close();
			result.close();
 
			if (result.toString().equals("0")) {
				System.out.println("SMS skickat korrekt!");
			} else {
				System.out.println("Fel vid anrop:" + result.toString());
			}
		} catch (Exception e) {
			System.out.println("Kunde inte ansluta till server");
		}
	}
	
	private String generateMessage(String giver, String reciver) {
		return "Hej " + giver + " det här är ett meddelande från tomtens underrättelsecenter. Vi har ett hemligt uppdrag åt dig. Du ska nämligen vara hemligtomte åt " + reciver + " i år. Det betyder att du ska ge " + reciver + " en julklapp (för ca 500kr) på julafton. /Agent Tomten";
	}
}