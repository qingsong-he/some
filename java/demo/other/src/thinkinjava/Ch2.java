package thinkinjava;

/**
 * A class comment
 *
 * @author ...
 * @version 0.1
 */
public class Ch2 {
    /**
     * A field comment
     * <pre>
     *     System.out.println(new Date());
     * </pre>
     */
    public boolean bnVal; // Boolean false
    char cVal; // Character '\u0000'
    byte bVal; // Byte 0
    short sVal; // Short 0
    int iVal; // Integer 0
    long lVal; // Long 0L
    float fVal; // Float 0.0f
    double dVal; // Double 0.0d
    Object o;

    /**
     * A method comment
     *
     * @param
     * @throws
     */
    protected void func1() {
        int iVal; // must be initialized before use
    }

    static int iVal1;

    static void func2() {
    }

    public static void main(String[] args) {
        var o1 = new Ch2();
        System.out.println(o1.dVal); // 0.0
        System.out.println(o1.o); // null
        {
            int x = 12;
            {
//                int x = 96; // Illegal
            }
        }

        System.out.println(iVal1); // 0
        func2();

        System.getProperties().list(System.out);
    }
}
