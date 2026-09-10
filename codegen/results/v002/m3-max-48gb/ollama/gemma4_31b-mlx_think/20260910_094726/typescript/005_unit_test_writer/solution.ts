/**
 * Calculates a bulk discount based on the quantity of items.
 * - Less than 10 items: No discount.
 * - 10 to 49 items: 10% discount.
 * - 50 or more items: 25% discount.
 */
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

interface TestCase {
    description: string;
    price: number;
    quantity: number;
    expected: number;
}

function runTests() {
    const tests: TestCase[] = [
        {
            description: "Boundary condition: quantity 9 (no discount)",
            price: 100,
            quantity: 9,
            expected: 900,
        },
        {
            description: "Boundary condition: quantity 10 (10% discount)",
            price: 100,
            quantity: 10,
            expected: 900, // 100 * 10 * 0.9
        },
        {
            description: "Boundary condition: quantity 49 (10% discount)",
            price: 100,
            quantity: 49,
            expected: 4410, // 100 * 49 * 0.9
        },
        {
            description: "Boundary condition: quantity 50 (25% discount)",
            price: 100,
            quantity: 50,
            expected: 3750, // 100 * 50 * 0.75
        },
        {
            description: "Standard case: small quantity",
            price: 10,
            quantity: 2,
            expected: 20,
        },
        {
            description: "Standard case: large quantity",
            price: 10,
            quantity: 100,
            expected: 750, // 10 * 100 * 0.75
        },
    ];

    tests.forEach((test) => {
        const actual = calculateDiscount(test.price, test.quantity);
        
        // We use a small epsilon for floating point comparisons to avoid precision issues
        if (Math.abs(actual - test.expected) < 0.000001) {
            console.log(`PASS: ${test.description}`);
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${actual})`);
        }
    });
}

// Run the test suite
runTests();