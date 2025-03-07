import { describe, expect, it } from 'bun:test'
const isPrime = require("../src/structy/002-is-prime/approch1")

describe('is the provided number Prime', () => {
    it.each([[2, true], [3, true], [4, false], [5, true], [6, false], [7, true]
        , [8, false], [25, false], [31, true], [2017, true], [2048, false], [1, false]
    ])('should be able to find if %i prime assert is %b', (val, res) => {
        expect(isPrime(val)).toBe(res);
    })
})