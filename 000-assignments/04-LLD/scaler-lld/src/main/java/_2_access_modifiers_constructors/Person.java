package _2_access_modifiers_constructors;

public class Person {
    int age;
    String name;

    public Person(int age, String name) {
        this.age = age;
        this.name = name;
    }

    public static void main(String[] args) {

        Person person = new Person(20, "John");
        System.out.println(person.age);
        System.out.println(person.name);
    }
}
