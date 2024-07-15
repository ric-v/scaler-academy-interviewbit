package _2_access_modifiers_constructors;

public class Point {
    int x;
    int y;

    public Point(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public Point(Point c) {
        this.x = c.x;
        this.y = c.y;
    }
}
