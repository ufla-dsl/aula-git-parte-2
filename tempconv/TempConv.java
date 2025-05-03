package tempconv;

public class TempConv {

    public static final double ABSOLUTE_ZERO_C = -273.15;
    public static final double FREEZING_C = 0.0;
    public static final double BOILING_C = 100.0;

    public static class Celsius {
        private final double value;

        public Celsius(double value) {
            this.value = value;
        }

        public double getValue() {
            return value;
        }

        @Override
        public String toString() {
            return String.format("%.2f°C", value);
        }
    }

    public static class Fahrenheit {
        private final double value;

        public Fahrenheit(double value) {
            this.value = value;
        }

        public double getValue() {
            return value;
        }

        @Override
        public String toString() {
            return String.format("%.2f°F", value);
        }
    }

    public static Fahrenheit cToF(Celsius c) {
        return new Fahrenheit(c.getValue() * 9 / 5 + 32);
    }

    public static Celsius fToC(Fahrenheit f) {
        return new Celsius((f.getValue() - 32) * 5 / 9);
    }
}
