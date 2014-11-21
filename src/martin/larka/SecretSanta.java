package martin.larka;

import java.io.IOException;
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

		SendSMS sendSMS = new SendSMS();

		while (!quitBool) {
			System.out.print("Enter new santas name: ");
			name = keyboard.nextLine();
			System.out.print("and phone number: ");
			phoneNumber = keyboard.nextLine();
			System.out.print("should santa not be santa to?: ");
			notSantaTo = keyboard.nextLine();
			if (name.equals("")) {
				quitBool = true;
			} else {
				santaList.add(new Santa(name, phoneNumber, notSantaTo));
			}
			System.out.println();
		}

		ArrayList<Santa> shuffledSantaList = new ArrayList<Santa>();
		Collections.shuffle(santaList);

		Random rand = new Random();

		shuffledSantaList.add(santaList.get(0));
		santaList.remove(0);

		while (!santaList.isEmpty()) {
			ArrayList<Santa> possibleSantas = (ArrayList<Santa>)santaList.clone();
			for (Santa santa: santaList) {
				if (santa.getName().equals(shuffledSantaList.get(shuffledSantaList.size()-1).getNotSantaTo())) {
					possibleSantas.remove(santa);
				}
			}
			int randInt = rand.nextInt(possibleSantas.size());
			shuffledSantaList.add(possibleSantas.get(randInt));
			santaList.remove(possibleSantas.get(randInt));
		}
		// Shuffle the two last santas if the last santa cant be secret santa to the first
		if (shuffledSantaList.get(0).getName().equals(shuffledSantaList.get(shuffledSantaList.size()-1).getNotSantaTo())) {
			shuffledSantaList.add(shuffledSantaList.remove(shuffledSantaList.size()-2));
		}

		for (int i=0; i<shuffledSantaList.size()-1; i++) {
			//System.out.println(shuffledSantaList.get(i).getName() + " is secret santa to " + shuffledSantaList.get(i+1).getName());
			try {
				sendSMS.doSendSMS(shuffledSantaList.get(i), shuffledSantaList.get(i+1));
			} catch (IOException e) {}
		}
		//System.out.println(shuffledSantaList.get(shuffledSantaList.size()-1).getName() + " is secret santa to " + shuffledSantaList.get(0).getName());
		try {
			sendSMS.doSendSMS(shuffledSantaList.get(shuffledSantaList.size()-1), shuffledSantaList.get(0));
		} catch (IOException e) {}

		keyboard.close();

	}	
	
}