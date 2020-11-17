class Node:
    """
    Node is the class where we have 
    the certain methos to build a tree
    such as :
        constructor (__init__)
        insert
        find_value
        numberOfLeaves
        depth
        findMin
    """
    def __init__(self, data):

        """
        __init__ is a default construct which takes 
                input data(any type of data)
                self.left :
                    class global variable which links 
                    up the left node
                self.data:
                    class global variable which links 
                    up the right node
                self.data :
                    class global variable which store
                    the data or the values
        """
        self.left = None
        self.right = None
        self.data = data


    def insert(self, data):
        """
        insert : is the class method, which appends the new
                value into the BSRT
                inputs :
                    data:
                        data is the value
        """
        if self.data:
            # if the current data is or node is greter than new value 
            #  attach it to left
            if data < self.data:
                if self.left is None:
                    self.left = Node(data)
                else:
                    self.left.insert(data)
            # if the current data is lesser than new value 
            #  attach it to right
            elif data > self.data:
                if self.right is None:
                    self.right = Node(data)
                else:
                    self.right.insert(data)
        else:
            self.data = data


    def find_value(self, value):

        if value < self.data:
            if self.left is None:
                return str(value) + "value is not found"
            return self.left.find_value(value)
        elif value > self.data:
            if self.right is None:
                return str(value) + "value is not found"
            return self.right.find_value(value)
        else:
            return str(self.data) + " is found"


    def numberOfLeaves(self):

        # Q1
        #  finding the number of leaves
        couter = 0
        if self.left is None and self.right is None:
            couter += 1
        if self.left:
            couter += self.left.numberOfLeaves()
        if self.right:
            couter += self.right.numberOfLeaves()
        return couter


    def depth(self):

        # Q2
        # finding the depth of BST
        if self.right and self.left:
            return 1 + max(self.left.depth(), self.right.depth())
        elif self.left:
            return 1 + self.left.depth()
        elif self.right:
            return 1 + self.right.depth()
        else:
            return 1


    def findMin(self):

        # Q3
        if self.left:
            return self.left.findMin()
        else:
            return self.data

    def postorder(self):
        
        # Q4
        if self.left:
            self.left.postorder()
        if self.right:
            self.right.postorder()
        print (str(self.data))


if __name__ == "__main__":
    #  NODES_ is the accept list to collectively creat a Tree
    NODES_ = [29, 21, 43, 19, 24, 39, 56, 23, 25, 36, 40, 55, 76, 77]
    #  variable defined to intialize the class 
    root = None
    # looping over the all the nodes 
    for i, val in enumerate(NODES_):
        #  if the index of the i == 0 then lets consider it as root node 
        if i == 0:
            print("Adding the Root Node in to the ques", val)
            # Initializing the Root Node
            root = Node(val)
        else:
            #  Else it a child or descended Node
            print("Adding a new Node in the Binary Search", val)
            # Initializing the child Node
            root.insert(val)
    # print(root.find_value(27))
    print(" Number of leaves ", root.numberOfLeaves())
    print("Height of the tree is ", root.depth())
    print("Minimum value of the tree is", root.findMin())
    print("post Order values")
    root.postorder()