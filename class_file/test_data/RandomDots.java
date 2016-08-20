import java.util.Random;

public class RandomDots {
	private static final int w = 20;
	private static final int h = 7;
	private static Random rand;

	private static char getDot() {
		if (rand.nextInt(100) > 50) {
			return '#';
		}
		return ' ';
	}

	public static void main(String[] args) {
		rand = new Random();
		for (int y = 0; y < h; y++) {
			for (int x = 0; x < w; x++) {
				System.out.printf("%c", getDot());
			}
			System.out.print("\n");
		}
	}
}
