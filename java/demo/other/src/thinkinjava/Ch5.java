package thinkinjava;

public class Ch5 {
    int[] iValList = {1, 2, 3, 4, 5}; // init
    int[] iValList1 = new int[]{1, 2, 3, 4, 5};
    Object oVal = new Object(); // init

    int iVal = 0; // init

    Ch5() {
        this(0);

        // Init in Constructor
        iVal = 1;
    }

    // Overload Constructor
    Ch5(int iVal) {
        this.iVal = iVal;
    }

    void func1() {

    }

    // Overload func1
    void func1(int iVal) {

    }

    void func1(int... iValList) {
        for (int i : iValList) {
            System.out.println(i);
        }
    }

    public static void main(String[] args) {
        new Ch5().func1(1, 2, 3);
        Spiciness howHot = Spiciness.MEDIUM;
        for (Spiciness s : Spiciness.values()) {
            System.out.println(s + ", ordinal " + s.ordinal());
        }
    }
}

enum Spiciness {
    NOT, MILD, MEDIUM, HOT, FLAMING
}
