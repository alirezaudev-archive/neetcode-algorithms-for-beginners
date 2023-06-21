class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

    @classmethod
    def fromArray(cls, tree: list) -> 'TreeNode':
        if tree is None:
            return cls()
        root = cls(tree[0])
        layer = [root]
        i = 1

        while i < len(tree):
            tempLayer = []
            for node in layer:
                if i < len(tree) and tree[i] is not None:
                    node.left = cls(tree[i])
                    tempLayer.append(node.left)
                i += 1
                if i < len(tree) and tree[i] is not None:
                    node.right = cls(tree[i])
                    tempLayer.append(node.right)
                i += 1
            layer = tempLayer

        return root

    def print(self, sep='۰') -> 'TreeNode':
        height = self.height()
        col = self.__getCol(height)
        M = [[None for _ in range(col)] for __ in range(height)]
        self.__printTree(M, self, col // 2, 0, height)
        for i in M:
            for j in i:
                if j is None:
                    print(sep, end=sep)
                else:
                    print(j, end=sep)
            print("")
        return self

    def height(self) -> int:
        return self.__height(self)

    def __height(self, root) -> int:
        if root is None:
            return 0
        return max(self.__height(root.left), self.__height(root.right)) + 1

    def __getCol(self, height):
        if height == 1:
            return 1
        return self.__getCol(height - 1) + self.__getCol(height - 1) + 1

    def __printTree(self, M, root, col, row, height):
        if root is None:
            return
        M[row][col] = root.val
        self.__printTree(M, root.left, col - pow(2, height - 2), row + 1, height - 1)
        self.__printTree(M, root.right, col + pow(2, height - 2), row + 1, height - 1)
