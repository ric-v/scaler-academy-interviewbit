package _1_intro_to_oops;

public class Student {
    int age;
    String name;

    public void display() {
        System.out.println("My name is " + name + ". I am " + age + " years old.");
    }

    public void sayHello() {
        System.out.println("Hello " + name + "!");
    }

    public static void main(String[] args) {
        Student student = new Student();
        student.name = "John";
        student.age = 20;
        student.display();
        student.sayHello();
    }
}
