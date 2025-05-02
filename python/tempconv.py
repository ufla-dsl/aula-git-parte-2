class Celsius(float):
    def __str__(self):
        return f"{self:.1f}°C"

class Fahrenheit(float):
    def __str__(self):
        return f"{self:.1f}°F"

AbsoluteZeroC = Celsius(-273.15)
FreezingC = Celsius(0)
BoilingC = Celsius(100)

def CToF(c: Celsius) -> Fahrenheit:
    return Fahrenheit(c * 9 / 5 + 32)

def FToC(f: Fahrenheit) -> Celsius:
    return Celsius((f - 32) * 5 / 9)
