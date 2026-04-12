function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function testDiscount() {
    const tests = [
        { name: "Buy less than 10 items at price 10", price: 10, quantity: 9, expected: 90 },
        { name: "Buy exactly 10 items at price 10", price: 10, quantity: 10, expected: 90 },
        { name: "Buy 49 items at price 10 (boundary)", price: 10, quantity: 49, expected: 44.1 },
        { name: "Buy exactly 50 items at price 10 (boundary)", price: 10, quantity: 50, expected: 37.5 }
    ];

    tests.forEach(test => {
        const actual = calculateDiscount(test.price, test.quantity);
        const passed = Math.abs(actual - test.expected) < 0.0001; // Use tolerance for floating point
        if (passed) {
            console.log(`PASS: ${test.name}`);
        } else {
            console.log(`FAIL: ${test.name} (expected: ${test.expected}, got: ${actual})`);
        }
    });
}

testDiscount();