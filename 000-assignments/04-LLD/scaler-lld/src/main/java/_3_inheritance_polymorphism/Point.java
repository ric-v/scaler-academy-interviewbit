package _3_inheritance_polymorphism;

public class Point {
    protected int x;
    protected int y;
    /*
    When you add the static keyword to a variable, it means that the variable is no longer
    tied to an instance of the class. This means that only one instance of that static
    member is created which is shared across all instances of the class. Static variables
    are stored in heap memory just like non-static variables.
     */
    public static String name = "Point";

    public Point(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public void display() {
        System.out.println(x + " " + y);
    }

    public static void displayStatic() {
        System.out.println(name);
    }
}
