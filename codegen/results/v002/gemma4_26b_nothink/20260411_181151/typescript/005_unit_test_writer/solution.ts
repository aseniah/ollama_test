/**
 * Calculates a bulk discount based on the quantity of items.
 * - quantity < 10: No discount
 * - 10 <= quantity < 50: 10% discount (0.9 multiplier)
 * - quantity >= 50: 25% discount (0.75 multiplier)
 * 
 * @param price The price per unit
 * @param quantity The number of units
 * @returns The total price after applying discounts
 */
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

/**
 * Interface representing a single test case
 */
interface TestCase {
    description: string;
    price: number;
    quantity: number;
    expected: number;
}

/**
 * Main execution function to run the tests
 */
function runTests() {
    const testCases: TestCase[] = [
        {
            description: "Boundary condition: quantity 9 (No discount)",
            price: 100,
            quantity: 9,
            expected: 900
        },
        {
            description: "Boundary condition: quantity 10 (10% discount starts)",
            price: 100,
            quantity: 10,
            expected: 900 // 100 * 10 * 0.9
        },
        {
            description: "Boundary condition: quantity 49 (10% discount applies)",
            price: 100,
            quantity: 49,
            expected: 4410 // 100 * 49 * 0.9
        },
        {
            description: "Boundary condition: quantity 50 (25% discount starts)",
            price: 100,
            quantity: 50,
            expected: 3750 // 100 * 50 * 0.75
        },
        {
            description: "Standard case: quantity 100 (25% discount)",
            price: 10,
            quantity: 100,
            expected: 750 // 10 * 100 * 0.75
        }
    ];

    let passedCount = 0;
    let failedCount = 0;

    testCases.forEach((test) => {
        const result = calculateDiscount(test.price, test.quantity);
        
        // Using a small epsilon for floating point comparison
        const isPass = Math.abs(result - test.expected) < 0.000001;

        if (isPass) {
            console.log(`PASS: ${test.description}`);
            passedCount++;
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
            failedCount++;
        }
    });

    console.log(`\nTest Summary: ${passedCount} passed, ${failedCount} failed.`);
    
    // Exit with non-zero code if there are failures
    if (failedCount > 0) {
        process.exit(1);
    }
}

// Run the program
runTests();