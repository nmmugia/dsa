package com.dsa.assessments.glob_matching;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.DisplayName;
import static org.assertj.core.api.Assertions.*;

@DisplayName("Glob Matching Tests")
class SolutionTest {

    private final Solution solution = new Solution();

    @Test
    @DisplayName("Test Case 1: Basic wildcard and question mark")
    void testCase1() {
        String fileName = "abcdefg";
        String pattern = "a*e?g";
        assertThat(solution.globMatching(fileName, pattern)).isTrue();
    }

    @Test
    @DisplayName("Test Case 2: Pattern ending with wildcard and question mark")
    void testCase2() {
        String fileName = "abcdagvvvvag";
        String pattern = "*?ag";
        assertThat(solution.globMatching(fileName, pattern)).isTrue();
    }

    @Test
    @DisplayName("Test Case 3: Pattern not matching")
    void testCase3() {
        String fileName = "abcdagvvvvavg";
        String pattern = "*?ag";
        assertThat(solution.globMatching(fileName, pattern)).isFalse();
    }

    @Test
    @DisplayName("Test Case 4: Multiple wildcards and question marks")
    void testCase4() {
        String fileName = "abcdagvvgaavg";
        String pattern = "*?aa?g";
        assertThat(solution.globMatching(fileName, pattern)).isTrue();
    }

    @Test
    @DisplayName("Test Case 5: No wildcards, no match")
    void testCase5() {
        String fileName = "abcdefg";
        String pattern = "abcddefg";
        assertThat(solution.globMatching(fileName, pattern)).isFalse();
    }

    @Test
    @DisplayName("Test Case 6: Empty strings")
    void testCase6() {
        String fileName = "";
        String pattern = "";
        assertThat(solution.globMatching(fileName, pattern)).isTrue();
    }

    @Test
    @DisplayName("Test Case 7: Empty fileName with question mark")
    void testCase7() {
        String fileName = "";
        String pattern = "?";
        assertThat(solution.globMatching(fileName, pattern)).isFalse();
    }

    @Test
    @DisplayName("Test Case 8: Multiple wildcards not matching")
    void testCase8() {
        String fileName = "abcdefg";
        String pattern = "*c*cdefg";
        assertThat(solution.globMatching(fileName, pattern)).isFalse();
    }

    @Test
    @DisplayName("Test Case 9: Wildcard in middle")
    void testCase9() {
        String fileName = "abcdefg";
        String pattern = "abcd*efg";
        assertThat(solution.globMatching(fileName, pattern)).isTrue();
    }

    @Test
    @DisplayName("Test Case 10: Single question mark not matching")
    void testCase10() {
        String fileName = "abcdefg";
        String pattern = "?";
        assertThat(solution.globMatching(fileName, pattern)).isFalse();
    }

    @Test
    @DisplayName("Test Case 11: Complex pattern with many wildcards")
    void testCase11() {
        String fileName = "abcdefg";
        String pattern = "*****a*?f*********g";
        assertThat(solution.globMatching(fileName, pattern)).isTrue();
    }

    @Test
    @DisplayName("Test Case 12: Empty fileName with multiple wildcards")
    void testCase12() {
        String fileName = "";
        String pattern = "***";
        assertThat(solution.globMatching(fileName, pattern)).isTrue();
    }

    @Test
    @DisplayName("Test Case 13: Non-empty fileName with empty pattern")
    void testCase13() {
        String fileName = "abcdefg";
        String pattern = "";
        assertThat(solution.globMatching(fileName, pattern)).isFalse();
    }
}
