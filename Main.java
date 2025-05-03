import tempconv.TempConv;
import tempconv.TempConv.Celsius;
import tempconv.TempConv.Fahrenheit;

public class Main {
    public static void main(String[] args) {
        System.out.println("Que frio! " + new Celsius(TempConv.ABSOLUTE_ZERO_C));

        Celsius boiling = new Celsius(TempConv.BOILING_C);
        Fahrenheit boilingF = TempConv.cToF(boiling);
        System.out.println("Fervendo! " + boilingF);
    }
}
