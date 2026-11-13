package com.dsa.arrays.two_number_sum;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.DisplayName;
import static org.assertj.core.api.Assertions.*;

import java.util.Arrays;

@DisplayName("Two Number Sum Tests")
class SolutionTest {

    private final Solution solution = new Solution();

    @Test
    @DisplayName("Test Case 1: Basic case")
    void testCase1() {
        int[] array = {3, 5, -4, 8, 11, 1, -1, 6};
        int targetSum = 10;
        int[] result = solution.twoNumberSum(array, targetSum);
        Arrays.sort(result);
        assertThat(result).containsExactly(-1, 11);
    }

    @Test
    @DisplayName("Test Case 2: First two elements")
    void testCase2() {
        int[] array = {4, 6, 1, -3};
        int targetSum = 10;
        int[] result = solution.twoNumberSum(array, targetSum);
        Arrays.sort(result);
        assertThat(result).containsExactly(4, 6);
    }

    @Test
    @DisplayName("Test Case 3: No solution")
    void testCase3() {
        int[] array = {1, 2, 3, 4, 5};
        int targetSum = 20;
        int[] result = solution.twoNumberSum(array, targetSum);
        assertThat(result).isEmpty();
    }

    @Test
    @DisplayName("Test Case 4: Negative numbers")
    void testCase4() {
        int[] array = {-7, -5, -3, -1, 0, 1, 3, 5, 7};
        int targetSum = -8;
        int[] result = solution.twoNumberSum(array, targetSum);
        Arrays.sort(result);
        assertThat(result).containsExactly(-7, -1);
    }

    @Test
    @DisplayName("Test Case 5: Large numbers")
    void testCase5() {
        int[] array = {15, 2, 8, 11};
        int targetSum = 17;
        int[] result = solution.twoNumberSum(array, targetSum);
        Arrays.sort(result);
        assertThat(result).containsExactly(2, 15);
    }
}
