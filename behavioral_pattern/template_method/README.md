# Template method (Một khung mẫu cho các hành vi)
# Khi nào dùng:
1. Khi có từ 2 đối tượng trở lên, và tất cả các kiểu dlieu này có cùng hành mà chúng sẽ được sd trong một thuật giải.
--> nên sd design patter này để đưa các cách và thứ tự sd của các hành vi này vào 1 vị trí.

# Thành phần:
1. method interface thành phần chứa các hành vi chung của các kiểu dữ liệu mà chúng ta cài đặt cho DPatter này
2. concrete object 1 (mục tiêu các cái này sẽ cài đặt các method trong medthod interface)
3. concrete object 2
4. template method (quan trong) --> nó sẽ giúp chúng ta sd các hành vi trong hàm này theo thứ tự chúng ta mong muốn

# Ví dụ:
    1. IOtp -> chứa các hành vi khi thực hiện thanh toán chúng ta có mã opt, nhiều hình thức chúng ta nhận, nhưng chúng tạo ra như nhau.
        + genRandomOTP
        + saveOTPCache
        + getMessage
        + sendNotification --> gửi cho người sd.
        
    2. SMS
    3. Email
    4. OTP struct (Nó sẽ sữ dụng các hành vi của IOtp theo thứ tự) 
    