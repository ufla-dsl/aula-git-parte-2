-- Módulo principal
module Main where

-- Define o valor do zero absoluto em Celsius
absoluteZeroC :: Double
absoluteZeroC = -273.15

-- Define a temperatura de ebulição da água em Celsius
boilingC :: Double
boilingC = 100.0

-- Função que converte Celsius para Fahrenheit
cToF :: Double -> Double
cToF c = (c * 9 / 5) + 32

-- Função principal
main :: IO ()
main = do
    -- Imprime a temperatura do zero absoluto formatada
    putStrLn $ "  Que frio!     " ++ formatTemp absoluteZeroC "°C"
    
    -- Imprime a temperatura em Fahrenheit
    putStrLn $ " Fervendo!     " ++ formatTemp (cToF boilingC) "°F"

-- Função para formatar os valores com 2 casas decimais e a unidade
formatTemp :: Double -> String -> String
formatTemp value unit = show (fromIntegral (round (value * 100)) / 100) ++ unit
