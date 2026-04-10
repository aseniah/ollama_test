function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runDiscountTests() {
    const testCases: [number, number, number] = [
        [10, 10, 100], // Boundary condition: quantity = 9
        [10, 11, 110], // Quantity just above 10
        [50, 50, 2500], // Boundary condition: quantity = 49
        [50, 51, 2625], // Quantity just above 50
        [100, 10, 1000], // Large quantity
        [100, 100, 10000], // Large quantity
        [100, 50, 5000]  // Small quantity
    ];

    for (const [price, quantity, expected] of testCases) {
        const actual = calculateDiscount(price, quantity);
        console.log(
            actual === expected ? `PASS: Calculate discount for price ${price} and quantity ${quantity}` :
            `FAIL: Calculate discount for price ${price} and quantity ${quantity}. Expected: ${expected}, got: ${actual}`
        );
    }
}

// Run the tests
runDiscountTests();