package martin.larka;

import java.io.*;
import java.net.*;
import javax.servlet.*;
import javax.servlet.http.*;
 
public class SendSMS extends HttpServlet {
	private static final long serialVersionUID = 1L;
	private static String mosms_username = "ditt-användarnamn";
	private static String mosms_password = "ditt-lösenord";
	private static String mosms_url = "http://www.mosms.com/se/sms-send.php";
	private static String mosms_number = "0701234567";
	private static String mosms_type = "text";
	private static String mosms_data = "Hello world!";
 
	public void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
		PrintWriter out = response.getWriter();
		try {
			URL url = new URL(mosms_url + "?username=" + URLEncoder.encode(mosms_username, "ISO-8859-1")
					+ "&password=" + URLEncoder.encode(mosms_password, "ISO-8859-1")
					+ "&nr=" + mosms_number + "&type="
					+ mosms_type + "&data="	+ URLEncoder.encode(mosms_data, "ISO-8859-1"));
 
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
				out.println("SMS skickat korrekt!");
			} else {
				out.println("Fel vid anrop:" + result.toString());
			}
		} catch (Exception e) {
			out.println("Kunde inte ansluta till server");
		}
	}
}