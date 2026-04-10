function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Function to run tests and report results
function runTests() {
    const testCases = [
        { description: "Quantity 9", quantity: 9 },
        { description: "Quantity 10", quantity: 10 },
        { description: "Quantity 49", quantity: 49 },
        { description: "Quantity 50", quantity: 50 }
    ];

    for (const test of testCases) {
        const result = calculateDiscount(100, test.quantity); // Assuming price is 100 for simplicity
        const expected = calculateDiscount(100, test.quantity);
        
        if (result === expected) {
            console.log(`PASS: ${test.description}`);
        } else {
            console.log(`FAIL: ${test.description} (expected: ${expected}, got: ${result})`);
        }
    }
}

// Run the tests
runTests();