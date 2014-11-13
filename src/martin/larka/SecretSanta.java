package martin.larka;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Random;
import java.util.Scanner;

public class SecretSanta {

	public static void main(String[] args) {

		Scanner keyboard = new Scanner(System.in);
		String name = "";
		String phoneNumber = "";
		String notSantaTo = "";

		ArrayList<Santa> santaList = new ArrayList<Santa>();

		boolean quitBool = false;

		while (!quitBool) {
			System.out.print("Enter new santas name: ");
			name = keyboard.nextLine();
			System.out.print("and phone number: ");
			phoneNumber = keyboard.nextLine();
			System.out.print("should santa not be santa to any santa?: ");
			notSantaTo = keyboard.nextLine();
			if (name.equals("")) {
				quitBool = true;
			} else {
				santaList.add(new Santa(name, phoneNumber, notSantaTo));
			}
		}
		
		ArrayList<Santa> shuffledList = new ArrayList<Santa>();

		// Ta godtycklig santa
		Random rand = new Random();
		shuffledList.add(santaList.get(rand.nextInt(santaList.size())));
		
		// Hämta Alla santas som den kan bli snata till
		
		for (int i=0; i<santaList.size()-1; i++) {
			System.out.println(santaList.get(i).getName() + " is secret santa to " + santaList.get(i+1).getName());
		}
		System.out.println(santaList.get(santaList.size()-1).getName() + " is secret santa to " + santaList.get(0).getName());

		keyboard.close();

	}

}