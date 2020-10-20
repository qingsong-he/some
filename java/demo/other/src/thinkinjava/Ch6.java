package thinkinjava;

public class Ch6 extends foobar {
    // The access level cannot be lowered, but the access level can be increased
    @Override
    public void func1() {
        super.func1();
    }

    public static void main(String[] args) {
//        Sundae x = new Sundae();
        Sundae x = Sundae.makeASundae();
        new Ch6().func1();
    }
}

class Sundae {
    // If there is only this constructor, then the class cannot be inherited
    private Sundae() {
    }

    static Sundae makeASundae() {
        return new Sundae();
    }
}

class foobar {
    public int iVal;
    protected float fVal;
    double dVal;
    private String sVal;

    protected void func1() {
    }
}

class Soup1 {
    private Soup1() {
    }

    public static Soup1 makeSoup() {
        return new Soup1();
    }
}

class Soup2 {
    private Soup2() {
    }

    private static Soup2 ps1 = new Soup2();

    public static Soup2 access() {
        return ps1;
    }

    public void f() {
    }
}

