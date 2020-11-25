# Proxy (Người đại diện)
# Khi nào được sữ dụng nào ?
1. Khi có nhu cầu muốn xây dựng đối tượng được ủy quền hoặc thay thế cho các đối tượng khác
# Mục đích:
1. Dùng proxy đó để lưu tạm các dữ liệu trả về từ đối tượng khác --> được gọi là catch proxy
2. BẢo vệ cho các yêu cầu mà không được tín nhiệm. -->Security proxy
3. Tăng hiệu suất của ứng dụng -->Vitual proxy
4. Bổ xung thêm chức năng trước khi thực thi phương thức gốc
5. Tạo ra đối tượng mới có chức năng cao hơn đối tượng ban đầu
6. Giảm chi phí khi có nhiều truy cập vào đối tượng có chi phí khởi tạo ban đầu lớn

# Thành phần:
1. subject interface (Chứa các hành vi mà chúng ta truyền vào proxy)
2. proxy struct (Nó sẽ cài đăt tất cả các hành vi trong subject interface)
3.real subject (Nó là đối tượng được proxy ủy quyền hoặc thay thế)

--> Với cái real subject này thì chúng ta lúc nào cũng thông qua proxy trước, tức là chúng ta gọi một phương thức của real subject thì chúng ta gọi đến 1 phương thức của proxy, sau đó proxy sẽ thực hiện một số logic.
Tùy theo chức năng của proxy là gì, sau khi thực hiện xong hành vi đó nó ms chuyển tiếp đến lời gọi hành vi của real subject.


# Ví dụ: 
1. Chúng ta sẽ thực hiện một ví dụ minh họa xây dựng một máy server ảo, khi chúng ta truy cập web server thì nó sẽ có các đối tượng sau:
    

       i. Subject interfaace ->server
       ii. proxy -> ngix
       iii. real subject ->application
==> ví dụ như các browser nó  gửi cái http request lên server thì nó sẽ đi ra proxy(ngix, ...) sau khi nó xử lý xong nó sẽ forward qua cái application.
