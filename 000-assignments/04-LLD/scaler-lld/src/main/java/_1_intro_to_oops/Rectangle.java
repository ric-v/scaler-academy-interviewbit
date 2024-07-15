package _1_intro_to_oops;

public class Rectangle {
    Point topLeft;
    int height;
    int width;

    public Rectangle(Point topLeft, int height, int width) {
        this.topLeft = topLeft;
        this.height = height;
        this.width = width;
    }

    public int area() {
        return height * width;
    }

    public int perimeter() {
        return 2 * (height + width);
    }

    public Point getBottomRight() {
        Point bottomRight = new Point(1, 2);

        bottomRight.x = topLeft.x + width;
        bottomRight.y = topLeft.y + height;

        return bottomRight;
    }

    public static void main(String[] args) {

        Rectangle rectangle = new Rectangle(new Point(0, 0), 10, 20);
        System.out.println(rectangle.area());
        System.out.println(rectangle.perimeter());
    }
}
