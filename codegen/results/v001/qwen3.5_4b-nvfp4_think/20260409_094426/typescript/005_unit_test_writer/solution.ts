// Define the function provided in the prompt
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Read command-line arguments with process.argv
const args = process.argv.slice(2);
const defaultPriceStr = '100';
let testPriceString: string | undefined;
if (args.length > 0) {
    testPriceString = args[0];
} else {
    // Fallback to standard argument if none provided
    testPriceString = defaultPriceStr;
}

// Parse price into number type
const basePrice: number = parseFloat(testPriceString || defaultPriceStr);

// Define test cases covering boundary conditions specified in requirements
// Cases: Quantity 9 (no discount), 10 (min discount tier), 49 (max non-discount tier), 50 (new discount tier)
const testCases = [
    { quantity: 9, desc: 'quantity < 10 (No Discount Tier)' },
    { quantity: 10, desc: 'quantity >= 10 and < 50 (First Discount Tier)' },
    { quantity: 49, desc: 'quantity < 50 boundary (Second Discount Tier)' },
    { quantity: 50, desc: 'quantity >= 50 (Third Discount Tier)' }
];

// Run the test suite
console.log("Running discount calculation tests...\n");

testCases.forEach(({ quantity, desc }) => {
    const expectedValue = calculateDiscount(basePrice, quantity);
    const actualResult = calculateDiscount(basePrice, quantity);
    
    // Since we are testing a function we implemented directly from the spec above, 
    // theoretically expected should equal actual. We compare to verify stability.
    // To simulate potential failures in a generic test:
    let passed = true;
    // Note: We rely on the implementation of calculateDiscount being consistent with our internal 'expected' calculation 
    // However, since we copy the function source code, we can assert equality against ourselves.
    // A better simulation for testing robustness would be to check if the value matches mathematical expectations of the tiers.
    
    let expectedResult: number;
    if (quantity < 10) {
        expectedResult = basePrice * quantity;
    } else if (quantity < 50) {
        expectedResult = basePrice * quantity * 0.9;
    } else {
        expectedResult = basePrice * quantity * 0.75;
    }

    // We strictly check for float equality issues
    const isMatch = Math.abs(expectedResult - actualResult) < 1e-5;
    
    if (isMatch) {
        console.log(`PASS: ${desc}`);
    } else {
        console.log(`FAIL: ${desc} (expected: ${expectedResult}, got: ${actualResult})`);
    }
});

console.log("");
console.log("Test suite completed using process.argv for configuration.");