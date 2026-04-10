function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases: [price, quantity, expectedResult, description]
const testCases: Array<[number, number, number, string]> = [
    // Boundary condition: Just below 10 (no discount)
    [10, 9, 90, 'Boundary quantity 9 (No discount)'],
    
    // Boundary condition: Exactly 10 (10% discount)
    [10, 10, 90, 'Boundary quantity 10 (10% discount)'],
    
    // Boundary condition: Just below 50 (10% discount)
    [100, 49, 4410, 'Boundary quantity 49 (10% discount)'],
    
    // Boundary condition: Exactly 50 (25% discount)
    [100, 50, 3750, 'Boundary quantity 50 (25% discount)'],

    // Extra case: Large quantity to ensure logic holds
    [200, 100, 15000, 'Large quantity 100 (25% discount)']
];

console.log("Running bulk discount tests...\n");

testCases.forEach((testCase) => {
    const [price, quantity, expected, description] = testCase;
    const actual = calculateDiscount(price, quantity);

    // Using a small epsilon for floating point comparison safety, 
    // though these specific numbers are exact in JS.
    const isPass = Math.abs(actual - expected) < 0.00001;

    if (isPass) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${actual})`);
    }
});