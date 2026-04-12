function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
const testCases: [number, number, number] = [
    [50, 9, 450],  // Quantity 9, below 10
    [50, 10, 500], // Quantity 10, below 50
    [50, 49, 2450], // Quantity 49, below 50
    [50, 50, 2500], // Quantity 50, exactly 50
];

// Function to run tests and print results
function runTests(testCases: [number, number, number]) {
    for (const [price, quantity, expected] of testCases) {
        const result = calculateDiscount(price, quantity);
        if (result === expected) {
            console.log(`PASS: Correct discount calculated for price=${price}, quantity=${quantity}`);
        } else {
            console.log(`FAIL: Expected ${expected}, got ${result}`);
        }
    }
}

// Run the tests
runTests(testCases);