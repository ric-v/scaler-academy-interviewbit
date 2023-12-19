def dp(A, B):
    """Computes the maximum tastiness that can be achieved by partitioning the first `i` chocolates into `j` partitions.

    Args:
      A: A list of chocolates, where A[i] denotes the tastiness of the i-th chocolate.
      B: The number of partitions to make.

    Returns:
      The maximum tastiness that can be achieved.
    """

    n = len(A)
    dp = [[0] * (B + 1) for _ in range(n + 1)]

    # Initialize the base cases.
    for i in range(n + 1):
        dp[i][0] = 0

    for j in range(B + 1):
        dp[0][j] = 0

    # Compute the dynamic programming table.
    for i in range(1, n + 1):
        for j in range(1, B + 1):
            dp[i][j] = max(dp[i - 1][j - 1] + A[i - 1] + A[i], dp[i - 1][j])

    # Return the maximum tastiness that can be achieved.
    return dp[n][B]


def maximum_tastiness(A, B):
    """Finds the maximum tastiness that can be achieved by partitioning the array A into B partitions.

    Args:
      A: A list of chocolates, where A[i] denotes the tastiness of the i-th chocolate.
      B: The number of partitions to make.

    Returns:
      The maximum tastiness that can be achieved.
    """

    return dp(A, B)


print(maximum_tastiness([1, 2, 3, 4, 5], 2))
