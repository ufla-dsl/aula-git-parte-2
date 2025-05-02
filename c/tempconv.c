#include <stdio.h>

typedef double Celsius;
typedef double Fahrenheit;

const Celsius AbsoluteZeroC = -273.15;
const Celsius FreezingC = 0;
const Celsius BoilingC = 100;

Fahrenheit CToF(Celsius c)
{
  return c * 9.0 / 5.0 + 32;
}

Celsius FToC(Fahrenheit f)
{
  return (f - 32) * 5.0 / 9.0;
}

void printC(Celsius c)
{
  printf("%.1f°C\n", c);
}

void printF(Fahrenheit f)
{
  printf("%.1f°F\n", f);
}

int main()
{
  Celsius c = 25;
  Fahrenheit f = CToF(c);
  printF(f);

  f = 77;
  c = FToC(f);
  printC(c);
  return 0;
}
