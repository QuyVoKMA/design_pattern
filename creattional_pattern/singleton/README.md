# Định nghĩa
1. là một design pattern đảm bảo một class chỉ có duy nhất một instance(khởi tạo)
2. Cung cấp một cách toàn cục để truy cập tới instance đo
#Các ứng dụng:
1. Chúng ta muốn sử dụng lại kết nối đến database khi truy vấn
2. Khi chúng ta mở một kết nối SSH đến một server để thực hiện một số công việc và chúng ta không muốn mở một kết nối khác mà sử dụng lại kết nối trước đó
3. Khi chúng ta cần hạn chế sự truy cập đến một vài biến số, chúng ta sử dụng mẫu singleton là trung gian để truy cập biến này