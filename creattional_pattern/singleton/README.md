# Singleton (Khá phổ biến)
# Khi nào sữ dụng:
1. Khi chúng ta lập trình với những chương trình lớn đôi khi chúng ta có những struct mà chúng ta phải khỏi tạo các đối tượng cho nó tốn rất nhiều tài nguyên, khi chúng ta vận hành thì các hành vi khởi tạo đều giống nhau.
ví dụ như chúng ta có object giao tiếp với DB, objecct đọc ghi phai, object ghi log khi ứng dụng chúng ta chạy. Khi đó chúng ta có nhu cầu khỏi tạo một object duy nhất xuyên suốt cái chương trình chúng ta cho một struct.

 --> Singleton đáp ứng được nhu cầu đó của chúng ta

# Các thành phần:

 
# Định nghĩa
1. là một design pattern đảm bảo một class chỉ có duy nhất một instance(khởi tạo)
2. Cung cấp một cách toàn cục để truy cập tới instance đo
#Các ứng dụng:
1. Chúng ta muốn sử dụng lại kết nối đến database khi truy vấn
2. Khi chúng ta mở một kết nối SSH đến một server để thực hiện một số công việc và chúng ta không muốn mở một kết nối khác mà sử dụng lại kết nối trước đó
3. Khi chúng ta cần hạn chế sự truy cập đến một vài biến số, chúng ta sử dụng mẫu singleton là trung gian để truy cập biến này

