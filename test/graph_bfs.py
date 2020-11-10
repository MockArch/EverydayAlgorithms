# Breadth-First Search Algorithm
# Valid for any tree graph

from pprint import pprint
from collections import defaultdict, deque


class Tree:
    def __init__(self):  # Constructor function
        self.Tree = defaultdict(list)  # default dictionary to store graph

    def Branch(self, u, v):  # Build the required branches for the Tree
        self.Tree[u].append(v)

    def __str__(self):

        pprint(self.Tree)
        return ""

    def bfs(self, root):

        visited, queue = set(), deque([root])
        visited.add(root)

        while queue:

            # Dequeue a vertex from queue
            vertex = queue.popleft()
            print(str(vertex) + " ", end="")

            # If not visited, mark it as visited, and
            # enqueue it
            for neighbour in self.Tree[vertex]:
                if neighbour not in visited:
                    visited.add(neighbour)
                    queue.append(neighbour)
            print("Queue :", list(queue))


# print('\n')

# Build the Tree graph
if __name__ == "__main__":
    Net = Tree()
    Net.Branch(0, 1)
    Net.Branch(0, 2)
    Net.Branch(0, 3)
    Net.Branch(1, 4)
    Net.Branch(2, 5)
    Net.Branch(2, 6)
    Net.Branch(3, 7)
    Net.Branch(3, 8)
    Net.Branch(4, 4)
    Net.Branch(5, 1)
    Net.Branch(5, 4)
    Net.Branch(6, 9)
    Net.Branch(7, 10)
    Net.Branch(8, 10)
    Net.Branch(9, 8)
    Net.Branch(9, 11)
    Net.Branch(10, 10)
    Net.Branch(11, 11)

    print("The visited nodes starting from vertex 0", Net)
    Net.bfs(0)