#include <iostream>
#include <iomanip>

class Celsius
{
public:
  double value;
  Celsius(double v) : value(v) {}
  void print() const
  {
    std::cout << std::fixed << std::setprecision(1) << value << "°C\n";
  }
};

class Fahrenheit
{
public:
  double value;
  Fahrenheit(double v) : value(v) {}
  void print() const
  {
    std::cout << std::fixed << std::setprecision(1) << value << "°F\n";
  }
};

const Celsius AbsoluteZeroC(-273.15);
const Celsius FreezingC(0);
const Celsius BoilingC(100);

Fahrenheit CToF(const Celsius &c)
{
  return Fahrenheit(c.value * 9 / 5 + 32);
}

Celsius FToC(const Fahrenheit &f)
{
  return Celsius((f.value - 32) * 5 / 9);
}

int main()
{
  Celsius c(25);
  Fahrenheit f = CToF(c);
  f.print();

  f = Fahrenheit(77);
  c = FToC(f);
  c.print();

  return 0;
}
