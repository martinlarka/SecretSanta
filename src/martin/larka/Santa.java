package martin.larka;

public class Santa {
	private String name;
	private String phoneNumber;
	private String notSantaTo;
	
	public Santa(String name, String phoneNumber, String notSantaTo) {
		this.name = name;
		this.phoneNumber = phoneNumber;
		this.notSantaTo = notSantaTo;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public String getPhoneNumber() {
		return phoneNumber;
	}

	public void setPhoneNumber(String phoneNumber) {
		this.phoneNumber = phoneNumber;
	}

	public String getNotSantaTo() {
		return notSantaTo;
	}

	public void setNotSantaTo(String notSantaTo) {
		this.notSantaTo = notSantaTo;
	}
	
	public String toString() {
		return name + " " + phoneNumber;
	}
}
