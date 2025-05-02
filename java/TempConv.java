public class TempConv {

    public static class Celsius {
        double value;
        Celsius(double value) { this.value = value; }
        public String toString() { return String.format("%.1f°C", value); }
    }

    public static class Fahrenheit {
        double value;
        Fahrenheit(double value) { this.value = value; }
        public String toString() { return String.format("%.1f°F", value); }
    }

    public static final Celsius AbsoluteZeroC = new Celsius(-273.15);
    public static final Celsius FreezingC = new Celsius(0);
    public static final Celsius BoilingC = new Celsius(100);

    public static Fahrenheit CToF(Celsius c) {
        return new Fahrenheit(c.value * 9 / 5 + 32);
    }

    public static Celsius FToC(Fahrenheit f) {
        return new Celsius((f.value - 32) * 5 / 9);
    }

    public static void main(String[] args) {
        System.out.println(CToF(new Celsius(25)));  // 77.0°F
        System.out.println(FToC(new Fahrenheit(77)));  // 25.0°C
    }
}
