class Solution:
    # @param A : list of list of integers
    # @param B : list of strings
    # @param C : list of strings
    # @param D : integer
    # @return an integer
    def solve(self, A, B, C, D):
        graph = [[] for _ in range(D)]
        for edge in A:
            graph[edge[0]].append(edge[1])
            graph[edge[1]].append(edge[0])

        def can_form_word(word, start_node, visited):
            if not word:
                return True

            visited[start_node] = True
            for neighbor in graph[start_node]:
                if not visited[neighbor] and B[neighbor] == word[0]:
                    if can_form_word(word[1:], neighbor, visited):
                        return True

            visted[start_node] = False
            return False

        alice_score = bob_score = 0
        for word in C:
            for i in range(D):
                if can_form_word(word, i, [False] * D):
                    if len(word) % 2 == 1:
                        alice_score += 1
                    else:
                        bob_score += 1
                    break

        return int(alice_score > bob_score)


if __name__ == "__main__":
    A = [[0, 1], [1, 2], [2, 3], [3, 4]]
    B = ["a", "b", "c", "d", "e"]
    C = ["abcde", "abcde", "abcde", "abcde"]
    D = 5
    print(Solution().solve(A, B, C, D))
