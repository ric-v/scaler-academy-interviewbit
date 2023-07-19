class Rectangle:
    # Define properties here
    width = 0
    height = 0
    
    # Define constructor here
    def __init__(self, width, height):
        self.width = width
        self.height = height

    def perimeter(self):
        return 2 * (self.width + self.height)
    
    def area(self):
        return self.width * self.height
    