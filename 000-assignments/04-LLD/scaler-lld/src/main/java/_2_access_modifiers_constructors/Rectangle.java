package _2_access_modifiers_constructors;

public class Rectangle {
    Point topLeft;
    Point bottomRight;

    public Rectangle(int topLeftX, int topLeftY, int bottomRightX, int bottomRightY) {
        this.topLeft = new Point(topLeftX, topLeftY);
        this.bottomRight = new Point(bottomRightX, bottomRightY);
    }

    public Rectangle(Point topLeft, Point bottomRight) {
        this.topLeft = new Point(topLeft);
        this.bottomRight = new Point(bottomRight);
    }

    public Rectangle(Rectangle r) {
        this.topLeft = new Point(r.topLeft);
        this.bottomRight = new Point(r.bottomRight);
    }

    public static void main(String[] args) {

        Rectangle r = new Rectangle(0, 0, 10, 10);
        System.out.println(r.topLeft.x + " " + r.topLeft.y);
        System.out.println(r.bottomRight.x + " " + r.bottomRight.y);

        Rectangle r2 = new Rectangle(r);
        System.out.println(r2.topLeft.x + " " + r2.topLeft.y);
        System.out.println(r2.bottomRight.x + " " + r2.bottomRight.y);

        Rectangle r3 = new Rectangle(r.topLeft, r.bottomRight);
        System.out.println(r3.topLeft.x + " " + r3.topLeft.y);
        System.out.println(r3.bottomRight.x + " " + r3.bottomRight.y);
    }
}
