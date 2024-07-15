package _3_inheritance_polymorphism;

public class Client {

    public static void main(String[] args) {
        Point point = new Point(2, 3);
        point.display();
        Point.displayStatic();
        ThreedPoint threedPoint = new ThreedPoint(1, 2, 3);
        threedPoint.display();
    }
}
