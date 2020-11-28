# Memeto
# Khi nào dùng:
- Khi chúng ta có nhu cầu muốn lưu trữ các trạng thái thay đổi theo thời gian thì ta dùng cái design pattern này.

# Thành phần:
1. originator (Chính là 1 cái struct mà các đối tượng mà tạo ra từ cái struct này mà chúng ta muốn lưu trữ lại cái trạng thái của nó)
2. memeto (Nó lưu trữ các trạng thái (Cũ mới) của đối tượng originator thì các trạng thái cũ của nó sẽ lưu vào đây)
3. caretaker (CẦu nối giữa memeto vs originator)

# Ví dụ:
1. Tạo 1 cái struct lưu chuỗi, sau đó tạo các đối tượng cảu orig
2. SAu đó mỗi lần thay đổi chúng ta lưu trạng thái vào memeto
3. Dùng cái này lấy lại các trạng thái cũ mà orig thay đổi