export class Celsius {
  constructor(public value: number) { }
  toString(): string {
    return `${this.value.toFixed(1)}°C`;
  }
}

export class Fahrenheit {
  constructor(public value: number) { }
  toString(): string {
    return `${this.value.toFixed(1)}°F`;
  }
}

export const AbsoluteZeroC = new Celsius(-273.15);
export const FreezingC = new Celsius(0);
export const BoilingC = new Celsius(100);

export function CToF(c: Celsius): Fahrenheit {
  return new Fahrenheit(c.value * 9 / 5 + 32);
}

export function FToC(f: Fahrenheit): Celsius {
  return new Celsius((f.value - 32) * 5 / 9);
}
