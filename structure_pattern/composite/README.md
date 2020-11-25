# Composite (Tổng hợp)

# Khi nào sữ dụng.
1. Khi chúng ta có nhu cầu xử lý 1 nhóm các đối tượng nhưng chung ta muốn xữ lý dựa trên 1 đối tượng duy nhất ta dùng composite.

# Dùng
 1. Component inteface  ==>component  (chứa các hành vi chung con cái composite và cái leaf)
 2. Composite -->folder (Chính là một cái struct để chúng ta cài đặc các thành phần trong đó).
 3. Leaf -->file (Cài đặc các hành vi của component inteface tuy nhiên nó lại nằm trong composite như là 1 cái array hoăc map)

Khi mình muốn xử lý một hành vi nào đó mà mình định nghĩa trong component này. Thay vì mình tạo các đối tượng cho cái file này rồi gọi tất cả các file ra thì mình dùng cái composite(folder) để thao tác.