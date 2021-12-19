// RandomDotsSimple does not rely on any external libraries except for
// System.out.print(char).
public class RandomDotsSimple {
	private static final int w = 20;
	private static final int h = 7;
	private static long rngState = 1337;

	// Always returns a positive 32-bit value.
	private static long xorShift32() {
		long mask = 0xffffffffL;
		long x = rngState & mask;
		x ^= x << 13;
		x &= mask;
		x ^= x >> 17;
		x &= mask;
		x ^= x << 5;
		x &= mask;
		rngState = x;
		return x;
	}

	private static char getDot() {
		char toReturn = ' ';
		switch (((int) xorShift32()) % 11) {
			case 1:
				toReturn = '$';
				break;
			case 2:
				toReturn = '%';
				break;
			case 3:
				toReturn = '^';
				break;
			case 4:
				toReturn = '&';
				break;
			case 5:
				toReturn = '*';
				break;
			default:
				toReturn = ' ';
				break;
		}
		return toReturn;
	}

	public static void main(String[] args) {
		for (int y = 0; y < h; y++) {
			for (int x = 0; x < w; x++) {
				System.out.print(getDot());
			}
			System.out.print('\n');
		}
	}
}
