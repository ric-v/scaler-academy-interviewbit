package _3_inheritance_polymorphism;

public class ThreedPoint extends Point {
    private final int z;

    public ThreedPoint(int x, int y, int z) {
        super(x, y);
        this.z = z;
    }

    @Override
    public void display() {
        System.out.println(x + " " + y + " " + z);
    }
}
