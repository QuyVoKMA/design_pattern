# Visitor (Người viến thăm)
# Khi nào dùng:
1. KHi chúng ta có nhu cầu thêm vào hành vi hoặc phương thức mới cho một đối tượng hoặc struct mà ko muốn thay đổi cấu trúc của đối tượng hay struct đó thì dùng -->Visitor
# Cách dùng:
1. Thêm vào các visitor mới.
# Thành phần:
    1. element inteface:Chứa các hành vi chung của đối tượng mà chúng ta muốn thêm hành vi mới vào
        + accept(visitor interface)
    2. Concrete element A
    3. ncrete element B
    3. Concrete element c
    5. Visitor inteface (quan trọng) --> chứa các hành vi chung mà mỗi hành vi này nó sẽ thực hiện với các tham số đầu vào (A,B,C)
        + visit for element A
        + visit for element B
        + visit for element C
    6. Concrete visitor 1
    7. Concrete visitor 2

# Ví dụ: (Các số tương ứng với các thành phần ở trên)
    1. shape(hình học)
    2. A: square
    3. B: circle
    4. V: rectangle
    5. visitor
    6. areaCal (Tính diện tích)
    7. PerimeterCal (Tính chu vi)

# Thực hành.