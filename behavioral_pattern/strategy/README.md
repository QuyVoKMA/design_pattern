# Strategy (Chiến lược chiến thuật)
# Khi nào dùng.
1. Khi chúng ta có một context hoặc đối tương mà chúng ta muốn cài đặt các chiến lược khác nhau cho đối tượng đó, mỗi chiến lược lại có các hành vi giống nhau, tuy nhiên cách cài đặt lại khác nhau.
Lưu ý: có thể thay đổi các chiến lược tại lúc runtime.
# Các thành phần
1. Context
2. strately inteface --> Định nghĩa các hành vi giống nhau giữa tất cả các strategy phía sau
3. concrete strategy 1
4. concrete strategy 2
5. concrete strategy 3

# Ví dụ:
1. Cache -> Bộ nhớ tạm, mục tính tăng tốc độ ghi và truy xuất data, thông thường cache sẽ được lưu trữ trên RAM hoặc CACHE trên máy
.Cái này phải giới hạn cho nó. Nên khi đưa vào các phần tử mà cache đã đầy thì chúng ta phải có chiến lược đã lưu trữ bên trong đó.
2. FIFO: fisrt in first out
3. LRU: Lấy ra các phần tử nào mà sữ dụng ít nhất.
4. LFU: phần tử xuất hiện ít nhất.